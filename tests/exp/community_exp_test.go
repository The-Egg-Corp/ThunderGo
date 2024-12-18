package tests

import (
	TSGO "github.com/the-egg-corp/thundergo/experimental"

	//TSGOV1 "github.com/the-egg-corp/thundergo/v1"
	"testing"
)

// region Experimental Tests
func TestCommunityExp(t *testing.T) {
	comm, found, _ := TSGO.GetCommunity("lethal-company")

	if !found {
		t.Error("Could not find the specified community!")
	}

	_, err := comm.Categories()
	if err != nil {
		t.Fatal(err)
	}
	//util.PrettyPrint(categories)
}

func TestCommunitiesExp(t *testing.T) {
	_, err := TSGO.GetCommunities()

	if err != nil {
		t.Error("Could not get list of communities!")
	}

	//util.PrettyPrint(comms)
}

//endregion
