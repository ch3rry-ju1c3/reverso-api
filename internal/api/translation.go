package api

import (
	"encoding/json"
	//reverso_api "github.com/ch3rry-ju1c3/reverso-api"
)

type Translation struct {
}

func NewTranslation() *Translation {
	return &Translation{}
}

func (t *Translation) Translate(data *reverso_api.TranslationRequest) (*reverso_api.TranslationResponse, error) {
	res := reverso_api.TranslationResponse{}

	d, err := postRequest(translationPath, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(d, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
