package src

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetFileList(cmd *cobra.Command) map[string][]string {
	fmt.Println(viper.GetStringSlice("ignore"))
	fmt.Println(viper.GetStringSlice("pattern"))
	fmt.Println(cmd.Flags().GetString("root"))

	fileSystem, _ := cmd.Flags().GetString("root")
	rootDir, _ := filepath.Abs(fileSystem)
	collect := make(map[string][]string)
	fs.WalkDir(os.DirFS(fileSystem), ".", func(localpath string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		// ignore
		for _, ignorePattern := range viper.GetStringSlice("ignore") {
			isIgnore, err := path.Match(ignorePattern, localpath)
			if err != nil {
				log.Fatal(err)
			}
			if isIgnore {
				fmt.Println("ignore", ignorePattern, localpath)
				if d.IsDir() {
					return filepath.SkipDir
				} else {
					return nil
				}
			}
		}

		if !d.IsDir() {
			nowpath, _ := filepath.Abs(filepath.Join(fileSystem, localpath))
			rel, _ := filepath.Rel(rootDir, nowpath)
			dirname := filepath.Dir(rel)
			collect[dirname] = append(collect[dirname], filepath.Base(nowpath))
		}
		return nil
	})
	return collect
}

type Row struct {
	Index    int
	Filename string
	Desc     string
}

type DescriptionPatternItem struct {
	match string
	desc  string
}

func getFileListKeys(collect map[string][]string) []string {
	keys := make([]string, 0, len(collect))
	for key := range collect {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}

// https://stackoverflow.com/questions/35583735/unmarshaling-into-an-interface-and-then-performing-type-assertion
func TransformFileList(cmd *cobra.Command, collect map[string][]string) map[string][]Row {
	keys := getFileListKeys(collect)
	fmt.Println(keys)

	fileIndex := 0
	dataset := make(map[string][]Row)

	fmt.Println(viper.Get("descriptions"))
	Descriptions, ok := (viper.Get("descriptions")).([]map[string]string)
	fmt.Println(Descriptions)
	if !ok {
		log.Fatal("description invalid type")
	}

	for _, key := range keys {
		var rows []Row
		for _, filename := range collect[key] {
			fileIndex++
			var row Row
			row.Filename = filename
			row.Index = fileIndex
			row.Desc = ""
			// for _, obj := range Descriptions {
			// 	isMatch, err := path.Match(obj.match, filename)
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}
			// 	if isMatch {
			// 		row.Desc = obj.desc
			// 		break
			// 	}
			// }
			rows = append(rows, row)
		}
		dataset[key] = rows
	}

	return dataset
}
