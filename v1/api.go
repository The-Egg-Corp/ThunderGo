package v1

import (
	"errors"
	"sync"

	"golang.org/x/sync/errgroup"
)

// // The list of every package on Thunderstore across every community.
// func GetAllPackages() (PackageList, error) {
// 	communities, err := Exp.GetCommunities()

// 	if err != nil {
// 		return nil, err
// 	}

// 	identifiers := lop.Map(communities, func(c Exp.Community, _ int) string {
// 		return c.Identifier
// 	})

// 	return PackagesFromCommunities(NewCommunityList(identifiers...))
// }

// Returns a single slice with all packages from the specified communities.
func PackagesFromCommunities(communities []Community) (PackageList, error) {
	amt := len(communities)
	if amt == 0 {
		return nil, errors.New("empty input list")
	}

	g := errgroup.Group{}
	g.SetLimit(50)

	var list PackageList
	var mut sync.Mutex

	for i := 0; i < amt; i++ {
		i := i
		g.Go(func() error {
			comm := communities[i]
			pkgs, err := comm.AllPackages()

			if err != nil {
				return err
			}

			if pkgs.Size() > 0 {
				mut.Lock()
				list.AddFlat(pkgs)
				mut.Unlock()
			}

			return nil
		})
	}

	err := g.Wait()
	return list, err
}
