#!/bin/sh
# Runs the Python generator tests: the generated example's roundtrip suite,
# checks of the generator's input validation, and the interoperability tests
# against the Go reference implementation (webrpc-test): the generated client
# against the Go server, and the Go client against the generated WSGI server.
# Requires go and Python 3.11+ with httpx (see `make venv`), or set $PYTHON.
set -eu

cd "$(dirname "$0")"
ROOT=$(cd ../.. && pwd)
PYTHON=${PYTHON:-python3}
TMP=$(mktemp -d "${TMPDIR:-/tmp}/webrpc-python.XXXXXX")
PIDS=""

cleanup() {
    for pid in $PIDS; do kill "$pid" 2>/dev/null || true; done
    rm -rf "$TMP"
}
trap cleanup EXIT INT TERM

gen() {
    "$ROOT/bin/webrpc-gen" -target=python "$@"
}

echo "example roundtrip"
"$PYTHON" test_example.py

expect_rejected() {
    if err=$(gen -schema="$1" -client -server -out="$TMP/rejected.py" 2>&1); then
        echo "FAIL: $1 was not rejected" >&2
        exit 1
    fi
    case "$err" in
    *"not supported"*) ;;
    *)
        echo "FAIL: $1 failed for another reason: $err" >&2
        exit 1
        ;;
    esac
}

echo "wire format"
"$PYTHON" test_wire.py

echo "streaming and file schemas are rejected"
for schema in streaming.ridl file.ridl file-field.ridl; do
    expect_rejected $schema
done

echo "keyword identifiers generate valid Python"
gen -schema=keywords.ridl -client -server -out="$TMP/keywords.py" >/dev/null
"$PYTHON" -m py_compile "$TMP/keywords.py"

GO_SERVER_PORT=9994
PY_SERVER_PORT=9995

wait_for_port() {
    tries=0
    until nc -z 127.0.0.1 "$1" 2>/dev/null; do
        tries=$((tries + 1))
        if [ $tries -gt 300 ]; then
            echo "test server on port $1 did not start" >&2
            exit 1
        fi
        sleep 0.1
    done
}

go build -o "$TMP/webrpc-test" "$ROOT/cmd/webrpc-test"

echo "Python client -> Go server"
"$TMP/webrpc-test" -server -port=$GO_SERVER_PORT -timeout=10m &
PIDS="$PIDS $!"
wait_for_port $GO_SERVER_PORT
"$PYTHON" test_interop.py "http://127.0.0.1:$GO_SERVER_PORT"

echo "Go client -> Python server"
"$PYTHON" serve.py $PY_SERVER_PORT &
PIDS="$PIDS $!"
wait_for_port $PY_SERVER_PORT
# The file type is not supported by the Python generator.
"$TMP/webrpc-test" -client -url="http://127.0.0.1:$PY_SERVER_PORT" -skip=files
echo "GO CLIENT OK"
