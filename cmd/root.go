package cmd

import (
	"fmt"

	"github.com/cvetkovski98/zvax-qrcode/internal/config"
	"github.com/spf13/cobra"
)

var cfgFile string
var root = &cobra.Command{
	Short: "QR code microservice",
	Long:  `QR code microservice`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("QR code microservice")
	},
}

func init() {
	cobra.OnInitialize(configure)
	root.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.dev.yaml", "config file name")
	root.AddCommand(runCommand)
	root.AddCommand(migrateCommand)
	root.AddCommand(createBucketCommand)
}

func configure() {
	if err := config.LoadConfig(cfgFile); err != nil {
		panic(err)
	}
}

func Execute() error {
	return root.Execute()
}
