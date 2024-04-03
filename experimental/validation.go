package experimental

import (
	"thundergo/util"
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

func ValidateIcon(data Base64String) (ValidatorResponse, error) {
	body := IconValidatorParams{
		IconData: data.String(),
	}

	endpoint := "api/experimental/submission/validate/icon"
	return util.JsonPostRequest[ValidatorResponse](endpoint, body)
}
