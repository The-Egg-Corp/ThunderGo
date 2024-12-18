package v1

import (
	"errors"
	"sync"

	"golang.org/x/sync/errgroup"
)

func PackagesFromCommunity(identifier string) (PackageList, error) {
	comm := Community{
		Identifier: identifier,
	}

	return comm.AllPackages()
}

func PackageFromCommunity(identifier string, owner string, packageName string) *Package {
	comm := Community{
		Identifier: identifier,
	}

	return comm.GetPackage(owner, packageName)
}

// Returns a single slice with all packages from the specified communities.
func PackagesFromCommunities(communities []Community) (PackageList, error) {
	amt := len(communities)
	if amt == 0 {
		return nil, errors.New("empty input list")
	}

	g := errgroup.Group{}
	g.SetLimit(100)

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
