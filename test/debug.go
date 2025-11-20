package main

import (
	"fmt"
	"strings"

	kpsclient "github.com/Ctere1/kpsclient"
)

// PrintKPSResultBox pretty-prints the KPS query result in a colorized, boxed format to stdout.
func PrintKPSResultBox(res kpsclient.Result) {
	// Color definitions
	const (
		red       = "\x1b[31m"
		green     = "\x1b[32m"
		yellow    = "\x1b[33m"
		blue      = "\x1b[34m"
		magenta   = "\x1b[35m"
		cyan      = "\x1b[36m"
		reset     = "\x1b[0m"
		bold      = "\x1b[1m"
		italic    = "\x1b[3m"
		underline = "\x1b[4m"
	)

	// Box width
	const boxWidth = 68
	boxContentWidth := boxWidth - 4

	var output strings.Builder
	output.WriteString(fmt.Sprintf("\n%s┌%s┐%s\n", magenta, strings.Repeat("─", boxWidth-2), reset))
	plain := fmt.Sprintf("%-*s", boxContentWidth, "KPS QUERY RESULT")
	output.WriteString(fmt.Sprintf("%s│ %s%s %s│%s\n", magenta, bold+yellow, plain, magenta, reset))
	output.WriteString(fmt.Sprintf("%s├%s┤%s\n", magenta, strings.Repeat("─", boxWidth-2), reset))

	statusColor := green
	if !res.Status {
		statusColor = red
	}
	statusLine := fmt.Sprintf("Status: %v", res.Status)
	codeLine := fmt.Sprintf("Code: %d", res.Code)
	messageLine := fmt.Sprintf("Message: %s", res.Message)
	personLine := fmt.Sprintf("Person Type: %s", res.Person)

	output.WriteString(fmt.Sprintf("%s│ %s%-*s %s│%s\n", magenta, statusColor, boxContentWidth, statusLine, magenta, reset))
	output.WriteString(fmt.Sprintf("%s│ %s%-*s %s│%s\n", magenta, cyan, boxContentWidth, codeLine, magenta, reset))
	output.WriteString(fmt.Sprintf("%s│ %s%-*s %s│%s\n", magenta, yellow, boxContentWidth, messageLine, magenta, reset))
	output.WriteString(fmt.Sprintf("%s│ %s%-*s %s│%s\n", magenta, blue, boxContentWidth, personLine, magenta, reset))
	output.WriteString(fmt.Sprintf("%s├%s┤%s\n", magenta, strings.Repeat("─", boxWidth-2), reset))

	output.WriteString(fmt.Sprintf("%s│ %s%-*s %s│%s\n", magenta, bold+blue, boxContentWidth, "Extra Fields", magenta, reset))
	for k, v := range res.Extra {
		fieldLine := fmt.Sprintf("  %s: %s", k, v)
		output.WriteString(fmt.Sprintf("%s│ %s%-*s %s│%s\n", magenta, cyan, boxContentWidth, fieldLine, magenta, reset))
	}
	output.WriteString(fmt.Sprintf("%s├%s┤%s\n", magenta, strings.Repeat("─", boxWidth-2), reset))

	// Show raw response header, but truncate unless verbose
	output.WriteString(fmt.Sprintf("%s│ %s%-*s %s│%s\n", magenta, bold+yellow, boxContentWidth, "Raw SOAP/XML (first 500 chars)", magenta, reset))
	raw := res.Raw
	if len(raw) > 500 {
		raw = raw[:500] + "..."
	}
	for len(raw) > 0 {
		line := raw
		if len(line) > boxContentWidth {
			line = line[:boxContentWidth]
		}
		output.WriteString(fmt.Sprintf("%s│ %-*s %s│%s\n", magenta, boxContentWidth, line, magenta, reset))
		raw = raw[len(line):]
	}
	output.WriteString(fmt.Sprintf("%s└%s┘%s\n", magenta, strings.Repeat("─", boxWidth-2), reset))

	fmt.Print(output.String())
}
