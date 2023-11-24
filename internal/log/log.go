package logHandler

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

// Define logs folder
func init() {
	_, err := os.Stat("logs")
	if os.IsNotExist(err) {
		err = os.Mkdir("logs", os.ModePerm)
		if err != nil {
			fmt.Print("Error creating logs folder: ", err)
			panic(err)
		}
	}
}

var logger *slog.Logger

// customizeLogAttributes is a function that takes in a slice of strings and an attribute of type slog.Attr.
// It returns a slog.Attr with customized attributes based on the input attribute's key.
// If the key is slog.TimeKey, it replaces the key with "Fecha" and returns the original value.
// If the key is slog.SourceKey, it replaces the file attribute with the short file name and returns the original value.
// If the key is slog.LevelKey, it replaces the key with "Nivel" and returns the original value.
// If the key is slog.MessageKey, it replaces the key with "Mensaje" and returns the original value.
// If the key is not any of the above, it returns the original attribute.
func customizeLogAttributes(groups []string, a slog.Attr) slog.Attr {
	// Replace the time attribute with a value in Spanish
	if a.Key == slog.TimeKey {
		return slog.Attr{Key: "Fecha", Value: a.Value}
	}
	// Replace the file attribute with the short file name
	if a.Key == slog.SourceKey {
		source := a.Value.Any().(*slog.Source)
		pathSegments := strings.Split(source.File, "/")
		if len(pathSegments) > 2 {
			source.File = strings.Join(pathSegments[len(pathSegments)-2:], "/")
		}
		// Replace the file attribute with a value in Spanish
		return slog.Attr{Key: "Ruta", Value: a.Value}
	}

	// Replace the level attribute with a value in Spanish
	if a.Key == slog.LevelKey {
		return slog.Attr{Key: "Nivel", Value: a.Value}
	}

	// Replace the message attribute with a value in Spanish
	if a.Key == slog.MessageKey {
		return slog.Attr{Key: "Mensaje", Value: a.Value}
	}

	return a
}

// GetInstance returns a pointer to a log.Logger instance that writes to both the log file and os.Stdout.
// The log file is located at "logs/logs.log" and has read, write, and append permissions.
// The logger prefix is "WebSocket: " and it includes the date, time, and short file name in the log message.
func GetInstance() *slog.Logger {
	if logger == nil {
		logFile, cantCreateInfoLogFile := os.OpenFile("logs/logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if cantCreateInfoLogFile != nil {
			fmt.Print("Error creando archivo de logs: ", cantCreateInfoLogFile)
			panic(cantCreateInfoLogFile)
		}
		logger = slog.New(slog.NewTextHandler(io.MultiWriter(logFile, os.Stdout), &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true, ReplaceAttr: customizeLogAttributes}))
	}
	return logger
}
