package command

import (
	"fmt"
	"github.com/amanvir/enumerator/scanner"
	"github.com/urfave/cli"
	"os"
)

func CmdDomain(c *cli.Context) error {
	domain := c.Args().Get(0)
	wordList := c.Args().Get(1)

	if domain == "" {
		fmt.Println("Please specify a domain")
		os.Exit(2)
	}

	if wordList == "" {
		fmt.Println("Please specify a wordlist")
		os.Exit(2)
	}

	subdomain.Bruteforce(domain, wordList)

	return nil
}
