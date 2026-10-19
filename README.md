# starter-golang

A personal repo for learning Go (Golang) step by step. Each numbered folder is a separate lesson with its own small examples and exercises.

## Structure

```
.
├── 01/          # Lesson 1: variables, functions, printing, float types
│   ├── go.mod
│   └── main.go
├── 02/          # Lesson 2 (in progress)
└── README.md
```

## Requirements

- [Go](https://go.dev/dl/) 1.27 or newer

Check your version:

```sh
go version
```

## Running a lesson

Go into the lesson folder and run the whole package:

```sh
cd 01
go run .
```

`go run .` compiles every `.go` file in the folder together. Use it instead of `go run main.go` so that functions in other files (for example `exercise.go`) can be found.

To build a binary:

```sh
go build -o app .
./app
```

## Notes

- Every runnable program needs `package main` and exactly one `func main()`.
- Other files in the same folder can define helper functions, which `main()` then calls.

## Resources

- [A Tour of Go](https://go.dev/tour/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
