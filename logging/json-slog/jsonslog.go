package main

import (
	"log/slog"
	"os"
	"time"
)

func main() {

	sloghander := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			//match value if you want
			if a.Key == slog.TimeKey {
				a.Key = "date"
				a.Value = slog.Int64Value(int64(time.Now().Unix()))
			}
			return a
		},
	}).WithAttrs([]slog.Attr{ //second way to format message
		slog.String("name", "tifa"),
		slog.Int64("joson", 42)})
	//init slog for example
	Logger := slog.New(sloghander)
	//info level

	/*//Logger.Info("Info level"),and format message of slog by slog.Group, first way to format message
	Logger.Info("Info level", slog.Group(
		"votes",
		slog.Int64("joson", 42),
		slog.String("name", "tifa"),
	))
	*/
	slog.Info("my logger")
	slog.SetDefault(Logger)
	Logger.Debug("votes")
}
