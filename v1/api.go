package v1

import (
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"

	lop "github.com/samber/lo/parallel"
	Exp "github.com/the-egg-corp/thundergo/experimental"
)

// The list of every package on Thunderstore across every community.
func GetAllPackages() (PackageList, error) {
	communities, err := Exp.GetCommunities()

	if err != nil {
		return nil, err
	}

	identifiers := lop.Map(communities, func(c Exp.Community, _ int) string {
		return c.Identifier
	})

	return PackagesFromCommunities(NewCommunityList(identifiers...))
}

// Returns a single slice with all packages from the specified communities.
func PackagesFromCommunities(communities []Community) (PackageList, error) {
	var list PackageList
	var mut sync.Mutex

	g := errgroup.Group{}
	g.SetLimit(300)

	fmt.Println(len(communities))

	for i := 0; i < len(communities); i++ {
		i := i
		g.Go(func() error {
			comm := communities[i]

			pkgs, err := comm.AllPackages()
			//fmt.Println(pkgs.Size())

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
