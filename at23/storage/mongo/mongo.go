package mongo

import (
	"at23/messages"
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Proizvod struct {
	Identifikator string
	Kolicina      int
}

func DodajArtikleProto(artikli []*messages.Proizvod) {
	for _, a := range artikli {
		var id string
		id = a.Identifikator
		var kol int
		kol = int(a.Kolcicina)
		DodajAritkal(Proizvod{Identifikator: id, Kolicina: kol})
	}
}

// dodati kanale?
func DodajArtikle(artikli []Proizvod) {
	for _, a := range artikli {

		DodajAritkal(a)
	}
}

// https://christiangiacomi.com/posts/mongo-go-driver-update/
func DodajAritkal(artikal Proizvod) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("golang_master").Collection("Proizvodi")
	filter := bson.M{"identifikator": artikal.Identifikator}
	count, err := col.CountDocuments(context.TODO(), filter)
	if count < 1 {
		oneDoc := Proizvod{Identifikator: artikal.Identifikator, Kolicina: artikal.Kolicina}
		_, insertErr := col.InsertOne(context.TODO(), oneDoc)
		if insertErr != nil {
			os.Exit(1)
		}
		// } else {
		// 	newId := result.InsertedID
		// 	fmt.Println("InsertedId", newId)
		// }
	} else {
		update := bson.M{"$inc": bson.M{"kolicina": artikal.Kolicina}}
		col.FindOneAndUpdate(ctx, filter, update)

	}
}
