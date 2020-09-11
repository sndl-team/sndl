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

var mirror bool

// favoritesCmd represents the favorites command
var favoritesCmd = &cobra.Command{
	Use:   "favorites",
	Short: "A comment to get the media saved in your favorites",
	Long: `A command to download all the media in your favorites
and save it to the local filesystem`,
	Run: func(cmd *cobra.Command, args []string) {
		var client = client.FavoritesClient{Mirror: mirror, Verbose: verbose}
		fmt.Println(client.Mirror)
	},
}

func init() {
	rootCmd.AddCommand(favoritesCmd)

	favoritesCmd.Flags().BoolVarP(&mirror, "mirror", "m", false, "Whether to download the media or not")
}
