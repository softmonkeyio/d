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
	"context"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var (
	contextBackground = context.Background()
	dockerCli         client.APIClient
)

var DockerCmd = &cobra.Command{
	Use:   "d",
	Short: "DinG -> (d)ocker CLI the fun way ",
	Long: `Shortcutted Docker CLI commands.

Wrapper for Docker CLI with compatibility to original Docker CLI
with some syntactic sugar for ease of use.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func init() {
	DockerCmd.AddCommand(PsCmd)
	initDockerSdk()
}

func initDockerSdk() {
	var err error
	dockerCli, err = client.NewEnvClient()
	if err != nil {
		panic(err)
	}
}
