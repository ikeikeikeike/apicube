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
    "query": {
      "bool": {
          "should": [
                {"terms": {"name": ["アイス"] }},
                {"terms": {"name_ruby": ["チェリーアイスサンド"] }},
                {"terms": {"name_anything" : ["チェリーアイスサンド"]}},
                {"terms": {"description_detail": ["アイス"] }},
                {"terms": {"search_word": ["アイス"] }}
          ]
      }
    }
}'
