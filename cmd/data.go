package cmd

import "github.com/xacnio/trs/utils"

type TranslateInfo struct {
	TargetLang string
	SourceLang string
	SourceText string
	TargetText string
}

func (t *TranslateInfo) SetSourceLang(lang string) {
	if _, ok := utils.Langs[lang]; ok {
		t.SourceLang = lang
	}
}

func (t *TranslateInfo) SetTargetLang(lang string) {
	if lang == "auto" {
		return
	}
	if _, ok := utils.Langs[lang]; ok {
		t.TargetLang = lang
	}
}
