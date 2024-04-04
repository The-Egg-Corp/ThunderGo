package common

import "thundergo/util"

type BasePackageMetadata struct {
	Name        string        `json:"name"`
	FullName    string        `json:"full_name"`
	DateCreated util.DateTime `json:"date_created"`
}

type PackageCategory struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Response received for markdown files like README etc.
type MarkdownResponse struct {
	Markdown string `json:"markdown"`
}
