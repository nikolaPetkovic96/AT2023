package main

import (
	"at23/messages"
	"fmt"
	"time"

	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
)

type ConsumerActor struct {
}
type SimulationActor struct {
}

func (state *SimulationActor) Receive(context actor.Context) {
	switch context.Message().(type) {
	case *messages.Simulate:
		fmt.Println("Starting simulation")
    // get random products from state managment

		//code to send message
		consumerProps := actor.PropsFromProducer(func() actor.Actor { return &ConsumerActor{} })
		consumerPID := context.Spawn(consumerProps)
		context.Send(consumerPID, &messages.StartSimulation{
			ProduktId: "Name",
			Kolicina:  1,
		})

		// or some random value to start simulations
		time.Sleep(10 * time.Second)
		context.Send(context.Self(), &messages.Simulate{
			Name: "Name",
		})

	}

}

func (state *ConsumerActor) Receive(context actor.Context) {
	switch context.Message().(type) {
	case *messages.StartSimulation:
		fmt.Println("Simulating..")
		spawnResponse, _ := remoting.SpawnNamed("127.0.0.1:8080", "whatever", "consumer", 5*time.Second)
		context.Send(spawnResponse.Pid, &messages.BuyProduct{
			TransactionId: "123",
			Stavke: []*messages.Stavka{
				{
					ProduktId: "123",
					Kolicina:  2,
				},
				{
					ProduktId: "442",
					Kolicina:  4,
				},
			},
		},
		)

	}
}

var remoting *remote.Remote

func main() {
	system := actor.NewActorSystem()
	remoteConfig := remote.Configure("127.0.0.1", 8090)
	remoting = remote.NewRemote(system, remoteConfig)
	remoting.Start()
	context := system.Root
	props := actor.PropsFromProducer(func() actor.Actor { return &SimulationActor{} })
	pid := context.Spawn(props)
	message := &messages.Simulate{Name: "Name"}

	context.ActorSystem().Root.Send(pid, message)

	console.ReadLine()
}
