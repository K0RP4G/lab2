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

func main() {
    library := make(map[Author][]Book)

    // Приклад авторів
    author1 := Author{"Тарас Шевченко", 25}
    author2 := Author{"Леся Українка", 15}

    // Додавання книг
    library[author1] = []Book{
        {"Кобзар", author1, 1840},
        {"Гайдамаки", author1, 1841},
    }

    library[author2] = []Book{
        {"Лісова пісня", author2, 1911},
        {"Камінний господар", author2, 1912},
    }

    // Пошук книг за автором
    searchAuthor(author1, library)
}

// Функція пошуку книг певного автора
func searchAuthor(a Author, lib map[Author][]Book) {
    books, exists := lib[a]
    if exists {
        fmt.Printf("Книги автора %s:\n", a.Name)
        for _, book := range books {
            fmt.Printf("- %s (%d)\n", book.Title, book.Year)
        }
    } else {
        fmt.Printf("Книг автора %s не знайдено.\n", a.Name)
    }
}
