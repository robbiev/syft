/*
Package javascript provides a concrete Cataloger implementation for JavaScript ecosystem files (yarn and npm).
*/
package javascript

import (
	"github.com/anchore/syft/syft/pkg/cataloger/generic"
)

// NewPackageCataloger returns a new cataloger object for NPM.
func NewPackageCataloger() *generic.Cataloger {
	return generic.NewCataloger("javascript-package-cataloger").
		WithParserByGlobs(parsePackageJSON, "**/package.json")
}

// NewLockCataloger returns a new cataloger object for NPM (and NPM-adjacent) lock files.
func NewLockCataloger() *generic.Cataloger {
	return generic.NewCataloger("javascript-lock-cataloger").
		WithParserByGlobs(parsePackageLock, "**/package-lock.json").
		WithParserByGlobs(parseYarnLock, "**/yarn.lock").
		WithParserByGlobs(parsePnpmLock, "**/pnpm-lock.yaml")
}
