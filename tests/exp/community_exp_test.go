package tests

import (
	"testing"

	TSGO "github.com/the-egg-corp/thundergo/experimental"
)

func TestCommunityExp(t *testing.T) {
	comm, found, _ := TSGO.GetCommunity("lethal-company")

	if !found {
		t.Error("Could not find the specified community!")
	}

	_, err := comm.Categories()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCommunitiesExp(t *testing.T) {
	_, err := TSGO.GetCommunities()

	if err != nil {
		t.Error("Could not get list of communities!")
	}
}
