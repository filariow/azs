package fzf

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/filariow/azs/pkg/az"
	"github.com/ktr0731/go-fuzzyfinder"
)

var ErrorAbort = errors.New("Operation aborted by user")

func ChooseSubscription(p *az.Profile) (*az.Subscription, error) {
	ss := p.Subscriptions
	sort.Slice(ss, func(i, j int) bool {
		if ss[i].IsDefault {
			return true
		}
		if ss[j].IsDefault {
			return false
		}
		return strings.Compare(ss[i].Name, ss[j].Name) > 0
	})

	return fzfSubscription(ss)
}

func fzfSubscription(ss []az.Subscription) (*az.Subscription, error) {
	idx, err := fuzzyfinder.Find(
		ss,
		func(i int) string {
			s := ss[i]
			if s.IsDefault {
				return fmt.Sprintf("\u2713 %s", s.Name)
			}
			return ss[i].Name
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			s := ss[i]
			d := "\u274C"
			if s.IsDefault {
				d = "\u2713"
			}

			return fmt.Sprintf(
				`Name            %s
Subscription    %s
User            %s
Tenant          %s
IsDefault       %s
State           %s`,
				s.Name, s.ID, s.User.Name, s.TenantId, d, s.State)
		}))
	if err != nil {
		if errors.Is(err, fuzzyfinder.ErrAbort) {
			return nil, ErrorAbort
		}
		return nil, err
	}
	return &ss[idx], nil
}
