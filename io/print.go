package io

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"os"
)

type logger struct {
	b *bytes.Buffer
	w io.Writer
}

func (i logger) Write(p []byte) (int, error) {
	i.b.Reset()
	json.Indent(i.b, p, "", "\t")
	n, err := i.b.WriteTo(i.w)
	return int(n), err
}

var log *slog.Logger = slog.New(slog.NewJSONHandler(&logger{new(bytes.Buffer), os.Stdout}, &slog.HandlerOptions{
	ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			return slog.String("time", a.Value.Time().Format("01/02/2006 15:04:05.000 PM"))
		}
		return a
	},
}))

// Puts writes a to standard output, and adds a newline at the end.
//
// It returns amount of bytes written and any write error encountered.
func Log(msg string, a ...any) {
	log.Info(msg, a...)
}

// Warn writes a to standard error, and adds a newline at the end.
//
// It returns amount of bytes written and any write error encountered.
func Warn(msg string, a ...any) {
	log.Warn(msg, a...)
}

func Err(msg string, a ...any) {
	log.Error(msg, a...)
}
