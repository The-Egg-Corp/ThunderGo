package experimental

import (
	"bytes"
	"errors"
	"image"
	_ "image/png"
)

func ValidateReadme(data []byte) (bool, error) {
	return false, nil
}

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
