package experimental

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	_ "image/png"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/hashicorp/go-version"
	"github.com/the-egg-corp/thundergo/util"
)

const MAX_MARKDOWN_SIZE = 1000 * 100
const MAX_ICON_SIZE = 1024 * 1024 * 6

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

// TODO: Implement this. Should take an auth key which the user gathers from 'Service Accounts'
// func SubmitPackage(data []byte) (bool, error) {
// 	return false, nil
// }

func ValidateReadme(data []byte) (bool, error) {
	if !utf8.Valid(data) {
		return false, errors.New("error parsing readme: file is not UTF-8 compatible")
	}

	bom := []byte{0xEF, 0xBB, 0xBF}
	if bytes.HasPrefix(data, bom) {
		return false, fmt.Errorf("readme cannot begin with a UTF-8 BOM")
	}

	str := string(data)
	if len(str) > MAX_MARKDOWN_SIZE {
		return false, fmt.Errorf("readme is too large: max file size is 100kb")
	}

	return true, nil
}

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

// Decodes image data and validates that the image passes the following requirements:
//
// - Max file size does not exceed 6MB.
//
// - Is in the PNG format.
//
// - Dimensions match 256x256.
func ValidateIcon(fileName string, data []byte) (bool, error) {
	// Check bytes dont exceed
	if len(data) > MAX_ICON_SIZE {
		return false, errors.New("invalid icon: max file size is 6MB")
	}

	// Decode data into the Image type.
	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return false, err
	}

	if format != "png" {
		return false, errors.New("invalid icon: must be in PNG format")
	}

	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// Verify dimensions
	if width != 256 && height != 256 {
		return false, errors.New("invalid icon: dimensions must match 256x256")
	}

	return true, nil
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
