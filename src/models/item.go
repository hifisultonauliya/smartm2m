package models

import (
	"context"
	"errors"
	"log"
	"net/url"
	"smartm2m/src/helper"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemNFT struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID        primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	Name          string             `json:"name" bson:"name"`
	Rating        int                `json:"rating" bson:"rating"`
	Category      string             `json:"category" bson:"category"`
	Image         string             `json:"image" bson:"image"`
	Reputation    int                `json:"reputation" bson:"reputation"`
	ReputationStr string             `json:"reputationstr" bson:"reputationstr"`
	Price         int                `json:"price" bson:"price"`
	Availability  int                `json:"availability" bson:"availability"`
}

func (d *ItemNFT) TableName() string {
	return "itemsNFT"
}

func CreateItem(item ItemNFT) (*ItemNFT, error) {
	if err := item.Validate(); err != nil {
		return nil, err
	}

	_, err := helper.GetDB().Collection(item.TableName()).InsertOne(context.Background(), item)
	if err != nil {
		log.Printf("Error while inserting item: %v\n", err)
		return nil, errors.New("failed to create item")
	}
	return &item, nil
}

func GetItems(userID string) ([]*ItemNFT, error) {
	filters := bson.M{}

	if userID != "" {
		objID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			log.Printf("Invalid userID: %v\n", err)
			return nil, errors.New("invalid user ID")
		}

		filters = bson.M{"userId": objID}
	}

	// Find items that match the filter
	cursor, err := helper.GetDB().Collection(new(ItemNFT).TableName()).Find(context.Background(), filters)
	if err != nil {
		log.Printf("Error while fetching items: %v\n", err)
		return nil, errors.New("failed to fetch items")
	}
	defer cursor.Close(context.Background())

	var items []*ItemNFT
	// Iterate through the cursor and decode each item
	for cursor.Next(context.Background()) {
		var item ItemNFT
		if err := cursor.Decode(&item); err != nil {
			log.Printf("Error while decoding item: %v\n", err)
			return nil, errors.New("failed to decode item")
		}
		items = append(items, &item)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v\n", err)
		return nil, errors.New("cursor error")
	}

	return items, nil
}

func GetItem(id string) (*ItemNFT, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ObjectID: %v\n", err)
		return nil, errors.New("invalid ID")
	}

	var item ItemNFT
	err = helper.GetDB().Collection(new(ItemNFT).TableName()).FindOne(context.Background(), bson.M{"_id": objID}).Decode(&item)
	if err != nil {
		log.Printf("Error while fetching item: %v\n", err)
		return nil, errors.New("item not found")
	}

	return &item, nil
}

func UpdateItem(id string, newItem ItemNFT) (*ItemNFT, error) {
	if err := newItem.Validate(); err != nil {
		return nil, err
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ObjectID: %v\n", err)
		return nil, errors.New("invalid ID")
	}

	update := bson.M{
		"$set": bson.M{
			"userId":        newItem.UserID,
			"name":          newItem.Name,
			"rating":        newItem.Rating,
			"category":      newItem.Category,
			"image":         newItem.Image,
			"reputation":    newItem.Reputation,
			"reputationStr": newItem.ReputationStr,
			"price":         newItem.Price,
			"availability":  newItem.Availability,
		},
	}

	_, err = helper.GetDB().Collection(new(ItemNFT).TableName()).UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		log.Printf("Error while updating item: %v\n", err)
		return nil, errors.New("failed to update item")
	}

	return &newItem, nil
}

func DeleteItem(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ObjectID: %v\n", err)
		return errors.New("invalid ID")
	}

	_, err = helper.GetDB().Collection(new(ItemNFT).TableName()).DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		log.Printf("Error while deleting item: %v\n", err)
		return errors.New("failed to delete item")
	}

	return nil
}

func PurchaseItem(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ObjectID: %v\n", err)
		return errors.New("invalid ID")
	}

	// Get the item
	item, err := GetItem(id)
	if err != nil {
		return err
	}

	// Check availability
	if item.Availability <= 0 {
		return errors.New("no availability for this item")
	}

	// Update availability
	update := bson.M{"$inc": bson.M{"availability": -1}}
	_, err = helper.GetDB().Collection(new(ItemNFT).TableName()).UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		log.Printf("Error while updating item availability: %v\n", err)
		return errors.New("failed to purchase item")
	}

	return nil
}

func (item *ItemNFT) Validate() error {
	// Item name validation
	if len(item.Name) <= 10 {
		return errors.New("item name should be longer than 10 characters")
	}
	for _, word := range []string{"Sex", "Gay", "Lesbian"} {
		if strings.Contains(item.Name, word) {
			return errors.New("item name cannot contain restricted words")
		}
	}

	// Rating validation
	if item.Rating < 0 || item.Rating > 5 {
		return errors.New("rating must be between 0 and 5")
	}

	// Category validation
	validCategories := map[string]bool{
		"photo":     true,
		"sketch":    true,
		"cartoon":   true,
		"animation": true,
	}
	if !validCategories[item.Category] {
		return errors.New("invalid category")
	}

	// Image validation
	_, err := url.ParseRequestURI(item.Image)
	if err != nil {
		return errors.New("image must be a valid URL")
	}

	// Reputation validation
	switch {
	case item.Reputation <= 500:
		item.ReputationStr = "red"
	case item.Reputation <= 799:
		item.ReputationStr = "yellow"
	default:
		item.ReputationStr = "green"
	}

	// Price and availability validation
	if item.Price < 0 {
		return errors.New("price must be a non-negative integer")
	}
	if item.Availability < 0 {
		return errors.New("availability must be a non-negative integer")
	}

	return nil
}
