package mongo

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// POLJA STRUKTURE KOJA SE UPIUSUJE U BAZU MORAJU BITI VELIKA, VEROVATNO SLICAN RAZLOG KAO ZA FUNKCIJU
type StorageList struct {
	AdresaSkladista string
}

// DA BI SE IMPORTOVALAU U MAIN, NAZIV FUNKCIJE MORA POCINJATI VELIKIM SLOVOM
func DodajUStorage(adresa string) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		os.Exit(1)
	}
	//ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("golang_master").Collection("Storage_Coll")

	//check if exists
	count, err := col.CountDocuments(context.TODO(), bson.M{"adresaskladista": adresa})
	if count < 1 {
		oneDoc := StorageList{
			AdresaSkladista: adresa,
		}
		result, insertErr := col.InsertOne(context.TODO(), oneDoc)
		if insertErr != nil {
			os.Exit(1)
		} else {
			newId := result.InsertedID
			fmt.Println("InsertedId", newId)
		}
	}

}
