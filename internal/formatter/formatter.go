package formatter

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/logslice/logslice/internal/parser"
)

// Format controls the output format of log entries.
type Format string

const (
	// FormatPretty renders a human-readable, colourised single line.
	FormatPretty Format = "pretty"
	// FormatJSON re-encodes the entry as compact JSON.
	FormatJSON Format = "json"
	// FormatRaw passes the original raw line through unchanged.
	FormatRaw Format = "raw"
)

// levelColors maps log level strings to ANSI colour codes.
var levelColors = map[string]string{
	"debug": "\033[36m",  // cyan
	"info":  "\033[32m",  // green
	"warn":  "\033[33m",  // yellow
	"warning": "\033[33m",
	"error": "\033[31m",  // red
	"fatal": "\033[35m",  // magenta
}

const colorReset = "\033[0m"

// Formatter writes formatted log entries to an io.Writer.
type Formatter struct {
	format  Format
	color   bool
	writer  io.Writer
	timeFmt string
}

// New returns a Formatter configured with the given options.
func New(w io.Writer, format Format, color bool, timeFmt string) *Formatter {
	if timeFmt == "" {
		timeFmt = time.RFC3339
	}
	return &Formatter{format: format, color: color, writer: w, timeFmt: timeFmt}
}

// Write formats and writes a single parsed entry to the underlying writer.
func (f *Formatter) Write(e parser.Entry) error {
	var line string
	switch f.format {
	case FormatJSON:
		line = f.renderJSON(e)
	case FormatRaw:
		line = e.Raw
	default:
		line = f.renderPretty(e)
	}
	_, err := fmt.Fprintln(f.writer, line)
	return err
}

func (f *Formatter) renderPretty(e parser.Entry) string {
	var sb strings.Builder

	if !e.Timestamp.IsZero() {
		sb.WriteString(e.Timestamp.Format(f.timeFmt))
		sb.WriteByte(' ')
	}

	level := strings.ToUpper(e.Level)
	if f.color {
		if code, ok := levelColors[strings.ToLower(e.Level)]; ok {
			level = code + level + colorReset
		}
	}
	if level != "" {
		fmt.Fprintf(&sb, "[%s] ", level)
	}

	sb.WriteString(e.Message)

	for k, v := range e.Fields {
		fmt.Fprintf(&sb, " %s=%v", k, v)
	}
	return sb.String()
}

func (f *Formatter) renderJSON(e parser.Entry) string {
	m := make(map[string]any, len(e.Fields)+3)
	for k, v := range e.Fields {
		m[k] = v
	}
	if !e.Timestamp.IsZero() {
		m["time"] = e.Timestamp.Format(f.timeFmt)
	}
	if e.Level != "" {
		m["level"] = e.Level
	}
	if e.Message != "" {
		m["msg"] = e.Message
	}
	b, err := json.Marshal(m)
	if err != nil {
		return e.Raw
	}
	return string(b)
}
