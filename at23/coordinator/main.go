package main

import (
	"at23/messages"
	"fmt"
	"math"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/google/uuid"

	//mongo "at23/coordinator/mongo"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/cluster"
	"github.com/asynkron/protoactor-go/cluster/clusterproviders/automanaged"
	"github.com/asynkron/protoactor-go/cluster/identitylookup/disthash"
	"github.com/asynkron/protoactor-go/remote"
)

const (
	pingPongAgent = "pingPong_"
)

// var klasteri *[]cluster
var skladista []string
var cluster_system *actor.ActorSystem
var skladista_clusteri []cluster.Cluster
var remoting *remote.Remote

var suppliers []string

type PingActor struct {
	system *actor.ActorSystem
}
type PingActor2 struct {
	system *actor.ActorSystem
}
type ConsumerActor struct {
	system *actor.ActorSystem
}
type StateActor struct {
}
type SupplierActor struct {
}
type SupplierRegisterActor struct {
}

func (p *PingActor2) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case struct{}:
		ch := make(chan *actor.Future, len(skladista_clusteri))

		for _, skl := range skladista_clusteri {
			ping := &messages.Ping{}

			grainPid := skl.Get("pp"+skl.Config.Name, "Pingpong")
			//future := ctx.RequestFuture(grainPid, ping, time.Second)
			ch <- ctx.RequestFuture(grainPid, ping, 3*time.Second)
			// result, err := future.Result()
			// if err != nil {
			// 	fmt.Println(err.Error(), skl)
			// 	return
			// }
			// fmt.Printf("Received %v", result)
		}
		time.Sleep(15 * time.Second)
		close(ch)
		for temp := range ch {
			//temp := <-ch
			result, err := temp.Result()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("Received", result)
		}

	}

}

func (state *PingActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.Ping: //dodaje skladiste prema nazivu u cluster, sem ako vec ne postoji node zaduzen za to skladiste
		fmt.Println("Coordinator Pingovan od strane :", msg.Ime, msg.Adresa, msg.GetPort())
		dodajSkladiste(msg.Adresa, int(msg.Port), msg.Ime, int(msg.Port2))
		fmt.Println("Skladiste dodato!")
	case *messages.Pong:
	}
}

func (state *ConsumerActor) Receive(context actor.Context) {
	fmt.Println("Got request from a consumer")
	switch msg := context.Message().(type) {
	case *messages.BuyProduct:
		fmt.Println("Requested items:", msg.Items)
		//context.Send(msg.Sender, &messages.CompletedTransaction{TransactionId: msg.TransactionId})
		ch := make(chan *actor.Future, len(skladista_clusteri))
		//var pids []*actor.PID
		for i, skl := range skladista_clusteri {
			pid := skl.Get("porudzbina_"+skl.Config.Name+strconv.Itoa(i), "Storage")
			msg.Sender = pid
			//future := context.RequestFuture(pid, msg, 10*time.Second)
			ch <- context.RequestFuture(pid, msg, 10*time.Second)
			// result, err := future.Result()
			// if err != nil {
			// 	fmt.Println(err.Error(), skl)
			// 	return
			// }
			// fmt.Printf("Odgovor sa skladista %v", result)
		}

		time.Sleep(15 * time.Second)
		close(ch)
		stanjaMap := make(map[string]*messages.BuyProduct2)
		var poruciIz string
		for temp := range ch {
			result, err := temp.Result()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("Received", result)
			stanje, ok := result.(*messages.BuyProduct2)
			if ok {
				stanjaMap[stanje.Skladiste] = stanje
			}
		}
		for key, elem := range stanjaMap {
			mapaProizvoda := make(map[string]int)
			for _, proizvod := range elem.Items {
				mapaProizvoda[proizvod.ItemId] = int(proizvod.Amount)
			}
			imaSve := true
			for _, item := range msg.Items {
				if item.Amount > int32(mapaProizvoda[item.ItemId]) {
					imaSve = false
					break
				}
			}
			if imaSve == true {
				poruciIz = key
				break
			}
		}
		var temp cluster.Cluster
		for _, sklad := range skladista_clusteri {
			if sklad.Config.Name == poruciIz {
				temp = sklad
			}
		}
		pid := temp.Get("por_"+msg.TransactionId, "Kupac")
		future := context.RequestFuture(pid, msg, 10*time.Second)
		result, err := future.Result()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("Odgovor sa kupovinu %v", result)

		// result, error := future.Result()
		// if error != nil {
		// 	fmt.Println(err.Error())
		// 	return
		// }
		// odgovor, ok := result.(*messages.BuyProduct2)
		// if ok {
		// 	context.Respond(odgovor)
		// }

	}
}

func (state *StateActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.GetAllProductsState:
		fmt.Println("Pulling data..")
		//TODO write code that pulls all items from databases and sums them up
		// context.Send(msg.Sender, &messages.ReturnAllProductsState{
		// 	Items: []*messages.Item{
		// 		{
		// 			ItemId: "1",
		// 			Amount: 2000,
		// 		},
		// 		{
		// 			ItemId: "2",
		// 			Amount: 2000,
		// 		},
		// 		{
		// 			ItemId: "3",
		// 			Amount: 2000,
		// 		},
		// 	},
		// })
		ch := make(chan *actor.Future, len(skladista_clusteri))
		var retItems []*messages.Item
		for _, skl := range skladista_clusteri {
			grainPid := skl.Get("pp"+skl.Config.Name, "Storage")
			ch <- context.RequestFuture(grainPid, msg, 10*time.Second)
		}
		time.Sleep(5 * time.Second)
		close(ch)
		for temp := range ch {
			//temp := <-ch

			result, err := temp.Result()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			lista, ok := result.(*messages.ReturnAllProductsState)
			if ok {
				retItems = append(retItems, lista.Items...)
			}
		}
		fmt.Println("Duzina liste: ", len(retItems))
		fmt.Println("RetItems:", retItems)
		context.Send(msg.Sender, &messages.ReturnAllProductsState{Items: CollectData(retItems)})
	}
}

func (state *SupplierActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	// maybe remove this case in future
	case *messages.GetItems:
		// check prices first
		ch := make(chan *actor.Future, len(skladista_clusteri))
		smallestPrice := math.Inf(1)
		smallestPriceAddress := ""
		for _, element := range suppliers {
			spawnResponse1, _ := remoting.SpawnNamed(element, "sup", "supplier", 10*time.Second)
			ch <- context.RequestFuture(spawnResponse1.Pid, &messages.CheckPrice{
				Items:  msg.Items,
				Sender: context.Self(),
			}, 10*time.Second)
		}
		time.Sleep(15 * time.Second)
		close(ch)
		for temp := range ch {
			result, err := temp.Result()
			if err != nil {
				price, ok := result.(*messages.ReturnPrice)
				if ok {
					fmt.Println("Response:", price.Address, price.Price)
					if smallestPrice > float64(price.Price) {
						smallestPrice = float64(price.Price)
						smallestPriceAddress = price.Address
					}
				} else {
					fmt.Println("Unexpected response type")
				}
			}
		}

		fmt.Println("Got smallestPrice of", smallestPrice, smallestPriceAddress)

		if smallestPrice == math.Inf(1) || smallestPriceAddress == "" {
			fmt.Println("Something went wrong, data was corrupted")
			return
		}

		// send request to cheapest supplier
		spawnResponse, _ := remoting.SpawnNamed(smallestPriceAddress, "sup", "supplier", 10*time.Second)
		context.Send(spawnResponse.Pid, &messages.GetItems{
			Items:         msg.Items,
			TransactionId: uuid.NewString(),
			Sender:        context.Self(),
		})
	case *messages.ReturnItems:
		fmt.Println("GOT ITEMS BACK! YAY!", msg.Items, msg.TransactionId)
		//var  isporuka messages.DostavaOdSuppliera
		//pozovi agenta za storage i posalji mu dostavu, takodje iz gloablne StorageNarucio ukloni ili celu transakciju, ako nije cela zavrsena
	}
}

// Register Suppliers
func (state SupplierRegisterActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.RegisterSupplier:
		suppliers = append(suppliers, msg.Address)
		fmt.Println(msg.Address)
	}
}

func main() {
	//inicjalizacija klastera kojem storage-i salju zahteve
	system := actor.NewActorSystem()
	cluster_system = system
	config := remote.Configure("127.0.0.1", 8079)
	clusterProvider := automanaged.NewWithConfig(1*time.Second, 6330, "localhost:6330")
	lookup := disthash.New()
	DistributorKind := dodajKindDistrubutorActor(system)

	clusterConfig := cluster.Configure("distributor_cluster", clusterProvider, lookup, config, cluster.WithKinds(&DistributorKind))
	ClusterDist = cluster.New(system, clusterConfig)
	ClusterDist.StartMember()
	defer ClusterDist.Shutdown(false)

	// //konfiguracija servera- bez clustera
	system2 := actor.NewActorSystem()
	remoteConfig := remote.Configure("127.0.0.1", 8080)
	remoting := remote.NewRemote(system2, remoteConfig)
	remoting.Start()

	// //registracija agenta kojem ce se clusteri obracati pri inicijalizaciji
	remoting.Register("ping", actor.PropsFromProducer(func() actor.Actor { return &PingActor{system: system2} }))
	remoting.Register("consumer", actor.PropsFromProducer(func() actor.Actor { return &ConsumerActor{} }))
	remoting.Register("state", actor.PropsFromProducer(func() actor.Actor { return &StateActor{} }))
	remoting.Register("supplier-register", actor.PropsFromProducer(func() actor.Actor { return &SupplierRegisterActor{} }))

	// START OF SECTOR
	// for test purposes only for now, actor needs to be spawned elsewhere
	// props := actor.PropsFromProducer(func() actor.Actor { return &SupplierActor{} })
	// pid := system2.Root.Spawn(props)
	// spawnResponse, _ := remoting.SpawnNamed("127.0.0.1:8093", "sup", "supplier", 10*time.Second)
	// system2.Root.Send(spawnResponse.Pid, &messages.GetItems{
	// 	Items: []*messages.Item{
	// 		{
	// 			ItemId: "123",
	// 			Amount: 2,
	// 		},
	// 		{
	// 			ItemId: "442",
	// 			Amount: 4,
	// 		},
	// 	},
	// 	TransactionId: uuid.NewString(),
	// Sender: pid,
	// })
	// END OF SECTOR

	ping2Prop := actor.PropsFromProducer(func() actor.Actor {
		return &PingActor2{
			system: system2,
		}
	})
	pingPid := system2.Root.Spawn(ping2Prop)

	//testiraj slanje zahteva na storage
	time.Sleep(10 * time.Second)
	// kupacProp := actor.PropsFromProducer(func() actor.Actor {
	// 	return &ConsumerActor{
	// 		system: system2,
	// 	}
	// })
	// KupacPid := system2.Root.Spawn(kupacProp)
	// system2.Root.Send(KupacPid, &messages.BuyProduct{
	// 	Sender:        KupacPid,
	// 	TransactionId: "1",
	// 	Items:         []*messages.Item{{ItemId: "1", Amount: 2000}, {ItemId: "1", Amount: 20}, {ItemId: "42", Amount: 2000}, {ItemId: "2", Amount: 2}},
	// })
	//////////kraj test slanje zahteva na storage
	//
	// Run till a signal comes
	finish := make(chan os.Signal, 1)
	signal.Notify(finish, os.Interrupt, os.Kill)
	//<-finish
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C: //salje ping svim klasterima
			system2.Root.Send(pingPid, struct{}{})
		case <-finish:
			return
		}
	}
}

func dodajSkladiste(adresa string, port int, nazivSkladista string, port2 int) {
	system := actor.NewActorSystem()
	freePortA, _ := GetFreePort()
	config := remote.Configure("127.0.0.1", freePortA)
	freePort, _ := GetFreePort()
	clusterProvider := automanaged.NewWithConfig(1*time.Second, freePort, "localhost:"+strconv.Itoa(port2))
	lookup := disthash.New()
	clusterConfig := cluster.Configure(nazivSkladista, clusterProvider, lookup, config)
	c := cluster.New(system, clusterConfig)
	c.StartClient()
	defer c.Shutdown(false)
	skladista_clusteri = append(skladista_clusteri, *c)
}

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func dodajKindDistrubutorActor(system *actor.ActorSystem) cluster.Kind {
	fmt.Println("dodajKind Distributor")
	kindPingPong := cluster.NewKind(
		"Distributor",
		actor.PropsFromProducer(func() actor.Actor {
			return &DistributorActor{
				system: system,
			}
		}))
	return *kindPingPong
}

type DistributorActor struct {
	system *actor.ActorSystem
}

func (p *DistributorActor) Receive(ctx actor.Context) {
	fmt.Println("Got request from a STORAGE")
	switch msg := ctx.Message().(type) {
	case *messages.PorudzbinaZaSuppliera: //storage porucuje
		fmt.Println("Pogodjen DISTIRBUTOR od strana skladista", msg)
	case *messages.DostavaOdSuppliera: //supplier dobavio -prosledi za storage

	}

}

var StorageNarucio []*messages.PorudzbinaZaSuppliera
var ClusterDist *cluster.Cluster

func CollectData(list []*messages.Item) []*messages.Item {
	var returnList []*messages.Item
	for _, item := range list {
		ind := false
		for _, item2 := range returnList {
			if item.ItemId == item2.ItemId {
				item2.Amount += item.Amount
				ind = true
			}
		}
		if ind == false {
			returnList = append(returnList, &messages.Item{ItemId: item.ItemId, Amount: item.Amount})
		}
	}
	return returnList
}
