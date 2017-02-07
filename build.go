package main

import "bytes"

var (
	// Date build time
	Date string
	// Commit latest git commit id
	Commit string
)

// Version 生成版本信息
func Version(prefix string) string {
	var buf bytes.Buffer
	if prefix != "" {
		buf.WriteString(prefix)
	}
	if Date != "" {
		buf.WriteByte('\n')
		buf.WriteString("date: ")
		buf.WriteString(Date)
	}
	if Commit != "" {
		buf.WriteByte('\n')
		buf.WriteString("commit: ")
		buf.WriteString(Commit)
	}
	return buf.String()
}
