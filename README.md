# Incomplete CRUD / RBAC Service in Go

The repository name means nothing. But your task is to complete this repository on your own to be a functional CRUD/RBAC service in Go programming language.

The stack for this service is simple:

- Go
- Chi for HTTP routing
- SQLite3 for database

## Use cases

The service is about storing food to a refrigerator. The `packages/migration` package will insert some family members
into the database which all has some access to do something.  

The features to be made on this service are:
1. Family member interaction to the refrigerator (with the concept of CRUD)
2. A family member access to do things with the refrigerator (via role-based access control)

### HTTP endpoints

- `GET /`: list all food items
- `POST /`: add a new food item
- `PATCH /`: update a food item
- `DELETE /`: delete a food item

Accepted query parameters:
- `/?id={id}`: specify a food item by ID (for `GET`, `PATCH`, and `DELETE`)

### Directory structure

```
.
├── business              - Business logic
│  ├── access.go
│  ├── access_test.go
│  ├── food.go
│  ├── enums.go
│  └── model.go
├── go.mod                - Module declaration
├── go.sum                - Checksum for module
├── handlers              - HTTP handlers
│  ├── delete.go
│  ├── get.go
│  ├── handlers.go
│  ├── middleware.go
│  ├── patch.go
│  └── post.go
├── house.db              - SQLite3 database
├── main.go               - Main entry point
├── packages              - Packages
│  ├── jwt
│  │  ├── jwt.go
│  │  └── jwt_test.go
│  └── migration
│     └── migration.go
└── README.md
```

- `/business` contains the main logic without any presentation layer. It could be tested just by supplying the required parameters. No need to do a full HTTP request to such handlers.
- `/handlers` contains both the HTTP handlers and the presentation layer.
- `/packages` contains the logic of external packages that might and can be used by the business or handlers package.

## Setup

### Clone the repository

[Clone](https://docs.github.com/en/repositories/creating-and-managing-repositories/cloning-a-repository) the repository via Terminal:

```sh
git clone https://github.com/teknologi-umum/friendly-enigma
```

### Set up your local machine

Make sure to have Go and SQLite3 installed.

- Go: https://golang.org/dl/
- SQLite3: https://www.servermania.com/kb/articles/install-sqlite/

### Resources to begin with

If you are new and not really familiar with Go, you should try:

- If you like reading and hands-on code: 
  - [Go Tour](https://tour.golang.org/) - Very recommended
  - [Go By Examples](https://gobyexample.com/)
  - [Effective Go](https://golang.org/doc/effective_go)
- If you like watching tutorials:
  - [Free Code Camp](https://www.youtube.com/watch?v=YS4e4q9oBaU)
  - [Golang Dojo](https://www.youtube.com/playlist?list=PLve39GJ2D71xX0Ham0WoPaYfl8oTzZfN6)
  - [Tech With Tim](https://www.youtube.com/playlist?list=PLzMcBGfZo4-mtY_SE3HuzQJzuj4VlUG0q)
  - [Programmer Zaman Now](https://www.youtube.com/playlist?list=PL-CtdCApEFH_t5_dtCQZgWJqWF45WRgZw) - Indonesian
- If you like to pay for courses:
  - [Learn How To Code: Google's Go (golang) Programming Language](https://www.udemy.com/course/learn-how-to-code/) - Udemy

Then, some other things that the core concepts will be used here:

- [Algorithm: Bit Manipulation](https://www.youtube.com/watch?v=NLKQEOgBAnw)
- Role-based access control
  - [Wikipedia](https://en.wikipedia.org/wiki/Role-based_access_control)
  - [UpGuard](https://www.upguard.com/blog/rbac)
  - [Cloudflare Learning](https://www.cloudflare.com/learning/access-management/role-based-access-control-rbac/)

You can start working on the `main.go` file. Good luck!