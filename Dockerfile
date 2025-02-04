# -----------------------------------------------------------------
# Builder
# -----------------------------------------------------------------
FROM golang:1.23-alpine3.21 as builder
ARG VERSION

RUN apk add --update git

ADD ./ /src

WORKDIR /src
RUN go build -ldflags="-s -w -X github.com/webrpc/webrpc.VERSION=${VERSION}" -o /usr/bin/webrpc-gen ./cmd/webrpc-gen

# -----------------------------------------------------------------
# Runner
# -----------------------------------------------------------------
FROM alpine:3.21

ENV TZ=UTC

RUN apk add --no-cache --update ca-certificates

COPY --from=builder /usr/bin/webrpc-gen /usr/bin/

ENTRYPOINT ["/usr/bin/webrpc-gen"]
