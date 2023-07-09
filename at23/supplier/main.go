package main

import (
	"at23/messages"
	"fmt"
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
	}
}

var remoting *remote.Remote

func main() {
	system := actor.NewActorSystem()
	remoteConfig := remote.Configure("127.0.0.1", 8093)
	remoting = remote.NewRemote(system, remoteConfig)
	remoting.Start()

	remoting.Register("supplier", actor.PropsFromProducer(func() actor.Actor { return &SupplierActor{} }))
  remoting.SpawnNamed("127.0.0.1:8093", "sup", "supplier", 3*time.Second)

	console.ReadLine()
}
