# ベースイメージとして公式のGoイメージを使用
FROM golang:1.18-alpine

# ワーキングディレクトリを設定
WORKDIR /app

# go.modとgo.sumをコピーして依存関係をダウンロード
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o /app/main ./cmd

# デフォルトのポートを公開
EXPOSE 8080

# コンテナ起動時に実行されるコマンド
CMD ["/app/main"]
