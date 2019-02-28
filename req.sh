#!/bin/bash

arg=$1


# { "terms": { "name_ruby": ["'$arg'"] }},
# { "terms": { "name_anything" : ["'$arg'"]}},
curl -H "Content-Type: application/json" -XPOST 'http://localhost:9200/products/products/_search?pretty=true' -d '{
  "min_score": 0.1,
  "query": {
    "bool": {
      "should": [
        { "terms": { "name": ["'$arg'"] }},
        { "terms": { "description_detail": ["'$arg'"] }},
        { "terms": { "search_word": ["'$arg'"] }}
      ],
      "filter": [
        { "terms": { "product_status_id": [1] }}
      ]
    }
  }
}'

