package gen

// webrpc-gen CLI version, available as {{.WebrpcGenVersion}} in generator templates.
//
// The value is automatically updated with the latest git tag when building official
// webrpc-gen builds in the CI pipeline.
//
// Update major/minor version by hand when changing Template functions API
// or when making any other backward-incompatible changes to webrpc-gen.
//
// Templates can check against {{.WerbpcGenVersion}} to require minimal major/minor
// version that guarantees certain Template functions API.
var VERSION = "v0.8.x-dev"
