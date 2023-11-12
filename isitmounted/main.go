package main

import (
	"fmt"
	"os"
	"flag"
	"github.com/moby/sys/mountinfo"
)

const True = 0
const False = 1

var helpFlag bool

func init() {
	flag.BoolVar(&helpFlag, "h", false, "Print help message")
	flag.BoolVar(&helpFlag, "help", false, "Print help message")
}

func main() {
	flag.Parse()

	if helpFlag || flag.NArg() == 0 {
		printHelp()
		os.Exit(False)
	}

    dir := flag.Arg(0)
	isMounted := isMounted(dir)

	if isMounted {
		os.Exit(True)
	} else {
		os.Exit(False)
	}
}

func usage() {
	fmt.Println("usage: isitmounted [-h] /directory")
	os.Exit(False)
}

func printHelp() {
	fmt.Println("Check if the specified directory is a mountpoint")
	fmt.Println("Sets bash $? to zero if it is a mountpoint")
	fmt.Println(" ")
	usage()
	os.Exit(False)
}

func isMounted(dir string) bool {
	isit, err := mountinfo.Mounted(dir)
	if err != nil {
		fmt.Printf("Error calling Mounted: %s", err)
		os.Exit(False)
	}

	return isit
}

