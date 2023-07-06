package main

import (
	"at23/messages"
	"fmt"
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

func (*PingPongActor) PingStorage(r *messages.Ping) *messages.Pong {
	fmt.Println("PING")
	var adr []string
	adr = append(adr, "test")
	pong := &messages.Pong{Adrese: adr}
	return pong

}
func (*PingPongActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.Ping:
		fmt.Println("PING")
		var adr []string
		adr = append(adr, "test")
		pong := &messages.Pong{Adrese: adr}
		context.Respond(pong)
	case *messages.Pong:
		fmt.Println("Pong primljen sa : ", msg.Adrese)
		pong := &messages.Pong{Adrese: msg.Adrese}
		context.Respond(pong)
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
	clusterConfig := cluster.Configure(naziv_skladista, provider, lookup, config, cluster.WithKinds(&clusterKinds))
	c := cluster.New(system, clusterConfig)
	c.StartMember()
	defer c.Shutdown(false)
	//javljanje na coordinatora
	for _, member := range c.MemberList.Members().Members() {
		fmt.Println("MEMBER:", member)
	}
	for {
		// for _, member := range c.MemberList.Members().Members() {
		// 	fmt.Println("MEMBER:", member)
		// }
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

const (
// pingPongAgent = "pingPong_"
)
