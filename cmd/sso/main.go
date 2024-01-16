package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sso/internal/app"
	"sso/internal/config"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger()
	log.Info("Starting SSO server")

	application := app.New(log, cfg.GRPC.Port)
	go application.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop
	log.Info("Received stop signal", slog.String("signal", sign.String()))
	application.GRPCSrv.Stop()
}

func setupLogger() *slog.Logger {
	var log *slog.Logger
	log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	return log
}
