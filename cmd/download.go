/*
The MIT License (MIT)

Copyright © 2020 sndl-team jayrad.security@protonmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/sndl-team/sndl/client"

	"github.com/spf13/cobra"
)

var resolution string

var server int

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A Command to download media",
	Long:  `A Command that given a query will allow the user to download media`,
	Run: func(cmd *cobra.Command, args []string) {
		var client = client.DownloadClient{
			Query:      query,
			Resolution: resolution,
			Server:     server,
			Movies:     movies,
			Series:     series}

		fmt.Println(client.Query)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringVarP(&query, "query", "q", "", "The search term")
	downloadCmd.MarkFlagRequired("query")

	downloadCmd.Flags().StringVarP(&resolution, "resolution", "r", "720p", "The resolution to download media at (720p or 1080p")

	// 0 should default to suggested...
	downloadCmd.Flags().IntVarP(&server, "server", "x", 0, "The server to use by ID? || Name?")

	downloadCmd.Flags().BoolVarP(&movies, "movies", "m", false, "Whether to search movies")
	downloadCmd.Flags().BoolVarP(&series, "series", "s", false, "Whether to search series")
}
