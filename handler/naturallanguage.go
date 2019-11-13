package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/languagetranslatorv3"
	"github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
)

var (
	languageTranslator, languageTranslatorErr = languagetranslatorv3.NewLanguageTranslatorV3(&languagetranslatorv3.LanguageTranslatorV3Options{
		Version: os.Getenv("LANGUAGETRANSLATOR_VERSION"),
		Authenticator: &core.IamAuthenticator{
			ApiKey: os.Getenv("LANGUAGETRANSLATOR_APIKEY"),
		},
		URL: os.Getenv("LANGUAGETRANSLATOR_URL"),
	})

	nlu, nluErr = naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
		Version: os.Getenv("NATURALLANGUAGEUNDERSTANDING_VERSION"),
		Authenticator: &core.IamAuthenticator{
			ApiKey: os.Getenv("NATURALLANGUAGEUNDERSTANDING_APIKEY"),
		},
		URL: os.Getenv("NATURALLANGUAGEUNDERSTANDING_URL"),
	})
)

// GetEmotion from Watson Natural Language Understanding
func GetEmotion(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	text := r.URL.Query().Get("text")
	if text == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "No text provided. Send again you request with a valid text."})
		return
	}

	languageTranslatorResult, _, err := languageTranslator.Translate(&languagetranslatorv3.TranslateOptions{
		Text:    []string{text},
		ModelID: core.StringPtr("pt-en"),
	})

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Fail to extract the translation from your text. Try again later."})
		return
	}

	if len(languageTranslatorResult.Translations) == 0 {
		json.NewEncoder(w).Encode(map[string]string{"error": "No translation found. Try again later."})
		return
	}

	nluResult, _, err := nlu.Analyze(&naturallanguageunderstandingv1.AnalyzeOptions{
		Text: languageTranslatorResult.Translations[0].Translation,
		Features: &naturallanguageunderstandingv1.Features{
			Emotion:   &naturallanguageunderstandingv1.EmotionOptions{},
			Sentiment: &naturallanguageunderstandingv1.SentimentOptions{},
		},
	})
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Fail to extract the emotion from your text. Try again later."})
		return
	}
	json.NewEncoder(w).Encode(nluResult)
}
