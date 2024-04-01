package main

import (
	"fmt"
	"testing"
	TSGO "thundergo/experimental"
	"thundergo/util"
)

func TestPkg(t *testing.T) {
	var err error
	var pkg *TSGO.Package

	pkg, err = TSGO.GetPackage("Owen3H", "CSync")

	if err != nil {
		t.Fatalf(err.Error())
	}

	cl, _ := pkg.Latest.GetChangelog()
	fmt.Println(util.Prettify(cl))
}

func TestCommunity(t *testing.T) {
	comm, found := TSGO.GetCommunity("riskofrain2")

	if found == false {
		t.Error("Could not find the specified community!")
	}

	categories, _ := comm.Categories()
	fmt.Println(util.Prettify(categories))
}
