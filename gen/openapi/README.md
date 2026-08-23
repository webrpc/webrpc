# webrpc OpenAPI 3.0.0 (Swagger) generator  <!-- omit in toc -->

This repo contains the templates used by the [webrpc-gen](https://github.com/webrpc/webrpc) cli to code-generate OpenAPI documentation from webrpc schema.

- [Generate OpenAPI 3.0.0 YAML file](#generate-openapi-300-yaml-file)
  - [Set custom template variables](#set-custom-template-variables)
- [Open in Swagger UI](#open-in-swagger-ui)
- [Build static HTML documentation with Redocly](#build-static-html-documentation-with-redocly)
- [Generate client/docs via OpenAPI generator](#generate-clientdocs-via-openapi-generator)
  - [Generate C# client for .NET 6.0](#generate-c-client-for-net-60)
  - [List of OpenAPI generators (as of Jan 5 2023)](#list-of-openapi-generators-as-of-jan-5-2023)
- [Authors](#authors)
- [License](#license)

# Generate OpenAPI 3.0.0 YAML file
```bash
webrpc-gen -schema=./proto.ridl -target=openapi -out petstore.gen.yaml
```

or
```bash
webrpc-gen -schema=./proto.ridl -target=github.com/webrpc/gen-openapi@v0.7.0 -out petstore.gen.yaml
```

## Set custom template variables
Change any of the following default values by passing `-option="Value"` CLI flag to webrpc-gen.

| webrpc-gen -option     | Default value            | Example value                                                                                                                        |
|------------------------|--------------------------|--------------------------------------------------------------------------------------------------------------------------------------|
| `-title`               | `{Services[0].Name} API` | `"Example API"`                                                                                                                      |
| `-apiVersion`          | `""`                     | `v22.10.25`                                                                                                                          |
| `-serverUrl`           | `""`                     | `https://api.example.com`                                                                                                            |
| `-serverDescription`   | `""`                     | `"Staging API"`                                                                                                                      |
| `-servers`             | `""`                     | `http://localhost:8080;description,http://localhost:8081;description`                                                                |
| `-securityAnnotation`  | `""`                     | `@auth`                                                                                                                              |
| `-securitySchemes`     | `""`                     | `{"ApiKeyAuth":{"type":"apiKey", "in":"header", "description":"Access key for authenticating requests", "name":"X-Access-Key"}}`     | 


Example:
- server url and server description will become part of the servers format in the end to keep it backward compatible
- means that the result will be `server="http://localhost:8080;description,http://localhost:8081;description,https://api.example.com;Production"`
```
webrpc-gen \ 
  -schema=./petstore.ridl \ 
  -target=github.com/webrpc/gen-openapi@v0.7.0 \ 
  -out petstore.gen.yaml \ 
  -title="Example webrpc API" \ 
  -apiVersion="v22.11.8" \ 
  -serverUrl=https://api.example.com \ 
  -serverDescription="Production"
  -servers="http://localhost:8080;description,http://localhost:8081;description"
```

- `securityAnnotation` must match the custom annotation you are using in your project to describe ACL for specific service methods
- `securitySchemes` give you possibility to pass an openapi security schemes which are later matched by your custom `acl annotation`. In this example it's an `@auth` annotation. 
Example with security annotation:
```
// proto.ridl
service ExampleService
  @auth:"ApiKeyAuth,ServiceAuth"
  - GetUserV2(header: map<string,string>, userID: uint64) => (profilePicture: string)

// Makefile
SECURITY_SCHEMES="{ \
  'ApiKeyAuth': { \
    'type': 'apiKey', \
    'in': 'header', \
    'description': 'Project access key for authenticating requests', \
    'name': 'X-Access-Key' \
  }, \
}"; \
webrpc-gen \
  -schema=./proto.ridl \
  -target=./ \
  -out=./openapi.gen.yaml \
  -title="Example webrpc API" \
  -apiVersion="v22.11.8"  \
  -serverUrl=https://api.example.com \
  -serverDescription="Production" \
  -securityAnnotation="@auth" \
  -securitySchemes="$$SECURITY_SCHEMES"
```

# Open in Swagger UI

Open generated documentation in Swagger editor:
```
docker run -p 8088:8080 -v $(pwd):/app -e SWAGGER_JSON=/app/petstore.gen.yaml swaggerapi/swagger-ui
```

Or in Swagger editor:
```
docker run -p 8088:8080 -v $(pwd):/app -e SWAGGER_FILE=/app/petstore.gen.yaml swaggerapi/swagger-editor
```

And go to http://localhost:8088

# Build static HTML documentation with Redocly
You can use [Redocly](https://redocly.com/docs/redoc/deployment/cli/) tool to build static HTML documentation. Customize the look and feel through [theming options](https://redocly.com/docs/api-reference-docs/configuration/theming/).

Example:
```
docker run --rm -v $(pwd):/app -w /app ghcr.io/redocly/redoc/cli build petstore.gen.yaml
```

# Generate client/docs via OpenAPI generator
You can use [OpenAPI generator](https://openapi-generator.tech/) to generate API clients in different languages or documentation from the generated OpenAPI yaml file.

## Generate C# client for .NET 6.0
```bash
docker run --rm -v "${PWD}:/app" openapitools/openapi-generator-cli:v6.2.1 generate -i /app/openapi.gen.yaml -g csharp-netcore -p targetFramework=net6.0 -o /app/csharp-net-7.0
```

## List of OpenAPI generators (as of Jan 5 2023)
- [ada](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/ada.md)
- [ada-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/ada-server.md)
- [android](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/android.md)
- [apache2](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/apache2.md)
- [apex](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/apex.md)
- [asciidoc](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/asciidoc.md)
- [aspnetcore](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/aspnetcore.md)
- [avro-schema](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/avro-schema.md)
- [bash](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/bash.md)
- [crystal](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/crystal.md)
- [c](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/c.md)
- [clojure](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/clojure.md)
- [cwiki](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/cwiki.md)
- [cpp-qt-client](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/cpp-qt-client.md)
- [cpp-qt-qhttpengine-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/cpp-qt-qhttpengine-server.md)
- [cpp-pistache-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/cpp-pistache-server.md)
- [cpp-restbed-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/cpp-restbed-server.md)
- [cpp-restbed-server-deprecated](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/cpp-restbed-server-deprecated.md)
- [cpp-restsdk](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/cpp-restsdk.md)
- [cpp-tiny](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/cpp-tiny.md)
- [cpp-tizen](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/cpp-tizen.md)
- [cpp-ue4](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/cpp-ue4.md)
- [csharp](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/csharp.md)
- [csharp-netcore](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/csharp-netcore.md)
- [csharp-dotnet2](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/csharp-dotnet2.md)
- [csharp-netcore-functions](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/csharp-netcore-functions.md)
- [dart](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/dart.md)
- [dart-dio](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/dart-dio.md)
- [eiffel](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/eiffel.md)
- [elixir](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/elixir.md)
- [elm](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/elm.md)
- [erlang-client](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/erlang-client.md)
- [erlang-proper](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/erlang-proper.md)
- [erlang-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/erlang-server.md)
- [fsharp-functions](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/fsharp-functions.md)
- [fsharp-giraffe-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/fsharp-giraffe-server.md)
- [go](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/go.md)
- [go-echo-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/go-echo-server.md)
- [go-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/go-server.md)
- [go-gin-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/go-gin-server.md)
- [graphql-schema](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/graphql-schema.md)
- [graphql-nodejs-express-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/graphql-nodejs-express-server.md)
- [groovy](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/groovy.md)
- [kotlin](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/kotlin.md)
- [kotlin-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/kotlin-server.md)
- [kotlin-spring](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/kotlin-spring.md)
- [kotlin-vertx](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/kotlin-vertx.md)
- [ktorm-schema](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/ktorm-schema.md)
- [haskell-http-client](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/haskell-http-client.md)
- [haskell](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/haskell.md)
- [haskell-yesod](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/haskell-yesod.md)
- [java](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/java.md)
- [jaxrs-cxf-client](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/jaxrs-cxf-client.md)
- [java-helidon-client](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/java-helidon-client.md)
- [java-helidon-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/java-helidon-server.md)
- [java-inflector](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/java-inflector.md)
- [java-micronaut-client](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/java-micronaut-client.md)
- [java-micronaut-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/java-micronaut-server.md)
- [java-msf4j](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/java-msf4j.md)
- [java-pkmst](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/java-pkmst.md)
- [java-play-framework](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/java-play-framework.md)
- [java-undertow-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/java-undertow-server.md)
- [java-vertx](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/java-vertx.md)
- [java-vertx-web](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/java-vertx-web.md)
- [java-camel](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/java-camel.md)
- [jaxrs-cxf](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/jaxrs-cxf.md)
- [jaxrs-cxf-extended](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/jaxrs-cxf-extended.md)
- [jaxrs-cxf-cdi](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/jaxrs-cxf-cdi.md)
- [jaxrs-jersey](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/jaxrs-jersey.md)
- [jaxrs-resteasy](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/jaxrs-resteasy.md)
- [jaxrs-resteasy-eap](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/jaxrs-resteasy-eap.md)
- [jaxrs-spec](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/jaxrs-spec.md)
- [javascript](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/javascript.md)
- [javascript-apollo-deprecated](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/javascript-apollo-deprecated.md)
- [javascript-flowtyped](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/javascript-flowtyped.md)
- [javascript-closure-angular](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/javascript-closure-angular.md)
- [jmeter](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/jmeter.md)
- [k6](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/k6.md)
- [lua](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/lua.md)
- [markdown](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/markdown.md)
- [mysql-schema](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/mysql-schema.md)
- [nim](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/nim.md)
- [nodejs-express-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/nodejs-express-server.md)
- [objc](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/objc.md)
- [ocaml](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/ocaml.md)
- [openapi](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/openapi.md)
- [openapi-yaml](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/openapi-yaml.md)
- [plantuml](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/plantuml.md)
- [perl](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/perl.md)
- [php](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/php.md)
- [php-laravel](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/php-laravel.md)
- [php-lumen](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/php-lumen.md)
- [php-slim-deprecated](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/php-slim-deprecated.md)
- [php-slim4](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/php-slim4.md)
- [php-symfony](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/php-symfony.md)
- [php-mezzio-ph](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/php-mezzio-ph.md)
- [php-dt](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/php-dt.md)
- [powershell](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/powershell.md)
- [protobuf-schema](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/protobuf-schema.md)
- [python-legacy](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/python-legacy.md)
- [python](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/python.md)
- [python-fastapi](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/python-fastapi.md)
- [python-prior](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/python-prior.md)
- [python-flask](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/python-flask.md)
- [python-aiohttp](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/python-aiohttp.md)
- [python-blueplanet](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/python-blueplanet.md)
- [r](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/r.md)
- [ruby](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/ruby.md)
- [ruby-on-rails](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/ruby-on-rails.md)
- [ruby-sinatra](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/ruby-sinatra.md)
- [rust](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/rust.md)
- [rust-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/rust-server.md)
- [scalatra](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/scalatra.md)
- [scala-akka](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/scala-akka.md)
- [scala-akka-http-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/scala-akka-http-server.md)
- [scala-finch](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/scala-finch.md)
- [scala-httpclient-deprecated](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/scala-httpclient-deprecated.md)
- [scala-gatling](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/scala-gatling.md)
- [scala-lagom-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/scala-lagom-server.md)
- [scala-play-server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/scala-play-server.md)
- [scala-sttp](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/scala-sttp.md)
- [scalaz](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/scalaz.md)
- [spring](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/spring.md)
- [dynamic-html](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/dynamic-html.md)
- [html](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/html.md)
- [html2](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/html2.md)
- [swift5](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/swift5.md)
- [typescript](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/typescript.md)
- [typescript-angular](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/typescript-angular.md)
- [typescript-aurelia](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/typescript-aurelia.md)
- [typescript-axios](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/typescript-axios.md)
- [typescript-fetch](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/typescript-fetch.md)
- [typescript-inversify](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/typescript-inversify.md)
- [typescript-jquery](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/typescript-jquery.md)
- [typescript-nestjs](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/typescript-nestjs.md)
- [typescript-node](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/typescript-node.md)
- [typescript-redux-query](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/typescript-redux-query.md)
- [typescript-rxjs](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/typescript-rxjs.md)
- [wsdl-schema](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/wsdl-schema.md)

# Authors
- [golang.cz](https://www.golang.cz)
- [Vojtech Vitek](https://github.com/VojtechVitek)

# License
[MIT LICENSE](./LICENSE)
