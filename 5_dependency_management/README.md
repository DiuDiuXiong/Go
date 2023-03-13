# Golang's Dependency Management

GOPATH --> GOVENDOR --> go mod (most mature one, better than the other two)

The other two are still remain in historical projects

## GOPATH
Whatever dependency should be installed under GOPATH, whatever not found in code will be looked under GOPATH.
Which is a bad idea, since eventually GOPATH will explode. (Deprecated, check 幕客 5-2 for example)
- All dependency for all your go projects under ${GOPATH/src}
- Can set go path under Setting/Language & Frameworks/Go/GOPATH
- To allow GOPATH, set `export GO111MODULE=off`
- `go get ...` will pull the library from remote place, each go project will have documentation
- The sequence of find dependency:
  - project/vendor
  - GOROOT
  - GOPATH

## GOVENDOR
Every GO project can have its own `vendor`, and go project will look for dependency under `vendor` first.
There are dependency management tools such as `Glide` & `Dep`

## GO mod
Canva use go.mod as well.
Try `go env` on terminal will get: `GO111MODULE="on"`
To get something, still `go get ...`. Here experiment on: [uber go zap](https://github.com/uber-go/zap).
```agsl
go get -u go.uber.org/zap
```
Will download:

Terminal:
```agsl
go: downloading go.uber.org/zap v1.24.0
go: downloading go.uber.org/atomic v1.7.0
go: downloading go.uber.org/multierr v1.6.0
go: downloading go.uber.org/multierr v1.10.0
go: downloading go.uber.org/atomic v1.10.0
go: added go.uber.org/atomic v1.10.0
go: added go.uber.org/multierr v1.10.0
go: added go.uber.org/zap v1.24.0
```

Go.mod have extra (and go.sum to make sure library is desired):
```
require (
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
)
```
- Its under `${GOPATH}/pkg/...` 
- The advantage is as long as you `go get <pkg_name>@v<version_number>` you can change the version of dependency.
- `go mod tidy` can clean all unused dependency
- for old projects not using `go.mod`, use `go mod init` or `go mod download` or `go build ./...` will cause all stuff written into go.mod
- to install all dependencies: `go get ./...` ./... means all recursive directories
 
# More about Go build
1. Each Directory can only have one main.main() function, but can create subdirectory to allow more main functions
2. go build `./...` will check if everything compile is fine
3. go install `./...` will install all file contain executable(main()) under current directory (executable) to `${GOPATH}/bin`