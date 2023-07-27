package webrpc

// VERSION defines version of webrpc tools, such as webrpc-gen.
//
// Please update major/minor version by hand when changing webrpc-gen
// Template Functions API or when making any backward-incompatible changes.
// The patch value is automatically updated with the latest git tag in CI.
//
// Available as {{.WebrpcGenVersion}} in generator templates. Templates can
// check against {{.WebrpcGenVersion}} to require minimal major/minor version
// to guarantees certain Template functions API.
var VERSION = "v0.13.0-dev"
