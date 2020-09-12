package client

import "github.com/sndl-team/sndl/lib"

type DownloadClient struct {
	Query      string
	Resolution string

	Server int

	Movies  bool
	Series  bool
	Verbose bool
}

func (downloadClient DownloadClient) Run(username string, password string) {
	var client = lib.NewSenturionClient()

	client.Init()
	client.Login(username, password)
}
