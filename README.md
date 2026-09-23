# Tag

Magefile support for tagging and releasing at Invopop.

## Installation

```bash
go get github.com/invopop/tag
```

## Releasing

Example `mage.go` file:

```go
//go:build mage
// +build mage

package main

import (
	"github.com/invopop/tag"
)

// Release a new version based on the current branch and timestamp.
func Release() error {
	return tag.Release()
}
```

Running `mage release` will:

1. Fail if there are uncommitted changes.
2. Tag the current commit with a timestamp-based version, unless it's already tagged.
3. Push tags to the remote.

Versions use the UTC time of the release. Commits on `main` get the bare timestamp, and other branches are prefixed with the branch name. Characters that aren't valid in git or Docker tags, such as `/`, are replaced with `-`:

| Branch           | Tag                             |
| ---------------- | ------------------------------- |
| `main`           | `v20260923T1504`                |
| `fix-login`      | `fix-login-v20260923T1504`      |
| `feat/new-thing` | `feat-new-thing-v20260923T1504` |

## Build Information

`tag.Now()` collects details about the current build from git, and can be used to embed version information in binaries at build time:

```go
// Build the application with version details embedded.
func Build() error {
	return sh.Run("go", "build", "-ldflags", tag.Now().LDFlags(), "./cmd/app")
}
```

This sets the `BuildVersion` and `BuildTime` variables in your `main` package, which must be declared:

```go
package main

var (
	BuildVersion = "dev"
	BuildTime    = ""
)
```

`Version()` returns the current tag if the commit has one. Otherwise it returns `<branch>-<short-commit>`, with a `-WIP` suffix if there are uncommitted changes. Outside a git repository, such as in some CI environments, it falls back to the `VERSION` environment variable.
