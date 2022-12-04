package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"filelist-generator/src"
)

var (
	cfgFile     string
	scanRootDir string
	rootCmd     = &cobra.Command{
		Use:   "filelist-generator",
		Short: "A tool.",
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Println(viper.GetString("root"))
			// fmt.Println(viper.GetStringSlice("ignore"))
			// fmt.Println(viper.GetStringSlice("pattern"))
			// fmt.Println(cmd.Flags().GetString("root"))

			collect := src.GetFileList(cmd)
			for k, v := range collect {
				fmt.Println(k, "value is ", v)
			}
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
		"ignore": [],
		"pattern": []
	}`)
	viper.ReadConfig(bytes.NewBuffer(defaultConfig))

	if cfgFile != "" {
		fmt.Println("從 --config 獲取設定檔案路徑")
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".cobra" (without extension).
		fmt.Println("在當前工作目錄底下尋找 filelist-generator")
		viper.AddConfigPath(".")
		viper.SetConfigType("json")
		viper.SetConfigName("filelist-generator")
	}

	viper.AutomaticEnv()

	if err := viper.MergeInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Fprintln(os.Stderr, err)
		fmt.Println("找不到使用者自訂設定檔，使用預設設定")
	}
	fmt.Println("將預設設定與使用者設定 merge")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
