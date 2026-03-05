package main

import (
	"context"
	"fmt"
	"log"
	"os"

	code "github.com/bkoshelev/go-project-244/src/gendiff"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "filepath1",
				Value:     "",
				UsageText: "path to first file",
			},
			&cli.StringArg{
				Name:      "filepath2",
				Value:     "",
				UsageText: "path to second file",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output format",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			filepath1 := cmd.StringArg("filepath1")
			filepath2 := cmd.StringArg("filepath1")

			format := cmd.String("format")

			if filepath1 == "" || filepath2 == "" {
				return fmt.Errorf("you need to write paths to configuration files")
			}

			result, err := code.GenDiff(filepath1, filepath2, format)

			if err != nil {
				return fmt.Errorf("fail to generate diff: %w", err)
			}
			fmt.Println(result)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
