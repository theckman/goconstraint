# goconstraint

 Package goconstraint is a lightweight way for you to declare that your code
 requires a minimum Go runtime version, using a blank import.

## License

This package is licensed under the `MIT License`. Please see
the [LICENSE](https://github.com/theckman/goconstraint/blob/master/LICENSE) file
for more details.

## Usage

 Say you wanted to require that your code only build with Go 1.9+, you would
 include the following at the top of the file that has the runtime dependency
 on the specific version:

```Go
import _ "github.com/theckman/goconstraint/go1.9/gte"
```

 All go versions, that can be restricted by build tag, have been supported (Go
 1.1+). There's also a package to import to always force the latest minor
 version of the runtime:

```Go
import _ "github.com/theckman/goconstraint/latest"
```

 While this may be enticing to use, please remember that if you do use it a new
 Go release may break your software's ability to build. Using the
 version-specific constraints is recommended.

 It's also recommended that you include a comment with the import, to indicate
 which functionality/feature is required from that Go version (while also
 including a Golang issue link if possible).

 If you try to build a project against a runtime that's too old, you'll see a
 build failure similar to this with the version number changing based on the
 requirement:

```Go
goconstraint/go1.9/gte/constraint.go:10: undefined: __SOFTWARE_REQUIRES_GO_VERSION_1_9__
```
