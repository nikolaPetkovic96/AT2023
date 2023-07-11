package main

import (
	"at23/messages"
	mongo "at23/storage/mongo"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/cluster"
	"github.com/asynkron/protoactor-go/cluster/clusterproviders/automanaged"
	"github.com/asynkron/protoactor-go/cluster/identitylookup/disthash"
	"github.com/asynkron/protoactor-go/remote"
)

var naziv_skladista string
var port int
var adresa string
var port2 int

type PingPongActor struct {
	system *actor.ActorSystem
}
type KupacActor struct {
	system *actor.ActorSystem
}
type StorageActor struct {
	system *actor.ActorSystem
}

func (*KupacActor) Receive(context actor.Context) {
	fmt.Println("Actor kupca pogodjen:")
	switch msg := context.Message().(type) {
	case *messages.BuyProduct: //
		//msg.Items je lista ali pretpostavi da ima jedan clan
		for _, artikal := range msg.Items {
			_, kupljeno := mongo.KupiArtikal(artikal)
			var porucenaKol int32
			if kupljeno {
				porucenaKol = artikal.Amount
				mongo.SacuvajTransakciju(artikal)
			} else {
				porucenaKol = 0
				//obavesti storage aktora da smo pri kraju sa zalihama za trazeni proizvod
				//proceni
				porucenaKol = mongo.ProceniPotrebnuKolicinu(artikal.ItemId)
				//sacuvaj porudzenicu
				//posalji zahtev
			}
			context.Respond(&messages.ArtikalPorucen{
				TransactionId:  msg.TransactionId,
				Nazivskladista: naziv_skladista,
				Identifikator:  artikal.ItemId,
				Kolicina:       porucenaKol,
			})
		}
	}
}
func (*StorageActor) Receive(context actor.Context) {
	fmt.Println("Actor Storage pogodjen:")
	switch msg := context.Message().(type) {
	case *messages.BuyProduct: //Storage actor proveri da li ima potrebnih kolicina i vrati kolicinu za trazene artikle u skladistu
		fmt.Println("Provera prozizvoda u skladistu :", naziv_skladista)
		for _, artikal := range msg.Items {
			artikal.Amount = int32(mongo.ProveriKolicine(artikal.ItemId))
			fmt.Println("Pronadjena kolicina za id:", artikal.ItemId, "=", artikal.Amount)
		}
		time.Sleep(3 * time.Second)
		context.Respond(msg)

	case string:
		{ //pogadjace se iz 48. linije, pocinje forimranje zahteva koji treba poslati supplieru
			//identifikator := msg
			//probaj napraviti metodu koja ce za identifikator proveriti koliko je prodato komada u nekom zadnjem periodu(1min-2)
			//pa na osnovu toga sracunati

		}
	}
}

// po pravilu bi trebao biti definsian endpoint ali spolja se uglavnom gadja odrejedni Kind agenta
func (*PingPongActor) PingStorage(r *messages.Ping) *messages.Pong {
	fmt.Println("PING")
	var adr []string
	adr = append(adr, "funkcija endpoint")
	pong := &messages.Pong{Adrese: adr}
	return pong

}
func (*PingPongActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.Ping:
		fmt.Println("Pogodjeno skladiste", naziv_skladista)
		var adr []string
		adr = append(adr, naziv_skladista)
		pong := &messages.Pong{Adrese: adr}
		context.Respond(pong)
	case *messages.Pong:
		fmt.Println("Pogodjeno skladiste", naziv_skladista)
		pong := &messages.Pong{Adrese: append(msg.Adrese, naziv_skladista)}
		context.Respond(pong)
	case *messages.NoviArtikli:
		mongo.DodajArtikleProto(msg.Artikli)
	}
}
func main() {
	naziv_skladista = os.Args[1]        //"ryzen"
	port, _ := strconv.Atoi(os.Args[3]) //8080
	adresa = os.Args[2]
	port2, _ := strconv.Atoi(os.Args[3])
	fmt.Println(naziv_skladista, adresa, port)
	system := actor.NewActorSystem()

	config := remote.Configure(adresa, port)
	//var cluster_host string = adresa + ":" + strconv.Itoa(port)
	provider := automanaged.NewWithConfig(1*time.Second, port2, "localhost"+":"+strconv.Itoa(port2))
	lookup := disthash.New()
	clusterKinds := dodajKindPingPong(system)
	KupacKind := dodajKindKupacActor(system)
	StorageKind := dodajKindStorageActor(system)
	clusterConfig := cluster.Configure(naziv_skladista, provider, lookup, config, cluster.WithKinds(&clusterKinds, &KupacKind, &StorageKind))
	c := cluster.New(system, clusterConfig)
	c.StartMember()
	defer c.Shutdown(false)
	// for _, member := range c.MemberList.Members().Members() {
	// 	fmt.Println("MEMBER:", member)
	// }
	//javljanje na coordinatora
	for {
		//testiranje dodavanje novog proizvoda ili dodavanje novih kolicina postojeceg
		var artikli []mongo.Proizvod
		artikli = append(artikli, mongo.Proizvod{Identifikator: "1", Kolicina: rand.Intn(7)})
		artikli = append(artikli, mongo.Proizvod{Identifikator: "2", Kolicina: rand.Intn(7)})
		artikli = append(artikli, mongo.Proizvod{Identifikator: "3", Kolicina: rand.Intn(7)})
		mongo.DodajArtikle(artikli)

		//javljanje na coordinatora- radi se samo jednom
		spawnResponse, err := c.Remote.SpawnNamed("127.0.0.1:8080", "coord", "ping", time.Second)
		if err != nil {
			continue
		}
		system.Root.Send(spawnResponse.Pid, &messages.Ping{Ime: naziv_skladista, Adresa: adresa, Port: int32(port), Port2: int32(port2)})
		break
	}
	// Run till a signal comes
	finish := make(chan os.Signal, 1)
	signal.Notify(finish, os.Interrupt, os.Kill)
	<-finish
}

func dodajKindPingPong(system *actor.ActorSystem) cluster.Kind {
	fmt.Println("dodajKind")
	kindPingPong := cluster.NewKind(
		"Pingpong",
		actor.PropsFromProducer(func() actor.Actor {
			return &PingPongActor{
				system: system,
			}
		}))
	return *kindPingPong
}
func dodajKindKupacActor(system *actor.ActorSystem) cluster.Kind {
	fmt.Println("dodajKind")
	kindPingPong := cluster.NewKind(
		"Kupac",
		actor.PropsFromProducer(func() actor.Actor {
			return &KupacActor{
				system: system,
			}
		}))
	return *kindPingPong
}
func dodajKindStorageActor(system *actor.ActorSystem) cluster.Kind {
	fmt.Println("dodajKind")
	kindPingPong := cluster.NewKind(
		"Storage",
		actor.PropsFromProducer(func() actor.Actor {
			return &StorageActor{
				system: system,
			}
		}))
	return *kindPingPong
}

const (
// pingPongAgent = "pingPong_"
)
