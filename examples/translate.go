package main

import (
	"fmt"

	reverso_api "github.com/ch3rry-ju1c3/reverso-api"
	api "github.com/ch3rry-ju1c3/reverso-api/internal/api"
)

func main() {
	r := reverso_api.NewReversoApi()

	req := api.TranslationRequest{
		Format: "text",
		From:   reverso_api.ENG,
		To:     reverso_api.RU,
		Input:  "Hello world!",
		Options: api.Options{
			SentenceSplitter:  true,
			Origin:            "translation.web",
			ContextResults:    false,
			LanguageDetection: false,
		},
	}

	t, err := r.Translation.Translate(&req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(t)
}
