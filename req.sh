#!/bin/bash


# {
		# "_source": ["_id"],
		# "query": {
		  # "bool": {
			  # "must": {
				  # "term": { "_id": [%s] }
			  # }
		  # }
		# }

    # }

curl -H "Content-Type: application/json" -XPOST 'http://localhost:9200/products/products/_search?pretty=true' -d '{
  "min_score": 0.1,
  "query": {
    "bool": {
      "should": [
        { "terms": { "name": ["チョコレート"] }},
        { "terms": { "name_ruby": ["チョコレート"] }},
        { "terms": { "name_anything" : ["チョコレート"]}},
        { "terms": { "description_detail": ["チョコレート"] }},
        { "terms": { "search_word": ["チョコレート"] }}
      ],
      "filter": [
        { "terms": { "product_status_id": [1] }}
      ]
    }
  }
}'
