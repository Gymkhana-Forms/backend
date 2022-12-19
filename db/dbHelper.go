package db

import (
	"context"

	"github.com/Gymkhana-Forms/backend/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FormData struct {
	ID           primitive.ObjectID
	Form_ID      string
	FormType     string
	Submitted_by string
	Data         map[string]interface{}
}

var formCollection *mongo.Collection = db.OpenCollection(db.Client, "forms")

func ExtractFormByID(myform *FormData, id string) error {
	var result bson.M
	err := formCollection.FindOne(context.TODO(), bson.D{{"form_id", id}}).Decode(&result)

	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, myform)
	return err
}

func ExtractAllUserForms(allforms *[]FormData, user_mail string) error {
	cursor, err := formCollection.Find(context.TODO(), bson.D{{"submitted_by", user_mail}})
	if err != nil {
		return err
	} else {
		err = cursor.Decode(allforms)
	}
	return err
}

func ExtractAllForms(allforms []FormData) error {
	return nil
}
