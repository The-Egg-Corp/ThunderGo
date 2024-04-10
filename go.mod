module github.com/the-egg-corp/thundergo

replace github.com/the-egg-corp/thundergo => ./

go 1.22.1

require (
	github.com/samber/lo v1.39.0
	github.com/sanity-io/litter v1.5.5
)

require golang.org/x/net v0.22.0 // indirect

require (
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/go-resty/resty/v2 v2.12.0
	golang.org/x/exp v0.0.0-20220303212507-bbda1eaf7a17 // indirect
)
