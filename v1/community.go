package v1

import (
	"fmt"
	"thundergo/util"
)

var PackageCache PackageList

type PackageList []Package

func (list PackageList) Filter(predicate func(pkg Package) bool) PackageList {
	arr := make(PackageList, 0, len(list))

	for _, v := range list {
		if predicate(v) {
			arr = append(arr, v)
		}
	}

	return arr
}

func (list PackageList) Size() int {
	return len(list)
}

type Community struct {
	Identifier string
}

func (comm Community) AllPackages(predicate ...func(item Package, index int) bool) (PackageList, error) {
	endpoint := fmt.Sprint("c/", comm.Identifier, "/api/v1/package")
	return util.JsonGetRequest[PackageList](endpoint)
}

func (comm Community) GetPackage(name string) {

}
