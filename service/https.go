package service

import (
	"github.com/goccy/go-json"
	"github.com/minio/minio-go/v7"
	"github.com/pnptcn/chief/investigation"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cache"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/etag"
	"github.com/gofiber/fiber/v3/middleware/idempotency"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

/*
HTTPS wraps the fiber app and the mongo service.
*/
type HTTPS struct {
	app      *fiber.App
	services map[string]Handler
}

/*
NewHTTPS sets up the fiber app and the mongo service.
*/
func NewHTTPS(minioClient *minio.Client, bucket string) *HTTPS {
	return &HTTPS{
		app: fiber.New(fiber.Config{
			CaseSensitive:            true,
			StrictRouting:            true,
			EnableSplittingOnParsers: true,
			ServerHeader:             "Fiber",
			AppName:                  "Integration",
			JSONEncoder:              json.Marshal,
			JSONDecoder:              json.Unmarshal,
		}),
		services: map[string]Handler{
			"investigation": investigation.NewService(minioClient, bucket),
		},
	}
}

/*
Up starts the fiber app, adding the middleware and routes.
*/
func (https *HTTPS) Up() error {
	https.app.Use(
		logger.New(),
		recover.New(),
		cache.New(),
		etag.New(),
		compress.New(),
		idempotency.New(),
		requestid.New(),
		cors.New(),
	)

	https.app.Get("/", func(ctx fiber.Ctx) error {
		ctx.Status(fiber.StatusOK)
		return ctx.SendString("OK")
	})

	https.app.All("/:service/:operation", func(ctx fiber.Ctx) (err error) {
		ctx.Response().Header.Set("Content-Type", "application/json")
		srv := https.services[ctx.Params("service")]

		switch ctx.Params("operation") {
		case "find":
			return srv.Find(ctx)
		case "create":
			return srv.Create(ctx)
		case "update":
			return srv.Update(ctx)
		case "delete":
			return srv.Delete(ctx)
		default:
			return fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	})

	// Preforking allows us to bind multiple Go processes to a single port.
	// This enables significant performance gains, next to the already added
	// benefits of fasthttp.
	return https.app.Listen(":5050", fiber.ListenConfig{
		EnablePrefork: true,
	})
}
