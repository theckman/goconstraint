// The contents of this file has been released in to the Public Domain.

// Package goconstraint is a lightweight way for you to declare that your code
// requires a minimum Go runtime version, using a blank import.
//
// Say you wanted to require that your code only build with Go 1.9+, you would
// include the following at the top of the file that has the runtime dependency
// on the specific version:
//
//		import _ "github.com/theckman/goconstraint/go1.9/gte"
//
// All go versions, that can be restricted by build tag, have been supported (Go
// 1.1+). It's also recommended that you include a comment with the import, to
// indicate which functionality/feature is required from that Go version (while
// also including a Golang issue link if possible).
//
// If you try to build a project against a runtime that's too old, you'll see a
// build failure similar to this with the version number changing based on the
// requirement:
//
//		goconstraint/go1.9/gte/constraint.go:10: undefined: __SOFTWARE_REQUIRES_GO_VERSION_1_9__
//
package goconstraint
