package main

//TODO: Redo and use git
import (
	"context"
	"fmt"
	"time"

	"github.com/mboufous/k-island/internal/kube"
	"github.com/mboufous/k-island/pkg/healthcheck"
	"github.com/mboufous/k-island/pkg/server"

	"github.com/mboufous/k-island/utils/config"
	"github.com/mboufous/k-island/utils/log"
	"github.com/mboufous/k-island/utils/signal"

	_ "go.uber.org/automaxprocs/maxprocs"
)

const appName = "KISLAND"

func main() {

	if err := run(); err != nil {
		log.Fatalf("%s", err)
	}

}

func run() error {

	//----------------------------
	// Configuration
	conf, err := config.LoadConfig(appName)
	if err != nil {
		return err
	}

	//----------------------------
	// Logger

	err = log.NewZapLogger(&conf.Log)
	if err != nil {
		return err
	}

	//----------------------------
	// API
	kubeService, err := kube.NewService()
	if err != nil {
		return err
	}

	kubeApi := kube.NewKubeAPI(kubeService)
	healthcheck.Setup(kubeApi.Router)
	kubeApi.V1()

	apiServer := server.New(&server.Config{
		Host:    conf.Server.Host,
		Port:    conf.Server.Port,
		Handler: kubeApi.APIHandler(),
	})

	apiServer.Start()

	//----------------------------
	// Register signal handler and wait
	signal.Stop.OnSignal(signal.DefaultStopSignals...)
	<-signal.Stop.Chan()

	// Shutdown the server gracefully
	shutdownTimeout, _ := time.ParseDuration(conf.Server.ShutdownTimeout)
	log.Infof("start gracefull shutdown ...timeout after %s", shutdownTimeout)

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer shutdownCancel()

	if err := apiServer.Shutdown(shutdownCtx); err != nil {
		apiServer.Close()
		log.Warn("Forcing shutdown after timeout")
		log.Flush()
		return fmt.Errorf("shutdown: %w", err)
	}

	log.Flush()
	log.Info("server shutdown successfully")
	return nil

}
