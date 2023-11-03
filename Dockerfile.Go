# 使用するGoのバージョンを指定
FROM golang:1.21.3 AS builder

# ワーキングディレクトリの設定
WORKDIR /app

# Goモジュールを有効にする
ENV GO111MODULE=on

# 依存関係をインストール
COPY go.mod .
COPY go.sum .
RUN go mod download

# ソースコードをコピー
COPY . .

# air のインストール
RUN go install github.com/cosmtrek/air@latest

# $GOPATH/bin を $PATH に追加する
ENV PATH="${PATH}:${GOPATH}/bin"

# アプリケーションの実行（air を使用してホットリロードを有効にする）
CMD ["air"]
