# goconstraint

 Package goconstraint is a lightweight way for you to declare that your code
 requires a minimum Go runtime version, using a blank import.

## License

This package is released under two software licenses (MIT | Public Domain):

* The `scripts/` directory is released under the `MIT License`. This license can
  be found in the `LICENSES` directory as `mit.txt`
  ([here](https://github.com/theckman/goconstraint/blob/master/LICENSES/mit.txt)).
* All other files are released to the Public Domain using The Unlicense . This
  license can be found in the `LICENSES` directory as `unlicense.txt`
  ([here](https://github.com/theckman/goconstraint/blob/master/LICENSES/unlicense.txt)).

## Rationale

While functionality of a file can be guarded with build tags, it has the side
effect of all code in that file no longer being compiled. While this initially
seems like a desirable experience, this means that none of the items in that
file get compiled.

So when the building of the project fails, it's not because the Go version is
too old, but because a referenced thing is missing (variable, constant, or
function). This can be confusing to developers, as it looks like source code is
missing and not that the Go runtime version is too old.

Also, depending on what Go functionality you depend on, it may not automatic
result in a build failure. For example, `time.Now()` works with
`time.Time.Sub()` across leap seconds as of Go 1.9+. This was not the case in
older versions, and would cause unexpected behaviors. This required no API
changes, so the bug will be silently compiled in to your program if using an
older Go toolchain.

The purpose of this project is to be able to easily enforce Go runtime version
constraints, while providing a useful error to developers without requiring that
developers redeclare functions, variables, or constants.

You can read the blog post that was written about the problem that inspired this
package here:

* https://medium.com/@theckman/version-constraints-and-go-c9309be15773

You can also find the experimental code from that blog post here:

* https://github.com/theckman/constraint-test

## Usage

 Say you wanted to require that your code only build with Go 1.9+, you would
 include the following at the top of the file that has the runtime dependency
 on the specific version:

```Go
import _ "github.com/theckman/goconstraint/go1.9/gte"
```

 All go versions, that can be restricted by build tag, are supported (Go 1.1+).

 It's recommended that you include a comment with the import, to indicate
 which functionality/feature is required from that Go version (while also
 including a Golang issue link if possible).

 If you try to build a project against a runtime that's too old, you'll see a
 build failure similar to this with the version number changing based on the
 requirement:

```Go
goconstraint/go1.9/gte/constraint.go:10: undefined: __SOFTWARE_REQUIRES_GO_VERSION_1_9__
```
