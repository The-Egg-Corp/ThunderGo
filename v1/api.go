package v1

import (
	"thundergo/util"
)

func GetAllPackages() (PackageList, error) {
	return util.JsonGetRequest[PackageList]("api/v1/package")
}
