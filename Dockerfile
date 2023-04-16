FROM opensearchproject/opensearch:latest
# 日本語の検索をするために必要なプラグインをインストール。
# see: https://opensearch.org/docs/latest/opensearch/plugins/
# see: https://subro.mokuren.ne.jp/0930.html
RUN /usr/share/opensearch/bin/opensearch-plugin install analysis-kuromoji 
RUN /usr/share/opensearch/bin/opensearch-plugin install analysis-icu
