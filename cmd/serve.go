package cmd

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pnptcn/chief/service"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var minioClient *minio.Client

		if minioClient, err = minio.New("your-minio-endpoint", &minio.Options{
			Creds:  credentials.NewStaticV4("your-access-key", "your-secret-key", ""),
			Secure: true,
		}); err != nil {
			log.Fatalf("Error creating MinIO client: %v", err)
		}

		return service.NewHTTPS(minioClient, "panopticon").Up()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
