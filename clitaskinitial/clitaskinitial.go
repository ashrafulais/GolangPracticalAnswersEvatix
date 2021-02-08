package createinitialtask

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"time"
)

// "Title,Message 1,Message 2,Stream Delay,Run Times\nCLI Invoker Name,First Message,Second Msg,2,10"
type CliStreamerRecord struct {
	Title       string `csv:"Title"`
	Message1    string `csv:"Message 1"`
	Message2    string `csv:"Message 2"`
	StreamDelay int    `csv:"Stream Delay"`
	RunTimes    int    `csv:"Run Times"`
}

type CliRunnerRecord struct {
	// How many streamer will run.
	Run         string `csv:"Run"`
	Title       string `csv:"Title"`
	Message1    string `csv:"Message 1"`
	Message2    string `csv:"Message 2"`
	StreamDelay int    `csv:"Stream Delay"`
	RunTimes    int    `csv:"Run Times"`
}

func (cliRunnerRecord CliRunnerRecord) CliStreamerRecord() CliStreamerRecord {
	return CliStreamerRecord{
		Title:       cliRunnerRecord.Title,
		Message1:    cliRunnerRecord.Message1,
		Message2:    cliRunnerRecord.Message2,
		StreamDelay: cliRunnerRecord.StreamDelay,
		RunTimes:    cliRunnerRecord.RunTimes,
	}
}

func (cliRunnerRecord CliRunnerRecord) CliStreamerRecordCsv() string {
	cliStreamerRecords := []CliStreamerRecord{cliRunnerRecord.CliStreamerRecord()}

	out, err := gocsv.MarshalString(cliStreamerRecords)

	if err != nil {
		panic(err)
	}

	return out
}

func Csv(cliRunners *[]CliRunnerRecord) string {
	out, err := gocsv.MarshalString(cliRunners)

	if err != nil {
		panic(err)
	}

	return out
}

func AsyncRunnerRecord(cliRecord CliRunnerRecord) {
	for i := 1; i <= cliRecord.RunTimes; i++ {
		fmt.Println(cliRecord.Title, "->", cliRecord.Message1)
		time.Sleep(time.Duration(cliRecord.StreamDelay) * 1000 * time.Millisecond)
		fmt.Println(cliRecord.Message2)
		time.Sleep(time.Duration(cliRecord.StreamDelay) * 1000 * time.Millisecond)
	}

}

func ExecuteInitialTask() {
	args := "Title,Message 1,Message 2,Stream Delay,Run Times\nCLI Invoker Name,First Message,Second Msg,2,10"
	var cliStreamers []CliStreamerRecord
	gocsv.UnmarshalString(
		args,
		&cliStreamers)

	fmt.Print(gocsv.MarshalString(cliStreamers))
}

func ExecuteInitialTask_Detailed() {
	//args := os.Args[1] requires os package to import
	args := "Run,Title,Message 1,Message 2,Stream Delay,Run Times\n1,CLI Invoke1,First Msg 1,Second Msg 2,2,5\n2,CLI Invoke2,First Msg 1,Second Msg 2,2,3"
	var cliRunners []CliRunnerRecord
	gocsv.UnmarshalString(
		args,
		&cliRunners)

	fmt.Print(Csv(&cliRunners))
	fmt.Println("---------------------------------")
	for _, runner := range cliRunners {
		//fmt.Println(i, ":")
		//fmt.Println(runner.RunTimes)
		go func(cliRecord CliRunnerRecord) {
			for i := 1; i <= cliRecord.RunTimes; i++ {
				fmt.Println(cliRecord.Title, "->", cliRecord.Message1)
				time.Sleep(time.Duration(cliRecord.StreamDelay) * 1000 * time.Millisecond)
				fmt.Println(cliRecord.Message2)
				time.Sleep(time.Duration(cliRecord.StreamDelay) * 1000 * time.Millisecond)
			}
		}(runner)
	}
}
