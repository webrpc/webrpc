"""Serves the generated TestApi over WSGI, for the Go reference client.

Usage: python serve.py <port>
"""
import sys
from wsgiref.simple_server import WSGIRequestHandler, make_server

from common import COMPLEX, ENUM_LIST, ENUM_MAP, ONE, SCHEMA_ERRORS, THREE, TWO, api


def _expect(name, want, got):
    if want != got:
        raise api.UnexpectedValue(cause=f"{name}: want {want!r}, got {got!r}")


class TestApiImpl:
    def get_empty(self):
        return None

    def get_error(self):
        raise RuntimeError("internal error")

    def get_one(self):
        return ONE

    def send_one(self, one):
        _expect("one", ONE, one)

    def get_multi(self):
        return ONE, TWO, THREE

    def send_multi(self, one, two, three):
        _expect("one", ONE, one)
        _expect("two", TWO, two)
        _expect("three", THREE, three)

    def get_complex(self):
        return COMPLEX

    def send_complex(self, complex):
        _expect("complex", COMPLEX, complex)

    def get_enum_list(self):
        return ENUM_LIST

    def get_enum_map(self):
        return ENUM_MAP

    def get_schema_error(self, code):
        for want, cls, cause in SCHEMA_ERRORS:
            if want == code:
                raise cls(cause=cause)


class QuietHandler(WSGIRequestHandler):
    def log_message(self, format, *args):
        pass


if __name__ == "__main__":
    server = make_server("127.0.0.1", int(sys.argv[1]), api.TestApiServer(TestApiImpl()), handler_class=QuietHandler)
    server.serve_forever()
