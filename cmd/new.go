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
var format = fmt.Sprintf
var COMPONENTS_PATH = "./src/components/"
var SCREENS_PATH = "./src/screens/"

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "generate a new file",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {		
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

func generateScreen(screenName string) string {
	content := templates.Screen(screenName)
	dumpToFile(SCREENS_PATH, "index.tsx",  content)

	log("generated "+screenName)
	return content
}

func generateComponent(componentName string) string {
	
	component := templates.Component(componentName)
	styles := templates.Styled("")
	dumpToFile(COMPONENTS_PATH+componentName, "/index.tsx", component)
	dumpToFile(COMPONENTS_PATH+componentName, "/style.ts", styles)

	log("generated "+ componentName)
	return component
}
// filePath example: ./src/components/CompA/index.tsx
func dumpToFile(path, fileName, content string)  {
	
	os.MkdirAll(path, os.ModePerm)
	file, _ := os.Create(path+fileName)
	fmt.Fprint(file, content)
	file.Close()
}

func init() {
	rootCmd.AddCommand(newCmd)
}
