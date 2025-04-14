package main

import "testing"

func TestSearchAuthorBooks(t *testing.T) {
    author1 := Author{"Тарас Шевченко", 25}
    author2 := Author{"Іван Франко", 30}

    library := map[Author][]Book{
        author1: {
            {"Кобзар", author1, 1840},
            {"Гайдамаки", author1, 1841},
        },
    }

    tests := []struct {
        name      string
        input     Author
        wantFound bool
        wantLen   int
    }{
        {
            name:      "Author exists",
            input:     author1,
            wantFound: true,
            wantLen:   2,
        },
        {
            name:      "Author does not exist",
            input:     author2,
            wantFound: false,
            wantLen:   0,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            books, found := SearchAuthorBooks(tt.input, library)

            if found != tt.wantFound {
                t.Errorf("expected found=%v, got %v", tt.wantFound, found)
            }

            if len(books) != tt.wantLen {
                t.Errorf("expected %d books, got %d", tt.wantLen, len(books))
            }
        })
    }
}
