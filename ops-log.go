// https://flaviocopes.com/go-shell-pipes/
// https://blog.golang.org/json-and-go

package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

var red = color.New(color.FgRed).SprintfFunc()
var green = color.New(color.FgGreen).SprintfFunc()
var yellow = color.New(color.FgYellow).SprintfFunc()
var white = color.New(color.FgWhite).SprintfFunc()
var black = color.New(color.FgHiBlack).SprintfFunc()

type loglevel int

const (
	logError loglevel = 50
	logWarn  loglevel = 40
	logInfo  loglevel = 30
	logDebug loglevel = 20
	logTrace loglevel = 10
)

func main() {
	var (
		flagNoColor       = flag.Bool("no-color", false, "Disable color output")
		flagFallThrough   = flag.Bool("fall-through", false, "Print unhandled lines")
		flagIgnoreAllMeta = flag.Bool("ignore-all-meta", false, "Do not print any metadata")
	)

	flag.Parse()

	if *flagNoColor {
		color.NoColor = true
	}

	info, err := os.Stdin.Stat()

	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: cmd | ops-log")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line, err := formatLogLine(scanner.Text(), flagIgnoreAllMeta)

		if err == nil {
			fmt.Println(line)
		} else {
			if *flagFallThrough {
				fmt.Println(scanner.Text())
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func formatLogLine(logLine string, flagIgnoreAllMeta *bool) (string, error) {
	var lineRaw interface{}

	err := json.Unmarshal([]byte(logLine), &lineRaw)

	if err != nil {
		return "", errors.New("can't parse line")
	}

	line := lineRaw.(map[string]interface{})

	var (
		timeFormatted = ""
		level         = ""
		message       = ""
		stack         = ""
		color         = white
		parts         []string
	)

	for k, v := range line {
		switch k {
		case "time":
			timeRaw := v.(float64)
			timeParsed, err := msToTime(strconv.FormatFloat(timeRaw, 'f', 0, 64))

			if err != nil {
				return "", fmt.Errorf("bad value for time: %f", timeRaw)
			}

			timeFormatted = timeParsed.UTC().Format(time.RFC3339)
		case "level":
			levelRaw := loglevel(v.(float64))

			switch levelRaw {
			case logError:
				level = "ERROR"
				color = red
			case logWarn:
				level = "WARN"
				color = yellow
			case logInfo:
				level = "INFO"
				color = green
			case logDebug:
				level = "DEBUG"
			case logTrace:
				level = "TRACE"
			}
		case "msg":
			message = v.(string)
		case "stack":
			stackRaw := v.(string)
			stack = fmt.Sprintf("\n%s", stackRaw)
		default:
			if *flagIgnoreAllMeta {
				continue
			}

			value := ""

			switch vv := v.(type) {
			case string:
				value = fmt.Sprintf("%s=%s", k, vv)
			case float64:
				parsed := strconv.FormatFloat(vv, 'f', 0, 64)
				value = fmt.Sprintf("%s=%s", k, parsed)
			default:
				continue
			}

			parts = append(parts, value)
		}
	}

	if timeFormatted != "" && level != "" {
		partsJoined := strings.Join(parts[:], " ")

		pieces := []string{
			black(timeFormatted),
			color("[%s]", level),
			message,
			black(partsJoined),
			stack,
		}

		return strings.Join(pieces[:], " "), nil
	}

	return "", errors.New("nothing to log")
}

func msToTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)

	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(0, msInt*int64(time.Millisecond)), nil
}
