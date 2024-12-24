package experimental

import (
	"fmt"

	"github.com/the-egg-corp/thundergo/common"
	"github.com/the-egg-corp/thundergo/util"
)

//var commCache CommunityList

type Category = common.PackageCategory
type CommunityList []Community

func (list CommunityList) Size() int {
	return len(list)
}

type CommunitiesResponse struct {
	Results CommunityList `json:"results"`
}

type CommunityCategories struct {
	Results    []Category `json:"results"`
	Pagination struct {
		NextLink     any `json:"next_link"`
		PreviousLink any `json:"previous_link"`
	} `json:"pagination"`
}

type Community struct {
	Identifier              string  `json:"identifier"`
	Name                    string  `json:"name"`
	DiscordURL              *string `json:"discord_url"`
	WikiURL                 *string `json:"wiki_url"`
	PackageApprovalRequired bool    `json:"require_package_listing_approval"`
}

func (community Community) Categories() ([]Category, error) {
	endpoint := fmt.Sprint("api/experimental/community/", community.Identifier, "/category")
	res, _, err := util.JsonGetRequest[CommunityCategories](endpoint)

	if res == nil {
		return nil, err
	}

	return res.Results, err
}
