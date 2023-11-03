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
RUN go get -u github.com/cosmtrek/air

# ホストマシンのソースコードディレクトリをコンテナのワーキングディレクトリにマウントする
# Docker コマンドで実行する際に指定します。

# アプリケーションの実行（air を使用してホットリロードを有効にする）
CMD ["air"]
