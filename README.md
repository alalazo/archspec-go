# archspec-go
Test repository in Go to try a layout based on git submodules.

To get a command that prints the JSON file to `stdout` (assuming that `$GOPATH/bin` is in your `PATH`):
```console
$ GO111MODULE="on" go get -v github.com/alalazo/archspec-go/cmd/archspec
[ ... ]

$ archspec
```
The JSON submodule seems not to be cloned by `go get`:
```console
$ tree /home/culpo/go/pkg/mod/github.com/alalazo/archspec-go@v0.0.0-20200112130302-009033d8041b/
/home/culpo/go/pkg/mod/github.com/alalazo/archspec-go@v0.0.0-20200112130302-009033d8041b/
├── archspec
│   ├── pkged.go
│   └── read.go
├── cmd
│   └── archspec
│       └── main.go
├── go.mod
├── go.sum
└── version.go
```
but this doesn't matter much since the approach to embed the files requires the developer to generate a `pkged.go` 
file with the byte stream of the resources that will be embedded.
