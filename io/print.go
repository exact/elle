package io

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"os"
)

type lgr struct {
	b *bytes.Buffer
	w io.Writer
}

func (i lgr) Write(p []byte) (int, error) {
	i.b.Reset()
	json.Indent(i.b, p, "", "\t")
	n, err := i.b.WriteTo(i.w)
	return int(n), err
}

var Log *slog.Logger = slog.New(slog.NewJSONHandler(&lgr{new(bytes.Buffer), os.Stdout}, nil))
