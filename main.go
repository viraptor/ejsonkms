package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// version information. This will be overridden by the ldflags
var version = "dev"

func main() {
	app := cli.NewApp()
	app.Usage = "manage encrypted secrets using EJSON & AWS KMS"
	app.Version = version
	app.Author = "Steve Hodgkiss"
	app.Email = "steve@envato.com"
	app.Commands = []cli.Command{
		{
			Name:  "encrypt",
			Usage: "(re-)encrypt one or more EJSON files",
			Action: func(c *cli.Context) {
				if err := encryptAction(c.Args()); err != nil {
					fmt.Fprintln(os.Stderr, "Encryption failed:", err)
					os.Exit(1)
				}
			},
		},
		{
			Name:  "decrypt",
			Usage: "decrypt an EJSON file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "o",
					Usage: "print output to the provided file, rather than stdout",
				},
				cli.StringFlag{
					Name:  "aws-region",
					Usage: "AWS Region",
				},
			},
			Action: func(c *cli.Context) {
				if err := decryptAction(c.Args(), c.String("aws-region"), c.String("o")); err != nil {
					fmt.Fprintln(os.Stderr, "Decryption failed:", err)
					os.Exit(1)
				}
			},
		},
		{
			Name:  "keygen",
			Usage: "generate a new EJSON keypair",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "kms-key-id",
					Usage: "KMS Key ID to encrypt the private key with",
				},
				cli.StringFlag{
					Name:  "aws-region",
					Usage: "AWS Region",
				},
				cli.StringFlag{
					Name:  "o",
					Usage: "write EJSON file to a file rather than stdout",
				},
			},
			Action: func(c *cli.Context) {
				if err := keygenAction(c.Args(), c.String("kms-key-id"), c.String("aws-region"), c.String("o")); err != nil {
					fmt.Fprintln(os.Stderr, "Key generation failed:", err)
					os.Exit(1)
				}
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "Unexpected failure:", err)
		os.Exit(1)
	}
}