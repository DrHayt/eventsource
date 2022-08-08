package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/vancelongwill/eventsource"
	"github.com/vancelongwill/eventsource/_examples/fullproto/models/orders"
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

	//store, err := dynamodbstore.New("es_orders")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	store, err := boltdbstore.New("orders")
	check(err)

	repo := eventsource.New(&orders.Aggregate{},
		eventsource.WithStore(store),
		eventsource.WithSerializer(orders.NewSerializer()),
	)

	id := *OrderId
	ctx := context.Background()

	t0, err := repo.Apply(ctx,
		&orders.CreateOrder{
			Id: id,
			By: "Elon Musk",
		},
	)
	check(err)
	spew.Dump(t0)

	for i := 0; i < 20000; i++ {
		if _, err := repo.Apply(ctx, &orders.ChangeName{
			Id:        id,
			FirstName: "First",
			LastName:  "Last",
		}); err != nil {
			panic(err)
		}
	}

	//t1, err := repo.Apply(ctx, &orders.ChangeName{
	//	Id:        id,
	//	FirstName: "First",
	//	LastName:  "First",
	//})
	//check(err)
	//spew.Dump(t1)
	//
	//t2, err := repo.Apply(ctx, &orders.CancelOrder{
	//	Id:     id,
	//	Reason: "I dont like you",
	//})
	//check(err)
	//spew.Dump(t2)
	//
	//t3, err := repo.Apply(ctx, &orders.ChangeName{
	//	Id:        id,
	//	FirstName: "First",
	//	LastName:  "Second",
	//})
	//check(err)
	//spew.Dump(t3)
	//
	{
		t0 := time.Now()
		aggregate, err := repo.Load(ctx, id)
		check(err)
		fmt.Printf("Too: %s\n", time.Now().Sub(t0))
		found := aggregate.(*orders.Aggregate)
		spew.Dump(found)
	}

	{
		t1 := time.Now()
		aggregate, err := repo.Load(ctx, id)
		check(err)
		fmt.Printf("Too: %s\n", time.Now().Sub(t1))
		found := aggregate.(*orders.Aggregate)
		spew.Dump(found)
	}

	//found := aggregate.(*orders.Aggregate)
	//fmt.Printf("Order %v [version %v] %v %v\n", found.ID, found.Version, found.State, found.UpdatedAt.Format(time.RFC822))
	//spew.Dump(found)
}
