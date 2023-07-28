# MegaverseAPIClient

[![codecov](https://codecov.io/gh/MihaiLupoiu/MegaverseAPIClient/branch/main/graph/badge.svg?token=drKg4oMLPh)](https://codecov.io/gh/MihaiLupoiu/MegaverseAPIClient)


## Usage

Help menu for make options:
```
$ make help
```
Build the project:
```
$ make build
```


## Testing
```
$ make test
```

Example of output: 
```
mlupoiu@DESKTOP-PMSP4AL:~/code/MegaverseAPIClient$ make test
go clean -testcache && go test ./... -cover -race
?       github.com/MihaiLupoiu/MegaverseAPIClient/cmd   [no test files]
ok      github.com/MihaiLupoiu/MegaverseAPIClient/megaverse     0.020s  coverage: 14.8% of statements
ok      github.com/MihaiLupoiu/MegaverseAPIClient/megaverse/astral      0.020s  coverage: 100.0% of statements
```

## Strucure

```
.
├── Makefile
├── README.md
├── cmd  // Entry point for the binary
│   └── main.go // Main CLI folder
├── go.mod
├── go.sum
└── megaverse
    ├── astral
    │   ├── astral.go
    │   ├── cometh.go
    │   ├── cometh_test.go
    │   ├── map.go
    │   ├── polyanet.go
    │   ├── polyanet_test.go
    │   ├── soloon.go
    │   └── soloon_test.go
    ├── astral.go
    ├── astral_test.go
    ├── megaverse.go
    └── megaverse_test.go
```