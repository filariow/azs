package fzf

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
)

func ChooseProfile() (string, error) {
	cmd := exec.Command("fzf", "--ansi", "--preview", "az account show -s {} -o yaml")
	var out bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = &out

	cmd.Env = append(os.Environ(),
		`FZF_DEFAULT_COMMAND=az account list --query '[].id' -o tsv`)
	if err := cmd.Run(); err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			return "", err
		}
	}
	choice := strings.TrimSpace(out.String())
	if choice == "" {
		return "", errors.New("you did not choose any of the options")
	}

	return out.String(), nil
}
