package tests

import (
	"fmt"
	"testing"
	TSGOExp "thundergo/experimental"
	"thundergo/util"
	TSGOV1 "thundergo/v1"
)

func TestPackage(t *testing.T) {
	var err error
	var pkg *TSGOExp.Package

	pkg, err = TSGOExp.GetPackage("Owen3H", "CSync")

	if err != nil {
		t.Fatalf(err.Error())
	}

	cl, _ := pkg.Latest.GetReadme()
	fmt.Println(util.Prettify(cl))
}

func TestAllPackages(t *testing.T) {
	var err error
	var pkgs TSGOV1.PackageList

	pkgs, err = TSGOV1.GetAllPackages()
	if err != nil {
		t.Fatalf(err.Error())
	}

	pkg := pkgs.Filter(func(item TSGOV1.Package) bool {
		return item.Name == "CSync"
	})

	fmt.Println(util.Prettify(pkg))
}

func TestCommunityPackages(t *testing.T) {
	comm := TSGOV1.Community{
		Identifier: "lethal-company",
	}

	pkgs, _ := comm.AllPackages()
	pkgs = pkgs.Filter(func(pkg TSGOV1.Package) bool {
		return pkg.Owner == "Owen3H"
	})

	fmt.Println(util.Prettify(pkgs))
}
