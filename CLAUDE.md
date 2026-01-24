# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## プロジェクト概要

Go言語で実装された書店API。クリーンアーキテクチャとドメイン駆動設計（DDD）を採用したREST APIプロジェクト。

## アーキテクチャ

このプロジェクトは**クリーンアーキテクチャ**の4層構造を採用しており、依存関係は外側から内側へ向かう：

```
Handler層 → UseCase層 → Domain層
    ↓
Infrastructure層（Domain層のインターフェースを実装）
```

### 依存性注入の流れ

`cmd/api/main.go`で全ての依存関係を組み立て、各層に注入：

1. **Config**: 環境変数から設定を読み込み（`config.Load()`）
2. **Infrastructure**: データベース接続を初期化（`database.NewDB(cfg)`）
3. **Repository**: DAOとExternal実装をリポジトリとして作成
4. **UseCase**: リポジトリを注入してビジネスロジックを組み立て
5. **Handler**: UseCaseを注入してHTTPハンドラーを作成

### 重要な設計原則

- **Domain層は他の層に依存しない**: `domain/`配下はインターフェースとモデルのみ
- **Repository パターン**: `domain/repository/`でインターフェースを定義し、`infra/dao/`と`infra/external/`で実装
- **DTOの使用**: `handler/request/`と`handler/response/`でHTTP層とドメイン層を分離

## 開発コマンド

### Docker環境の操作

```bash
# コンテナの起動
docker compose up -d

# コンテナの停止
docker compose down

# コンテナとデータベースボリュームを完全削除
docker compose down -v

# アプリケーションの再起動（コード変更後）
docker compose restart app

# ログの確認
docker compose logs -f app
docker compose logs -f db
```

### テストの実行

```bash
# 全てのテストを実行
docker compose exec app go test ./...

# 特定のパッケージのテストを実行
docker compose exec app go test ./domain/model/

# 詳細な出力でテストを実行
docker compose exec app go test -v ./domain/model/

# 単一のテスト関数を実行
docker compose exec app go test -v -run TestNewBook ./domain/model/
```

### データベース操作

```bash
# MySQLに接続
docker compose exec db mysql -u bookstore_user -pbookstore_password bookstore

# テーブル構造の確認
docker compose exec db mysql -u bookstore_user -pbookstore_password bookstore -e "SHOW CREATE TABLE books\G"

# データの確認
docker compose exec db mysql -u bookstore_user -pbookstore_password bookstore -e "SELECT * FROM books;"
```

### ビルドとモジュール管理

```bash
# アプリケーションのビルド（コンテナ内）
docker compose exec app go build -o ./tmp/main ./cmd/api

# 依存関係の追加
docker compose exec app go get <package-name>

# go.modの整理
docker compose exec app go mod tidy
```

## モジュール構成

- **モジュール名**: `bookstore-api`
- **エントリーポイント**: `cmd/api/main.go`
- **設定管理**: `config/config.go`（環境変数から設定を読み込み）

### 環境変数

Docker Compose経由で以下の環境変数が設定される（`compose.yml`参照）：

- `DB_HOST`: データベースホスト（コンテナ間通信では `db`）
- `DB_PORT`: データベースポート（デフォルト: `3306`）
- `DB_USER`: データベースユーザー
- `DB_PASSWORD`: データベースパスワード
- `DB_NAME`: データベース名

## テスト作成のガイドライン

- **テーブル駆動テスト**を使用
- 変数名は`actual`（実際の値）と`expected`（期待する値）を使用
- 構造体の比較は`==`演算子で直接比較（フィールド追加時のテスト漏れを防ぐ）
- テスト関数名は`Test<対象>_<メソッド名>`形式

## Git/コミットのガイドライン

- **コミットメッセージは日本語で記述する**
- コミットメッセージの最後に `Co-Authored-By: Claude Sonnet 4.5 <noreply@anthropic.com>` を追加（Claudeとの共同作業時）

## API エンドポイント

- `GET /hello`: ヘルスチェック
- `GET /bookstore/api/books`: 書籍一覧の取得
- `POST /bookstore/api/checkouts`: 書籍のチェックアウト

## ホットリロード

Airによるホットリロードが有効。`.air.toml`で設定されており、以下のディレクトリの変更を監視：
- `cmd/`, `config/`, `domain/`, `handler/`, `infra/`, `usecase/`

ソースコードを変更すると自動的にアプリケーションが再ビルド・再起動される。

## データベーススキーマ

初期化スクリプト: `init.sql`

**booksテーブル**:
- `id` (INT, PRIMARY KEY, AUTO_INCREMENT)
- `name` (VARCHAR(255))
- `price` (INT)
- `created_at` (TIMESTAMP)
- `updated_at` (TIMESTAMP)
