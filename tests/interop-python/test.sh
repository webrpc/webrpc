#!/bin/sh
# Runs the Python generator tests: the generated example's roundtrip suite and
# checks of the generator's input validation. Requires go and Python 3.9+ with
# httpx (see `make venv`), or set $PYTHON.
set -eu

cd "$(dirname "$0")"
ROOT=$(cd ../.. && pwd)
PYTHON=${PYTHON:-python3}
TMP=$(mktemp -d "${TMPDIR:-/tmp}/webrpc-python.XXXXXX")

cleanup() {
    rm -rf "$TMP"
}
trap cleanup EXIT INT TERM

gen() {
    "$ROOT/bin/webrpc-gen" -target=python "$@"
}

echo "example roundtrip"
"$PYTHON" test_example.py

echo "streaming schemas are rejected"
if gen -schema=streaming.ridl -client -server -out="$TMP/streaming.py" >/dev/null 2>&1; then
    echo "FAIL: streaming schema was not rejected" >&2
    exit 1
fi

echo "keyword identifiers generate valid Python"
gen -schema=keywords.ridl -client -server -out="$TMP/keywords.py" >/dev/null
"$PYTHON" -m py_compile "$TMP/keywords.py"
