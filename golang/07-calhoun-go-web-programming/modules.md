# Go Modules

## 1. Dependency management

Make sure other developers can build your code using a similar version

Go uses SIV (Semantic Import Versioning) which is based on Semantic Versioning
- See <semver.org>

Highlights:
- Composed of 3 numbers. eg `v1.12.4`
- Major, minor, patch
- Major bumps == breaking change. Manual upgrade is necessary
- Other bumps can be automated by tooling if needed

Examples:
- Install `v1.12.4`, other devs may end up building with `1.13.0` if it releases
- If releasing a new version from `v1.13.6` with breaking change, next version will probably be `v2.0.0`
- If `v2.0.0` releases our tooling will NOT upgrade to that automatically

Special case:
- `v0.x.x` allows for any breaking changes. It is a special version
- `go get github.com/joncalhoun/somelib` will NOT get the latest version. If there are v1,v2,v3,v4, probably will grab v1.
  - We need to run `go get github.com/joncalhoun/somelib/v4` or similar
  - Tooling around this isn't always perfect. Will hopefully improve over time.

## 2. Working outside of `GOPATH`

All Go code used to live inside a single directory in your computer - the `GOPATH`.

Go modules allow us to run code from anywhere, as long as we initialize a module.

## Setting up our module

```bash
go mod init github.com/joncalhoun/lenslocked
```

This creates the `go.mod` file in our project. It also adds a `go.sum` file which is used to verify the integrity of our dependencies.

`go mod tidy` add missing and remove unused modules.

From this point onwards we don't need to think too much about modules. If we `go get` a library, the tooling will update our `go.mod` and `go.sum` files for us.

*Gotcha: In some/most editors you need to open the code from the directory with the `go.mod` file in it. Otherwise the editor won't know about the module. Eg `cd <dir-with-go.mod>` then `code .` to open the dir.*