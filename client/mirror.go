package client

import "github.com/sndl-team/sndl/lib"

type MirrorClient struct {
	Path string

	Delta   bool
	Verbose bool
}

func (mirrorClient MirrorClient) Run(username string, password string) {
	var client = lib.NewSenturionClient()

	client.Init()
	client.Login(username, password)
}
