package logger

import (
	"log/slog"
	"os"
)

func MustInitLogger() (*slog.Logger, *slog.Logger) {
	logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("не удалось создать или открыть файл лога:" + err.Error())
	}
	fileHandler := slog.NewTextHandler(logFile, nil)
	textHandler := slog.NewTextHandler(os.Stdout, nil)

	logger := slog.New(textHandler)
	loggerFile := slog.New(fileHandler)

	return logger, loggerFile
}
