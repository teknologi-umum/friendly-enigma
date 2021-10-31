package handlers

import "database/sql"

// If you don't understand what this is, read this:
// https://www.freecodecamp.org/news/a-quick-intro-to-dependency-injection-what-it-is-and-when-to-use-it-7578c84fa88f/
//
// It's called a "dependency injection", where you inject something
// as a dependency to something else.
type Deps struct {
	DB *sql.DB
}
