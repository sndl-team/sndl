package client

import "github.com/sndl-team/sndl/lib"

type FavoritesClient struct {
	Mirror  bool
	Verbose bool
}

func (favoritesClient FavoritesClient) Run(username string, password string) {
	var client = lib.NewSenturionClient()

	client.Init()
	client.Login(username, password)
}
