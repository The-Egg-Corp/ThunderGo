package experimental

import (
	"bytes"
	"errors"
	"image"
	_ "image/png"
)

type ValidatorResponse struct {
	Success bool `json:"success" must:"true"`
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

// func ValidateIcon(data Base64String) (ValidatorResponse, error) {
// 	body := IconValidatorParams{
// 		IconData: data.String(),
// 	}

// 	endpoint := "api/experimental/submission/validate/icon"
// 	return util.JsonPostRequest[ValidatorResponse](endpoint, body)
// }

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
