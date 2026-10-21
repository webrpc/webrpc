#!/bin/sh
# Runs the C# client interoperability tests against the webrpc test servers
# (webrpc-test), the CSharp test server (./server) and the golang-basics
# example server. Requires go and the .NET SDK (dotnet, or set $DOTNET).
set -eu

cd "$(dirname "$0")"
ROOT=$(cd ../.. && pwd)
DOTNET=${DOTNET:-dotnet}
TMP=$(mktemp -d)
PIDS=""

cleanup() {
    for pid in $PIDS; do kill "$pid" 2>/dev/null || true; done
    rm -rf "$TMP"
}
trap cleanup EXIT INT TERM

TESTAPI_PORT=9991
TESTAPI_REST_PORT=9992
CSHARP_PORT=9993
EXAMPLE_PORT=9994

# Build the servers first, so the started processes are the servers themselves.
go build -o "$TMP/webrpc-test" "$ROOT/cmd/webrpc-test"
go build -o "$TMP/csharp-server" ./server
(cd "$ROOT/_examples/golang-basics" && go build -o "$TMP/example-server" .)

"$TMP/webrpc-test" -server -port=$TESTAPI_PORT -timeout=10m &
PIDS="$PIDS $!"
"$TMP/webrpc-test" -server -rest -port=$TESTAPI_REST_PORT -timeout=10m &
PIDS="$PIDS $!"
"$TMP/csharp-server" -port=$CSHARP_PORT &
PIDS="$PIDS $!"
PORT=$EXAMPLE_PORT "$TMP/example-server" >/dev/null 2>&1 &
PIDS="$PIDS $!"

for port in $TESTAPI_PORT $TESTAPI_REST_PORT $CSHARP_PORT $EXAMPLE_PORT; do
    tries=0
    until nc -z 127.0.0.1 $port 2>/dev/null; do
        tries=$((tries + 1))
        if [ $tries -gt 300 ]; then
            echo "test server on port $port did not start" >&2
            exit 1
        fi
        sleep 0.1
    done
done

# The generated code must compile for every target framework, including
# netstandard2.0 with C# 8 (eg. Unity), without warnings.
"$DOTNET" build --nologo -v quiet lib/stj
"$DOTNET" build --nologo -v quiet lib/newtonsoft

status=0
for json in stj newtonsoft; do
    "$DOTNET" run --project test/$json -- \
        "http://127.0.0.1:$TESTAPI_PORT" \
        "http://127.0.0.1:$TESTAPI_REST_PORT" \
        "http://127.0.0.1:$CSHARP_PORT" \
        "http://127.0.0.1:$EXAMPLE_PORT" || status=1
done
exit $status
