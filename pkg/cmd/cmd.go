package cmd

import (
	"flag"
	"log"
	"os"
	"fmt"

	"github.com/kevinschoon/bankocr/pkg/account"
	"github.com/kevinschoon/bankocr/pkg/parser"
)

func maybe(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error() + "\n")
		os.Exit(1)
	}
}

func dump(path string) {
	p := parser.New()
	fp, err := os.Open(path)
	defer fp.Close()
	maybe(err)
	results, err := p.ReadAll(fp)
	maybe(err)
	log.Printf("loaded %d accounts", len(results))
	for _, number := range results {
		ac := account.Number(number)
		fmt.Printf("%s\n", ac)
	}
}

// Run is the main entrypoint
func Run() {
	var (
		path = flag.String("path", "", "path to an account data file")
	)
	flag.Parse()
	if *path == "" {
		maybe(fmt.Errorf("no path specified"))
	}
	dump(*path)
}
