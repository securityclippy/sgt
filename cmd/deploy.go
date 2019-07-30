// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/oktasecuritylabs/sgt/handlers/deploy"
	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy named components",
	Long: `components specifies which components to deploy.  It will automatically take care of ordering
them properly so that any dependent changes will be deployed in the correct order`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(DeployComponents) < 1 {
			log.Fatal(cmd.Help())
		}
		err := deploy.ValidateComponets(DeployComponents, UseElastic)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("deploying: %+v\n", DeployComponents)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.Flags().StringSliceVarP(&DeployComponents, "components", "", DeployComponents, "")
	err :=	deployCmd.MarkFlagRequired("components")
	if err != nil {
		log.Fatal(err)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
