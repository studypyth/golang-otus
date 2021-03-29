package main

import (
	"context"
	"flag"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/studypyth/hw12_13_14_15_calendar/internal/app"
	"github.com/studypyth/hw12_13_14_15_calendar/internal/config"
	"github.com/studypyth/hw12_13_14_15_calendar/internal/logger"
	"github.com/studypyth/hw12_13_14_15_calendar/internal/storage"
)

var configFile string
var version bool

func init() {
	flag.StringVar(&configFile, "config", "", "Path to configuration file")
	flag.BoolVar(&version, "version", false, "Print version")
}

func main() {
	flag.Parse()

	if version {
		printVersion()
		return
	}

	cfg := config.New(configFile)
	if os.Getenv("DEBUG") == "TRUE" {
		cfg.Logger.Level = 0
	}
	logg := logger.New(cfg.Logger)
	stor, err := storage.New(cfg.Storage)
	if err != nil {
		logg.ErrorMsg(err.Error())
	}
	calendar := app.New(logg, stor)
	calendar.CreateEvent(context.Background(), uuid.New().String(), "BDay", "Bday Dilya", "damir", time.Now(), 12*time.Hour, 1*time.Hour)
	calendar.CreateEvent(context.Background(), uuid.New().String(), "BDay", "Bday Damir", "damir", time.Now().Add(-1*(3*(time.Hour))), 12*time.Hour, 1*time.Hour)

	// Тестирование  Inmemorystorage
	//empJSON, _ := json.MarshalIndent(calendar.Storage, "", "  ")
	//fmt.Printf("%s\n", string(empJSON))
	//fmt.Printf("%v\n", calendar.Storage)

	//server := internalhttp.NewServer(calendar)
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//
	//go func() {
	//	signals := make(chan os.Signal, 1)
	//	signal.Notify(signals, syscall.SIGINT, syscall.SIGHUP)
	//
	//	<-signals
	//	signal.Stop(signals)
	//	cancel()
	//
	//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	//	defer cancel()
	//
	//	if err := server.Stop(ctx); err != nil {
	//		log.Error().Msg("failed to stop http server: " + err.Error())
	//	}
	//}()
	//
	//log.InfoMsg("calendar is running...")
	//
	//if err := server.Start(ctx); err != nil {
	//	log.Error().Msg("failed to start http server: " + err.Error())
	//	os.Exit(1)
	//}
}
