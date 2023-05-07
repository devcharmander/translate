# google-translate CLI

A command line interface to translate text to target language.


> This CLI needs you to create an API key in cloud.google.com

## Install 

``` go
go install github.com/devcharmander/translate@latest
```

then,

``` sh
translate "hello world"

# result
Hej världen
```

Alternatively, 

#### Run the project 

```go
go run cmd/* "Hello"
```

### Env variables

- `GOOGLE_TRANSLATE_API_KEY` : cloud translation API key 
- `GOOGLE_TRANSLATE_TARGET_LANG` : target language - default Swedish(sv)
- `GOOGLE_TRANSLATE_SOURCE_LANG` : source language - default English(en)

#### Flags

- `key` : cloud translation api key 
- `tl` : target language
- `sl` : source language


``` go
go run cmd/* -tl en -sl ja "こんにちは世界"
# result
Hello world
```
