// Package format implements standard formatting of Go source.
//
// Note that formatting of Go source code changes over time, so tools relying on
// consistent formatting should execute a specific version of the gofmt binary
// instead of using this package. That way, the formatting will be stable, and
// the tools won't need to be recompiled each time gofmt changes.
//
// For example, pre-submit checks that use this package directly would behave
// differently depending on what Go version each developer uses, causing the
// check to be inherently fragile.
package format
