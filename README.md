# Tag

Magefile support for tagging and releasing at Invopop.

## Usage

Example `mage.go` file:

```
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
