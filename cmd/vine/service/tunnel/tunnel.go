// MIT License
//
// Copyright (c) 2020 Lack
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package tunnel

import (
	"os"
	"strings"
	"time"

	"github.com/lack-io/cli"
	"github.com/lack-io/vine"
	"github.com/lack-io/vine/core/client"
	"github.com/lack-io/vine/core/client/mucp"
	"github.com/lack-io/vine/core/registry/memory"
	rr "github.com/lack-io/vine/core/router"
	"github.com/lack-io/vine/core/router/registry"
	"github.com/lack-io/vine/core/server"
	smucp "github.com/lack-io/vine/core/server/mucp"
	log "github.com/lack-io/vine/lib/logger"
	tun "github.com/lack-io/vine/lib/network/tunnel"
	"github.com/lack-io/vine/lib/network/tunnel/transport"
	"github.com/lack-io/vine/lib/proxy"
	pmucp "github.com/lack-io/vine/lib/proxy/mucp"
	"github.com/lack-io/vine/util/muxer"
)

var (
	// Name of the tunnel service
	Name = "go.vine.tunnel"
	// Address is the tunnel address
	Address = ":8083"
	// Tunnel is the name of the tunnel
	Tunnel = "tun:0"
	// The tunnel token
	Token = "vine"
)

// Run runs the vine server
func Run(ctx *cli.Context, svcOpts ...vine.Option) {
	if len(ctx.String("server-name")) > 0 {
		Name = ctx.String("server-name")
	}
	if len(ctx.String("address")) > 0 {
		Address = ctx.String("address")
	}
	if len(ctx.String("token")) > 0 {
		Token = ctx.String("token")
	}
	if len(ctx.String("id")) > 0 {
		Tunnel = ctx.String("id")
		// We need host:port for the Endpoint value in the proxy
		parts := strings.Split(Tunnel, ":")
		if len(parts) == 1 {
			Tunnel = Tunnel + ":0"
		}
	}
	var nodes []string
	if len(ctx.String("server")) > 0 {
		nodes = strings.Split(ctx.String("server"), ",")
	}

	// Initialise service
	svc := vine.NewService(
		vine.Name(Name),
		vine.RegisterTTL(time.Duration(ctx.Int("register-ttl"))*time.Second),
		vine.RegisterInterval(time.Duration(ctx.Int("register-interval"))*time.Second),
	)

	// local tunnel router
	r := registry.NewRouter(
		rr.Id(svc.Server().Options().Id),
		rr.Registry(svc.Client().Options().Registry),
	)

	// start the router
	if err := r.Start(); err != nil {
		log.Errorf("Tunnel error starting router: %s", err)
		os.Exit(1)
	}

	// create a tunnel
	t := tun.NewTunnel(
		tun.Address(Address),
		tun.Nodes(nodes...),
		tun.Token(Token),
	)

	// start the tunnel
	if err := t.Connect(); err != nil {
		log.Errorf("Tunnel error connecting: %v", err)
	}

	log.Infof("Tunnel [%s] listening on %s", Tunnel, Address)

	// create tunnel client with tunnel transport
	tunTransport := transport.NewTransport(
		transport.WithTunnel(t),
	)

	// local server client talks to tunnel
	localSrvClient := mucp.NewClient(
		client.Transport(tunTransport),
	)

	// local proxy
	localProxy := pmucp.NewProxy(
		proxy.WithClient(localSrvClient),
		proxy.WithEndpoint(Tunnel),
	)

	// create new muxer
	mux := muxer.New(Name, localProxy)

	// init server
	svc.Server().Init(
		server.WithRouter(mux),
	)

	// local transport client
	svc.Client().Init(
		client.Transport(svc.Options().Transport),
	)

	// local proxy
	tunProxy := pmucp.NewProxy(
		proxy.WithRouter(r),
		proxy.WithClient(svc.Client()),
	)

	// create memory registry
	memRegistry := memory.NewRegistry()

	// local server
	tunSrv := smucp.NewServer(
		server.Address(Tunnel),
		server.Transport(tunTransport),
		server.WithRouter(tunProxy),
		server.Registry(memRegistry),
	)

	if err := tunSrv.Start(); err != nil {
		log.Errorf("Tunnel error starting tunnel server: %v", err)
		os.Exit(1)
	}

	if err := svc.Run(); err != nil {
		log.Errorf("Tunnel %s failed: %v", Name, err)
	}

	// stop the router
	if err := r.Stop(); err != nil {
		log.Errorf("Tunnel error stopping tunnel router: %v", err)
	}

	// stop the server
	if err := tunSrv.Stop(); err != nil {
		log.Errorf("Tunnel error stopping tunnel server: %v", err)
	}

	if err := t.Close(); err != nil {
		log.Errorf("Tunnel error stopping tunnel: %v", err)
	}
}

func Commands(options ...vine.Option) []*cli.Command {
	command := &cli.Command{
		Name:  "tunnel",
		Usage: "Run the vine network tunnel",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "address",
				Usage:   "Set the vine tunnel address :8083",
				EnvVars: []string{"VINE_TUNNEL_ADDRESS"},
			},
			&cli.StringFlag{
				Name:    "id",
				Usage:   "Id of the tunnel used as the internal dial/listen address.",
				EnvVars: []string{"VINE_TUNNEL_ID"},
			},
			&cli.StringFlag{
				Name:    "server",
				Usage:   "Set the vine tunnel server address. This can be a comma separated list.",
				EnvVars: []string{"VINE_TUNNEL_SERVER"},
			},
			&cli.StringFlag{
				Name:    "token",
				Usage:   "Set the vine tunnel token for authentication",
				EnvVars: []string{"VINE_TUNNEL_TOKEN"},
			},
		},
		Action: func(ctx *cli.Context) error {
			Run(ctx, options...)
			return nil
		},
	}

	return []*cli.Command{command}
}
