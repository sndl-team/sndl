package client

import (
	"github.com/rs/zerolog/log"
)

type SearchClient struct {
	Query   string
	Series  bool
	Movies  bool
	Verbose bool
}

// Sample method to test logging mechanism
func (searchClient SearchClient) Search() {
	if searchClient.Verbose {
		log.Print("hello world")
	}
}
