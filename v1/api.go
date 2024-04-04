package v1

import (
	"thundergo/util"
)

func GetAllPackages() ([]PackageListing, error) {
	return util.JsonGetRequest[[]PackageListing]("api/v1/package")
}
