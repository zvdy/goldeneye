package options

import (
    "flag"
    "fmt"
)

func ParseOptions() *Options {
    opts := &Options{}

    flag.StringVar(&opts.UserAgents, "useragents", "randomly generated", "File with user-agents to use (default: randomly generated)")
    flag.StringVar(&opts.UserAgents, "u", "randomly generated", "File with user-agents to use (default: randomly generated)")
    flag.IntVar(&opts.Workers, "workers", 50, "Number of concurrent workers (default: 50)")
    flag.IntVar(&opts.Workers, "w", 50, "Number of concurrent workers (default: 50)")
    flag.IntVar(&opts.Sockets, "sockets", 30, "Number of concurrent sockets (default: 30)")
    flag.IntVar(&opts.Sockets, "s", 30, "Number of concurrent sockets (default: 30)")
    flag.StringVar(&opts.Method, "method", "get", "HTTP Method to use 'get' or 'post' or 'random' (default: get)")
    flag.StringVar(&opts.Method, "m", "get", "HTTP Method to use 'get' or 'post' or 'random' (default: get)")
    flag.BoolVar(&opts.Debug, "debug", false, "Enable Debug Mode [more verbose output] (default: False)")
    flag.BoolVar(&opts.Debug, "d", false, "Enable Debug Mode [more verbose output] (default: False)")
    flag.BoolVar(&opts.NoSSLCheck, "nosslcheck", true, "Do not verify SSL Certificate (default: True)")
    flag.BoolVar(&opts.NoSSLCheck, "n", true, "Do not verify SSL Certificate (default: True)")
    flag.BoolVar(&opts.Help, "help", false, "Shows this help")
    flag.BoolVar(&opts.Help, "h", false, "Shows this help")

    flag.Parse()

    if len(flag.Args()) > 0 {
        opts.URL = flag.Args()[0]
    }

    return opts
}

func PrintHelp() {
    fmt.Println(`USAGE: ./goldeneye <url> [OPTIONS]

OPTIONS:
    Flag           Description                     Default
    -u, --useragents   File with user-agents to use                     (default: randomly generated)
    -w, --workers      Number of concurrent workers                     (default: 50)
    -s, --sockets      Number of concurrent sockets                     (default: 30)
    -m, --method       HTTP Method to use 'get' or 'post'  or 'random'  (default: get)
    -d, --debug        Enable Debug Mode [more verbose output]          (default: False)
    -n, --nosslcheck   Do not verify SSL Certificate                    (default: True)
    -h, --help         Shows this help`)
}