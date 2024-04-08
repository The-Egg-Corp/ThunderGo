package v1

import (
	"github.com/the-egg-corp/thundergo/util"
)

// The list of every package on Thunderstore across every community.
func GetAllPackages() (PackageList, error) {
	return util.JsonGetRequest[PackageList]("api/v1/package")
}
