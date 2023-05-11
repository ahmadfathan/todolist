package helper

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	defaultGraceTimeout = 10 * time.Second
)

var (
	// ErrGraceShutdownTimeout happens when the server graceful shutdown exceed the given grace timeout.
	ErrGraceShutdownTimeout = errors.New("server shutdown timed out")
)

// HTTPServer represents an HTTP server
type HTTPServer interface {
	Shutdown(ctx context.Context) error
	Serve(l net.Listener) error
}

// GRPCServer represents interface for grpc server
type GRPCServer interface {
	GracefulStop()
	Stop()
	Serve(l net.Listener) error
}

// ServeGRPC start the grpc server on the given address and add graceful shutdown handler
// graceTimeout specify how long we want to wait for the graceful stop to run.
// if exceed the duration, we will forcefully stop the gRPC server.
// if graceTimeout = 0, we use default value: 30 second
func ServeGRPC(server GRPCServer, address string, graceTimeout time.Duration) error {
	lis, err := Listen(address)
	if err != nil {
		return err
	}

	stoppedCh := WaitTermSig(func(ctx context.Context) error {
		if graceTimeout == 0 {
			graceTimeout = defaultGraceTimeout
		}
		stopped := make(chan struct{})
		go func() {
			server.GracefulStop()
			close(stopped)
		}()

		select {
		case <-time.After(graceTimeout):
			server.Stop()
			return ErrGraceShutdownTimeout
		case <-stopped:
		}
		return nil
	})

	log.Printf("GRPC server running on adress %v", address)

	if err := server.Serve(lis); err != nil {
		// Error starting or closing listener:
		return err
	}

	<-stoppedCh
	log.Println("GRPC server stopped")
	return nil
}

// ServeHTTP start the http server on the given address and add graceful shutdown handler
// graceTimeout specify how long we want to wait for the Shutdown to run.
// if graceTimeout = 0, we use default value: 30 second
func ServeHTTP(srv HTTPServer, address string, graceTimeout time.Duration) error {
	// start graceful listener
	lis, err := Listen(address)
	if err != nil {
		return err
	}

	stoppedCh := WaitTermSig(func(ctx context.Context) error {
		if graceTimeout == 0 {
			graceTimeout = defaultGraceTimeout
		}

		stopped := make(chan struct{})
		ctx, cancel := context.WithTimeout(ctx, graceTimeout)
		defer cancel()
		go func() {
			srv.Shutdown(ctx)
			close(stopped)
		}()

		select {
		case <-ctx.Done():
			return ErrGraceShutdownTimeout
		case <-stopped:

		}

		return nil
	})

	log.Printf("http server running on address: %v", address)

	// start serving
	if err := srv.Serve(lis); err != http.ErrServerClosed {
		return err
	}

	<-stoppedCh
	log.Println("HTTP server stopped")
	return nil
}

// WaitTermSig wait for termination signal and then execute the given handler
// when the signal received
//
// The handler is usually service shutdown, so we can properly shutdown our server upon SIGTERM.
//
// It returns channel which will be closed after the signal received and the handler executed.
// We can use the signal to wait for the shutdown to be finished.
func WaitTermSig(handler func(context.Context) error) <-chan struct{} {
	stoppedCh := make(chan struct{})
	go func() {
		signals := make(chan os.Signal, 1)

		// wait for the sigterm
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-signals

		// We received an os signal, shut down.
		if err := handler(context.Background()); err != nil {
			log.Printf("graceful shutdown  failed: %v", err)
		} else {
			log.Println("gracefull shutdown succeed")
		}

		close(stoppedCh)

	}()
	return stoppedCh
}

// Listen listens to the given port or to file descriptor as specified by socketmaster.
//
// This method is taken from tokopedia/grace repo and  modified to work with
// socketmaster's -wait-child-notif option.
func Listen(port string) (net.Listener, error) {
	var l net.Listener

	// see if we run under socketmaster
	fd := os.Getenv("EINHORN_FDS")
	if fd != "" {
		sock, err := strconv.Atoi(fd)
		if err != nil {
			return nil, err
		}
		log.Println("detected socketmaster, listening on", fd)
		file := os.NewFile(uintptr(sock), "listener")
		fl, err := net.FileListener(file)
		if err != nil {
			return nil, err
		}
		l = fl
	}

	if l != nil { // we already have the listener, which listen on EINHORN_FDS
		notifSocketMaster()
		return l, nil
	}

	// we are not using socketmaster, no need to notify

	return net.Listen("tcp4", port)
}

// notifSocketMaster notify socket master about our readyness
// we should remove this func after we are fully moved to tokopedia new platform
func notifSocketMaster() {
	go func() {
		err := NotifyMaster()
		if err != nil {
			log.Printf("failed to notify socketmaster: %v, ignore if you don't use `wait-child-notif` option", err)
		} else {
			log.Println("successfully notify socketmaster")
		}
	}()
}
