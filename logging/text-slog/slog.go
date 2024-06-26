package main

import (
	"log/slog"
	"os"
)

func main() {
	//slog.Config
	sloghander := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})
	//init slog for example
	Logger := slog.New(sloghander)
	//info level

	//Logger.Info("Info level"),and format message of slog by slog.Group
	Logger.Info("Info level", slog.Group(
		"votes",
		slog.Int64("count", 42),
		slog.String("name", "tifa"),
	))

	Logger.Warn("Warn level")

	Logger.Error("Error level")

	//Logger.Debug("debug level")

	Logger.Debug("what do you means?", slog.String("answer", "42"))

}
