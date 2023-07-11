package main

import (
	"at23/messages"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
)

type SupplierActor struct {
}

func (state *SupplierActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.GetItems:
		fmt.Println("Returning items..", msg.Items)
		context.Send(msg.Sender, &messages.ReturnItems{
			TransactionId: msg.TransactionId,
			Items:         msg.Items,
		})

	case *messages.CheckPrice:
		fmt.Println("Checking price for items..", msg.Items)
    // returning random price
    context.Respond(&messages.ReturnPrice{
      Price: float32(randFloat(100, 3000)),
    })
	}
}

var remoting *remote.Remote

func main() {
	system := actor.NewActorSystem()
  address := "127.0.0.1"
  port := 8093
	remoteConfig := remote.Configure(address, port)

	remoting = remote.NewRemote(system, remoteConfig)
	remoting.Start()

	remoting.Register("supplier", actor.PropsFromProducer(func() actor.Actor { return &SupplierActor{} }))
  remoting.SpawnNamed(address+":"+strconv.Itoa(port), "sup", "supplier", 3*time.Second)

  // regisering on coordinator
  spawnResponse, err := remoting.SpawnNamed("127.0.0.1:8080", "supplier-register", "supplier-register", 3*time.Second)
  if err == nil{
    system.Root.Send(spawnResponse.Pid, &messages.RegisterSupplier{Address: address+":"+strconv.Itoa(port)})
  }

	console.ReadLine()
}

func randFloat(min, max float64) float64 {
    return min + rand.Float64()*(max-min)
}

