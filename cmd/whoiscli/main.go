package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"whois-ip-cli/internal/whois"
)

const version = "1.1.0"

var (
	flagJSON    = flag.Bool("json", false, "Ausgabe als JSON")
	flagVersion = flag.Bool("v", false, "Version anzeigen")
	flagHelp    = flag.Bool("h", false, "Hilfe anzeigen")
)

func main() {
	flag.Parse()

	if *flagHelp {
		fmt.Println("Nutzung: whoiscli [Optionen] <domain|ip>")
		fmt.Println("Optionen:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *flagVersion {
		fmt.Println("whoiscli Version", version)
		os.Exit(0)
	}

	if flag.NArg() < 1 {
		fmt.Println("Fehler: kein Ziel angegeben. Nutzung mit -h anzeigen.")
		os.Exit(1)
	}

	target := flag.Arg(0)

	if whois.IsIP(target) {
		whois.LookupIP(target, *flagJSON)
	} else {
		valid := regexp.MustCompile(`^[a-zA-Z0-9.-]+$`).MatchString
		if !valid(target) {
			fmt.Println("Ungültige Eingabe.")
			os.Exit(1)
		}
		whois.LookupDomain(target, *flagJSON)
	}
}
