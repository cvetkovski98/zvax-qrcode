package cmd

import (
	"log"
	"net"

	"github.com/cvetkovski98/zvax-common/gen/pbqr"
	"github.com/cvetkovski98/zvax-common/pkg/postgresql"
	"github.com/cvetkovski98/zvax-qrcode/internal/config"
	"github.com/cvetkovski98/zvax-qrcode/internal/delivery"
	"github.com/cvetkovski98/zvax-qrcode/internal/model/migrations"
	"github.com/cvetkovski98/zvax-qrcode/internal/repository"
	"github.com/cvetkovski98/zvax-qrcode/internal/service"
	"github.com/cvetkovski98/zvax-qrcode/pkg/minio"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	runCommand = &cobra.Command{
		Use:   "run",
		Short: "Run QR code microservice",
		Long:  `Run QR code microservice`,
		Run:   run,
	}
	network string
	address string
)

func init() {
	runCommand.Flags().StringVarP(&network, "network", "n", "tcp", "network to listen on")
	runCommand.Flags().StringVarP(&address, "address", "a", ":50052", "address to listen on")
}

func run(cmd *cobra.Command, args []string) {
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s://%s...", network, address)
	cfg := config.GetConfig()
	db, err := postgresql.NewPgDb(&cfg.PostgreSQL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()
	minIOClient, err := minio.NewMinioClient(&cfg.MinIO)
	if err != nil {
		log.Fatalf("failed to connect to minio: %v", err)
	}
	if err := postgresql.Migrate(cmd.Context(), db, migrations.Migrations); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	qrObjStore := repository.NewMinioObjectStore(minIOClient)
	if err := qrObjStore.CreateBucket(cmd.Context(), cfg.MinIO.BucketName); err != nil {
		log.Fatalf("failed to create a bucket: %v", err)
	}
	qrRepository := repository.NewPgQRCodeRepository(db)
	qrService := service.NewQRCodeService(qrRepository, qrObjStore)
	qrGrpc := delivery.NewQRCodeServer(qrService)
	server := grpc.NewServer()
	pbqr.RegisterQRCodeServer(server, qrGrpc)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
