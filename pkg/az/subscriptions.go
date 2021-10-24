package az

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/dimchansky/utfbom"
)

type Profile struct {
	Subscriptions []Subscription `json:"subscriptions"`
}

type Subscription struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	State           string `json:"state"`
	User            User   `json:"user"`
	IsDefault       bool   `json:"isDefault"`
	TenantId        string `json:"tenantId"`
	EnvironmentName string `json:"environmentName"`
	HomeTenantId    string `json:"homeTenantId"`
}

type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

const azureProfilePath = ".azure/azureProfile.json"

func ReadProfiles(ctx context.Context) (*Profile, error) {
	h, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error reading home dir: %w", err)
	}
	fp := path.Join(h, azureProfilePath)
	f, err := os.Open(fp)
	if err != nil {
		return nil, fmt.Errorf("error opening path %s: %w", fp, err)
	}

	pb, err := ioutil.ReadAll(utfbom.SkipOnly(f))
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", fp, err)
	}

	var p Profile
	if err := json.Unmarshal(pb, &p); err != nil {
		return nil, fmt.Errorf("error unmarshaling json data: %w", err)
	}

	return &p, nil
}

func WriteProfiles(ctx context.Context) error {
	panic("not implemented")
}

func ChangeProfile(ctx context.Context, subscriptionID string) error {
	cmd := exec.CommandContext(ctx, "az", "account", "set", "--subscription", subscriptionID)

	if err := cmd.Run(); err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			return err
		}
	}

	return nil
}
