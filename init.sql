-- データベースの使用
USE bookstore;

-- booksテーブルの作成
CREATE TABLE IF NOT EXISTS books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- サンプルデータの挿入（オプション）
INSERT INTO books (name, price) VALUES
    ('Go言語による並行処理', 3200),
    ('Dockerで始めるコンテナ開発', 2800),
    ('実践MySQL入門', 3500);
