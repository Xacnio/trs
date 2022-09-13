package cmd

import (
	"bufio"
	"log"
	"os"
	"strings"

	gt "github.com/bas24/googletranslatefree"
)

var Input = TranslateInfo{"en", "auto", "", ""}

func ParseCmd(cmd string) {
	if strings.HasPrefix(cmd, "s ") {
		a := strings.SplitN(cmd, " ", 2)
		Input.SetSourceLang(a[1])
		PrintScreen(false)
	} else if strings.HasPrefix(cmd, "t ") {
		a := strings.SplitN(cmd, " ", 2)
		Input.SetTargetLang(a[1])
		PrintScreen(false)
	} else {
		if cmd != "" {
			Input.SourceText = cmd
			UpdateTextProcess()
		} else {
			PrintScreen(false)
		}
	}
}

func UpdateTextProcess() {
	PrintScreen(true)
	result, _ := gt.Translate(Input.SourceText, Input.SourceLang, Input.TargetLang)
	Input.TargetText = result
	PrintScreen(false)
}

func CmdRun() {
	PrintScreen(false)
	for {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			ParseCmd(s.Text())
		}
		if s.Err() != nil {
			log.Fatal(s.Err())
		}
	}
}
