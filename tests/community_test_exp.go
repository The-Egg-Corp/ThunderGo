package tests

import (
	TSGO "github.com/the-egg-corp/thundergo/experimental"
	"github.com/the-egg-corp/thundergo/util"

	//TSGOV1 "github.com/the-egg-corp/thundergo/v1"
	"testing"
)

// region Experimental Tests
func TestCommunityExp(t *testing.T) {
	comm, found, _ := TSGO.GetCommunity("lethal-company")

	if !found {
		t.Error("Could not find the specified community!")
	}

	categories, _ := comm.Categories()
	util.PrettyPrint(categories)
}

func TestCommunitiesExp(t *testing.T) {
	comms, err := TSGO.GetCommunities()

	if err != nil {
		t.Error("Could not get list of communities!")
	}

	util.PrettyPrint(comms.Results)
}

//endregion
