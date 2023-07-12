package mongo

import (
	"at23/messages"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"

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
		Kolicina:      artikal.Amount,
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
	Kolicina      int32
	Vreme         time.Time
	Cena          int
}

func ProceniPotrebnuKolicinu(identifikator string) (potrebno int32) {
	//proveri porudzbine za identifkator u zadnjem periodu(npr 2min),
	//sumiraj i vrati tu kolicinu * 1.2 (Poruci 20 posto vise )
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("golang_master").Collection("Transakcija")
	filter := bson.M{"idartikla": identifikator, "vreme": bson.M{"$gt": time.Now().Add(-time.Hour * 20)}}
	cur, err := col.Find(ctx, filter)
	var transakcije []Transakcija
	for cur.Next(ctx) {
		var transakcija Transakcija
		cur.Decode(&transakcija)
		transakcije = append(transakcije, transakcija)
	}
	potrebno = 0
	if len(transakcije) != 0 {
		for _, trans := range transakcije {
			potrebno += trans.Kolicina
		}
		potrebno += potrebno / 5
		fmt.Println("POTREBNO JE ", potrebno, "ARTIKALA SA IDENTIFIKATOROM", identifikator)

	}
	return potrebno
}

// salje suplijeru
func SacuvajPorudzbinu(identifikator string, kolicina int32) {
	//sacuvati u posebnu kolekciju, tip PotrebanProizvod, dodati sadasnji trenutak i status poruceno
	if kolicina == 0 {
		return
	}
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("golang_master").Collection("Porudzbina")
	oneDoc := Porudzbina{
		Identifikator: uuid.NewString(),
		IdArtikla:     identifikator,
		Kolicina:      kolicina,
		Vreme:         time.Now(),
		Status:        Poruceno,
	}
	_, insertErr := col.InsertOne(ctx, oneDoc)
	if insertErr != nil {
		fmt.Println("PORUDZBINA NIJE USPESNO SACUVANO")
		os.Exit(1)
	}
	fmt.Println("PORUCENO JE ", kolicina, "ARTIKALA SA IDENTIFIKATOROM", identifikator)
	fmt.Println(oneDoc)
}

type Status int

const (
	Poruceno Status = iota
	Zavrseno
)

type Porudzbina struct {
	Identifikator string
	IdArtikla     string
	Kolicina      int32
	Status        Status
	Vreme         time.Time
}

// func SacuvajDostavuOdSuppliera(dostava *messages.DostavaOdSuppliera) {
// 	//sacivati u posebnu kolekciju dostavu, preuzeti listu stavki, naziv skladista, i dodati trenutni datum- datum isporuke
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		os.Exit(1)
// 	}
// 	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
// 	col := client.Database("golang_master").Collection("Dostava")

// 	var proizvodi []Proizvod
// 	for _, dost := range dostava.Stavke {
// 		var proizvod = Proizvod{
// 			Identifikator: dost.Identifikator,
// 			Kolicina:      int(dost.Kolcicina),
// 			//Cena:          dost.Cena,
// 		}
// 		NabaviArtikal(&messages.Item{ItemId: proizvod.Identifikator, Amount: int32(proizvod.Kolicina)})
// 		proizvodi = append(proizvodi, proizvod)
// 	}

// 	var oneDoc = Dostava{
// 		Supplier_id:     dostava.SupplierId,
// 		Naziv_skladista: dostava.NazivSkladista,
// 		Proizvodi:       proizvodi,
// 	}
// 	_, insertErr := col.InsertOne(ctx, oneDoc)
// 	if insertErr != nil {
// 		fmt.Println("NIJE USPESNO SACUVANO")
// 		os.Exit(1)
// 	}
// 	fmt.Println("DOSTAVA USPESNO SACUVANA")
// 	fmt.Println(oneDoc)
// }

type Dostava struct {
	//Supplier_id     string
	Naziv_skladista string
	Proizvodi       []Proizvod
}

func NabaviArtikal(artikal *messages.Item) {
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
		update := bson.M{"$inc": bson.M{"kolicina": artikal.Amount}}
		col.FindOneAndUpdate(ctx, filter, update)
		var updatedPr Proizvod
		col.FindOne(ctx, filter).Decode(&updatedPr)
		fmt.Println("USPESNO POVECANA KOLICINA ARTIKLA", artikal.ItemId, "ZA", artikal.Amount)
		return
	} else {
		fmt.Println("IDENTIFIKATOR", artikal.ItemId, "NE POSTOJI")
		return
	}
}

func GetAllProducts() (items []*messages.Item) {
	var retValue []Proizvod
	var retPointers []*messages.Item
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		os.Exit(1)
	}
	//ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("golang_master").Collection("Proizvodi")
	cursor, _ := col.Find(context.TODO(), bson.D{{}})
	if err = cursor.All(context.TODO(), &retValue); err != nil {
		panic(err)
	}
	for _, r := range retValue {
		retPointers = append(retPointers, &messages.Item{ItemId: r.Identifikator, Amount: int32(r.Kolicina)})
	}
	fmt.Println("Ukupno razlicitih proizvoda", len(retPointers))
	return retPointers

}

func SacuvajDostavuOdSuppliera(dostava *messages.ReturnItems, skladiste string) {
	//sacivati u posebnu kolekciju dostavu, preuzeti listu stavki, naziv skladista, i dodati trenutni datum- datum isporuke
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("golang_master").Collection("Dostava")

	var proizvodi []Proizvod
	for _, dost := range dostava.Items {
		var proizvod = Proizvod{
			Identifikator: dost.ItemId,
			Kolicina:      int(dost.Amount),
			//Cena:          dost.Cena,
		}
		NabaviArtikal(&messages.Item{ItemId: proizvod.Identifikator, Amount: int32(proizvod.Kolicina)})
		proizvodi = append(proizvodi, proizvod)
	}

	var oneDoc = Dostava{
		//Supplier_id:     dostava.SupplierId,
		Naziv_skladista: skladiste,
		Proizvodi:       proizvodi,
	}
	_, insertErr := col.InsertOne(ctx, oneDoc)
	if insertErr != nil {
		fmt.Println("NIJE USPESNO SACUVANO")
		os.Exit(1)
	}
	fmt.Println("DOSTAVA USPESNO SACUVANA")
	fmt.Println(oneDoc)
}
