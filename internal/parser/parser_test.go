package parser_test

import (
	"testing"
	"time"

	"github.com/logslice/logslice/internal/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse_JSONEntry(t *testing.T) {
	line := `{"level":"info","msg":"server started","time":"2024-03-01T12:00:00Z","port":8080}`
	entry, err := parser.Parse(line)
	require.NoError(t, err)

	assert.Equal(t, "info", entry.Level)
	assert.Equal(t, "server started", entry.Message)
	assert.False(t, entry.Timestamp.IsZero())
	assert.Equal(t, float64(8080), entry.Fields["port"])
	assert.Equal(t, line, entry.Raw)
}

func TestParse_AlternativeKeys(t *testing.T) {
	line := `{"severity":"error","message":"disk full","ts":"2024-03-01T12:00:00Z"}`
	entry, err := parser.Parse(line)
	require.NoError(t, err)

	assert.Equal(t, "error", entry.Level)
	assert.Equal(t, "disk full", entry.Message)
	assert.False(t, entry.Timestamp.IsZero())
}

func TestParse_PlainText(t *testing.T) {
	line := "plain text log line"
	entry, err := parser.Parse(line)
	require.NoError(t, err)

	assert.Equal(t, "plain text log line", entry.Message)
	assert.Equal(t, "", entry.Level)
	assert.True(t, entry.Timestamp.IsZero())
}

func TestParse_EmptyLine(t *testing.T) {
	_, err := parser.Parse("")
	assert.Error(t, err)
}

func TestParse_TimestampParsing(t *testing.T) {
	line := `{"msg":"ok","time":"2024-06-15T08:30:00.123456789Z"}`
	entry, err := parser.Parse(line)
	require.NoError(t, err)

	expected := time.Date(2024, 6, 15, 8, 30, 0, 123456789, time.UTC)
	assert.Equal(t, expected, entry.Timestamp)
}

func TestParse_ExtraFieldsPreserved(t *testing.T) {
	line := `{"level":"debug","msg":"trace","traceID":"abc123","user":"alice"}`
	entry, err := parser.Parse(line)
	require.NoError(t, err)

	assert.Equal(t, "abc123", entry.Fields["traceID"])
	assert.Equal(t, "alice", entry.Fields["user"])
	// level and msg should not appear in Fields
	_, hasLevel := entry.Fields["level"]
	assert.False(t, hasLevel)
}
