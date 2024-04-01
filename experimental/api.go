package experimental

import (
	"fmt"
	"thundergo/util"

	"github.com/samber/lo"
)

func GetCommunities() (CommunityList, error) {
	return util.JsonRequest[CommunityList]("api/experimental/community")
}

func GetCommunity(nameOrId string) (*Community, bool) {
	communities, err := GetCommunities()

	if err != nil {
		return nil, false
	}

	return lo.Find(communities.Results, func(comm *Community) bool {
		return comm.Name == nameOrId || comm.Identifier == nameOrId
	})
}

func GetPackage(author string, name string) (*Package, error) {
	endpoint := fmt.Sprint("api/experimental/package/", author, "/", name)
	return util.JsonRequest[*Package](endpoint)
}
