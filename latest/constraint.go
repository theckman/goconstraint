// Copyright (c) 2017 Tim Heckman
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file at the root of this repository.

// Package latestgo should only be used as a blank import. If imported, it
// will only compile if the Go runtime version is == the latest version.
//
// Instead of using this package, you should use one of the versioned
// constraints. Otherwise your code may suddenly fail to build.
package latestgo

import _ "github.com/theckman/goconstraint/go1.9/gte"
