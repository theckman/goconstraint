// Copyright (c) 2017 Tim Heckman
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file at the root of this repository.

// Package gtego12 should only be used as a blank import. If imported, it
// will only compile if the Go runtime version is >= 1.2.
package gtego12

// This will fail to compile if the Go runtime version isn't >= 1.2.
var _ = __SOFTWARE_REQUIRES_GO_VERSION_1_2__
