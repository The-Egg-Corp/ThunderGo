package tests

import (
	"fmt"
	"testing"
	TSGO "thundergo/experimental"
	"thundergo/util"
)

func TestCommunity(t *testing.T) {
	comm, found := TSGO.GetCommunity("riskofrain2")

	if found == false {
		t.Error("Could not find the specified community!")
	}

	categories, _ := comm.Categories()
	fmt.Println(util.Prettify(categories))
}
