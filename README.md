# Zed-Gen-Go

A program to generate types, enums, and builder functions for working with [Authzed](https://github.com/authzed/authzed-go) clients in Go.

Given a SpiceDB schema file (e.g. `schema/schema.zed`), parses the file into a list of Resource definitions, each with its set of Relations and Permissions.

Generates a new file `schema.zed.go` with Object types for each resource, enum values for each resources Relations and Permissions, and type safe Relation Builder functions for building request objects for the `authzed.v1.PermissionsService.WriteRelationships` RPC.

Notes:
- Does not yet support Caveats.
- Does not currently parse the types of the Relation definitions. Future versions may parse those, and set up relation builder functions enforcing the correct resource types being assigned for a given relation.

### Usage:

Command 

```sh
<excutable> gen <path-to-schema.zed> -o <output-dir>
```
Example
```sh
go build -o ./build/zedgen ./cmd/zed-gen-go

./build/zedgen gen schema/schema.zed -o generated
```



