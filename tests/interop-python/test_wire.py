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


if __name__ == "__main__":
    for name, fn in sorted(globals().items()):
        if name.startswith("test_") and callable(fn):
            fn()
    print("WIRE OK")
