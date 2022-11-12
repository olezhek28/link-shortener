package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/olezhek28/link-shortener/internal/app/api/linkShortenerV1"
	"github.com/olezhek28/link-shortener/internal/logger"
	"github.com/olezhek28/link-shortener/internal/metrics"
	"github.com/olezhek28/link-shortener/internal/middleware"
	desc "github.com/olezhek28/link-shortener/pkg/link_shortener/v1"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

const (
	timeFormat = "2006-02-01"
)

type App struct {
	linkShortenerImpl *linkShortenerV1.Implementation
	serviceProvider   *serviceProvider

	grpcServer *grpc.Server
	httpServer *http.Server
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
	err := logger.Init(zapcore.DebugLevel, timeFormat)
	if err != nil {
		return err
	}

	inits := []func(context.Context) error{
		metrics.Init,
		a.initServiceProvider,
		a.initServer,
		a.initGRPCServer,
		a.initPublicHTTPHandlers,
	}

	for _, f := range inits {
		err = f(ctx)
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
	a.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)

	desc.RegisterLinkShortenerV1Server(a.grpcServer, a.linkShortenerImpl)

	return nil
}

func (a *App) initPublicHTTPHandlers(ctx context.Context) error {
	mux := runtime.NewServeMux()
	// TODO will be make auth
	// nolint
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// TODO will be make auth
	// nolint
	a.httpServer = &http.Server{
		Addr: a.serviceProvider.GetHTTPConfig().Host(),
		// add handler with middleware
		Handler: middleware.AddLogger(middleware.AddMetrics(mux)),
	}

	err := desc.RegisterLinkShortenerV1HandlerFromEndpoint(ctx, mux, a.serviceProvider.GetGRPCConfig().Host(), opts)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runGRPC(wg *sync.WaitGroup) error {
	list, err := net.Listen("tcp", a.serviceProvider.GetGRPCConfig().Host())
	if err != nil {
		return err
	}

	go func() {
		defer wg.Done()

		if err = a.grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to process gRPC server: %s", err.Error())
		}
	}()

	log.Printf("Run gRPC server on %s port\n", a.serviceProvider.GetGRPCConfig().Host())
	return nil
}

func (a *App) runPublicHTTP(wg *sync.WaitGroup) error {
	go func() {
		defer wg.Done()

		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("failed to process muxer: %s", err.Error())
		}
	}()

	log.Printf("Run public http handler on %s port\n", a.serviceProvider.GetHTTPConfig().Host())
	return nil
}
