package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/vancelongwill/eventsource"
	"github.com/vancelongwill/eventsource/_examples/fullproto/orders"
	"github.com/vancelongwill/eventsource/boltdbstore"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var (
		OrderId = flag.String("o", "1", "What order number are we working with")
	)
	flag.Parse()

	store, err := boltdbstore.New("orders")
	check(err)

	repo := eventsource.New(&orders.Order{},
		eventsource.WithStore(store),
		eventsource.WithSerializer(eventsource.NewProtoSerializer(
			orders.OrderCreated{},
			orders.OrderShipped{},
		)),
	)

	id := *OrderId
	ctx := context.Background()

	//if _, err := repo.Load(ctx, id); err != nil {
	//	// How to distinguish between a missing ID and a systems failure.
	//	panic(err)
	//}

	_, err = repo.Apply(ctx, &orders.CreateOrder{
		CommandModel: eventsource.CommandModel{ID: id},
	})
	check(err)

	_, err = repo.Apply(ctx, &orders.ShipOrder{
		CommandModel: eventsource.CommandModel{ID: id},
	})
	check(err)

	aggregate, err := repo.Load(ctx, id)
	check(err)

	found := aggregate.(*orders.Order)
	fmt.Printf("Order %v [version %v] %v %v\n", found.ID, found.Version, found.State, found.UpdatedAt.Format(time.RFC822))
}
