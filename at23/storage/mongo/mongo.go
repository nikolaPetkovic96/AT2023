package mongo

import (
	"at23/messages"
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Proizvod struct {
	Identifikator string `bson:"identifikator"`
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

// OBRADITI SLUCAJ ZA COUNT !=1, tj za count>1 baza je nekonzistenta
func ProveriKolicine(artikal string) (kol int) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("golang_master").Collection("Proizvodi")
	filter := bson.M{"identifikator": artikal}
	var pr Proizvod
	count, err := col.CountDocuments(context.TODO(), filter)
	if count == 1 {
		col.FindOne(ctx, filter).Decode(&pr)
		fmt.Println("PRONADJENA KOL:", pr.Kolicina)
		return pr.Kolicina
	} else {
		return 0
	}
}

func KupiArtikal(artikal *messages.Item) (identifikator string, uspesnoKupljen bool) {
	//zeljenaKolicina := artikal.Amount
	//trazeniArtikal := artikal.ItemId
	//proveri da li ima potreben kolicine u skladistu za trazeniArtikal, mozes iskoristiti funkciju iznad, ako ima
	//ako ima, smanji kolicinu u bazi, (vidi na liniji 60-61)
	//vrati identifikator i true/false, zavisno da li je uspesno smanjena kolicina proizvoda

	return "temp", false
}
func sacuvajPorudzbinu(*messages.BuyProduct) { //cuvanje porudzbine koja je obradjena uspesno za consumera
	//sacuvaj porudzbinu(transactionId, id artikla, kolicina, ) + datum ,moze trenutni moment
}

type Porudzbina struct {
	//todo
}

func ProceniPotrebnuKolicinu(identifikator string) (potrebno int) {
	//proveri porudzbine za identifkator u zadnjem periodu(npr 2min),
	//sumiraj i vrati tu kolicinu * 1.2 (Poruci 20 posto vise )
	return 0
}

func SacuvajPorudzbinuPoslatuSupplieru(identifikator string, kolcicina int) {
	//sacuvati u posebnu kolekciju, tip PotrebanProizvod, dodati sadasnji trenutak i status poruceno
}

type PotrebanProizvod struct {
	identifikator string
	kolicina      int
	//dodaj datum
	//dodaj status, inicjalno poruceno, i moze se promeniti u zavrseno
}

func SacuvajDostavuOdSuppliera(dostava *messages.DostavaOdSuppliera) {
	//sacivati u posebnu kolekciju dostavu, preuzeti listu stavki, naziv skladista, i dodati trenutni datum- datum isporuke
}
