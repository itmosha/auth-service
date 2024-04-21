package logger

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/itmosha/auth-service/internal/config"
)

type Logger struct {
	log *slog.Logger
}

func NewLogger(filePath string, env config.EnvType) *Logger {
	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}

	var log *slog.Logger
	switch env {
	case config.EnvLocal:
		log = slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case config.EnvProd:
		log = slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return &Logger{log}
}

func (l *Logger) LogRequest(r *http.Request, status int, err error) {
	requestInfo := []interface{}{
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.Int("status", status),
	}

	if status >= 400 {
		if err != nil {
			requestInfo = append(requestInfo, slog.String("error", err.Error()))
		}
		l.log.Error("API request", requestInfo...)
	} else {
		l.log.Info("API request", requestInfo...)
	}
}

func (l *Logger) LogError(op string, err error) {
	l.log.Error(op, slog.String("error", err.Error()))
}
