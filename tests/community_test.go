package tests

import (
	"fmt"
	TSGO "github.com/The-Egg-Corp/ThunderGo/experimental"
	"github.com/The-Egg-Corp/ThunderGo/util"
	"testing"
)

func TestCommunity(t *testing.T) {
	comm, found := TSGO.GetCommunity("lethal-company")

	if found == false {
		t.Error("Could not find the specified community!")
	}

	categories, _ := comm.Categories()
	fmt.Println(util.Prettify(categories))
}

func TestCommunities(t *testing.T) {
	comms, err := TSGO.GetCommunities()

	if err != nil {
		t.Error("Could not get list of communities!")
	}

	fmt.Println(util.Prettify(comms.Results))
}
