package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/PlanktoScope/device-portal/internal/app/deviceportal"
	"github.com/PlanktoScope/device-portal/internal/app/deviceportal/conf"
)

const shutdownTimeout = 5 // sec

func main() {
	e := echo.New()

	// Get config
	config, err := conf.GetConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Prepare server
	server, err := deviceportal.NewServer(config, e.Logger)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err = server.Register(e); err != nil {
		e.Logger.Fatal(err)
	}

	// Run server
	ctxRun, cancelRun := signal.NotifyContext(
		context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT,
	)
	go func() {
		if err = server.Run(e); err != nil {
			e.Logger.Error(err)
		}
		cancelRun()
	}()
	<-ctxRun.Done()
	cancelRun()

	// Shut down server
	ctxShutdown, cancelShutdown := context.WithTimeout(
		context.Background(), shutdownTimeout*time.Second,
	)
	defer cancelShutdown()
	e.Logger.Infof("attempting to shut down gracefully within %d sec", shutdownTimeout)
	if err := server.Shutdown(ctxShutdown, e); err != nil {
		e.Logger.Warn("forcibly closing http server due to failure of graceful shutdown")
		if closeErr := server.Close(e); closeErr != nil {
			e.Logger.Error(closeErr)
		}
	}
	e.Logger.Info("finished shutdown")
}
