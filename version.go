package webrpc

// VERSION defines version of webrpc tools, such as webrpc-gen.
//
// Please update major/minor version when making any incompatible changes
// to the webrpc schema or webrpc-gen Template Functions API.
// The patch value is automatically updated with the latest git tag in CI.
//
// Version is available as {{.WebrpcGenVersion}} in generator templates.
var VERSION = "v0.15.0-dev"
