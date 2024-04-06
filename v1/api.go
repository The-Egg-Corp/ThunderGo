package v1

import (
	"github.com/the-egg-corp/thundergo/util"
)

func GetAllPackages() (PackageList, error) {
	return util.JsonGetRequest[PackageList]("api/v1/package")
}
