package config

import (
	"io"
	"os"
	"time"

	"golang.org/x/exp/slog"
)

func SetupSlogOutputFile() *os.File {

	date := time.Now().Format(time.DateOnly)
	serviceName := os.Getenv("SERVICE_NAME")
	fileName := "./logs/" + serviceName + "_" + date + "_log.json"
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	wr := io.MultiWriter(os.Stdout, logFile)
	logger := slog.New(slog.NewJSONHandler(wr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	return logFile

}

// WriteFunc convert func to io.Writer.
type WriteFunc func([]byte) (int, error)

func (fn WriteFunc) Write(data []byte) (int, error) {
	return fn(data)
}

func NewSlogInfoWriter() io.Writer {
	return WriteFunc(func(data []byte) (int, error) {
		slog.Info(string(data))
		return 0, nil
	})
}

func NewSlogErrorWriter() io.Writer {
	return WriteFunc(func(data []byte) (int, error) {
		slog.Error(string(data))
		return 0, nil
	})
}
