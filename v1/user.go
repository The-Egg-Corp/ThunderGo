package v1

import "thundergo/util"

type UserMedia struct {
	UUID        string        `json:"uuid"`
	FileName    string        `json:"filename"`
	Size        uint64        `json:"size"`
	DateCreated util.DateTime `json:"datetime_created"`
	Expiry      *string       `json:"expiry"`
	Status      []string      `json:"status"`
}

type UserTeam struct {
	Name        string `json:"name"`
	Role        string `json:"role"`
	MemberCount uint8  `json:"member_count"`
}

type UserProfile struct {
	Username      string                 `json:"username"`
	Capabilities  []*string              `json:"capabilities"`
	Connections   []SocialAuthConnection `json:"connections"`
	Subscription  SubscriptionStatus     `json:"subscription"`
	RatedPackages []*string              `json:"rated_packages"`
	Teams         []*string              `json:"teams"`
	TeamsFull     []UserTeam             `json:"teams_full"`
}

type SocialAuthConnection struct {
	Provider string `json:"provider"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type SubscriptionStatus struct {
	Expires string `json:"expires"`
}
