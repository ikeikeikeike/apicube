#!/bin/bash

arg=$1

curl -H "Content-Type: application/json" -XPOST 'http://localhost:9200/products/products/_search?pretty=true' -d '{
  "query": {
    "bool": {
      "must": {
        "more_like_this" : {
            "fields" : ["name", "description_detail"],
            "like" : "'$arg'",
            "min_term_freq": 1,
            "min_doc_freq": 1
        }
      },
      "filter": [
        { "terms": { "product_status_id": [1] }}
      ]
    }
  }
}
'

# "min_term_freq" : 1,
# "max_query_terms" : 12

