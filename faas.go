// Package faas implements function as a Services
package faas

import (
	"net/http"

	"github.com/blacknaml/hello-api/handlers/rest"
	"github.com/blacknaml/hello-api/translation"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)
	translateHandler.TranslateHandler(w, r)
}
