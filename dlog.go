// Package dlog implements a simple data logging package. 
package dlog

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Logger struct {
	mu        sync.Mutex
	metadata  map[string]interface{}
	out       *bufio.Writer
	sortOrder []string
}

func New(out io.Writer, metadata map[string]interface{}) *Logger {
	return &Logger{
		out:      bufio.NewWriter(out),
		metadata: metadata,
	}
}

func (l *Logger) Output(data map[string]interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for k, v := range l.metadata {
		data[k] = v
	}

	keys := make([]string, len(data))
	i := 0
	for k, _ := range data {
		keys[i] = k
		i++
	}
	sort.Sort(ByKeys{keys, l.sortOrder})

	l.writeData(keys, data)
	l.out.WriteByte('\n')
	l.out.Flush()
}

func (l *Logger) SortOrder(order []string) {
	l.sortOrder = order
}

func (l *Logger) writeData(keys []string, data map[string]interface{}) {
	fields := len(data)

	for i, k := range keys {
		v := data[k]
		s := stringEncode(v)

		if s != "" {
			if strings.Contains(s, " ") {
				s = strconv.Quote(s)
			}

			l.out.WriteString(fmt.Sprintf("%s=%v", k, s))
			if i < fields-1 {
				l.out.WriteByte(' ')
			}
		}
	}
}

func stringEncode(val interface{}) string {
	switch val.(type) {
	case time.Time:
		return val.(time.Time).UTC().Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", val)
}
