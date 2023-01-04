package cmd

import (
	"fmt"

	"github.com/cvetkovski98/zvax-qrcode/internal/config"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Short: "QR code microservice",
	Long:  `QR code microservice`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("QR code microservice")
	},
}

func init() {
	cobra.OnInitialize(configure)
	root.AddCommand(runCommand)
	root.AddCommand(migrateCommand)
	root.AddCommand(createBucketCommand)
}

func configure() {
	if err := config.LoadConfig("config.dev.yaml"); err != nil {
		panic(err)
	}
}

func Execute() error {
	return root.Execute()
}
