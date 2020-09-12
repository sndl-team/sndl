package client

import "github.com/sndl-team/sndl/lib"

type SearchClient struct {
	Query   string
	Series  bool
	Movies  bool
	Verbose bool
}

func (searchClient SearchClient) Run(username string, password string) {
	var client = lib.NewSenturionClient()

	// These function calls should likely be chained. This will require some renaming.
	client.Init()
	client.Login(username, password)
}
