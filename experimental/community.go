package experimental

import (
	"fmt"

	"github.com/the-egg-corp/thundergo/common"
	"github.com/the-egg-corp/thundergo/util"
)

//var commCache CommunityList

type Category = common.PackageCategory

type CommunityList []Community

type CommunitiesResponse struct {
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  CommunityList `json:"results"`
}

type CommunityCategories struct {
	Results    []Category `json:"results"`
	Pagination struct {
		NextLink     any `json:"next_link"`
		PreviousLink any `json:"previous_link"`
	} `json:"pagination"`
}

// region Community Struct
type Community struct {
	Identifier              string  `json:"identifier"`
	Name                    string  `json:"name"`
	DiscordURL              *string `json:"discord_url"`
	WikiURL                 *string `json:"wiki_url"`
	PackageApprovalRequired bool    `json:"require_package_listing_approval"`
}

func (community Community) AllPackages() ([]Package, error) {
	return nil, nil
}

func (community Community) Categories() ([]Category, error) {
	endpoint := fmt.Sprint("api/experimental/community/", community.Identifier, "/category")
	res, err := util.JsonGetRequest[CommunityCategories](endpoint)

	return res.Results, err
}

//endregion
