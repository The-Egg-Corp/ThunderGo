package thundergo

var Experimental = createClient("experimental")
var Legacy = createClient("v1")

func createClient(str string) Client {
	return Client{
		ApiType: str,
	}
}

type Client struct {
	ApiType string
}
