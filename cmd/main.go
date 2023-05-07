package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	envKeyTranslateAPI = "GOOGLE_TRANSLATE_API_KEY"

	key    = flag.String("key", envString(envKeyTranslateAPI, ""), "google translate API key")
	target = flag.String("tl", envString("GOOGLE_TRANSLATE_TARGET_LANG", "sv"), "target language")
	source = flag.String("sl", envString("GOOGLE_TRANSLATE_SOURCE_LANG", "en"), "source language")
)

// Response represents the http response struct
type Response struct {
	Data struct {
		Translations []Translation `json:"translations"`
	} `json:"data"`
}

// Translation represents sub-struct of response
type Translation struct {
	TranslatedText string `json:"translatedText"`
}

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if *key == "" {
		log.Panicf("env variable %q not set", envKeyTranslateAPI)
	}

	v := make(url.Values)
	v.Set("key", *key)
	v.Set("target", *target)
	v.Set("source", *source)
	v.Set("q", strings.Join(flag.Args(), " "))

	resp, err := http.Get("https://www.googleapis.com/language/translate/v2?" + v.Encode())
	if err != nil {
		log.Panic("could not fetch translation")
	}

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Panic("could not translate, error code", resp.StatusCode)
	}

	var response Response
	err = decoder.Decode(&response)
	if err != nil {
		log.Panic("could not unmarshal json response")
	}

	for _, translation := range response.Data.Translations {
		fmt.Println(translation.TranslatedText)
	}
}

func envString(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
