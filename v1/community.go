package v1

import (
	"fmt"

	"github.com/the-egg-corp/thundergo/util"
)

type Community struct {
	Identifier string
}

// Returns a list of all packages (mods) within this community.
func (comm Community) AllPackages(predicate ...func(item Package, index int) bool) (PackageList, error) {
	endpoint := fmt.Sprint("c/", comm.Identifier, "/api/v1/package")
	pkgs, err := util.JsonGetRequest[PackageList](endpoint)

	if err != nil {
		return nil, err
	}

	//pkgCache = pkgs
	return pkgs, nil
}

// Gets a single package from this community given the owner and package name.
func (comm Community) GetPackage(author string, name string) *Package {
	// if pkgCache != nil {
	// 	return pkgCache.Get(author, name)
	// }

	pkgs, err := comm.AllPackages()

	if err != nil {
		return nil
	}

	return pkgs.Get(author, name)
}
