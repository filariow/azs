package main

import (
	"context"
	"fmt"
	"os"

	"github.com/filariow/azs/pkg/az"
	"github.com/filariow/azs/pkg/fzf"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p, err := az.ReadProfiles(ctx)
	if err != nil {
		fatal(err)
	}

	s, err := fzf.ChooseSubscription(p)
	if err != nil {
		fatal(err)
	}

	if err := az.ChangeProfile(ctx, s.ID); err != nil {
		fatal(err)
	}

	fmt.Printf("Changed subscription to '%s' (%s)\n", s.Name, s.ID)
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
