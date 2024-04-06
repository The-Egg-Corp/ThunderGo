package tests

import (
	"fmt"
	TSGOExp "github.com/The-Egg-Corp/ThunderGo/experimental"
	"github.com/The-Egg-Corp/ThunderGo/util"
	TSGOV1 "github.com/The-Egg-Corp/ThunderGo/v1"
	"testing"
)

// region Experimental Tests
func TestPackage(t *testing.T) {
	var err error
	var pkg *TSGOExp.Package

	pkg, err = TSGOExp.GetPackage("Owen3H", "CSync")

	if err != nil {
		t.Fatalf(err.Error())
	}

	cl, _ := pkg.Latest.Readme()
	fmt.Println(util.Prettify(cl))
}

//endregion

// region V1 Tests
func TestAllPackages(t *testing.T) {
	var err error
	var pkgs TSGOV1.PackageList

	pkgs, err = TSGOV1.GetAllPackages()
	if err != nil {
		t.Fatalf(err.Error())
	}

	pkg := pkgs.Get("Owen3H", "CSync")
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

//endregion
