"""Wire-format tests of the generated wire.gen.py against the Go and TypeScript
generators' contract."""
import importlib.util
import io
import json
import pathlib
import sys

import httpx

_path = pathlib.Path(__file__).parent / "wire.gen.py"
_spec = importlib.util.spec_from_file_location("wire_gen", _path)
api = importlib.util.module_from_spec(_spec)
sys.modules[_spec.name] = api
_spec.loader.exec_module(api)


def _call(impl, method, body):
    """Posts a JSON body to the generated server and returns (status, JSON)."""
    raw = json.dumps(body).encode()
    status = []
    environ = {
        "REQUEST_METHOD": "POST",
        "PATH_INFO": f"/rpc/Wire/{method}",
        "CONTENT_TYPE": "application/json",
        "CONTENT_LENGTH": str(len(raw)),
        "wsgi.input": io.BytesIO(raw),
    }
    out = b"".join(api.WireServer(impl)(environ, lambda s, h: status.append(int(s.split()[0]))))
    return status[0], json.loads(out)


def _client(respond):
    """A generated client whose requests are answered by respond(request)."""
    http = httpx.Client(transport=httpx.MockTransport(respond))
    return api.WireClient("http://x", client=http)


def test_enum_wire_values():
    data = api.Data(k=api.Kind.Admin, c=api.Country.CA)
    assert data.to_dict() == {"k": "ADMIN", "c": "Canada"}
    assert api.Data.from_dict({"k": "ADMIN", "c": "Canada"}) == data
    assert api.Kind.from_dict("User") is api.Kind.User


def test_succinct_method_uses_bare_structs():
    seen = []

    def respond(request):
        seen.append(json.loads(request.content))
        return httpx.Response(200, json={"b": "out"})

    got = _client(respond).flatten(api.FlattenRequest(a="in"))
    assert seen == [{"a": "in"}]
    assert got == api.FlattenResponse(b="out")

    class Impl:
        def flatten(self, flattenRequest):
            return api.FlattenResponse(b=flattenRequest.a + "!")

    assert _call(Impl(), "Flatten", {"a": "in"}) == (200, {"b": "in!"})


def test_succinct_method_with_keyword_named_argument():
    seen = []

    def respond(request):
        seen.append(json.loads(request.content))
        return httpx.Response(200, json={"b": "out"})

    assert _client(respond).go(api.Import(a="in")) == api.FlattenResponse(b="out")
    assert seen == [{"a": "in"}]

    class Impl:
        def go(self, import_):
            return api.FlattenResponse(b=import_.a + "!")

    assert _call(Impl(), "Go", {"a": "in"}) == (200, {"b": "in!"})


def test_aliases_of_user_types():
    wire = {"users": [{"id": 1}], "nested": [{"id": 2}], "kind": "User"}
    team = api.Team.from_dict(wire)
    assert team == api.Team(users=[api.User(id=1)], nested=[api.User(id=2)], kind=api.Kind.User)
    assert team.to_dict() == wire
    assert api.K2 is api.Kind


def test_unset_optionals_are_omitted():
    assert api.Opt(need="n").to_dict() == {"need": "n"}
    assert api.Opt(need="n", may="m").to_dict() == {"need": "n", "may": "m"}
    assert api.Opt.from_dict({"need": "n", "may": None}) == api.Opt(need="n")

    seen = []

    def respond(request):
        seen.append(json.loads(request.content))
        return httpx.Response(200, json={"z": None, "w": "w"})

    assert _client(respond).maybe(None, "y") == (None, "w")
    assert seen == [{"y": "y"}]

    class Impl:
        def maybe(self, x, y):
            assert x is None
            return None, "w"

    assert _call(Impl(), "Maybe", {"x": None, "y": "y"}) == (200, {"w": "w"})


def test_reserved_names_in_methods_and_arguments():
    seen = []

    def respond(request):
        seen.append(json.loads(request.content))
        return httpx.Response(200, json={"ok": True})

    client = _client(respond)
    assert client.from_("c") is True
    assert client.import_() is None
    assert client.get("s") is True
    assert seen == [{"class": "c"}, {}, {"self": "s"}]

    class Impl:
        def from_(self, class_):
            return class_ == "c"

        def import_(self):
            pass

        def get(self, self_):
            return self_ == "s"

    assert _call(Impl(), "From", {"class": "c"}) == (200, {"ok": True})
    assert _call(Impl(), "Import", {}) == (200, {})
    assert _call(Impl(), "Get", {"self": "s"}) == (200, {"ok": True})


def test_methods_named_like_server_internals():
    class Impl:
        def respond(self):
            return "r"

        def error(self):
            return "e"

        def headers(self):
            return "h"

    for method in ("Respond", "Error", "Headers"):
        want = {"v": method[0].lower()}
        assert _call(Impl(), method, {}) == (200, want), method


def test_unknown_error_code_keeps_what_the_server_sent():
    body = {"error": "Custom", "code": 777, "msg": "m", "cause": "c", "status": 422}
    try:
        _client(lambda request: httpx.Response(422, json=body)).get("s")
    except api.WebRPCError as err:
        got = (type(err), err.name, err.code, err.message, err.cause, err.status)
    else:
        raise AssertionError("expected an error")
    assert got == (api.WebRPCError, "Custom", 777, "m", "c", 422)


def test_unusable_responses_are_bad_responses():
    cases = {
        "redirect without body": httpx.Response(302),
        "error status with a JSON array": httpx.Response(500, json=["x"]),
        "error status with HTML": httpx.Response(502, content=b"<html>bad gateway</html>"),
        "error status with a JSON object without a code": httpx.Response(500, json={"msg": "x"}),
        "success with a JSON array": httpx.Response(200, json=[1]),
        "success with HTML": httpx.Response(200, content=b"<html>"),
        "success with an undecodable value": httpx.Response(200, json={"kind": "NOPE"}),
    }
    for name, response in cases.items():
        try:
            _client(lambda request: response).get_kind()
        except api.WebrpcBadResponse as err:
            assert err.cause, name
        else:
            raise AssertionError(f"{name}: expected WebrpcBadResponse")


def test_invalid_request_arguments_are_bad_requests():
    class Impl:
        def check(self, kind, at, code, tags, ok, f):
            return True

    valid = {"kind": "ADMIN", "at": "2020-01-01T00:00:00Z", "code": 1, "tags": ["a"], "ok": True, "f": 1.5}
    assert _call(Impl(), "Check", valid) == (200, {"ok": True})
    invalid = {
        "unknown enum value": {"kind": "NOPE"},
        "missing timestamp": {"at": None},
        "malformed timestamp": {"at": "yesterday"},
        "string for int": {"code": "abc"},
        "float for int": {"code": 1.5},
        "bool for int": {"code": True},
        "number for string list": {"tags": [1]},
        "string for list": {"tags": "abc"},
        "string for bool": {"ok": "yes"},
        "string for float": {"f": "1.5"},
    }
    for name, change in invalid.items():
        status, body = _call(Impl(), "Check", {**valid, **change})
        assert status == 400 and body["error"] == "WebrpcBadRequest" and body["cause"], (name, status, body)


def test_handler_exceptions_are_endpoint_errors():
    class Impl:
        def respond(self):
            raise RuntimeError("boom")

        def error(self):
            raise api.WebrpcBadRoute(cause="kept")

    status, body = _call(Impl(), "Respond", {})
    assert (status, body["error"], body["code"], body["cause"]) == (400, "WebrpcEndpoint", 0, "boom"), body
    status, body = _call(Impl(), "Error", {})
    assert (status, body["error"], body["cause"]) == (404, "WebrpcBadRoute", "kept"), body


if __name__ == "__main__":
    for name, fn in sorted(globals().items()):
        if name.startswith("test_") and callable(fn):
            fn()
    print("WIRE OK")
