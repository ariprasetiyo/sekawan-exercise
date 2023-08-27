package unittest

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/go-routeros/routeros"
)

var (
	command  = flag.String("command", "/queue/simple/print", "RouterOS command")
	address  = flag.String("address", "change_with_ip_mikrotik:8728", "RouterOS address and port")
	username = flag.String("username", "change_with_username_mikrotik", "User name")
	password = flag.String("password", "change_with_password_mikrotik", "Password")
	async    = flag.Bool("async", false, "Use async code")
	useTLS   = flag.Bool("tls", false, "Use TLS")
)

func dial() (*routeros.Client, error) {
	if *useTLS {
		return routeros.DialTLS(*address, *username, *password, nil)
	}
	return routeros.Dial(*address, *username, *password)
}

func TestMikrotik(t *testing.T) {
	flag.Parse()
	c, err := dial()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	if *async {
		c.Async()
	}

	r, err := c.RunArgs(strings.Split(*command, " "))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
	log.Print(r)
}
