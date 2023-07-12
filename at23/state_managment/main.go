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
	items []*messages.Item
}
type ClockActor struct {
}

func (state *ClockActor) Receive(context actor.Context) {
	switch context.Message().(type) {
	case *messages.ClockPing:
		fmt.Println("Clock cycle..")

		spawnedPid, _ := remoting.SpawnNamed("127.0.0.1:8091", "state", "state", 10*time.Second)
		context.Send(spawnedPid.Pid, &messages.ClockPing{})

		// or some random value to start simulations
		time.AfterFunc(10*time.Second, func() {
			context.Send(context.Self(), &messages.ClockPing{})
		})
	}

}

func (state *StateManagmentActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.ClockPing:
		fmt.Println("Pulling data..")
		spawnResponse, _ := remoting.SpawnNamed("127.0.0.1:8080", "state-distributor", "state", 15*time.Second)
		context.Send(spawnResponse.Pid, &messages.GetAllProductsState{Sender: context.Self()})
		fmt.Println(state.items)
	case *messages.GetAllProductsState:
		fmt.Println("Transaction completed!")
		spawnResponse, _ := remoting.SpawnNamed("127.0.0.1:8090", "state-consumer", "simulation", 10*time.Second)
		context.Send(spawnResponse.Pid, &messages.ReturnAllProductsState{Items: state.items})
	case *messages.ReturnAllProductsState:
		fmt.Println("Got items!", msg.Items)
		state.items = msg.Items
	}
}

var remoting *remote.Remote

func main() {
	system := actor.NewActorSystem()
	remoteConfig := remote.Configure("127.0.0.1", 8091)
	remoting = remote.NewRemote(system, remoteConfig)
	remoting.Start()

	remoting.Register("state", actor.PropsFromProducer(func() actor.Actor { return &StateManagmentActor{} }))

	context := system.Root
	props := actor.PropsFromProducer(func() actor.Actor { return &ClockActor{} })
	pid := context.Spawn(props)
	message := &messages.ClockPing{}

	context.ActorSystem().Root.Send(pid, message)

	console.ReadLine()
}
