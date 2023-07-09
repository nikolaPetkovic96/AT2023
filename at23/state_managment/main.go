package main

import (
	"at23/messages"
	"fmt"
	"time"

	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
)

type StateManagmentActor struct {
}
type ClockActor struct {
}

func (state *ClockActor) Receive(context actor.Context) {
	switch context.Message().(type) {
	case *messages.ClockPing:
		fmt.Println("Clock cycle..")
		// get random products from state managment

    fmt.Println("Pulling data..")
		spawnResponse, _ := remoting.SpawnNamed("127.0.0.1:8080", "state-distributor", "state", 10*time.Second)
		context.Send(spawnResponse.Pid, &messages.GetAllProductsState{Sender: context.Self()})

		// or some random value to start simulations
		time.Sleep(10 * time.Second)
		context.Send(context.Self(), &messages.ClockPing{})
	}

}

func (state *StateManagmentActor) Receive(context actor.Context) {
	fmt.Println("test")
	fmt.Println("got something", context.Message())
	switch msg := context.Message().(type) {
	case *messages.GetAllProductsState:
		fmt.Println("Transaction completed!", msg.Sender)
		context.Send(msg.Sender, messages.Ping{})
	case *messages.ReturnAllProductsState:
		fmt.Println("Got items!", msg.Items)
		items = msg.Items
	}
}

var remoting *remote.Remote
var items []*messages.Item

func main() {
	system := actor.NewActorSystem()
	remoteConfig := remote.Configure("127.0.0.1", 8091)
	remoting = remote.NewRemote(system, remoteConfig)
	remoting.Start()

	remoting.Register("state-man-actor", actor.PropsFromProducer(func() actor.Actor { return &StateManagmentActor{} }))

	context := system.Root
	props := actor.PropsFromProducer(func() actor.Actor { return &ClockActor{} })
	pid := context.Spawn(props)
	message := &messages.ClockPing{}

	context.ActorSystem().Root.Send(pid, message)

	console.ReadLine()
}
