/*
Copyright © 2020 Davi William Moraes Suga <daviciencia1@gmail.com>

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
	"fmt"
	"os"
	"github.com/davisuga/rn-gen/cmd/templates"
	"github.com/spf13/cobra"
)
var log = fmt.Println
// newCmd represents the new command

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "generate a new file",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		log("new called")
		
		targetType, targetName := args[0], args[1:]

		

		if targetType == "component"{
			mapArray(targetName, generateComponent)
		}
		// #TODO: generate screens and service files
		// if targetType == "screen"{

		// }
		// if targetType == "service"{

		// }
	},
	Args:cobra.MinimumNArgs(1),
}

func mapArray(array []string, funcToApply func(string) string) []string {
	newArray := []string{}
	for _, arrayElement := range array {
		newArray = append(newArray, funcToApply(arrayElement))
	}
	return newArray
}

func generateComponent(componentName string) string {
	content := templates.Component(componentName)
	componentPath := "./src/components/"+componentName
	os.MkdirAll(componentPath, os.ModePerm)
	componentFile, _ := os.Create(componentPath+"/index.tsx")
	fmt.Fprint(componentFile, content)
	componentFile.Close()
	return content
}

func generateScreen(screenName, content, style string){
	screenPath := "./src/screens/"+screenName
	os.MkdirAll(screenPath, os.ModePerm)
	screenFile, _ := os.Create(screenPath+"/index.tsx")
	screenStyleFile, _ := os.Create(screenPath+"/styles.ts")

	fmt.Fprint(screenFile, content)
	fmt.Fprint(screenStyleFile, style)
	screenFile.Close()
}

func init() {
	rootCmd.AddCommand(newCmd)
}
