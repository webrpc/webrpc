"""Wire-format tests of the generated wire.gen.py against the Go and TypeScript
generators' contract."""
import importlib.util
import inspect
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


def _call(impl, method, body, content_length=True):
    """Posts a JSON body to the generated server and returns (status, JSON)."""
    raw = json.dumps(body).encode()
    status = []
    environ = {
        "REQUEST_METHOD": "POST",
        "PATH_INFO": f"/rpc/Wire/{method}",
        "CONTENT_TYPE": "application/json",
        "wsgi.input": io.BytesIO(raw),
    }
    if content_length:
        environ["CONTENT_LENGTH"] = str(len(raw))
    out = b"".join(api.WireServer(impl)(environ, lambda s, h: status.append(s)))
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

    assert _call(Impl(), "Flatten", {"a": "in"}) == ("200 OK", {"b": "in!"})


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

    assert _call(Impl(), "Go", {"a": "in"}) == ("200 OK", {"b": "in!"})


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

    assert _call(Impl(), "Maybe", {"x": None, "y": "y"}) == ("200 OK", {"w": "w"})


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

    assert _call(Impl(), "From", {"class": "c"}) == ("200 OK", {"ok": True})
    assert _call(Impl(), "Import", {}) == ("200 OK", {})
    assert _call(Impl(), "Get", {"self": "s"}) == ("200 OK", {"ok": True})


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
        assert _call(Impl(), method, {}) == ("200 OK", want), method


def test_unknown_error_code_keeps_what_the_server_sent():
    body = {"error": "Custom", "code": 777, "msg": "m", "cause": "c", "status": 422}
    try:
        _client(lambda request: httpx.Response(422, json=body)).get("s")
    except api.WebRPCError as err:
        got = (type(err), err.name, err.code, err.message, err.cause, err.status)
    else:
        raise AssertionError("expected an error")
    assert got == (api.WebRPCError, "Custom", 777, "m", "c", 422)


def test_error_body_without_a_code_keeps_its_message():
    body = {"error": "Custom", "msg": "m", "cause": "c", "status": 500}
    try:
        _client(lambda request: httpx.Response(500, json=body)).get("s")
    except api.WebRPCError as err:
        got = (type(err), err.name, err.code, err.message, err.cause, err.status)
    else:
        raise AssertionError("expected an error")
    assert got == (api.WebRPCError, "Custom", 0, "m", "c", 500)


def test_unusable_responses_are_bad_responses():
    cases = {
        "redirect without body": httpx.Response(302),
        "error status with a JSON array": httpx.Response(500, json=["x"]),
        "error status with HTML": httpx.Response(502, content=b"<html>bad gateway</html>"),
        "error status with a boolean code": httpx.Response(500, json={"code": True, "msg": "x"}),
        "error status with a string code": httpx.Response(500, json={"code": "1", "msg": "x"}),
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
    assert _call(Impl(), "Check", valid) == ("200 OK", {"ok": True})
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
        assert status == "400 Bad Request" and body["error"] == "WebrpcBadRequest" and body["cause"], (name, status, body)


def test_handler_exceptions_are_endpoint_errors():
    class Impl:
        def respond(self):
            raise RuntimeError("boom")

        def error(self):
            raise api.WebrpcBadRoute(cause="kept")

    status, body = _call(Impl(), "Respond", {})
    assert (status, body["error"], body["code"], body["cause"]) == ("400 Bad Request", "WebrpcEndpoint", 0, "boom"), body
    status, body = _call(Impl(), "Error", {})
    assert (status, body["error"], body["cause"]) == ("404 Not Found", "WebrpcBadRoute", "kept"), body


def test_required_struct_fields_are_validated():
    for name, wire in {"missing field": {}, "null field": {"need": None}}.items():
        try:
            api.Opt.from_dict(wire)
        except (TypeError, ValueError):
            continue
        raise AssertionError(f"{name}: expected the required value to be rejected")


def test_null_struct_alias_is_accepted():
    assert api.AliasHolder.from_dict({"user": None}) == api.AliasHolder(user=None)

    seen = []

    def respond(request):
        seen.append(json.loads(request.content))
        return httpx.Response(200, json={"user": None})

    assert _client(respond).alias_struct(None) is None
    assert seen == [{"user": None}]

    class Impl:
        def alias_struct(self, user):
            assert user is None
            return None

    assert _call(Impl(), "AliasStruct", {"user": None}) == ("200 OK", {"user": None})


def test_required_primitive_alias_is_validated():
    class Impl:
        def alias_text(self, text):
            return text

    assert _call(Impl(), "AliasText", {"text": "x"}) == ("200 OK", {"text": "x"})
    for name, body in {"missing input": {}, "null input": {"text": None}}.items():
        status, out = _call(Impl(), "AliasText", body)
        assert status == "400 Bad Request" and out["error"] == "WebrpcBadRequest", (name, status, out)

    for name, body in {"missing output": {}, "null output": {"text": None}}.items():
        try:
            _client(lambda request: httpx.Response(200, json=body)).alias_text("x")
        except api.WebrpcBadResponse:
            continue
        raise AssertionError(f"{name}: expected the required output to be rejected")


def test_required_method_outputs_are_validated():
    for name, body in {"missing output": {}, "null output": {"ok": None}}.items():
        try:
            _client(lambda request: httpx.Response(200, json=body)).get("x")
        except api.WebrpcBadResponse:
            continue
        raise AssertionError(f"{name}: expected the required output to be rejected")


def test_required_method_inputs_are_validated():
    class Impl:
        def from_(self, class_):
            return class_ == "c"

    assert _call(Impl(), "From", {"class": "c"}) == ("200 OK", {"ok": True})
    for name, body in {"missing input": {}, "null input": {"class": None}}.items():
        status, out = _call(Impl(), "From", body)
        assert status == "400 Bad Request" and out["error"] == "WebrpcBadRequest" and out["cause"], (name, status, out)


def test_bigint_inputs():
    assert api.Big.from_dict({"n": "123"}).n == 123
    assert api.Big.from_dict({"n": 123}).n == 123
    assert api.Big(n=5).to_dict() == {"n": "5"}
    for name, value in {"bool": True, "float": 1.5, "list": [1], "bad string": "x"}.items():
        try:
            api.Big.from_dict({"n": value})
        except (TypeError, ValueError):
            continue
        raise AssertionError(f"{name}: expected _parse_bigint to reject {value!r}")


def test_optional_method_input_annotations():
    def annotations(cls, name):
        return inspect.signature(getattr(cls, name)).parameters

    client = annotations(api.WireClient, "maybe")
    assert client["x"].annotation == "Optional[str]", client["x"].annotation
    assert client["y"].annotation == "str"

    protocol = annotations(api.Wire, "maybe")
    assert protocol["x"].annotation == "Optional[str]", protocol["x"].annotation
    assert protocol["y"].annotation == "str"


def test_status_lines_carry_reason_phrases():
    class Impl:
        def import_(self):
            pass

        def from_(self, class_):
            return class_ == "c"

    assert _call(Impl(), "Import", {})[0] == "200 OK"
    assert _call(Impl(), "From", {})[0] == "400 Bad Request"
    assert api._status_line(200) == "200 OK"
    assert api._status_line(599) == "599 Unknown"


def test_body_is_read_without_content_length():
    class Impl:
        def from_(self, class_):
            return class_ == "c"

    assert _call(Impl(), "From", {"class": "c"}, content_length=False) == ("200 OK", {"ok": True})


def test_empty_success_body_is_accepted_only_without_outputs():
    empty = lambda request: httpx.Response(200)
    assert _client(empty).import_() is None
    for name, call in {
        "get": lambda c: c.get("x"),
        "flatten": lambda c: c.flatten(api.FlattenRequest(a="a")),
    }.items():
        try:
            call(_client(empty))
        except api.WebrpcBadResponse:
            continue
        raise AssertionError(f"{name}: expected WebrpcBadResponse")


if __name__ == "__main__":
    for name, fn in sorted(globals().items()):
        if name.startswith("test_") and callable(fn):
            fn()
    print("WIRE OK")
