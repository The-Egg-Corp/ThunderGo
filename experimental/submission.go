package experimental

import (
	"bytes"
	"encoding/json"
	"errors"
	"image"
	_ "image/png"
	"regexp"
	"strings"

	"github.com/hashicorp/go-version"
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
	WebsiteURL    *string  `json:"website_url"`
	Description   string   `json:"description"`
	Dependencies  []string `json:"dependencies"`
}

type IconValidatorParams struct {
	FileName  string
	ImageData []byte
}

func NewErr(msg string) error {
	return errors.New(msg)
}

// TODO: Implement this
func SubmitPackage(data []byte) (bool, error) {
	return false, nil
}

// TODO: Implement this
func ValidateReadme(data []byte) (bool, error) {
	return false, nil
}

func ValidateManifest(author string, data []byte) (bool, []string, error) {
	var manifest ManifestMetadata
	var errors []string

	err := json.Unmarshal(data, &manifest)
	if err != nil {
		errors = append(errors, err.Error())
	}

	pkg, _ := GetPackage(author, manifest.Name)
	if pkg == nil {
		return false, nil, NewErr("package not found under the specified author")
	}

	AddIfEmpty(&errors, &manifest.Name, "required property 'name' is empty or unspecified")
	AddIfInvalid(&errors, &manifest.Name, "property 'name' must contain only valid characters (a-z A-Z 0-9 _)")
	AddIfEmpty(&errors, &manifest.Description, "required property 'description' is empty or unspecified")

	verEmpty := AddIfEmpty(&errors, &manifest.VersionNumber, "required property 'version_number' is empty or unspecified")
	if !verEmpty {
		sv, _ := util.CheckSemVer(manifest.VersionNumber)
		AddIfFalse(&errors, &sv, "property 'version_number' does not follow semantic versioning (major.minor.patch)")

		if !sv {
			return false, errors, nil
		}

		verA, _ := version.NewSemver(manifest.VersionNumber)
		verB, _ := version.NewSemver(pkg.Latest.VersionNumber)

		if verA.LessThanOrEqual(verB) {
			errors = append(errors, "property 'version_number' must be higher than the latest")
		}
	}

	if manifest.WebsiteURL == nil {
		errors = append(errors, "required property 'website_url' is unspecified")
	} else {
		url := strings.ToLower(*manifest.WebsiteURL)
		if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
			errors = append(errors, "property 'website_url' must be a valid URL")
		}
	}

	if manifest.Dependencies == nil {
		errors = append(errors, "manifest property 'dependencies' is required")
	}

	return len(errors) < 1, errors, nil
}

func AddIfEmpty(arr *[]string, str *string, errStr string) bool {
	empty := *str == "" || str == nil
	if empty {
		*arr = append(*arr, errStr)
	}

	return empty
}

func AddIfFalse(arr *[]string, val *bool, errStr string) {
	if !*val {
		*arr = append(*arr, errStr)
	}
}

func AddIfInvalid(arr *[]string, str *string, errStr string) {
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, *str)
	if !matched {
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
