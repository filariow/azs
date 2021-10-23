package main

import (
	"context"
	"fmt"
	"log"

	"github.com/filariow/azs/pkg/az"
	"github.com/filariow/azs/pkg/fzf"
)

func main() {
	ctx := context.Background()
	p, err := az.ReadProfiles(ctx)
	if err != nil {
		log.Fatal(err)
	}

	c := len(p.Subscriptions)
	ii := make([]string, c, c)
	for j, i := range p.Subscriptions {
		ii[j] = i.ID
	}

	cp, err := fzf.ChooseProfile()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Chosen profile: %s", cp)
	if err := az.ChangeProfile(cp); err != nil {
		log.Fatal(err)
	}
}
