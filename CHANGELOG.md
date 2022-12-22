# webrpc v0.9.0

Towards reaching webrpc@v1.0.0, we have decided to make some changes to webrpc schema and RIDL file format.

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

## JSON schema v0.9.0 changes

Field "messages" was renamed to "types".

```diff
{

}
```

## Migration guide

### RIDL v0.9.0 migration guide

Run this command to migrate your RIDL files to webrpc@v0.9.0+:

```bash
#!/bin/bash

find . -name '*.ridl' -exec sed -i -e 's/^message /struct /g' {} \;
```

### JSON schema v0.9.0 migration guide

Run this command to migrate your `.json` schema files to webrpc@v0.9.0+:

```bash
#!/bin/bash

find . -name '*.json' -exec sed -i -e 's/"messages":/\"types\":/g' {} \;
```

### Generator templates v0.9.0 migration guide

Run this command to migrate your `.go.tmpl` templates to webrpc@v0.9.0+:

```bash
#!/bin/bash

find . -name '*.go.tmpl' -exec sed -i -e 's/.Messages:/\".Types\":/g' {} \;
```

