package cmd

import (
	"fmt"

	ct "github.com/daviddengcn/go-colortext"
	"github.com/inancgumus/screen"
)

func HelpText() string {
	return "Cmds: s <source language> / t <target language> / <source text>"
}

func PrintScreen(loading bool) {
	screen.Clear()
	screen.MoveTopLeft()
	ct.Foreground(ct.Red, false)
	fmt.Printf("%s » %s\n\n", Input.SourceLang, Input.TargetLang)
	ct.Foreground(ct.Green, false)
	fmt.Printf("Source Text: %s\n", Input.SourceText)
	ct.Foreground(ct.Cyan, false)
	if loading {
		fmt.Printf("Target Text: ...\n")
	} else {
		fmt.Printf("Target Text: %s\n", Input.TargetText)
	}
	ct.Foreground(ct.Yellow, false)
	fmt.Printf("\n\n%s\n", HelpText())
	fmt.Print("> ")
}
