package controllers

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tsivxrev/cashdesk/database"
	"github.com/tsivxrev/cashdesk/models"
	"go.mongodb.org/mongo-driver/bson"
)

var entriesCollection = database.GetCollection(database.Client, "entries")
var validate = validator.New()

func GetAllEntries(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var entries []models.Entry
	results, err := entriesCollection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleEntry models.Entry
		err = results.Decode(&singleEntry)
		if err != nil {
			return err
		}

		entries = append(entries, singleEntry)
	}

	if len(entries) == 0 {
		return c.JSON([]models.Entry{})
	}

	return c.JSON(entries)
}

func GetEntry(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	entryId := c.Params("entryId")
	defer cancel()

	var entry models.Entry
	err := entriesCollection.FindOne(ctx, bson.M{"id": entryId}).Decode(&entry)
	if err != nil {
		return c.Next()
	}

	return c.JSON(&entry)
}

func CreateEntry(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var entry models.Entry
	defer cancel()

	err := c.BodyParser(&entry)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Code:   400,
			Detail: err.Error(),
		})
	}

	err = validate.Struct(&entry)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Code:   400,
			Detail: err.Error(),
		})
	}

	newEntry := models.Entry{
		Id:          uuid.New().String(),
		Name:        entry.Name,
		Description: entry.Description,
		Image:       entry.Image,
		Price:       entry.Price,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	_, err = entriesCollection.InsertOne(ctx, newEntry)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(&newEntry)
}

func EditEntry(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	entryId := c.Params("entryId")
	var entry models.Entry
	defer cancel()

	err := c.BodyParser(&entry)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Code:   400,
			Detail: err.Error(),
		})
	}

	err = validate.Struct(&entry)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Code:   400,
			Detail: err.Error(),
		})
	}

	update := bson.M{
		"name":        entry.Name,
		"description": entry.Description,
		"image":       entry.Image,
		"hidden":      entry.Hidden,
		"price":       entry.Price,
		"updated_at":  time.Now().Unix(),
	}

	result, err := entriesCollection.UpdateOne(ctx, bson.M{"id": entryId}, bson.M{"$set": update})
	if err != nil {
		return err
	}

	var updatedEntry models.Entry
	if result.MatchedCount == 1 {
		err := entriesCollection.FindOne(ctx, bson.M{"id": entryId}).Decode(&updatedEntry)
		if err != nil {
			return c.Next()
		}
	} else {
		return c.Next()
	}

	return c.JSON(updatedEntry)
}
