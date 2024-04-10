package experimental

import (
	"bytes"
	"errors"
	"image"
	_ "image/png"
)

type PackageSubmissionMetadata struct {
	Author              string   `json:"author_name"`
	Categories          []string `json:"categories"`
	Communities         []string `json:"communities"`
	HasNsfwContent      bool     `json:"has_nsfw_content"`
	UUID                string   `json:"upload_uuid"`
	CommunityCategories []string `json:"community_categories"`
}

// Not yet implemented.
func SubmitPackage() (bool, error) {
	return false, nil
}

// Not yet implemented.
func ValidateReadme(data []byte) (bool, error) {
	return false, nil
}

// Not yet implemented.
func ValidateManifest() (bool, error) {
	return false, nil
}

// Decodes image data and validates that the image is a PNG and the dimensions are 256x256.
func ValidateIcon(data []byte) (bool, error) {
	// Decode the image
	img, _, err := image.Decode(bytes.NewReader(data))
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
