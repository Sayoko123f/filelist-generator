package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	scanRootDir string
	rootCmd     = &cobra.Command{
		Use:   "filelist-generator",
		Short: "A tool.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(viper.GetString("root"))
			fmt.Println(viper.GetString("outputFilename"))
			fmt.Println(viper.GetStringSlice("ignore"))
			fmt.Println(viper.GetStringSlice("pattern"))
			fmt.Println(cmd.Flags().GetString("root"))

			// collect := src.GetFileList(cmd)
			// for k, v := range collect {
			// 	fmt.Println(k, "value is ", v)
			// }
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	rootCmd.PersistentFlags().StringVarP(&scanRootDir, "root", "r", "", "the folder path to scan")
	rootCmd.MarkPersistentFlagRequired("root")
}

func initConfig() {
	defaultConfig := []byte(`{
		"outputFilename": "filelist.json",
		"ignore": [
			"default"
		],
		"pattern": [
			"default pattern"
		]
	}`)
	viper.SetConfigType("json")
	err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		log.Fatal(err)
	}
	viper.AutomaticEnv()
	if cfgFile != "" {
		fmt.Println("從 --config 獲取設定檔案路徑")
		viper.SetConfigType("json")
		viper.SetConfigFile(cfgFile)
	} else {
		fmt.Println("在當前工作目錄搜尋 filelist-generator.json")
		viper.AddConfigPath(".")
		viper.SetConfigType("json")
		viper.SetConfigName("filelist-generator")
	}

	viper.MergeInConfig()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
