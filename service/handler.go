package service

import "github.com/gofiber/fiber/v3"

type Handler interface {
	Find(ctx fiber.Ctx) (err error)
	Create(ctx fiber.Ctx) (err error)
	Update(ctx fiber.Ctx) (err error)
	Delete(ctx fiber.Ctx) (err error)
}
