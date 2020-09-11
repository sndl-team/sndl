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

var new bool

// popularCmd represents the popular command
var popularCmd = &cobra.Command{
	Use:   "popular",
	Short: "A command to display the current popular media",
	Long: `A command to list the current most popular movies and 
series`,
	Run: func(cmd *cobra.Command, args []string) {
		var client = client.PopularClient{Series: series, Movies: movies, New: new}
		fmt.Println(client.New)
	},
}

func init() {
	rootCmd.AddCommand(popularCmd)

	popularCmd.Flags().BoolVarP(&series, "series", "s", false, "Whether to list series")
	popularCmd.Flags().BoolVarP(&movies, "movies", "m", true, "Whether to list movies")

	popularCmd.Flags().BoolVarP(&new, "new", "n", false, "List new episodes, ignores movies flag!")
}
