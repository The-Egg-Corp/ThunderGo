package experimental

type AuthOptions struct {
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
}

type AuthResponse struct {
	Username  string  `json:"username"`
	Email     *string `json:"email"`
	SessionID string  `json:"session_id"`
}

// TODO: Implement this. Edit: forgor why I needed this
// func LoginWithGithub(auth AuthOptions) (AuthResponse, error) {
// 	return AuthResponse{}, nil
// }
