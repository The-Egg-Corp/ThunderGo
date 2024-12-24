package v1

import (
	"errors"
	"fmt"
	"sync"

	"github.com/samber/lo"
	"github.com/the-egg-corp/thundergo/util"
	"golang.org/x/sync/errgroup"
)

type RateBody struct {
	TargetState RatingState `json:"target_state"`
}

// Tries to rate/unrate (set by targetState) a package given its unique identifier.
//
// Note that the listing could be nil given the following conditions:
//
// # - Bad or no status code
//
// # - Error unmarshalling response into listing
//
// # - Received empty/zero value as listing
func RatePackage(uuid string, targetState RatingState) (*PackageListing, error) {
	endpoint := fmt.Sprintf("api/v1/package/%s/rate/", uuid)
	reqBody := RateBody{TargetState: targetState}

	res, code, err := util.JsonPostRequest[PackageListing](endpoint, reqBody)
	if err != nil {
		return nil, err
	}

	if code == nil {
		err = fmt.Errorf("\nreceived no status code")
	} else {
		if *code >= 400 {
			err = fmt.Errorf("\nreceived bad status code: %v", *code)
		}
	}

	return lo.Ternary(util.Zero(res), nil, res), err
}

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
