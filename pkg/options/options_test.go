package options_test

import (
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/zvdy/goldeneye/pkg/options"
)

func TestParseOptions(t *testing.T) {
	tests := []struct {
		args     []string
		expected options.Options
	}{
		{
			args: []string{"-u", "useragents.txt", "-w", "100", "-s", "50", "-m", "post", "-d", "-n=false"},
			expected: options.Options{
				UserAgents: "useragents.txt",
				Workers:    100,
				Sockets:    50,
				Method:     "post",
				Debug:      true,
				NoSSLCheck: false,
			},
		},
		{
			args: []string{"-u", "agents.txt", "-w", "200", "-s", "60", "-m", "get", "-d", "-n=true"},
			expected: options.Options{
				UserAgents: "agents.txt",
				Workers:    200,
				Sockets:    60,
				Method:     "get",
				Debug:      true,
				NoSSLCheck: true,
			},
		},
	}

	for _, test := range tests {
		// Reset the flags
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		os.Args = append([]string{"cmd"}, test.args...)

		opts := options.ParseOptions()

		if opts.UserAgents != test.expected.UserAgents {
			t.Errorf("expected UserAgents %s, got %s", test.expected.UserAgents, opts.UserAgents)
		}
		if opts.Workers != test.expected.Workers {
			t.Errorf("expected Workers %d, got %d", test.expected.Workers, opts.Workers)
		}
		if opts.Sockets != test.expected.Sockets {
			t.Errorf("expected Sockets %d, got %d", test.expected.Sockets, opts.Sockets)
		}
		if opts.Method != test.expected.Method {
			t.Errorf("expected Method %s, got %s", test.expected.Method, opts.Method)
		}
		if opts.Debug != test.expected.Debug {
			t.Errorf("expected Debug %v, got %v", test.expected.Debug, opts.Debug)
		}
		if opts.NoSSLCheck != test.expected.NoSSLCheck {
			t.Errorf("expected NoSSLCheck %v, got %v", test.expected.NoSSLCheck, opts.NoSSLCheck)
		}
	}
}

func TestPrintHelp(t *testing.T) {
	// Capture the output of PrintHelp
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	options.PrintHelp()

	w.Close()
	os.Stdout = old

	var buf [1024]byte
	n, _ := r.Read(buf[:])
	output := string(buf[:n])

	expected := `USAGE: ./goldeneye <url> [OPTIONS]

OPTIONS:
    Flag           Description                     Default
    -u, --useragents   File with user-agents to use                     (default: randomly generated)
    -w, --workers      Number of concurrent workers                     (default: 50)
    -s, --sockets      Number of concurrent sockets                     (default: 30)
    -m, --method       HTTP Method to use 'get' or 'post'  or 'random'  (default: get)
    -d, --debug        Enable Debug Mode [more verbose output]          (default: False)
    -n, --nosslcheck   Do not verify SSL Certificate                    (default: True)
    -h, --help         Shows this help`

	// Trim the trailing newline character from the actual output
	output = strings.TrimSpace(output)

	if output != expected {
		t.Errorf("expected help output %q, got %q", expected, output)
	}
}
