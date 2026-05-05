package formatter_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/logslice/logslice/internal/formatter"
	"github.com/logslice/logslice/internal/parser"
)

var ts = time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)

func entry(level, msg string, fields map[string]any) parser.Entry {
	return parser.Entry{
		Timestamp: ts,
		Level:     level,
		Message:   msg,
		Fields:    fields,
		Raw:       `{"level":"` + level + `","msg":"` + msg + `"}`,
	}
}

func TestFormatter_Pretty(t *testing.T) {
	var buf bytes.Buffer
	f := formatter.New(&buf, formatter.FormatPretty, false, time.RFC3339)
	e := entry("info", "server started", map[string]any{"port": 8080})
	if err := f.Write(e); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "[INFO]") {
		t.Errorf("expected [INFO] in output, got: %s", out)
	}
	if !strings.Contains(out, "server started") {
		t.Errorf("expected message in output, got: %s", out)
	}
	if !strings.Contains(out, "port=8080") {
		t.Errorf("expected field in output, got: %s", out)
	}
}

func TestFormatter_JSON(t *testing.T) {
	var buf bytes.Buffer
	f := formatter.New(&buf, formatter.FormatJSON, false, time.RFC3339)
	e := entry("error", "boom", nil)
	if err := f.Write(e); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, `"level":"error"`) {
		t.Errorf("expected level field, got: %s", out)
	}
	if !strings.Contains(out, `"msg":"boom"`) {
		t.Errorf("expected msg field, got: %s", out)
	}
}

func TestFormatter_Raw(t *testing.T) {
	var buf bytes.Buffer
	f := formatter.New(&buf, formatter.FormatRaw, false, "")
	e := entry("debug", "raw line", nil)
	e.Raw = "original raw content"
	if err := f.Write(e); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "original raw content") {
		t.Errorf("expected raw content, got: %s", buf.String())
	}
}

func TestFormatter_NoTimestamp(t *testing.T) {
	var buf bytes.Buffer
	f := formatter.New(&buf, formatter.FormatPretty, false, time.RFC3339)
	e := parser.Entry{Level: "warn", Message: "no time", Fields: map[string]any{}}
	if err := f.Write(e); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if strings.HasPrefix(out, "2024") {
		t.Errorf("unexpected timestamp in output: %s", out)
	}
}
