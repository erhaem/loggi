package loggi

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

func log(level string, clr func(a ...any) string, msg string) {
	tstamp := time.Now().Format("2006/01/02 15:04:05")
	line := fmt.Sprintf("[%s] [%s] %s", tstamp, level, msg)

	out := os.Stdout
	if level == "ERR" {
		out = os.Stderr
	}
	fmt.Fprintln(out, clr(line))
}

func Info(msg string)  { log("INF", color.New(color.FgGreen).SprintFunc(), msg) }
func Warn(msg string)  { log("WRN", color.New(color.FgYellow).SprintFunc(), msg) }
func Error(msg string) { log("ERR", color.New(color.FgRed).SprintFunc(), msg) }
func Debug(msg string) { log("DBG", color.New(color.FgWhite).SprintFunc(), msg) }
