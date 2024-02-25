# Pokedex REPL

## Usage
```bash
$ ./pokedex-go
Welcome to Pokedex-CLI!

pokedex > help
...
pokedex > exit
```

## Development
Before a PR, please run the following:
```bash
# defaults to running formatters, tests and building application
make
```

### Formatting
Just do whatever `gofmt` tells you to.
```bash
# run formatters and tidy modfile
make tidy
```
NOTE: `gofmt` will sometimes provide suggestions for code simplification using the `-s` flag.
Recommend running the following to run additional linters:
```bash
make check
```

### Testing
100% test coverage is not required nor is it a goal. Test the important pieces of logic, but don't
worry about branch coverage.
```bash
# run all tests
make test
```
To just run tests for a particular subpackage, run one of the following (`make test` is a light
wrapper around `go test`):
```bash
# run tests for subpackage <pkg>
make test packages=./<pkg>
# OR just...
go test ./<pkg>
```
Unit testing is great, but also make sure the actual behavior is what you expect by running some
basic manual validation:
```bash
# starts the pokedex-cli
make run
```

## Planned
- [ ] Default region and generation (used for lookups, routes, etc.)
- [ ] Lookup command
    - [ ] Display basic info (base stats, types, etc)
    - [ ] Display weaknesses (NOTE: will require calculation from ability and types)
    - [ ] Fetch moveset (ideally, filtered by the generation)
- [ ] Lookup _with level_ (filters moveset by level)
- [ ] Installation script

## Resources
- [Rough project plan](https://www.boot.dev/learn/build-pokedex-cli) courtesy of boot.dev/
