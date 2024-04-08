package v1

import (
	"fmt"
	"github.com/the-egg-corp/thundergo/util"
)

type Community struct {
	Identifier string
}

// Returns a list of all packages (mods) within this Community.
func (comm Community) AllPackages(predicate ...func(item Package, index int) bool) (PackageList, error) {
	endpoint := fmt.Sprint("c/", comm.Identifier, "/api/v1/package")
	pkgs, err := util.JsonGetRequest[PackageList](endpoint)

	if err != nil {
		return nil, err
	}

	pkgCache = pkgs
	return pkgs, nil
}

func (comm Community) GetPackage(author string, name string) *Package {
	if pkgCache != nil {
		return pkgCache.Get(author, name)
	}

	return nil
}
