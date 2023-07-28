# MegaverseAPIClient

[![codecov](https://codecov.io/gh/MihaiLupoiu/MegaverseAPIClient/branch/main/graph/badge.svg?token=drKg4oMLPh)](https://codecov.io/gh/MihaiLupoiu/MegaverseAPIClient)

The MegaverseAPIClient is a Go package that provides a client library for interacting with the Megaverse API. The Megaverse is a virtual world where different types of entities, such as Polyanets, Comeths, and Soloons, can be generated and manipulated. This client library allows developers to easily create, retrieve, and delete these entities in the Megaverse using Go code.

## Usage

Help menu for make options:
```
$ make help
```
Build the project:
```
$ make build
```
Run the some basic test:
```
$ make run
```
Run phase1:
```
$ make run-phase1
```
Run phase2:
```
$ make run-phase2
```

Or run after building:
```
./target/bin/megaverse -phase=test/phase1/phase2
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
ok      github.com/MihaiLupoiu/MegaverseAPIClient/megaverse     0.023s  coverage: 80.2% of statements
ok      github.com/MihaiLupoiu/MegaverseAPIClient/megaverse/astral      0.020s  coverage: 100.0% of statements
```

## Strucure

```
.
├── Makefile
├── README.md
├── cmd                // Entry point for the binary
│   └── main.go        // Main CLI folder
├── go.mod
├── go.sum
└── megaverse
    ├── astral
    │   ├── astral.go         // Package defining common types and interfaces for Astral objects
    │   ├── cometh.go         // Package defining the Cometh Astral object
    │   ├── cometh_test.go    // Unit tests for Cometh Astral object
    │   ├── map.go            // Package defining the Map structure and related functions
    │   ├── polyanet.go       // Package defining the Polyanet Astral object
    │   ├── polyanet_test.go  // Unit tests for Polyanet Astral object
    │   ├── soloon.go         // Package defining the Soloon Astral object
    │   └── soloon_test.go    // Unit tests for Soloon Astral object
    ├── astral.go             // Package defining the Astral Service interface and common functions
    ├── astral_test.go        // Unit tests for Astral Service and common functions
    ├── megaverse.go          // Package defining the Megaverse API client and related functions
    └── megaverse_test.go     // Unit tests for Megaverse API client and related functions
```