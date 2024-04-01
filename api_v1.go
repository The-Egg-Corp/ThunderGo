package thundergo

import (
	"fmt"
	"thundergo/util"
)

func (client V1Client) GetPackage(author string, name string) (*Package, error) {
	endpoint := fmt.Sprint("/c/", , author, "/", name)
	return util.JsonRequest[*Package](endpoint)
}
