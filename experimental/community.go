package experimental

import (
	"fmt"
	"thundergo/common"
	"thundergo/util"
)

type Category = common.PackageCategory

type CommunityList struct {
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []*Community `json:"results"`
}

type Community struct {
	Identifier              string  `json:"identifier"`
	Name                    string  `json:"name"`
	DiscordURL              *string `json:"discord_url"`
	WikiURL                 *string `json:"wiki_url"`
	PackageApprovalRequired bool    `json:"require_package_listing_approval"`
}

type CommunityCategories struct {
	Results    []Category `json:"results"`
	Pagination struct {
		NextLink     any `json:"next_link"`
		PreviousLink any `json:"previous_link"`
	} `json:"pagination"`
}

func (community Community) AllPackages() ([]Package, error) {
	return nil, nil
}

func (community Community) Categories() ([]Category, error) {
	endpoint := fmt.Sprint("api/experimental/community/", community.Identifier, "/category")
	res, err := util.JsonRequest[CommunityCategories](endpoint)

	return res.Results, err
}