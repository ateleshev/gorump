// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build dragonfly linux netbsd openbsd rumprun solaris

package os

// supportsCloseOnExec reports whether the platform supports the
// O_CLOEXEC flag.
const supportsCloseOnExec = true
