// Package constraint implements parsing and evaluation of build constraint lines.
// See https://golang.org/cmd/go/#hdr-Build_constraints for documentation about build constraints themselves.
//
// This package parses both the original “// +build” syntax and the “//go:build” syntax that was added in Go 1.17.
// See https://golang.org/design/draft-gobuild for details about the “//go:build” syntax.
package constraint
