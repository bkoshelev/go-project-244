package main

import (
	"context"
	"fmt"
	"log"
	"os"

	code "github.com/bkoshelev/go-project-244/src"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			result := code.GenDiff()
			fmt.Println(result)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
