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
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kkirsche/hn/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// askCmd represents the ask command
var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			argsInt int
			err     error
		)

		logger = hnapi.NewLogger(viper.GetBool("verbose"))

		logger.VerbosePrintln("Checking if arguments were passed to `top`.")
		argsLen := len(args)
		if argsLen > 0 {
			logger.VerbosePrintfln("%d arguments were detected.", argsLen)
			logger.VerbosePrintln("Flattening arguments.")
			argsMerged := strings.Join(args, "")
			logger.VerbosePrintln("Converting flattened arguments into an integer.")
			argsInt, err = strconv.Atoi(argsMerged)
			if err != nil {
				log.Panicln(err.Error())
			}
		}

		if argsInt == 0 {
			argsInt = 25
		}

		storiesCh, err := hnapi.StreamNAskHNStories(argsInt, logger)
		if err != nil {
			log.Panicln(err)
		}

		i := 1
		for story := range storiesCh {
			fmt.Println(i, story.Title)
			i++
		}
	},
}

func init() {
	RootCmd.AddCommand(askCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// askCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// askCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
