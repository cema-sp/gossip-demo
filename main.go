package main

import (
	"github.com/stefankopieczek/gossip/base"
	"github.com/stefankopieczek/gossip/log"
	"time"
)

var (
	// Caller parameters
	caller = &endpoint{
		displayName: base.String{"Semyon"},
		username:    base.String{"2233553341@sip2sip.info"},
		host:        "127.0.0.1",
		port:        5060,
		transport:   base.String{"UDP"},
	}

	// Callee parameters
	callee = &endpoint{
		displayName: base.String{"Cema SP"},
		username:    base.String{"cema_sp"},
		host:        "sipnet.ru",
		port:        5060,
		transport:   base.String{"UDP"},
	}
)

func main() {
	log.SetDefaultLogLevel(log.DEBUG)
	err := caller.Start()
	if err != nil {
		log.Warn("Failed to start caller: %v", err)
		return
	}

	// Receive an incoming call.
	caller.ServeInvite()

	<-time.After(2 * time.Second)

	// Send the BYE
	caller.Bye(callee)
}
