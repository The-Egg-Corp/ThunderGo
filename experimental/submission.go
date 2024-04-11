package experimental

import (
	"bytes"
	"encoding/json"
	"errors"
	"image"
	_ "image/png"

	"github.com/the-egg-corp/thundergo/util"
)

type PackageSubmissionMetadata struct {
	UUID                string   `json:"upload_uuid"`
	Author              string   `json:"author_name"`
	Communities         []string `json:"communities"`
	CommunityCategories []string `json:"community_categories"`
	Categories          []string `json:"categories"`
	HasNsfwContent      bool     `json:"has_nsfw_content"`
}

type ManifestMetadata struct {
	Name          string   `json:"name"`
	VersionNumber string   `json:"version_number"`
	WebsiteURL    string   `json:"website_url"`
	Description   string   `json:"description"`
	Dependencies  []string `json:"dependencies"`
}

type IconValidatorParams struct {
	FileName  string
	ImageData []byte
}

// TODO: Implement this
func SubmitPackage(data []byte) (bool, error) {
	return false, nil
}

// TODO: Implement this
func ValidateReadme(data []byte) (bool, error) {
	return false, nil
}

// TODO: Implement this
func ValidateManifest(data []byte) (bool, []string) {
	var manifest ManifestMetadata
	json.Unmarshal(data, &manifest)

	var errors []string

	AddIfEmpty(&errors, &manifest.Name, "required manifest property 'name' is empty or unspecified")
	AddIfEmpty(&errors, &manifest.VersionNumber, "required manifest property 'version_number' is empty or unspecified")
	AddIfEmpty(&errors, &manifest.Description, "required manifest property 'description' is empty or unspecified")

	sv, _ := util.CheckSemVer(manifest.VersionNumber)
	AddIfFalse(&errors, &sv, "manifest version does not follow semantic versioning (major.minor.patch)")

	return len(errors) < 1, errors
}

func AddIfEmpty(arr *[]string, str *string, errStr string) {
	if *str == "" || str == nil {
		*arr = append(*arr, errStr)
	}
}

func AddIfFalse(arr *[]string, val *bool, errStr string) {
	if !*val {
		*arr = append(*arr, errStr)
	}
}

// Decodes image data and validates that the image is a PNG and the dimensions are 256x256.
//
// Additionally, if the file name is specified, it will validate that it is named correctly.
func ValidateIcon(params IconValidatorParams) (bool, error) {
	if params.FileName != "" && params.FileName != "icon.png" {
		return false, errors.New("image name provided did not match: icon.png")
	}

	// Decode data into the Image type.
	img, _, err := image.Decode(bytes.NewReader(params.ImageData))
	if err != nil {
		return false, err
	}

	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// Verify dimensions
	if width == 256 && height == 256 {
		return true, nil
	}

	return false, errors.New("image dimensions did not match: 256x256")
}
