package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func fatalf(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
	fmt.Println()
	os.Exit(1)
}

func main() {
	var format string
	flag.StringVar(&format, "f", time.RFC3339Nano, "parsing format string")
	flag.Usage = func() {
		fmt.Printf("gdd YYYY-MM-DDThh:mm:ss.ssssZ YYYY-MM-DDThh:mm:ss.ssssZ")
		fmt.Println("returns the difference between the supplied dates in RFC3339Nano format")
		flag.PrintDefaults()
	}
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		flag.Usage()
		os.Exit(0)
	}
	d1, err := time.Parse(format, args[0])
	if err != nil {
		fatalf("could not parse %s: %s", args[0], err)
	}
	d2, err := time.Parse(format, args[1])
	if err != nil {
		fatalf("could not parse %s: %s", args[1], err)
	}
	diff := d2.Sub(d1)
	fmt.Println(diff)
}
