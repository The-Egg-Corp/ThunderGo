package tests

import (
	"fmt"
	"testing"
	TSGO "thundergo"
	"thundergo/util"
)

func TestCommunity(t *testing.T) {
	comm, found := TSGO.Experimental.GetCommunity("riskofrain2")

	if found == false {
		t.Error("Could not find the specified community!")
	}

	categories, _ := comm.Categories()
	fmt.Println(util.Prettify(categories))
}
