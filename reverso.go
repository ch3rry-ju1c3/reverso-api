package reverso_api

import (
	api "github.com/ch3rry-ju1c3/reverso-api/internal/api"
)

type TranslateInterface interface {
	Translate(request *api.TranslationRequest) (*api.TranslationResponse, error)
}

type ReversoApi struct {
	Translation TranslateInterface
}

func NewReversoApi() *ReversoApi {

	return &ReversoApi{Translation: api.NewTranslation()}
}
