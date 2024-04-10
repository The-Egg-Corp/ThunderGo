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

// Not yet implemented.
func LoginWithGithub(auth AuthOptions) (AuthResponse, error) {
	return AuthResponse{}, nil
}
