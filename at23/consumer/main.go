package main

import (
	"at23/messages"
	"fmt"
	"time"

	//mongo "at23/coordinator/mongo"
  console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
)


type ConsumerActor struct {
	system *actor.ActorSystem
}
type SimulationActor struct {
	system *actor.ActorSystem
}

func (p *SimulationActor) Receive(context actor.Context) {
  switch msg := context.Message().(type) {
    case *messages.Simulate:
      fmt.Println("Starting sumulation")
      // conde to send message
      consumerProps := actor.PropsFromProducer(func() actor.Actor { return &ConsumerActor{} }) 
		  consumerPID := context.Spawn(consumerProps)
      
      context.Send(msg.consumerPID, &messages.StartSimulation{
            SomeValue: "result",
      })
      
      // or some random value to start simulations
      time.Sleep(10 * time.Second)
      context.Send(msg.Self(), &messages.Simulate{
            SomeValue: "result",
      })


	}

}

func (state *ConsumerActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
    case *messages.StartSimulation: //dodaje skladiste prema nazivu u cluster, sem ako vec ne postoji node zaduzen za to skladiste
    
  }
}

func main() {
  system := actor.NewActorSystem()
	remoteConfig := remote.Configure("127.0.0.1", 8092)
	remoting := remote.NewRemote(system, remoteConfig)
	remoting.Start()

	context := system.Root
	props := actor.PropsFromProducer(func() actor.Actor { return &SimulationActor{} })
	pid := context.Spawn(props)
	message := &messages.Echo{Message: "hej", Sender: pid}

	// this is to spawn remote actor we want to communicate with
	// spawnResponse, err := remoting.SpawnNamed("127.0.0.1:8090", "myactor", "hello", time.Second)
  // TODO spawn remote actor in distributor

  spawnResponse, err := remoting.SpawnNamed("127.0.0.1:8092", "SimulationActor", "hello", time.Second)
	if err != nil {
		panic(err)
		return
	}

	// get spawned PID
	spawnedPID := spawnResponse.Pid
	for i := 0; i < 10; i++ {
		context.Send(spawnedPID, message)
	}

	console.ReadLine()
	}
