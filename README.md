# Go Bookstore API

Go言語で実装された書店API。クリーンアーキテクチャと依存性注入のサンプルプロジェクトです。

## プロジェクト概要

書店のチェックアウト機能を提供するREST APIです。クリーンアーキテクチャの原則に従い、ドメイン駆動設計（DDD）を採用しています。

## 技術スタック

- **言語**: Go 1.23
- **Webフレームワーク**: chi v5
- **データベース**: MySQL 8.0
- **コンテナ**: Docker / Docker Compose
- **ホットリロード**: Air
- **アーキテクチャ**: Clean Architecture

## プロジェクト構成

```
.
├── cmd/
│   └── api/
│       └── main.go      # エントリーポイント
├── config/              # 設定管理
│   └── config.go        # 環境変数読み込み
├── domain/              # ドメイン層
│   ├── model/          # ドメインモデル
│   │   └── book.go     # Book エンティティ
│   └── repository/     # リポジトリインターフェース
│       ├── book.go     # Book リポジトリ
│       └── payment.go  # Payment リポジトリ
├── usecase/            # ユースケース層
│   └── book.go         # Book ビジネスロジック
├── handler/            # プレゼンテーション層
│   ├── book.go         # Book ハンドラー
│   ├── request/        # リクエストDTO
│   │   └── book.go
│   └── response/       # レスポンスDTO
│       └── book.go
├── infra/              # インフラストラクチャ層
│   ├── dao/            # データアクセス
│   │   └── book.go     # Book DAO
│   ├── external/       # 外部サービス
│   │   └── payment.go  # 決済サービス（Rakuten Pay）
│   └── database/       # データベース接続
│       └── db.go       # DB初期化
├── compose.yml         # Docker Compose 設定
├── Dockerfile          # Docker イメージ定義
├── init.sql            # データベース初期化スクリプト
├── .air.toml           # Air 設定ファイル
└── go.mod              # Go モジュール定義
```

## アーキテクチャ

このプロジェクトは以下の4層構造を採用しています：

1. **Domain層**: ビジネスロジックとエンティティを定義
2. **UseCase層**: アプリケーションのユースケースを実装
3. **Handler層**: HTTPリクエスト/レスポンスを処理
4. **Infrastructure層**: データベース、外部APIなどの実装

依存関係は外側から内側へ向かい、ドメイン層は他の層に依存しません。

## セットアップ

### 必要な環境

- Docker
- Docker Compose

### インストール

1. リポジトリをクローン

```bash
git clone https://github.com/yoshi-koyama/go-bookstore-api.git
cd go-bookstore-api
```

2. コンテナを起動

```bash
docker compose up -d
```

3. アプリケーションが起動したことを確認

```bash
docker compose ps
```

## 使用方法

### 動作確認

```bash
curl http://localhost:8080/hello
# 出力: hello world
```

### 書籍一覧取得API

```bash
curl http://localhost:8080/bookstore/api/books
```

### チェックアウトAPI

```bash
curl -X POST http://localhost:8080/bookstore/api/checkouts \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "amount_to_pay": 2000
  }'
```

## API エンドポイント

| メソッド | パス | 説明 |
|---------|------|------|
| GET | `/hello` | ヘルスチェック |
| GET | `/bookstore/api/books` | 書籍一覧の取得 |
| POST | `/bookstore/api/checkouts` | 書籍のチェックアウト |

## 環境変数

アプリケーションの設定は環境変数で管理されます。Docker Compose使用時は`compose.yml`で自動設定されます。

| 環境変数 | 説明 | デフォルト値 |
|---------|------|-------------|
| `DB_HOST` | データベースホスト | `db`（コンテナ間）/ `localhost`（ホストから） |
| `DB_PORT` | データベースポート | `3306` |
| `DB_USER` | データベースユーザー | `bookstore_user` |
| `DB_PASSWORD` | データベースパスワード | `bookstore_password` |
| `DB_NAME` | データベース名 | `bookstore` |

設定は`config/config.go`の`config.Load()`関数で読み込まれ、`cmd/api/main.go`で依存性注入に使用されます。

## データベース

### 接続情報

- **ホスト**: localhost（ホストから）/ db（コンテナ間）
- **ポート**: 3306
- **データベース名**: bookstore
- **ユーザー**: bookstore_user
- **パスワード**: bookstore_password

### テーブル構成

#### books テーブル

| カラム名 | 型 | 説明 |
|---------|-----|------|
| id | INT | 主キー（自動採番） |
| name | VARCHAR(255) | 書籍名 |
| price | INT | 価格 |
| created_at | TIMESTAMP | 作成日時 |
| updated_at | TIMESTAMP | 更新日時 |

### データベースへの接続

```bash
docker compose exec db mysql -u bookstore_user -pbookstore_password bookstore
```

## 開発

### ホットリロード

Airによるホットリロードが有効になっています。ソースコードを変更すると自動的にアプリケーションが再起動します。

### ログの確認

```bash
# 全てのコンテナのログ
docker compose logs -f

# アプリケーションのログのみ
docker compose logs -f app

# データベースのログのみ
docker compose logs -f db
```

### コンテナの停止

```bash
# コンテナを停止
docker compose down

# コンテナとデータを完全削除
docker compose down -v
```

## 依存関係

- [chi](https://github.com/go-chi/chi) - 軽量で高速なHTTPルーター
- [render](https://github.com/go-chi/render) - JSONレンダリング

## ライセンス

MIT
