package client

import "github.com/sndl-team/sndl/lib"

type PopularClient struct {
	Series bool
	Movies bool

	New     bool
	Verbose bool
}

func (popularClient PopularClient) Run(username string, password string) {
	var client = lib.NewSenturionClient()

	client.Init()
	client.Login(username, password)
}
