package experimental

import (
	"fmt"
	"thundergo/util"
)

type CommunityList struct {
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []*Community `json:"results"`
}

type Community struct {
	Identifier              string `json:"identifier"`
	Name                    string `json:"name"`
	DiscordURL              string `json:"discord_url"`
	WikiURL                 string `json:"wiki_url"`
	PackageApprovalRequired bool   `json:"require_package_listing_approval"`
}

type CommunityCategories struct {
	Pagination struct {
		NextLink     any `json:"next_link"`
		PreviousLink any `json:"previous_link"`
	} `json:"pagination"`
	Results []PackageCategory `json:"results"`
}

func (community Community) Categories() ([]PackageCategory, error) {
	endpoint := fmt.Sprint("api/experimental/community/", community.Identifier, "/category")
	res, err := util.JsonRequest[CommunityCategories](endpoint)

	return res.Results, err
}
