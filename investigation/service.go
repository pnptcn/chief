package investigation

import (
	"github.com/gofiber/fiber/v3"
	"github.com/minio/minio-go/v7"
	"github.com/pnptcn/chief/data"
)

type Service struct {
	artifactService *data.ArtifactService
}

func NewService(minioClient *minio.Client, bucket string) *Service {
	return &Service{
		artifactService: data.NewArtifactService(minioClient, bucket),
	}
}

func (srv *Service) Find(ctx fiber.Ctx) error {
	id := ctx.Query("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing id parameter")
	}

	artifact, err := srv.artifactService.GetArtifact(ctx.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve artifact")
	}

	return ctx.JSON(artifact)
}

func (srv *Service) Create(ctx fiber.Ctx) (err error) {
	artifact := data.New(data.JSON, data.INVESTIGATION)
	if _, err = artifact.Write(ctx.Body()); err != nil {
		return
	}

	return ctx.SendStream(artifact)
}

func (srv *Service) Update(ctx fiber.Ctx) (err error) {
	return
}

func (srv *Service) Delete(ctx fiber.Ctx) (err error) {
	return
}
