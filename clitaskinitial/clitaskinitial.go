package createinitialtask

import (
	"fmt"

	"github.com/gocarina/gocsv"
)

// "Title,Message 1,Message 2,Stream Delay,Run Times\nCLI Invoker Name,First Message,Second Msg,2,10"
type CliStreamerRecord struct {
	Title       string `csv:"Title"`
	Message1    string `csv:"Message 1"`
	Message2    string `csv:"Message 2"`
	StreamDelay int    `csv:"Stream Delay"`
	RunTimes    int    `csv:"Run Times"`
}

func ExecuteInitialTask() {
	args := "Title,Message 1,Message 2,Stream Delay,Run Times\nCLI Invoker Name,First Message,Second Msg,2,10"
	var cliStreamers []CliStreamerRecord
	gocsv.UnmarshalString(
		args,
		&cliStreamers)

	fmt.Print(gocsv.MarshalString(cliStreamers))
}
