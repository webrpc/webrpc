"""Wire-format tests of the generated wire.gen.py against the Go and TypeScript
generators' contract."""
import importlib.util
import pathlib
import sys

_path = pathlib.Path(__file__).parent / "wire.gen.py"
_spec = importlib.util.spec_from_file_location("wire_gen", _path)
api = importlib.util.module_from_spec(_spec)
sys.modules[_spec.name] = api
_spec.loader.exec_module(api)


def test_enum_wire_values():
    data = api.Data(k=api.Kind.Admin, c=api.Country.CA)
    assert data.to_dict() == {"k": "ADMIN", "c": "Canada"}
    assert api.Data.from_dict({"k": "ADMIN", "c": "Canada"}) == data
    assert api.Kind.from_dict("User") is api.Kind.User


if __name__ == "__main__":
    for name, fn in sorted(globals().items()):
        if name.startswith("test_") and callable(fn):
            fn()
    print("WIRE OK")
