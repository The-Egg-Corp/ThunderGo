package v1

import (
	lo "github.com/samber/lo"
	Exp "github.com/the-egg-corp/thundergo/experimental"
	"golang.org/x/sync/errgroup"
)

// The list of every package on Thunderstore across every community.
func GetAllPackages() (PackageList, error) {
	communities, err := Exp.GetCommunities()

	if err != nil {
		return nil, err
	}

	identifiers := lo.Map(communities, func(c Exp.Community, _ int) string {
		return c.Identifier
	})

	return PackagesFromCommunities(NewCommunityList(identifiers...))
}

// Returns a single slice with all packages from the specified communities.
func PackagesFromCommunities(communities []Community) (PackageList, error) {
	amt := len(communities)

	g := errgroup.Group{}
	g.SetLimit(300)

	var list PackageList

	for i := 0; i < amt; i++ {
		i := i
		g.Go(func() error {
			pkgs, err := communities[i].AllPackages()
			if err != nil {
				return err
			}

			list.AddFlat(pkgs)
			return nil
		})
	}

	g.Wait()

	return list, nil
}
