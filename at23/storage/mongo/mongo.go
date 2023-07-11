package mongo

import (
	"at23/messages"
	"context"
	"fmt"
	"github.com/google/uuid"
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
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("golang_master").Collection("Proizvodi")
	filter := bson.M{"identifikator": artikal.ItemId}
	var pr Proizvod
	count, err := col.CountDocuments(context.TODO(), filter)
	if count == 1 {
		col.FindOne(ctx, filter).Decode(&pr)
		fmt.Println("PRONADJENA KOL:", pr.Kolicina)
		fmt.Println("POTREBNA KOL:", artikal.Amount)

		if pr.Kolicina < int(artikal.Amount) {
			fmt.Println("NEDOVOLJNA KOL")
			return pr.Identifikator, false
		}
		update := bson.M{"$inc": bson.M{"kolicina": -artikal.Amount}}
		col.FindOneAndUpdate(ctx, filter, update)
		var updatedPr Proizvod
		col.FindOne(ctx, filter).Decode(&updatedPr)
		fmt.Println("SADASNJA KOL:", updatedPr.Kolicina)

		return pr.Identifikator, true
	} else {
		fmt.Println("IDENTIFIKATOR", artikal.ItemId, "NE POSTOJI")
		return "none", false
	}
	//zeljenaKolicina := artikal.Amount
	//trazeniArtikal := artikal.ItemId
	//proveri da li ima potreben kolicine u skladistu za trazeniArtikal, mozes iskoristiti funkciju iznad, ako ima
	//ako ima, smanji kolicinu u bazi, (vidi na liniji 60-61)
	//vrati identifikator i true/false, zavisno da li je uspesno smanjena kolicina proizvoda

	//return "temp", false
}
func SacuvajTransakciju(artikal *messages.Item) {
	//cuvanje porudzbine koja je obradjena uspesno za consumera
	//sacuvaj porudzbinu(transactionId, id artikla, kolicina, ) + datum ,moze trenutni moment
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("golang_master").Collection("Transakcija")
	oneDoc := Transakcija{
		TransactionId: uuid.NewString(),
		IdArtikla:     artikal.ItemId,
		Kolicina:      int(artikal.Amount),
		Vreme:         time.Now(),
		Cena:          0,
		//ovde promeni cenu kada namestis cenu u proizvodima
	}
	_, insertErr := col.InsertOne(ctx, oneDoc)
	if insertErr != nil {
		fmt.Println("NIJE USPESNO SACUVANO")
		os.Exit(1)
	}
	fmt.Println("TRANSAKCIJA USPESNO SACUVANA")
	fmt.Println(oneDoc)

}

type Transakcija struct {
	//todo
	TransactionId string
	IdArtikla     string
	Kolicina      int
	Vreme         time.Time
	Cena          int
}

func ProceniPotrebnuKolicinu(identifikator string) (potrebno int) {
	//proveri porudzbine za identifkator u zadnjem periodu(npr 2min),
	//sumiraj i vrati tu kolicinu * 1.2 (Poruci 20 posto vise )
	return 0
}

func SacuvajPorudzbinuPoslatuSupplieru(identifikator string, kolcicina int) {
	//sacuvati u posebnu kolekciju, tip PotrebanProizvod, dodati sadasnji trenutak i status poruceno
}

type Status int

const (
	Poruceno Status = iota
	Zavrseno
)

type PotrebanProizvod struct {
	identifikator string
	kolicina      int
	status        Status
	time          time.Time
}

func SacuvajDostavuOdSuppliera(dostava *messages.DostavaOdSuppliera) {
	//sacivati u posebnu kolekciju dostavu, preuzeti listu stavki, naziv skladista, i dodati trenutni datum- datum isporuke
}
