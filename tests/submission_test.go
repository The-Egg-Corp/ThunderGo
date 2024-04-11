package tests

import (
	"os"
	"testing"

	TSGO "github.com/the-egg-corp/thundergo/experimental"
	"github.com/the-egg-corp/thundergo/util"
)

func TestValidateIcon(t *testing.T) {
	var err error
	icon, err := os.ReadFile("../test_icon.png")

	if err != nil {
		t.Fatalf(err.Error())
	}

	valid, err := TSGO.ValidateIcon(TSGO.IconValidatorParams{
		ImageData: icon,
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	util.PrettyPrint(valid)
}

func TestValidateManifest(t *testing.T) {
	var errs []string
	data, err := os.ReadFile("../test_manifest.json")

	if err != nil {
		t.Fatalf(err.Error())
	}

	valid, errs := TSGO.ValidateManifest(data)

	if err != nil {
		t.Fatalf(err.Error())
	}

	util.PrettyPrint(valid)
	util.PrettyPrint(errs)
}
