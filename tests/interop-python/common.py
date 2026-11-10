"""Loads the generated TestApi module and holds the fixtures shared with the Go
reference implementation (tests/server/server.go and tests/client/client.go)."""
import importlib.util
import pathlib
import sys

_path = pathlib.Path(__file__).parent / "test.gen.py"
_spec = importlib.util.spec_from_file_location("test_gen", _path)
api = importlib.util.module_from_spec(_spec)
sys.modules[_spec.name] = api
_spec.loader.exec_module(api)

ONE = api.Simple(id=1, name="one")
TWO = api.Simple(id=2, name="two")
THREE = api.Simple(id=3, name="three")

_JOHN = api.User(id=1, USERNAME="John-Doe", role="admin")

COMPLEX = api.Complex(
    meta={"1": "23", "2": 24},
    metaNestedExample={"1": {"2": 1}},
    namesList=["John", "Alice", "Jakob"],
    numsList=[1, 2, 3, 4534643543],
    doubleArray=[["testing"], ["api"]],
    listOfMaps=[{"john": 1, "alice": 2, "Jakob": 251}],
    listOfUsers=[_JOHN],
    mapOfUsers={"admin": _JOHN},
    user=_JOHN,
    status=api.Status.AVAILABLE,
)

ENUM_LIST = [api.Status.AVAILABLE, api.Status.NOT_AVAILABLE]
ENUM_MAP = {api.Access.READ: 1, api.Access.WRITE: 2}

# code, error class, cause
SCHEMA_ERRORS = [
    (0, api.WebrpcEndpoint, "failed to read file: unexpected EOF"),
    (1, api.Unauthorized, "failed to verify JWT token"),
    (2, api.ExpiredToken, None),
    (3, api.InvalidToken, None),
    (4, api.Deactivated, None),
    (5, api.ConfirmAccount, None),
    (6, api.AccessDenied, None),
    (7, api.MissingArgument, None),
    (8, api.UnexpectedValue, None),
    (100, api.RateLimited, "1000 req/min exceeded"),
    (101, api.DatabaseDown, None),
    (102, api.ElasticDown, None),
    (103, api.NotImplemented, None),
    (200, api.UserNotFound, None),
    (201, api.UserBusy, None),
    (202, api.InvalidUsername, None),
    (300, api.FileTooBig, None),
    (301, api.FileInfected, None),
    (302, api.FileType, ".wav is not supported"),
]
