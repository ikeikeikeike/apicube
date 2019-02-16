package es

// Schema returns a defined schema mapping
func Schema(name string) string {
	switch name {
	default:
		return ""
	case ProductsName:
		return productsIndex
	// case FavoritesName:
	// return favoritesIndex
	}
}

const (
	commonSettings = `
  "settings": {
    "analysis": {
      "tokenizer": {
        "ngram_tokenizer": {
          "type": "nGram",
          "token_chars": [
            "letter",
            "digit"
          ],
          "min_gram": "2",
          "max_gram": "3"
        },
        "unigram_tokenizer": {
          "type": "nGram",
          "token_chars": [
            "letter",
            "digit"
          ],
          "min_gram": "1",
          "max_gram": "2"
        },
        "kuromoji_pserson_dic": {
          "type": "kuromoji_neologd_tokenizer"
        },
        "ja_tokenizer": {
          "type": "kuromoji_neologd_tokenizer",
          "mode": "search"
        }
      },
      "filter": {
        "split_delimiter": {
          "type": "word_delimiter",
          "stem_english_possessive": false,
          "split_on_numerics": false,
          "split_on_case_change": false,
          "preserve_original": false,
          "generate_word_parts": true,
          "generate_number_parts": false,
          "catenate_words": false,
          "catenate_numbers": false,
          "catenate_all": false
        },
        "katakana_readingform": {
          "use_romaji": false,
          "type": "kuromoji_neologd_readingform"
        },
        "ja_posfilter": {
          "type": "kuromoji_neologd_part_of_speech",
          "stoptags": [
            "助詞-格助詞-一般",
            "助詞-終助詞"
          ]
        }
      },
      "analyzer": {
        "rubytext_analyzer": {
          "type": "custom",
          "tokenizer": "kuromoji_pserson_dic",
		  "char_filter": ["html_strip"],  ` + /* TODO: html_stip seems like not working: will be made sure working */ `
          "filter": [
            "katakana_readingform",
            "split_delimiter"
          ]
        },
        "unigram_analyzer": {
          "type": "custom",
		  "char_filter": ["html_strip"],  ` + /* TODO: html_stip seems like not working: will be made sure working */ `
          "tokenizer": "unigram_tokenizer"
        },
        "ngram_analyzer": {
          "type": "custom",
		  "char_filter": ["html_strip"],  ` + /* TODO: html_stip seems like not working: will be made sure working */ `
          "tokenizer": "ngram_tokenizer"
        },
        "ja_analyzer": {
          "type": "custom",
          "tokenizer": "ja_tokenizer",
          "filter": [
            "kuromoji_neologd_baseform",
            "kuromoji_neologd_stemmer",
            "ja_posfilter",
            "cjk_width"
          ],
          "char_filter": [
            "html_strip",  ` + /* TODO: html_stip seems like not working: will be made sure working */ `
            "kuromoji_neologd_iteration_mark"
          ]
        }
      }
    }
  }
`
	// Uses for anything.
	//
	// Before definition, made prototype testing project that
	// it considered to smart mapping and settings behavior
	//
	ProductsName  = "products"
	productsIndex = `{
  ` + commonSettings + `
  ,
  "mappings": {
    "` + ProductsName + `": {
      "properties": {
        "id": {
          "type": "long"
        },
		` + /* Hybrid search between name, name_kana, name_anything */ `
        "name": {
          "type": "text",
          "boost": 4,
          "analyzer": "ja_analyzer"
        },
        "name_ruby": {
          "type": "text",
          "store": true,
          "boost": 4,
          "analyzer": "rubytext_analyzer"
        },
		` + /* Set boost=1 instead of boost=4, this is because just
		for prevent zero-result in elasticsearch's request */`
        "name_anything": {
          "type": "text",
          "boost": 1,
          "analyzer": "unigram_analyzer"
        },
        "description_detail": {
          "type": "text",
          "boost": 4,
          "analyzer": "ja_analyzer"
        },
        "search_word": {
          "type": "keyword",
          "boost": 6
        },
        "update_date": {
          "type": "date"
        }
      },
      "_all": {
        "enabled": false
      }
    }
  }
}`
)
