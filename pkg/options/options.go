package options

type Options struct {
	URL        string
	UserAgents string
	Workers    int
	Sockets    int
	Method     string
	Debug      bool
	NoSSLCheck bool
	Help       bool
}
