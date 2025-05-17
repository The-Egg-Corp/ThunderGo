package tests

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/the-egg-corp/thundergo/util"
	TSGOV1 "github.com/the-egg-corp/thundergo/v1"
)

var comm = TSGOV1.Community{Identifier: "lethal-company"}
var commPkgs TSGOV1.PackageList // Cache for all community tests. Set once in setup()

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		log.Fatal(errors.New("failed to get community packages. cannot use empty slice"))
		return
	}

	os.Exit(m.Run())
}

func setup() error {
	pkgs, err := comm.AllPackages()
	if err != nil {
		return errors.New("failed to get community packages.\n" + err.Error())
	}

	if pkgs.Size() < 1 {
		return errors.New("failed to get community packages. cannot use empty slice")
	}

	commPkgs = pkgs
	return nil
}

func TestPackagesFromList(t *testing.T) {
	pkgs, err := TSGOV1.PackagesFromCommunities(TSGOV1.NewCommunityList("subnautica", "valheim"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(pkgs.Size())
}

func TestCommunityPackages(t *testing.T) {
	pkgs, _ := comm.AllPackages()

	// pkgs = pkgs.Filter(func(pkg TSGOV1.Package) bool {
	// 	return pkg.Owner == "Owen3H"
	// })

	util.PrettyPrint(pkgs.Size())
}

func TestPackageGetMethods(t *testing.T) {
	pkg := commPkgs.GetExact("Owen3H-CSync")
	if pkg == nil {
		t.Fatal(errors.New("could not get package by its full name"))
	}

	pkg = commPkgs.GetByUUID("13d217b1-1e90-431a-a826-cd29c9eaea36")
	if pkg == nil {
		t.Fatal(errors.New("could not get package using UUID"))
	}

	pkg = commPkgs.Get("Owen3H", "CSync")
	if pkg == nil {
		t.Fatal(errors.New("could not get package given the name and author"))
	}

	ver := pkg.GetVersion("3.0.0")
	if ver == nil {
		t.Fatal(errors.New("could not find specific package version"))
	}
}

func TestMetrics(t *testing.T) {
	pkg := comm.GetPackage("Owen3H", "CSync")
	if pkg == nil {
		t.Fatal(errors.New("error retrieving metrics: package not found"))
	}

	metrics, err := pkg.Metrics()
	if err != nil {
		t.Fatal(err)
	}

	if util.Zero(metrics) {
		t.Fatal(errors.New("could not retreive metrics on existing package"))
	}

	//util.PrettyPrint(metrics)
}

func TestPackageDates(t *testing.T) {
	owner := "Owen3H"
	pkgName := "CSync"

	pkg := commPkgs.Get(owner, pkgName)
	if pkg == nil {
		t.Fatalf("error testing package dates. package %s-%s not found", owner, pkgName)
	}

	if pkg.DateCreated.IsZero() {
		t.Error("DateCreated should be valid and not its zero value.")
	}

	if pkg.DateUpdated.IsZero() {
		t.Error("DateUpdated should be valid and not its zero value.")
	}

	//util.PrettyPrint("Created " + pkg.DateCreated.Humanize())
	//util.PrettyPrint("Updated " + pkg.DateUpdated.Humanize())
}

func TestPackageFilter(t *testing.T) {
	filtered := commPkgs.ExcludeCategories("modpack", "modpacks")
	fmt.Println(filtered.Size())
}

func TestDownloadVersion(t *testing.T) {
	// if os.Getenv("GITHUB_ACTIONS") == "true" {
	// 	t.Skip()
	// }

	pkg := comm.GetPackage("Owen3H", "CSync")
	if pkg == nil {
		t.Fatal("error downloading version: package not found")
	}

	_, err := pkg.LatestVersion().Download()
	if err != nil {
		t.Fatalf("error downloading version:\n%v", err)
	}
}

func TestPackageFromCommunity(t *testing.T) {
	pkg := TSGOV1.PackageFromCommunity("lethal-company", "Megalophobia", "MEGALOPHOBIA")
	if pkg == nil {
		t.Fatal("package not found in community")
	}
}

func TestRatePackage(t *testing.T) {
	t.Skip() // Need to sort out auth before enabling this test.

	pkg := comm.GetPackage("Owen3H", "CSync")

	listing, err := pkg.Rate()
	if err != nil {
		t.Fatal(err)
	}

	util.PrettyPrint(listing)
}
