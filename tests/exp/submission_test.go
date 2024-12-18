package tests

import (
	"fmt"
	"os"
	"testing"

	TSGO "github.com/the-egg-corp/thundergo/experimental"
	"github.com/the-egg-corp/thundergo/util"
)

func TestValidateIcon(t *testing.T) {
	t.Skip()

	icon, err := os.ReadFile("../test-pkg/icon.png.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}

	valid, err := TSGO.ValidateIcon(TSGO.IconValidatorParams{ImageData: icon})
	if err != nil {
		t.Fatal(err.Error())
	}

	util.PrettyPrint(valid)
}

func TestValidateManifest(t *testing.T) {
	t.Skip()

	var errs []string

	data, err := os.ReadFile("../test-pkg/manifest.json")
	if err != nil {
		t.Fatal(err.Error())
	}

	valid, errs, err := TSGO.ValidateManifest("Owen3H", data)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println("Valid: ", valid)

	if len(errs) > 0 {
		util.PrettyPrint(errs)

		if valid {
			t.Fatal("errors were returned, but manifest is still valid")
		}

		return
	}

	if !valid {
		t.Fatal("manifest was marked as invalid despite empty errors array")
	}
}
