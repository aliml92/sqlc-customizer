/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/aliml92/sqlc-customizer/pkg/util"
	"github.com/spf13/cobra"
)

// modifyCmd represents the modify command
var modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("modify called")
		// 1. modify models
		for _, model := range cust.Config.Modify.Models {
			name := model.Name
			source := model.Source
			destination := model.Destination
			packageName := model.Package
			oldPackageName := model.OldPackage
			packagePath := model.PackagePath
			jsonOmitempty := model.JSONOmitempty

			// do better config validation here before proceeding in the future

			// 1.1. check if source file exists
			if !util.FileExists(source) {
				fmt.Printf("source file %s does not exist\n", source)
				continue
			}
			// 1.2. check if destination file exists
			if util.FileExists(destination) {
				// delete it
				if err := os.Remove(destination); err != nil {
					fmt.Printf("failed to remove file %s: %v\n", destination, err)
					continue
				}
			}

			// 1.3. move source file to destination file
			if err := util.MoveFile(source, destination); err != nil {
				fmt.Printf("failed to move file %s to %s: %v\n", source, destination, err)
				continue
			}

			// 1.4. modify package name in destination file
			if err := util.OverridePackageName(destination, packageName, oldPackageName); err != nil {
				fmt.Printf("failed to override package name in %s: %v\n", destination, err)
				continue
			}
			if jsonOmitempty {
				_, err := util.AppendOmitempty(destination)
				if err != nil {
					fmt.Printf("failed to append omitempty in %s: %v\n", destination, err)
				}	
			}

			directory := path.Dir(source)
			files, err := os.ReadDir(directory)
			if err != nil {
				fmt.Printf("failed to read directory %s: %v\n", directory, err)
				continue
			}
			for _, file := range files {
				file := path.Join(directory, file.Name())
				newWord := packageName + "." + name
				n, err := util.OverrideStuct(file, name, newWord)
				if err != nil {
					fmt.Printf("failed to override struct name %s in %s: %v\n", name, file, err)
					continue
				}
				if n == 0 {
					continue
				}
				
				if err := util.AppendImportEntry(file, packagePath); err != nil {
					fmt.Printf("failed to append import entry %s in %s: %v\n", packagePath, file, err)
					continue
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(modifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
