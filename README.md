# searcher

open search practice

## open search

- Docker file
  - https://dev.classmethod.jp/articles/how-to-build-opensearch-with-docker/

## cmd

``` sh
curl -XGET https://localhost:9200 -u 'admin:admin' --insecure
```

## dash-board

http://localhost:5601/

- admin:admin でログインできる

RDB と概念的に似てる部分が多い

- index: テーブル
- mapping: スキーマ、テーブルの定義
- document: レコード

```
GET _cat/indices
```

mapping

type: 文字列、数値、日時
text, keyword 
keyword 形態素解析しない
文字列のタイムスタンp、id など

fields はあっちが勝手に用意してくれた何か、同じデータをタイプ別ものとして追加してくれてる。
.keyword でアクセスできるようにしてる。

カメラ → 内視鏡カメラ、とかが
カメ → カラメでヒットするようになってる。


mapping は index 作成時に登録する。

## analyzers

analyzers settings とかに定義する。


response の構造体に入れたいやつだけ絞る

```
GET /pieen/_search
{
  "_source": "members.blood_type"
}
```

curl

``` sh
curl -XGET "https://localhost:9200/pieen/_search" -H 'Content-Type: application/json' -k -u admin:admin -d'
{
  "_source": "members.blood_type"
}'
```




```
POST /pieen3/_doc/afaf_id
```


``` 
GET /pieen3/_search
{
  "_source": [
    "name"
  ],
  "query": {
    "match": {
      "blood_type": "A"
    }
  },
  "highlight": {
    "fields": {
      "blood_type": {}
    }
  }
}


GET /pieen3/_search
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
      "blood_type": {}
    }
  }
}



GET /pieen4/_search
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


スタンダードの analyzer 

```
GET /_analyze
{
  "text": [
    "すたんだーどの analyzer"
    ]
}


GET /_analyze
{
  "text": [
    "第１回 Elastisearch 入門 インデックスを設計する際に知っておくべき事"
    ],
    "analyzer": "kuromoji"
}
```

fvh: 隣接する em は括るようにしたりできる。


- should は or 演算の意味
  - minimum なんかを設定させるべき
  - minimum_should_match はつけておいて損がない

