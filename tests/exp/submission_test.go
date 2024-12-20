package tests

import (
	"fmt"
	"os"
	"testing"

	TSGO "github.com/the-egg-corp/thundergo/experimental"
)

const iconPath = "../test-pkg/icon.png.jpg"
const readmePath = "../test-pkg/README.md"
const manifestPath = "../test-pkg/manifest.json"

func TestValidateIcon(t *testing.T) {
	t.Skip()

	icon, err := os.ReadFile(iconPath)
	if err != nil {
		t.Fatal(err.Error())
	}

	valid, err := TSGO.ValidateIcon(icon)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(valid)
}

func TestValidateReadme(t *testing.T) {
	t.Skip()

	data, err := os.ReadFile(readmePath)
	if err != nil {
		t.Fatal(err.Error())
	}

	valid, errs, err := TSGO.ValidateReadme(data)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Valid: ", valid)

	if len(errs) > 0 {
		fmt.Println(errs)

		if valid {
			t.Fatal("errors were returned, but readme is still valid")
		}

		return
	}
}

func TestValidateManifest(t *testing.T) {
	t.Skip()

	data, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatal(err.Error())
	}

	valid, errs, err := TSGO.ValidateManifest("Owen3H", data)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Valid: ", valid)

	if len(errs) > 0 {
		fmt.Println(errs)

		if valid {
			t.Fatal("errors were returned, but manifest is still valid")
		}

		return
	}

	if !valid {
		t.Fatal("manifest was marked as invalid despite empty errors array")
	}
}
