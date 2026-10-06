# sqlite-from-scratch

A personal project to build a minimal SQLite-like database engine in Go, from the ground up.

## Goals

- Understand how a relational database works internally by implementing one.
- Follow the general architecture of SQLite: tokenizer, parser, bytecode/VM, B-tree, pager, OS interface.

## Roadmap

- [ ] REPL with meta-commands (`.exit`, `.tables`)
- [ ] SQL tokenizer and parser (`INSERT`, `SELECT`)
- [ ] In-memory row storage
- [ ] Persistence to disk via a pager
- [ ] B-tree for indexed storage
- [ ] Cursors and table scans
- [ ] Node splitting and multi-level B-trees

## Running

```sh
go run ./cmd/sqlite
```

## Testing

```sh
go test ./...
```

## References

- [Let's Build a Simple Database](https://cstack.github.io/db_tutorial/)
- [SQLite Database File Format](https://www.sqlite.org/fileformat.html)
- [SQLite Architecture](https://www.sqlite.org/arch.html)
