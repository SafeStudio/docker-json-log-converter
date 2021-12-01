/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	dockerLogConverter "github.com/SafeStudio/djlc/lc"
	"github.com/spf13/cobra"
)

// djlcCmd represents the djlc command
var djlcCmd = &cobra.Command{
	Use:   "djlc",
	Short: "Docker JSON Log Converter",
	Long:  `Docker JSON Log Converter CLI Tool`,
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		outputFile, _ := cmd.Flags().GetString("output")

		dockerLogConverter.FromFile(inputFile).ToFile(outputFile).Convert()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(djlcCmd.Execute())
}

func init() {
	djlcCmd.PersistentFlags().StringP("output", "o", "log.txt", "Output log file")
}
