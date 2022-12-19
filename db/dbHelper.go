package db

import (
	"context"
	"time"

	"github.com/Gymkhana-Forms/backend/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FormData struct {
	ID       primitive.ObjectID
	Form_ID  string
	FormType string
	Data     map[string]interface{}
}

var formCollection *mongo.Collection = db.OpenCollection(db.Client, "forms")

func ExtractFormByID(myform *FormData, id string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err := formCollection.FindOne(ctx, bson.M{"form_id": id}).Decode(myform)
	return err
}

func ExtractAllForms(allforms *[]FormData) error {
	return nil
}
