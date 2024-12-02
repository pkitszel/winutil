package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func run(cmdName *string) error {
	if len(os.Args) < 2 {
		return errors.New("no command given, pick one of: server, wget, unzip")
	}

	var inName, inDesc, outDesc string
	outName := "out"
	var handler func(string, string) error

	*cmdName = os.Args[1]
	switch *cmdName {
	case "server":
		handler = serve_static_dir
		inDesc = `path to serve over HTTP`
		outDesc = `port to expose outside`
		inName = "dir"
		outName = "port"
	case "wget":
		handler = wget
		inDesc = `file location to download (with protocol, like 'http://')`
		outDesc = `where to save downloaded file, '-' for stdout`
		inName = "url"
	case "unzip":
		handler = unzip
		inDesc = `zip archive to extract`
		outDesc = `where to put extracted files, must be existing directory`
		inName = "in"
	default:
		return fmt.Errorf("unknown command %q", *cmdName)
	}

	flagSet := flag.NewFlagSet(os.Args[1], flag.ContinueOnError)
	inF := flagSet.String(inName, "", inDesc)
	outF := flagSet.String(outName, "", outDesc)
	err := flagSet.Parse(os.Args[2:])

	if err != nil {
		return err
	}
	if *inF == "" || *outF == "" {
		return fmt.Errorf(`two flags must be provided, see "-help" for details`)
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
