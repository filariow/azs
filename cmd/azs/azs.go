package main

import (
	"context"
	"fmt"
	"log"

	"github.com/filariow/azs/pkg/az"
)

func main() {
	ctx := context.Background()
	p, err := az.ReadProfiles(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range p.Subscriptions {
		fmt.Printf("%v\n", s)
	}
}
