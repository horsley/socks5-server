package main

import (
	"log"
	"net"

	"github.com/armon/go-socks5"
	"golang.org/x/net/context"
)

// LogTarget log connect target and source
func LogTarget(chain socks5.RuleSet) *targetLogger {
	return &targetLogger{chain}
}

type targetLogger struct {
	chain socks5.RuleSet
}

func (t *targetLogger) Allow(ctx context.Context, req *socks5.Request) (context.Context, bool) {
	log.Println("req:", req.RemoteAddr.String(), " => ", req.DestAddr.String())
	if t.chain != nil {
		return t.chain.Allow(ctx, req)
	}
	return ctx, true
}

func (t *targetLogger) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		log.Println("dns:", name, "err:", err)
		return ctx, nil, err
	}
	log.Println("dns:", name, "=>", addr.IP)
	return ctx, addr.IP, err
}
