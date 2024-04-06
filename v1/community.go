package v1

import (
	"fmt"
	"thundergo/util"
)

type Community struct {
	Identifier string
}

func (comm Community) AllPackages(predicate ...func(item Package, index int) bool) (PackageList, error) {
	endpoint := fmt.Sprint("c/", comm.Identifier, "/api/v1/package")
	pkgs, err := util.JsonGetRequest[PackageList](endpoint)

	if err != nil {
		return nil, err
	}

	PackageCache = pkgs
	return pkgs, nil
}

func (comm Community) GetPackage(author string, name string) *Package {
	if PackageCache != nil {
		return PackageCache.Get(author, name)
	}

	return nil
}
