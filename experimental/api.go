package experimental

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/the-egg-corp/thundergo/util"

	"github.com/samber/lo"
)

type Base64String string

func (b Base64String) String() string {
	return base64.StdEncoding.EncodeToString([]byte(b))
}

func GetCommunities() (CommunityList, error) {
	res, err := util.JsonGetRequest[CommunitiesResponse]("api/experimental/community")

	if err != nil {
		return CommunityList{}, err
	}

	return res.Results, nil
}

// Get a specific [Community] by it's identifier or short name.
//
// If the name/id does not match any existing community, the result will be nil.
func GetCommunity(nameOrId string) (*Community, bool, error) {
	communities, err := GetCommunities()

	if err != nil {
		return nil, false, err
	}

	comm, found := lo.Find(communities, func(c Community) bool {
		return strings.EqualFold(c.Name, nameOrId) || strings.EqualFold(c.Identifier, nameOrId)
	})

	if !found {
		return nil, false, nil
	}

	return &comm, true, nil
}

// Get a single [Package] given it's owner and package short name.
//
// If an error occurred or it was not found, the result will be nil.
func GetPackage(author string, name string) (*Package, error) {
	endpoint := fmt.Sprint("api/experimental/package/", author, "/", name)
	return util.JsonGetRequest[*Package](endpoint)
}
