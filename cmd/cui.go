// Copyright © 2016 Kevin Kirsche <kev.kirsche@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"

	"github.com/kkirsche/hn/cui"
	"github.com/spf13/cobra"
)

// cuiCmd represents the cui command
var cuiCmd = &cobra.Command{
	Use:   "cui",
	Short: "View Hacker News with the command line user interface",
	Long: `The cui command allows you to start a command line user interface in
which you may interact with Hacker News.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := hncui.Launch()
		if err != nil {
			log.Panicln(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(cuiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cuiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cuiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
