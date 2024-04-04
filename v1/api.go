package v1

import (
	"fmt"
	"thundergo/util"
)

func GetAllPackages() ([]PackageListing, error) {
	return util.JsonGetRequest[[]PackageListing]("api/v1/package")
}

func GetCommunityPackages(community string) ([]PackageListing, error) {
	endpoint := fmt.Sprint("c/", community, "/api/v1/package")
	return util.JsonGetRequest[[]PackageListing](endpoint)
}
