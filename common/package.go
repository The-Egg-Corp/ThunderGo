package common

type PackageCategory struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Response received for markdown files like README etc.
type MarkdownResponse struct {
	Markdown string `json:"markdown"`
}
