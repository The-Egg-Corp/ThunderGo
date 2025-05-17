package v1

import (
	"fmt"

	"github.com/the-egg-corp/thundergo/util"
)

type CommunityList []Community

type Community struct {
	Identifier string
}

func NewCommunityList(identifiers ...string) (list CommunityList) {
	for _, cur := range identifiers {
		list = append(list, Community{Identifier: cur})
	}

	return list
}

// Returns a list of all packages (mods) within this community.
func (comm Community) AllPackages(predicate ...func(item Package, index int) bool) (PackageList, error) {
	endpoint := fmt.Sprint("c/", comm.Identifier, "/api/v1/package")

	pkgs, code, err := util.JsonGetRequest[PackageList](endpoint)
	if err != nil {
		return nil, fmt.Errorf("[%s] could not get all packages. status code: %d\n%v", comm.Identifier, code, err)
	}

	return *pkgs, nil
}

// Gets a single package from this community given the owner and package name.
func (comm Community) GetPackage(owner string, name string) *Package {
	pkgs, err := comm.AllPackages()
	if err != nil {
		return nil
	}

	return pkgs.Get(owner, name)
}
