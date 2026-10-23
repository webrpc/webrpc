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


if __name__ == "__main__":
    for name, fn in sorted(globals().items()):
        if name.startswith("test_") and callable(fn):
            fn()
    print("WIRE OK")
