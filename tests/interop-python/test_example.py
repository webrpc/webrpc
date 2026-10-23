"""End-to-end test for the generated example.gen.py.

Covers type serde (incl. bigint wire-string, tz-normalized timestamps, maps,
enum-keyed maps, type aliases, Python-keyword field names), a full client <->
server roundtrip via httpx.WSGITransport, error mapping, and the server's
request-validation paths.
"""
import importlib.util
import io
import json
import pathlib
import sys
from datetime import datetime, timezone

import httpx

_path = pathlib.Path(__file__).parent / "example.gen.py"
_spec = importlib.util.spec_from_file_location("example_gen", _path)
api = importlib.util.module_from_spec(_spec)
sys.modules[_spec.name] = api  # required before exec for dataclass module resolution
_spec.loader.exec_module(api)


def _make_item():
    return api.Item(
        id="1",
        name="widget",            # ItemName alias -> str
        tier=api.ItemTier.PREMIUM,
        count=5,
        quantity=12,              # Quantity alias -> int
        balance=9007199254740993,  # > 2**53, must survive as a decimal string on the wire
        tags=["a", "b"],
        attributes={"k": "v"},
        class_="reserved-name",    # wire key "class"
        createdAt=datetime(2020, 1, 1),  # naive -> normalized to UTC on the wire
        lastUpdate=None,
        notes=["n1"],              # optional list present
        labels=None,               # optional map absent -> must stay null, not {}
    )


class Impl:
    def get_items(self):
        return [api.ItemSummary(id="1", name="widget")]

    def get_item(self, itemId):
        if itemId == "missing":
            raise api.NoSuchItem()
        return _make_item()

    def create_item(self, item):
        assert isinstance(item, api.CreateItemRequest)
        return None

    def put_one(self, itemId):
        return None

    def take_one(self, itemId):
        return None

    def delete_item(self, itemId):
        return None

    def count_by_tier(self):
        return {api.ItemTier.PREMIUM: 2, api.ItemTier.REGULAR: 1}


def test_wire_shape():
    d = _make_item().to_dict()
    assert d["quantity"] == 12                                # Quantity alias (uint32) -> int on the wire
    assert d["balance"] == "9007199254740993", d["balance"]   # bigint -> decimal string
    assert d["createdAt"] == "2020-01-01T00:00:00+00:00", d["createdAt"]  # tz-normalized
    assert d["class"] == "reserved-name"                      # keyword field uses wire key
    assert d["attributes"] == {"k": "v"}
    assert d["notes"] == ["n1"]
    assert d["labels"] is None                                # optional map None -> null, not {}


def test_type_serde_roundtrip():
    decoded = api.Item.from_dict(_make_item().to_dict())
    assert decoded.tier is api.ItemTier.PREMIUM
    assert decoded.quantity == 12 and isinstance(decoded.quantity, int)  # Quantity alias round-trips as int
    assert decoded.balance == 9007199254740993 and isinstance(decoded.balance, int)
    assert decoded.createdAt == datetime(2020, 1, 1, tzinfo=timezone.utc)
    assert decoded.class_ == "reserved-name"
    assert decoded.attributes == {"k": "v"}
    assert decoded.lastUpdate is None
    assert decoded.notes == ["n1"]
    assert decoded.labels is None


def _client():
    server = api.ExampleServiceServer(Impl())
    transport = httpx.WSGITransport(app=server)
    return api.ExampleServiceClient("http://testserver", client=httpx.Client(transport=transport)), server


def test_client_server_roundtrip():
    client, _ = _client()

    item = client.get_item("1")
    assert item.id == "1" and item.tier is api.ItemTier.PREMIUM
    assert item.balance == 9007199254740993
    assert item.createdAt == datetime(2020, 1, 1, tzinfo=timezone.utc)
    assert item.class_ == "reserved-name"

    assert client.get_items()[0].id == "1"
    assert client.create_item(api.CreateItemRequest(name="x", tier=api.ItemTier.REGULAR)) is None

    counts = client.count_by_tier()
    assert counts[api.ItemTier.PREMIUM] == 2 and counts[api.ItemTier.REGULAR] == 1


def test_error_roundtrip():
    client, _ = _client()
    try:
        client.get_item("missing")
    except api.NoSuchItem as err:
        assert err.code == 2 and err.status == 404
    else:
        raise AssertionError("expected NoSuchItem")


def _wsgi(server, method, path, body=b"", content_type="application/json"):
    captured = {}

    def start_response(status, headers):
        captured["status"] = status
        captured["headers"] = dict(headers)

    environ = {
        "REQUEST_METHOD": method,
        "PATH_INFO": path,
        "CONTENT_TYPE": content_type,
        "CONTENT_LENGTH": str(len(body)),
        "wsgi.input": io.BytesIO(body),
    }
    out = b"".join(server(environ, start_response))
    return captured, (json.loads(out) if out else None)


def test_server_request_validation():
    _, server = _client()
    path = "/v1/ExampleService/GetItems"

    cap, _ = _wsgi(server, "GET", path)
    assert cap["status"].startswith("405") and cap["headers"].get("Allow") == "POST"

    cap, body = _wsgi(server, "POST", "/v1/ExampleService/Nope")
    assert cap["status"].startswith("404") and body["error"] == "WebrpcBadRoute"

    cap, body = _wsgi(server, "POST", path, body=b"hi", content_type="text/plain")
    assert cap["status"].startswith("400") and body["error"] == "WebrpcBadRequest"

    cap, body = _wsgi(server, "POST", "/v1/ExampleService/GetItem", body=b"{not json")
    assert cap["status"].startswith("400") and body["error"] == "WebrpcBadRequest"

    # missing Content-Type is rejected (matches gen-golang)
    cap, body = _wsgi(server, "POST", path, body=b"{}", content_type="")
    assert cap["status"].startswith("400") and body["error"] == "WebrpcBadRequest"

    # Content-Type comparison is case-insensitive
    cap, _ = _wsgi(server, "POST", path, body=b"{}", content_type="Application/JSON")
    assert cap["status"].startswith("200"), cap["status"]

    # every response carries the Webrpc header
    assert cap["headers"].get("Webrpc")


if __name__ == "__main__":
    for name, fn in sorted(globals().items()):
        if name.startswith("test_") and callable(fn):
            fn()
    print("ROUNDTRIP OK")
