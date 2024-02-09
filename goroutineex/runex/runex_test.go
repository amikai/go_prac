package runex

import (
	"context"
	"log"
	"net/http"
	"syscall"

	"github.com/oklog/run"
)

func ExampleSignalHandler() {
	startServer := func(name string, srv *http.Server) func() error {
		return func() error {
			log.Printf("starting %s server", name)
			return srv.ListenAndServe()
		}
	}
	shoutdownServer := func(name string, srv *http.Server) func(err error) {
		return func(err error) {
			log.Printf("shutting down %s server", name)
			_ = srv.Shutdown(context.Background())
		}
	}

	ctx := context.Background()
	g := &run.Group{}
	g.Add(run.SignalHandler(ctx, syscall.SIGINT, syscall.SIGTERM))
	srv1, srv2 := &http.Server{Addr: ":9527"}, &http.Server{Addr: ":9528"}
	g.Add(startServer("server1", srv1), shoutdownServer("server1", srv2))
	g.Add(startServer("server2", srv1), shoutdownServer("server2", srv2))
	err := g.Run()
	log.Printf("run err: %v", err)
}

func ExampleContextHandler() {
	startServer := func(name string, srv *http.Server) func() error {
		return func() error {
			log.Printf("starting %s server", name)
			return srv.ListenAndServe()
		}
	}
	shoutdownServer := func(name string, srv *http.Server) func(err error) {
		return func(err error) {
			log.Printf("shutting down %s server", name)
			_ = srv.Shutdown(context.Background())
		}
	}

	ctx := context.Background()
	g := &run.Group{}
	g.Add(run.ContextHandler(ctx))
	srv1, srv2 := &http.Server{Addr: ":9527"}, &http.Server{Addr: ":9528"}
	g.Add(startServer("server1", srv1), shoutdownServer("server1", srv2))
	g.Add(startServer("server2", srv1), shoutdownServer("server2", srv2))
	err := g.Run()
	log.Printf("run err: %v", err)
}
