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

var path string
var delta bool

// mirrorCmd represents the mirror command
var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "A command to mirror from a sndl.txt file",
	Long: `A command that given an sndl.txt will attempt
to mirror the listings to the local filesystem`,
	Run: func(cmd *cobra.Command, args []string) {
		var client = client.MirrorClient{Path: path, Delta: delta, Verbose: verbose}
		fmt.Println(client.Path)
	},
}

func init() {
	rootCmd.AddCommand(mirrorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mirrorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mirrorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	mirrorCmd.Flags().StringVarP(&path, "path", "p", ".", "The path to save files to")
	mirrorCmd.MarkFlagRequired("path")

	// This only aids shell completion
	// This doesnt enforce that the path exists!
	mirrorCmd.MarkFlagDirname("path")

	mirrorCmd.Flags().BoolVarP(&delta, "delta", "d", false, "Show what has updated without dowloading")
}
