package v1

import (
	"github.com/The-Egg-Corp/ThunderGo/util"
)

func GetAllPackages() (PackageList, error) {
	return util.JsonGetRequest[PackageList]("api/v1/package")
}
