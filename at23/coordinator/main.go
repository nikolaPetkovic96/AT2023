package main

import (
	"at23/messages"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"

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

type PingActor struct {
	system *actor.ActorSystem
}
type PingActor2 struct {
	system *actor.ActorSystem
}

func (p *PingActor2) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case struct{}:
		//ch := make(chan *actor.Future, len(skladista_clusteri))

		for _, skl := range skladista_clusteri {
			ping := &messages.Ping{}

			grainPid := skl.Get("pp"+skl.Config.Name, "Pingpong")
			future := ctx.RequestFuture(grainPid, ping, time.Second)
			//ch <- ctx.RequestFuture(grainPid, ping, time.Second)
			result, err := future.Result()
			if err != nil {
				fmt.Println(err.Error(), skl)
				return
			}
			fmt.Printf("Received %v", result)
		}
		//close(ch)
		// for range ch {
		// 	temp := <-ch
		// 	result, err := temp.Result()
		// 	if err != nil {
		// 		fmt.Println(err.Error())
		// 		return
		// 	}
		// 	fmt.Printf("Received %v", result)
		// }

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

func main() {
	// system := actor.NewActorSystem()
	// cluster_system = system
	// config := remote.Configure("127.0.0.1", 8079)
	// clusterProvider := automanaged.NewWithConfig(1*time.Second, 6330, "localhost:6331")
	// lookup := disthash.New()
	// clusterConfig := cluster.Configure("skladiste", clusterProvider, lookup, config)
	// c := cluster.New(system, clusterConfig)
	// c.StartClient()
	// defer c.Shutdown(false)

	// //konfiguracija servera- bez clustera
	system2 := actor.NewActorSystem()
	remoteConfig := remote.Configure("127.0.0.1", 8080)
	remoting := remote.NewRemote(system2, remoteConfig)
	remoting.Start()

	// //registracija agenta kojem ce se clusteri obracati pri inicijalizaciji
	remoting.Register("ping", actor.PropsFromProducer(func() actor.Actor { return &PingActor{system: system2} }))

	ping2Prop := actor.PropsFromProducer(func() actor.Actor {
		return &PingActor2{
			system: system2,
		}
	})
	pingPid := system2.Root.Spawn(ping2Prop)

	// Run till a signal comes
	finish := make(chan os.Signal, 1)
	signal.Notify(finish, os.Interrupt, os.Kill)
	//<-finish

	ticker := time.NewTicker(5 * time.Second)
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
