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
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jaytaylor/html2text"
	"github.com/kkirsche/hn/api"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// storyCmd represents the story command
var storyCmd = &cobra.Command{
	Use:   "story",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			argsMerged string
			argsInt    int
			err        error
		)

		logger = hnapi.NewLogger(viper.GetBool("verbose"))

		logger.VerbosePrintln("Checking if arguments were passed to `top story`.")
		argsLen := len(args)
		if argsLen > 0 {
			logger.VerbosePrintfln("%d arguments were detected.", argsLen)
			logger.VerbosePrintln("Flattening arguments.")
			argsMerged = strings.Join(args, "")
			logger.VerbosePrintln("Converting flattened arguments into an integer.")
			argsInt, err = strconv.Atoi(argsMerged)
			if err != nil {
				if err.Error()[:25] == `strconv.ParseInt: parsing` {
					logger.Printfln(`Please provide "hn top story" with a number. For example: "hn top story 25" to
print the 25th top Hacker News story. After sanitization, you provided: %s`, argsMerged)
					return
				}

				log.Panicln(err)
			}

			if argsInt > 0 {
				argsInt--
			} else {
				logger.VerbosePrintln("Story number was 0 or below. If it was less than 0, it has been set to 0.")
				argsInt = 0
			}
		}

		stories := hnapi.TopStoriesItemNumbers()
		story := hnapi.GetItem(stories[argsInt])

		fmt.Println(story.Title)
		logger.Printfln("By: %s, When: %s", story.By, time.Unix(story.Time, 0).String())

		if story.URL != "" {
			fmt.Println()
			logger.Printfln("From URL: %s", story.URL)
		}

		if story.Text != "" {
			fmt.Println()
			story.Text, err = html2text.FromString(story.Text)
			if err != nil {
				log.Panicln(err)
			}
			fmt.Println(story.Text)
		} else {
			fmt.Println()
			fmt.Print("Post does not contain any text. Would you like to open the URL in your web browser? y/n: ")
			reader := bufio.NewReader(os.Stdin)
			resp, err := reader.ReadString('\n')
			if err != nil {
				log.Panicln(err)
			}

			openInBrowser := false
			validResp := false
			resp = strings.ToLower(strings.TrimSpace(string(resp)))
			for validResp != true {
				switch resp {
				case "y":
					validResp = true
					openInBrowser = true
				case "n":
					validResp = true
				default:
					fmt.Println("Invalid response. Please enter y for yes or n for no.")
					fmt.Print("Would you like to open the URL in your web browser? y/n: ")
					reader := bufio.NewReader(os.Stdin)
					resp, err = reader.ReadString('\n')
					if err != nil {
						log.Panicln(err)
					}
					resp = strings.ToLower(strings.TrimSpace(string(resp)))
				}
			}

			if openInBrowser {
				logger.VerbosePrintfln("Opening %s in default web browser.", story.URL)
				open.Run(story.URL)
			}
		}
	},
}

func init() {
	topCmd.AddCommand(storyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
