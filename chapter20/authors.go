package chapter20

var (
	authors = []map[string]any{
		{
			"id":   1,
			"name": "Virginia Woolf",
		},
		{
			"id":   2,
			"name": "Leo Tolstoy",
		},
		{
			"id":   3,
			"name": "Dr. Seuss",
		},
		{
			"id":   4,
			"name": "J. R. R. Tolkien",
		},
		{
			"id":   5,
			"name": "Mark Twain",
		},
	}
	books = []map[string]any{
		{
			"authorID": 3,
			"title":    "Hop on Pop",
		},
		{
			"authorID": 1,
			"title":    "Mrs. Dalloway",
		},
		{
			"authorID": 4,
			"title":    "The Fellowship of the Ring",
		},
		{
			"authorID": 1,
			"title":    "To the Lighthouse",
		},
		{
			"authorID": 2,
			"title":    "Anna Karenina",
		},
		{
			"authorID": 5,
			"title":    "The Adventures of Tom Sawyer",
		},
		// you get the idea...
	}
)

func ConnectBooksWithAuthors() []map[string]string {
	booksWithAuthors := make([]map[string]string, 0)
	for _, book := range books {
		for _, author := range authors {
			if book["authorID"] == author["id"] {
				booksWithAuthors = append(booksWithAuthors, map[string]string{
					"title":  book["title"].(string),
					"author": author["name"].(string),
				})
			}
		}
	}

	return booksWithAuthors
}

func ConnectBooksWithAuthorsWithMaps() []map[string]string {
	booksWithAuthors := make([]map[string]string, 0)
	authorHash := make(map[string]string)
	for _, author := range authors {
		authorHash[author["id"].(string)] = author["name"].(string)
	}

	for _, book := range books {
		booksWithAuthors = append(booksWithAuthors, map[string]string{
			"title":  book["title"].(string),
			"author": authorHash[book["authorID"].(string)],
		})
	}

	return booksWithAuthors
}
