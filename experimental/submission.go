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

// TODO: Implement this
// func SubmitPackage(data []byte) (bool, error) {
// 	return false, nil
// }

// TODO: Implement this
// func ValidateReadme(data []byte) (bool, error) {
// 	return false, nil
// }

func ValidateManifest(author string, data []byte) (valid bool, errs []string, err error) {
	var manifest ManifestMetadata

	err = json.Unmarshal(data, &manifest)
	if err != nil {
		return false, nil, errors.New("error deserializing manifest: \n" + err.Error())
	}

	AddIfEmpty(&errs, &manifest.Name, "required property 'name' is empty or unspecified")
	AddIfInvalid(&errs, &manifest.Name, "property 'name' must contain only valid characters (a-z A-Z 0-9 _)")
	AddIfEmpty(&errs, &manifest.Description, "required property 'description' is empty or unspecified")

	verEmpty := AddIfEmpty(&errs, &manifest.VersionNumber, "required property 'version_number' is empty or unspecified")
	if !verEmpty {
		matched, _ := util.CheckSemVer(manifest.VersionNumber)
		if matched {
			pkg, _ := GetPackage(author, manifest.Name)
			if pkg != nil {
				verA, _ := version.NewSemver(manifest.VersionNumber)
				verB, _ := version.NewSemver(pkg.Latest.VersionNumber)

				if verA.LessThanOrEqual(verB) {
					Add(&errs, "property 'version_number' must be higher than the latest")
				}
			}
		} else {
			Add(&errs, "property 'version_number' does not follow semantic versioning (major.minor.patch)")
		}
	}

	if manifest.WebsiteURL == nil {
		Add(&errs, "property 'website_url' is empty or unspecified")
	} else {
		url := strings.ToLower(*manifest.WebsiteURL)
		if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
			Add(&errs, "property 'website_url' must be a valid URL")
		}
	}

	if manifest.Dependencies == nil {
		Add(&errs, "required property 'dependencies' is empty or unspecified")
	} else {
		for _, dep := range manifest.Dependencies {
			fullName := author + "-" + manifest.Name
			if strings.Contains(strings.ToLower(dep), strings.ToLower(fullName)) {
				Add(&errs, "property 'dependencies' is invalid. cannot depend on self")
			}

			// TODO: Check multiple versions of same package
		}
	}

	return len(errs) < 1, errs, nil
}

// Decodes image data and validates that the image is a PNG and the dimensions are 256x256.
//
// Additionally, if the file name is specified, it will validate that it is named correctly.
func ValidateIcon(params IconValidatorParams) (bool, error) {
	// TODO: Why doesn't this work? "icon.png.jpg" should not work!
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

func Add(arr *[]string, errStr string) {
	*arr = append(*arr, errStr)
}

func AddIfEmpty(arr *[]string, str *string, errStr string) bool {
	empty := *str == "" || str == nil
	if empty {
		Add(arr, errStr)
	}

	return empty
}

func AddIfInvalid(arr *[]string, str *string, errStr string) {
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, *str)
	if !matched {
		Add(arr, errStr)
	}
}
