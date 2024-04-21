package models

import (
	"context"
	"errors"
	"smartm2m/src/helper"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `json:"username"`
	Password string             `json:"password"`
}

func (d *User) TableName() string {
	return "users"
}

func (user *User) Create() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	collection := helper.GetDB().Collection(new(User).TableName())

	// Check if username already exists
	var existingUser User
	err = collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		return errors.New("Username already exists")
	}

	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (user *User) Validate() bool {
	collection := helper.GetDB().Collection(new(User).TableName())

	var result User
	err := collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&result)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		return false
	}

	user.ID = result.ID
	return true
}
