package src

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetFileList(cmd *cobra.Command) {
	fmt.Println(viper.GetString("root"))
	fmt.Println(viper.GetStringSlice("ignore"))
	fmt.Println(viper.GetStringSlice("pattern"))
	fmt.Println(cmd.Flags().GetString("root"))

	fileSystem, _ := cmd.Flags().GetString("root")

	fs.WalkDir(os.DirFS(fileSystem), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		return nil
	})
}

func deepScan() {}
