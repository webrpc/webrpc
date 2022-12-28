# webrpc v0.9.0

Towards reaching webrpc@v1.0.0, we have decided to make some breaking changes to webrpc schema and RIDL file format.

## Breaking changes

### RIDL v0.9.0 changes

Keyword "message" was renamed to "struct".

```diff
 webrpc = v1
 
 name = your-app
 version = v0.1.0
 
- message User
+ struct User
   - id: uint64
   - username: string
   - createdAt?: timestamp
```

### JSON schema v0.9.0 changes

- Field "messages" was renamed to "types"
- Field "type" was renamed to "kind"

```diff
 {
   "webrpc": "v1",
   "name": "Test",
   "version": "v0.0.1",
   "imports": [],
-  "messages": [
+  "types": [
   {
     "name": "Status",
-    "type": "enum",
+    "kind": "enum",
     "fields": [
     {
       "name": "AVAILABLE",
-      "type": "uint32",
-      "optional": false,
       "value": "0",
-      "meta": null
     },
     {
       "name": "NOT_AVAILABLE",
-      "type": "uint32",
-      "optional": false,
       "value": "1",
-      "meta": null
     }
   },
   {
    "name": "Empty",
-   "type": "struct",
+   "kind": "struct",
-   "fields": null
   }]
 }
```

### Template changes

You might see the following error when running your webrpc generator templates against webrpc-gen@v0.9.0+:

```
template: main.go.tmpl:88:57: executing "main" at <.Messages>: can't evaluate field Messages in type struct { *schema.WebRPCSchema; SchemaHash string; WebrpcGenVersion string; WebrpcGenCommand string; WebrpcTarget string; Opts map[string]interface {} }
```

To fix this, rename `{{.Messages}}` variable to `{{.Types}}` in your `*.go.tmpl` template files.

## Migration guide

### RIDL v0.9.0 migration guide

Run this command to migrate your RIDL files to webrpc@v0.9.0+:

```bash
#!/bin/bash

find . -name '*.ridl' -exec sed -i -e 's/^message /struct /g' {} \;
```

### JSON schema v0.9.0 migration guide

Run this Node.js script to migrate your `*webrpc.json` schema files to webrpc@v0.9.0+:

`node migrate.js schema.webrpc.json`

Contents of `migrate.js` file:
```javascript
const fs = require("fs");

if (process.argv.length != 3) {
  throw Error(`Usage: node ${process.argv[1]} <webrpc-schema.json>`);
}

const filePath = process.argv[2];

console.log(filePath);

fs.readFile(filePath, "utf8", (e, data) => {
  if (e) {
    throw e;
  }

  let schema = JSON.parse(data);
  schema = {
    webrpc: schema.webrpc,
    name: schema.name,
    version: schema.version,
    types: schema.messages.map((orig) => {
      let type = {
        name: orig.name,
        kind: orig.type,
        fields: orig.fields,
      };

      if (type.kind == "enum") {
        type = {
          name: orig.name,
          kind: orig.type,
          type: orig.fields[0].type,
          fields: orig.fields.map((field) => {
            return { name: field.name, value: field.value };
          }),
        };
      }
      return type;
    }),
    services: schema.services,
  };

  schema.types = fs.writeFile(
    filePath,
    JSON.stringify(schema, null, "\t"),
    (err) => {
      if (err) {
        console.error(err);
        console.log(schema);
      }
    }
  );
});
```

### Generator templates v0.9.0 migration guide

Run this command to migrate your `.go.tmpl` templates to webrpc@v0.9.0+:

```bash
#!/bin/bash

find . -name '*.go.tmpl' -exec sed -i -e 's/\.Messages/.Types/g' {} \;
find . -name '*.go.tmpl' -exec sed -i -e 's/"Messages"/"Types"/g' {} \;
```

