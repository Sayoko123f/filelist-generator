package src

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetFileList(cmd *cobra.Command) map[string][]string {
	fmt.Println(viper.GetString("root"))
	fmt.Println(viper.GetStringSlice("ignore"))
	fmt.Println(viper.GetStringSlice("pattern"))
	fmt.Println(cmd.Flags().GetString("root"))

	fileSystem, _ := cmd.Flags().GetString("root")
	rootDir, _ := filepath.Abs(fileSystem)
	fmt.Println(rootDir)
	collect := make(map[string][]string)
	fs.WalkDir(os.DirFS(fileSystem), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !d.IsDir() {
			nowpath, _ := filepath.Abs(filepath.Join(fileSystem, path))
			rel, _ := filepath.Rel(rootDir, nowpath)
			dirname := filepath.Dir(rel)
			collect[dirname] = append(collect[dirname], filepath.Base(nowpath))
		}
		return nil
	})
	return collect
}
