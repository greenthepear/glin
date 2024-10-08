**glin** is a simple tool for adding links to Go documentation in markdown files.

It scans the file for strings enclosed in ``[`...`]`` not followed by `(...)` (ones that already have the link)*. It then simply searches for one or two full words within and appends the url with either the provided repository link or one automatically found in go.mod. For example:
```md
this: [`Cell`]
becomes:
[`Cell`](https://pkg.go.dev/github.com/greenthepear/egriden#Cell)

this: [`(*GridLayer).CellAtScreenPos`]
becomes:
[`(*GridLayer).CellAtScreenPos`](https://pkg.go.dev/github.com/greenthepear/egriden#GridLayer.CellAtScreenPos)

this: [`Cell.Anchor`](https://pkg.go.dev/github.com/greenthepear/egriden#Cell.Anchor)
doesn't change, because a link is already there
```

\* *subject to change, will probably add a way to use custom patterns later instead*

# Usage
Check with `glin -h`.

---
```
glin -repo="github.com/greenthepear/egriden" -in CHANGELOG.md -out CHANGELOG_new.md
```
Will add the links with the repository path `github.com/greenthepear/egriden` in `CHANGELOG.md` and save that to `CHANGELOG_new.md`

---
```
glin -in "CHANGELOG.md" -ow
```
Will do the same as the former but overwrite `CHANGELOG.md` instead. Also the omitted `-repo` will attempt to get the repository path from the go.mod file in the working directory.

---
```
glin
```
Will read from stdin and print to stdout, so if you're a fellow bash pipe enjoyer you can do this to achieve the same as the first snippet:
```
cat CHANGELOG.md | glin > CHANGELOG_new.md
```

# Install
For the development version, newest in this repo:

```
go install github.com/greenthepear/glin@main
```

Make sure your `$GOBIN` is in your path.