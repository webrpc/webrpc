# Webrpc changelog

## v0.9.0

Towards reaching webrpc@v1.0.0, we have decided to make some minor changes to RIDL file format. Here's what you need to do to migrate to webrpc@v0.9.0+.

## RIDL v0.9.0 changes

Keyword "message" was replaced by "struct".

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

## RIDL v0.9.0 migration guide

You can run this command to migrate your RIDL files to webrpc@v0.9.0+:

```bash
#!/bin/bash

find . -name '*.ridl' -exec sed -i -e 's/^message /struct /g' {} \;
```

## JSON schema v0.9.0 migration guide

You can run this command to migrate your .json schema files to webrpc@v0.9.0+:

```bash
#!/bin/bash

find . -name '*.json' -exec sed -i -e 's/^"messages":/\"structs\":/g' {} \;
```
