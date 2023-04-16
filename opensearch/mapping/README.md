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
    "match": {
      "name": "加藤"
    }
  },
  "highlight": {
    "fields": {
      "name": {}
    }
  }
}
```

