module github.com/the-egg-corp/thundergo

go 1.23.0
toolchain go1.24.2

replace github.com/the-egg-corp/thundergo => ./

require (
	github.com/dustin/go-humanize v1.0.1
	github.com/go-resty/resty/v2 v2.16.5
	github.com/hashicorp/go-version v1.7.0
	github.com/samber/lo v1.50.0
	github.com/sanity-io/litter v1.5.8
)

require (
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sync v0.14.0
	golang.org/x/text v0.25.0 // indirect
)
