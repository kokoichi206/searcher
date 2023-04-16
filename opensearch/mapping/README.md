## index & mapping 登録

``` json
PUT /members/
{
  "mappings" : {
    "properties" : {
      ...
  }
}
```

## document 登録

``` json
POST /members/_doc
{
  "id": 4,
  "name": "xx xx",
  "birthday": "2000年04月13日",
  "height": "155.0cm",
  "blood_type": "O型",
  "generation": "1期生",
  "blog_url": "https://hige",
  "img_url": "https://xxx.mydns.jp/paon.jpeg"
}
```

## mapping 定義取得

``` json
GET /members/_mapping
```

## 検索

``` json
GET /members/_search
{
  "_source": [
    "name"
  ],
  "query": {
    "bool": {
      "should": [
        {
          "match_phrase": {
            "name": "加藤"
          }
        },
        {
          "match": {
            "name": "加藤"
          }
        }
      ],
      "minimum_should_match": 1
    }
  },
  "highlight": {
    "fields": {
      "name": {}
    },
    "type": "fvh"
  }
}
```

- fvh で、連続する `<em>` を1つにまとめている
  - `"term_vector": "with_positions_offsets"` が必要
