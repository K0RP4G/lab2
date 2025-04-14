package main

import (
    "fmt"
)

// Author - структура автора книги
type Author struct {
    Name       string
    Experience int // Кількість років досвіду
}

// Book - структура книги
type Book struct {
    Title  string
    Author Author
    Year   int
}

// SearchAuthorBooks повертає список книг автора, якщо вони є
func SearchAuthorBooks(a Author, lib map[Author][]Book) ([]Book, bool) {
    books, exists := lib[a]
    return books, exists
}

func main() {
    library := make(map[Author][]Book)

    author1 := Author{"Тарас Шевченко", 25}
    author2 := Author{"Леся Українка", 15}

    library[author1] = []Book{
        {"Кобзар", author1, 1840},
        {"Гайдамаки", author1, 1841},
    }

    library[author2] = []Book{
        {"Лісова пісня", author2, 1911},
        {"Камінний господар", author2, 1912},
    }

    // Демонстрація пошуку
    books, found := SearchAuthorBooks(author1, library)
    if found {
        fmt.Printf("Книги автора %s:\n", author1.Name)
        for _, book := range books {
            fmt.Printf("- %s (%d)\n", book.Title, book.Year)
        }
    } else {
        fmt.Printf("Книг автора %s не знайдено.\n", author1.Name)
    }
}
