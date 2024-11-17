package routes

import (
	"backend-atlas/config"
	"backend-atlas/models"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func ProvinceRoutes(app *fiber.App) {
	collection := config.DB.Collection("provinces")

	app.Get("/provinces", func(c *fiber.Ctx) error {
		var provinces []models.Province

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var province models.Province
			if err := cursor.Decode(&province); err != nil {
				return c.Status(500).SendString(err.Error())
			}
			provinces = append(provinces, province)
		}

		return c.JSON(provinces)
	})

	app.Post("/provinces", func(c *fiber.Ctx) error {
		var province models.Province
		if err := c.BodyParser(&province); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, err := collection.InsertOne(ctx, province)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(province)
	})
}
