package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	linkShortenerV1 "github.com/olezhek28/link-shortener/internal/app/api/link_shortener_v1"
	desc "github.com/olezhek28/link-shortener/pkg/link_shortener/v1"
	"google.golang.org/grpc"
)

const (
	httpPort = ":80"
	grpcPort = ":50051"
)

type App struct {
	linkShortenerImpl *linkShortenerV1.Implementation
	serviceProvider   *serviceProvider

	grpcServer *grpc.Server
	mux        *runtime.ServeMux
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)

	return a, err
}

func (a *App) Run() error {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	err := a.runGRPC(wg)
	if err != nil {
		return err
	}

	err = a.runPublicHTTP(wg)
	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initServer,
		a.initGRPCServer,
		a.initPublicHTTPHandlers,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initServer(ctx context.Context) error {
	a.linkShortenerImpl = linkShortenerV1.NewLinkShortenerV1(
		a.serviceProvider.GetLinkShortenerService(ctx),
	)

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer()
	desc.RegisterLinkShortenerV1Server(a.grpcServer, a.linkShortenerImpl)

	return nil
}

func (a *App) initPublicHTTPHandlers(ctx context.Context) error {
	a.mux = runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterLinkShortenerV1HandlerFromEndpoint(ctx, a.mux, grpcPort, opts)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runGRPC(wg *sync.WaitGroup) error {
	list, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return err
	}

	go func() {
		defer wg.Done()

		if err = a.grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to process gRPC server: %s", err.Error())
		}
	}()

	log.Printf("Run gRPC server on %s port\n", grpcPort)
	return nil
}

func (a *App) runPublicHTTP(wg *sync.WaitGroup) error {
	go func() {
		defer wg.Done()

		if err := http.ListenAndServe(httpPort, a.mux); err != nil {
			log.Fatalf("failed to process muxer: %s", err.Error())
		}
	}()

	log.Printf("Run public http handler on %s port\n", httpPort)
	return nil
}
