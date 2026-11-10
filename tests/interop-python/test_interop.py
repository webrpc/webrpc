"""Runs the generated TestApi client against the Go reference server.

Usage: python test_interop.py <server url>
"""
import sys

from common import COMPLEX, ENUM_LIST, ENUM_MAP, ONE, SCHEMA_ERRORS, THREE, TWO, api


def check(name, want, got):
    assert want == got, f"{name}: want {want!r}, got {got!r}"


def main(url):
    client = api.TestApiClient(url)

    check("GetEmpty", None, client.get_empty())

    try:
        client.get_error()
    except api.WebRPCError:
        pass
    else:
        raise AssertionError("GetError: expected an error")

    check("GetOne", ONE, client.get_one())
    check("SendOne", None, client.send_one(ONE))

    check("GetMulti", (ONE, TWO, THREE), client.get_multi())
    check("SendMulti", None, client.send_multi(ONE, TWO, THREE))

    check("GetComplex", COMPLEX, client.get_complex())
    check("SendComplex", None, client.send_complex(COMPLEX))

    check("GetEnumList", ENUM_LIST, client.get_enum_list())
    check("GetEnumMap", ENUM_MAP, client.get_enum_map())

    for code, cls, cause in SCHEMA_ERRORS:
        try:
            client.get_schema_error(code)
        except api.WebRPCError as err:
            check(f"GetSchemaError({code}) class", cls, type(err))
            check(f"GetSchemaError({code}) code", code, err.code)
            check(f"GetSchemaError({code}) message", cls.message, err.message)
            check(f"GetSchemaError({code}) status", cls.status, err.status)
            check(f"GetSchemaError({code}) cause", cause, err.cause)
        else:
            raise AssertionError(f"GetSchemaError({code}): expected an error")

    print("INTEROP OK")


if __name__ == "__main__":
    main(sys.argv[1])
