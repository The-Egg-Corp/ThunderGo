package experimental

import (
	"encoding/base64"
	"fmt"
	"github.com/The-Egg-Corp/ThunderGo/util"

	"github.com/samber/lo"
)

type Base64String string

func (b Base64String) String() string {
	return base64.StdEncoding.EncodeToString([]byte(b))
}

func GetCommunities() (CommunityList, error) {
	return util.JsonGetRequest[CommunityList]("api/experimental/community")
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
	return util.JsonGetRequest[*Package](endpoint)
}
