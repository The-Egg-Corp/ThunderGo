package thundergo

import (
	"fmt"
	"thundergo/util"

	"github.com/samber/lo"
)

var Experimental = ExperimentalClient{
	Domain: "https://thunderstore.io/api/experimental",
}

var V1 = V1Client{
	Domain: "https://thunderstore.io",
}

type ExperimentalClient struct {
	Domain string
}

type V1Client struct {
	Domain string
}

func (client ExperimentalClient) GetCommunities() (CommunityList, error) {
	return util.JsonRequest[CommunityList]("/community")
}

func (client ExperimentalClient) GetCommunity(nameOrId string) (*Community, bool) {
	communities, err := client.GetCommunities()

	if err != nil {
		return nil, false
	}

	return lo.Find(communities.Results, func(comm *Community) bool {
		return comm.Name == nameOrId || comm.Identifier == nameOrId
	})
}

func (client ExperimentalClient) GetPackage(author string, name string) (*PackageExperimental, error) {
	endpoint := fmt.Sprint("/package", author, "/", name)
	return util.JsonRequest[*PackageExperimental](endpoint)
}
