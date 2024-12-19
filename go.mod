module github.com/the-egg-corp/thundergo

replace github.com/the-egg-corp/thundergo => ./

go 1.22.1

require (
	github.com/dustin/go-humanize v1.0.1
	github.com/go-resty/resty/v2 v2.16.2
	github.com/hashicorp/go-version v1.7.0
	github.com/samber/lo v1.39.0
	github.com/sanity-io/litter v1.5.5
)

require golang.org/x/sync v0.10.0

require (
	golang.org/x/exp v0.0.0-20241217172543-b2144cdd0a67 // indirect
	golang.org/x/net v0.33.0 // indirect
)
