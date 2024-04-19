package v1

import (
	"sync"

	Exp "github.com/the-egg-corp/thundergo/experimental"
)

// The list of every package on Thunderstore across every community.
func GetAllPackages() (PackageList, error) {
	communities, err := Exp.GetCommunities()

	if err != nil {
		return nil, err
	}

	len := communities.Size()

	var wg sync.WaitGroup
	wg.Add(len)

	var all PackageList
	worker := func(identifier string) {
		// get package
		comm := Community{
			Identifier: identifier,
		}

		pkgs, _ := comm.AllPackages()
		all = append(all, pkgs...)

		defer wg.Done()
	}

	for i := 0; i < len; i++ {
		go worker(communities[i].Identifier)
	}

	wg.Wait()

	return all, nil
}
