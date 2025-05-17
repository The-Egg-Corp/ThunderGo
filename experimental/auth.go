package experimental

type AuthOptions struct {
	Code        string  `json:"code"`
	RedirectURI *string `json:"redirect_uri"`
}

type AuthResponse struct {
	Username  string  `json:"username"`
	Email     *string `json:"email"`
	SessionID string  `json:"session_id"`
}

// // Mimicks the 'Login with GitHub' button on the Thunderstore website.
// //
// // If RedirectURI is ni or unspecified, it will default to "127.0.0.1".
// func LoginWithGithub(auth AuthOptions) (AuthResponse, error) {
// 	return SendAuthComplete(auth, "github")
// }

// // Mimicks the 'Login with GitHub' button on the Thunderstore website.
// //
// // If RedirectURI is nil or unspecified, it will default to "127.0.0.1".
// func LoginWithDiscord(auth AuthOptions) (AuthResponse, error) {
// 	return SendAuthComplete(auth, "discord")
// }

// TODO: Implement this.
// func SendAuthComplete(auth AuthOptions, provider string) (AuthResponse, error) {
// 	// send code and redirect uri to thunderstore

// 	if auth.RedirectURI == nil {
// 		*auth.RedirectURI = "127.0.0.1"
// 	}

// 	return AuthResponse{}, nil
// }
