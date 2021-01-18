package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func run(cmdName *string) error {
	if len(os.Args) < 2 {
		return errors.New("no command given")
	}

	var inName, inDesc, outDesc string
	var handler func(string, string) error

	switch os.Args[1] {
	case "wget":
		handler = wget
		inDesc = `file location to download (with protocol, like 'http://')`
		outDesc = `where to save downloaded file, '-' for stdout`
		inName = "url"
		*cmdName = os.Args[1]
	case "unzip":
		handler = unzip
		inDesc = `zip archive to extract`
		outDesc = `where to put extracted files, must be existing directory`
		inName = "in"
		*cmdName = os.Args[1]
	default:
		return fmt.Errorf("unknown command %q", os.Args[1])
	}

	flagSet := flag.NewFlagSet(os.Args[1], flag.ContinueOnError)
	inF := flagSet.String(inName, "", inDesc)
	outF := flagSet.String("out", "", outDesc)
	err := flagSet.Parse(os.Args[2:])

	if err != nil {
		return err
	}
	if *inF == "" || *outF == "" {
		return fmt.Errorf("both input and output arguments must be provided, see help for details")
	}

	return handler(*inF, *outF)
}

func main() {
	cmdName := "winutil"
	rc := 0
	err := run(&cmdName)
	if err != nil && err != flag.ErrHelp {
		fmt.Fprintf(os.Stderr, "%s: %s\n", cmdName, err)
		rc = 1
	}
	os.Exit(rc)
}
