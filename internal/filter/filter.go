// Package filter provides log entry filtering based on field matchers.
package filter

import (
	"regexp"
	"strings"
	"time"
)

// Entry represents a parsed log entry passed to the filter.
type Entry struct {
	Timestamp time.Time
	Level     string
	Message   string
	Fields    map[string]string
	Raw       string
}

// Filter holds compiled filter criteria.
type Filter struct {
	Levels    []string
	MsgRegexp *regexp.Regexp
	Since     time.Time
	Until     time.Time
	Fields    map[string]string
}

// Options configures a Filter.
type Options struct {
	Levels    []string
	MsgPattern string
	Since      time.Time
	Until      time.Time
	Fields     map[string]string
}

// New creates a Filter from Options, compiling any regex patterns.
func New(opts Options) (*Filter, error) {
	f := &Filter{
		Levels: opts.Levels,
		Since:  opts.Since,
		Until:  opts.Until,
		Fields: opts.Fields,
	}
	if opts.MsgPattern != "" {
		re, err := regexp.Compile(opts.MsgPattern)
		if err != nil {
			return nil, err
		}
		f.MsgRegexp = re
	}
	return f, nil
}

// Match returns true if the entry satisfies all filter criteria.
func (f *Filter) Match(e Entry) bool {
	if len(f.Levels) > 0 && !containsIgnoreCase(f.Levels, e.Level) {
		return false
	}
	if f.MsgRegexp != nil && !f.MsgRegexp.MatchString(e.Message) {
		return false
	}
	if !f.Since.IsZero() && e.Timestamp.Before(f.Since) {
		return false
	}
	if !f.Until.IsZero() && e.Timestamp.After(f.Until) {
		return false
	}
	for k, v := range f.Fields {
		if e.Fields[k] != v {
			return false
		}
	}
	return true
}

func containsIgnoreCase(list []string, val string) bool {
	for _, s := range list {
		if strings.EqualFold(s, val) {
			return true
		}
	}
	return false
}
