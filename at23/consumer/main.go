package main

import (
	"at23/messages"
	"fmt"
	"time"

	"github.com/google/uuid"

	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
)

type ConsumerActor struct {
}
type SimulationActor struct {
}

func (state *SimulationActor) Receive(context actor.Context) {
  switch msg := context.Message().(type) {
	case *messages.Simulate:
		fmt.Println("Starting simulation")
		spawnResponse, _ := remoting.SpawnNamed("127.0.0.1:8091", "state", "state", 5*time.Second)
    context.Send(spawnResponse.Pid, &messages.GetAllProductsState{})
		
    // or some random value to start simulations
		time.AfterFunc(10 * time.Second, func() {
      context.Send(context.Self(), &messages.Simulate{})
    })
	case *messages.ReturnAllProductsState:
		//code to send message
		// TODO get random products from state managment
		consumerProps := actor.PropsFromProducer(func() actor.Actor { return &ConsumerActor{} })
		consumerPID := context.Spawn(consumerProps)
		context.Send(consumerPID, &messages.StartSimulation{Items: msg.Items})
	}

}

func (state *ConsumerActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.StartSimulation:
		fmt.Println("Simulating..")
		spawnResponse, _ := remoting.SpawnNamed("127.0.0.1:8080", "consumer-distributor", "consumer", 5*time.Second)
		context.Send(spawnResponse.Pid, &messages.BuyProduct{
			TransactionId: uuid.NewString(),
			Items:         msg.Items,
			Sender:        context.Self(),
		},
		)
	case *messages.CompletedTransaction:
		fmt.Println("Transaction completed!", msg.TransactionId)
	}
}

var remoting *remote.Remote

func main() {
	system := actor.NewActorSystem()
	remoteConfig := remote.Configure("127.0.0.1", 8090)
	remoting = remote.NewRemote(system, remoteConfig)
	remoting.Start()
	context := system.Root
	
	remoting.Register("simulation", actor.PropsFromProducer(func() actor.Actor { return &SimulationActor{} }))

  props := actor.PropsFromProducer(func() actor.Actor { return &SimulationActor{} })
	pid := context.Spawn(props)
	message := &messages.Simulate{}

  

	context.ActorSystem().Root.Send(pid, message)

	console.ReadLine()
}
