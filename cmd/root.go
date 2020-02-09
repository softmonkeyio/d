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
package cmd

import (
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/softmonkeyio/d/docker"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var log = logrus.New()
var logNoColor bool
var logLevel bool

var rootCmd = &cobra.Command{
	Use: "d",
	Long: `DinG (Do in Go) - called (d) - a shortcutted and 1337ed developer toolbox. 

The project aims to provide each developer 
with a quick and intuitive set of tools 
that will optimize everyday repetitive work.

Optimization through standardization. Let's make CLI tools fun again!`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&logNoColor, "no-color", false, "Disable ANSI color output")
	rootCmd.PersistentFlags().BoolVarP(&logLevel, "debug", "v", false, "Increase verbosity of logger to debug")
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(docker.DockerCmd)
}

func initLogger() {
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors:          logNoColor,
		ForceColors:            true,
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
	})

	if logLevel {
		log.SetLevel(logrus.DebugLevel)
	}
}

func initConfig() {
	initLogger()
	log.Debug("Initialize config.")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			panic(err)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".d")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Debug("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Debug("Error during loading config file.", err)
	}
}
