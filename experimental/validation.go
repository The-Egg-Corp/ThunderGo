package experimental

import (
	"bytes"
	"errors"
	"image"
	_ "image/png"
)

type ValidatorResponse struct {
	Success bool `json:"success"`
}

type IconValidatorParams struct {
	IconData string `json:"icon_data"`
}

func ValidateReadme() (ValidatorResponse, error) {
	return ValidatorResponse{}, nil
}

func ValidateManifest() (ValidatorResponse, error) {
	return ValidatorResponse{}, nil
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
	} else {
		return false, errors.New("image dimensions did not match: 256x256")
	}
}
