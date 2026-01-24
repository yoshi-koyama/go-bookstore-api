package model

import "testing"

func TestNewBook(t *testing.T) {
	tests := []struct {
		name     string
		id       int
		bookName string
		price    int
		expected Book
	}{
		{
			name:     "正常系: 通常の書籍",
			id:       1,
			bookName: "Go言語による並行処理",
			price:    3200,
			expected: Book{id: 1, name: "Go言語による並行処理", price: 3200},
		},
		{
			name:     "正常系: 価格が0",
			id:       2,
			bookName: "無料の本",
			price:    0,
			expected: Book{id: 2, name: "無料の本", price: 0},
		},
		{
			name:     "正常系: 長いタイトル",
			id:       3,
			bookName: "これは非常に長いタイトルの書籍です。テストのために作成しました。",
			price:    5000,
			expected: Book{id: 3, name: "これは非常に長いタイトルの書籍です。テストのために作成しました。", price: 5000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewBook(tt.id, tt.bookName, tt.price)
			if actual != tt.expected {
				t.Errorf("NewBook() = %+v, expected %+v", actual, tt.expected)
			}
		})
	}
}

func TestBook_ID(t *testing.T) {
	tests := []struct {
		name       string
		book       Book
		expectedID int
	}{
		{
			name:       "ID取得: 正の値",
			book:       NewBook(123, "テスト本", 1000),
			expectedID: 123,
		},
		{
			name:       "ID取得: 0",
			book:       NewBook(0, "ID0の本", 500),
			expectedID: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.book.ID()
			if actual != tt.expectedID {
				t.Errorf("Book.ID() = %v, expected %v", actual, tt.expectedID)
			}
		})
	}
}

func TestBook_Name(t *testing.T) {
	tests := []struct {
		name         string
		book         Book
		expectedName string
	}{
		{
			name:         "名前取得: 通常の書籍名",
			book:         NewBook(1, "ドメイン駆動設計", 3500),
			expectedName: "ドメイン駆動設計",
		},
		{
			name:         "名前取得: 空文字列",
			book:         NewBook(2, "", 1000),
			expectedName: "",
		},
		{
			name:         "名前取得: 英語のタイトル",
			book:         NewBook(3, "Clean Architecture", 4000),
			expectedName: "Clean Architecture",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.book.Name()
			if actual != tt.expectedName {
				t.Errorf("Book.Name() = %v, expected %v", actual, tt.expectedName)
			}
		})
	}
}

func TestBook_Price(t *testing.T) {
	tests := []struct {
		name          string
		book          Book
		expectedPrice int
	}{
		{
			name:          "価格取得: 通常の価格",
			book:          NewBook(1, "テスト本", 2500),
			expectedPrice: 2500,
		},
		{
			name:          "価格取得: 0円",
			book:          NewBook(2, "無料本", 0),
			expectedPrice: 0,
		},
		{
			name:          "価格取得: 高額な書籍",
			book:          NewBook(3, "専門書", 15000),
			expectedPrice: 15000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.book.Price()
			if actual != tt.expectedPrice {
				t.Errorf("Book.Price() = %v, expected %v", actual, tt.expectedPrice)
			}
		})
	}
}
