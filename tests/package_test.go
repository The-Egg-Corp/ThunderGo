package tests

import (
	TSGOExp "github.com/the-egg-corp/thundergo/experimental"
	"github.com/the-egg-corp/thundergo/util"
	TSGOV1 "github.com/the-egg-corp/thundergo/v1"
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
	util.PrettyPrint(cl)
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

	util.PrettyPrint(pkgs)
}

func TestCommunityPackages(t *testing.T) {
	comm := TSGOV1.Community{
		Identifier: "lethal-company",
	}

	pkgs, _ := comm.AllPackages()
	pkgs = pkgs.Filter(func(pkg TSGOV1.Package) bool {
		return pkg.Owner == "Owen3H"
	})

	util.PrettyPrint(pkgs)
}

func TestPackageVersion(t *testing.T) {
	pkgs, err := TSGOV1.GetAllPackages()
	if err != nil {
		t.Fatalf(err.Error())
	}

	pkg := pkgs.Get("Owen3H", "CSync").GetVersion("2.0.0")
	util.PrettyPrint(pkg)
}

//endregion
