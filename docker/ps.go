/*
Copyright © 2020 Maciej "Cichy" Świderski <mmswiderski@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package docker

import (
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/softmonkeyio/d/helpers"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var PsCmd = &cobra.Command{
	Use:   "ps [grep-like-filter-string]",
	Short: "List containers - fun way " + helpers.GlassesOfDisapproval(),
	Long:  `[=== TODO ===]`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stype := types.ContainerListOptions{}
		stype.All = true
		containers, err := dockerCli.ContainerList(contextBackground, stype)
		if err != nil {
			panic(err)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "NAME", "IMAGE", "STATE"})
		table.SetBorder(false)
		host, _ := os.Hostname()
		table.SetCaption(true, "Running containers on: " + host)
		b, _ := json.Marshal(containers[0])
		logrus.Info(string(b))
		for _, container := range containers {
			table.Append([]string{container.ID[0:6], strings.TrimLeft(strings.Join(container.Names, "\n"), "/"), container.Image, container.State})
		}
		table.Render()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// psCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// psCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
