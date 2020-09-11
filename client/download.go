package client

type DownloadClient struct {
	Query      string
	Resolution string

	Server int

	Movies bool
	Series bool
}
