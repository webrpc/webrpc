FILE=sector.ridl
NAME=$(basename $FILE .ridl)

cd ../../
make build
mv ./bin/webrpc-gen ./cmd/webrpc-gen/webrpc-gen
cd cmd/webrpc-gen
sudo chmod 777 webrpc-gen

echo "webrpc-gen -schema=${FILE} -target=go -pkg=${NAME} -server -out=./services/${NAME}/webrpc.gen.go"

./webrpc-gen -schema=${FILE} -target=go -pkg=${NAME} -server -out=./services/${NAME}/webrpc.gen.go
./webrpc-gen -schema=${FILE} -target=ts -client -out=./sdk/${NAME}.rpc.gen.ts
