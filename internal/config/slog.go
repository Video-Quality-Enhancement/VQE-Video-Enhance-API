package config

import (
	"io"
	"os"
	"time"

	"golang.org/x/exp/slog"
)

func SlogSetupLogOutputFile() *os.File {

	date := time.Now().Format(time.DateOnly)
	fileName := "./logs/video_service_" + date + "_log.json"
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	wr := io.MultiWriter(os.Stdout, logFile)
	logger := slog.New(slog.NewJSONHandler(wr))
	slog.SetDefault(logger)

	return logFile

}
