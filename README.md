# gogid - Go, Get it Done!

gogid is a command line task tool loosely modeled as a GTD type system.

gogid was created as a learning project.

---

# Installation

Pre-built binaries are not currently provided. Use of `go get` is recommended.

```bash
go get github.com/shindakun/gogid/cmd/gogid
```

---

# Usage

Run with `go run ./cmd/gogid/gogid.go` or build with
  `go build ./cmd/gogid/gogid.go`

```bash
$ ./gogid -h
Usage of E:\Projects\Go\src\gogid\gogid.exe:
  -complete int
        Mark a single task as completed. (default -1)
  -id int
        ID of task in question, used to add notes (default -1)
  -new string
        Add a new task, enclose in quotes.
  -newnote string
        Add a new note to a task, enclose in quotes.
  -notcomplete int
        Mark a single task as not completed. (default -1)
  -print int
        Print a single task. (default -1)
  -printtasks
        Print entire task list to console.
```