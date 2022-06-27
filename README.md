# Alien Invasion

Alien Invasion was an interview assignment done by me, Gustavo Albino.

I took a few assumptions, like the following:
- Every city that connects to one, also have the reverse connection.
    - e.g. if foo west=bar, then bar east=foo
- The check to see if aliens are in the same space were made in the beginning of the round (before anyone moves)
- All aliens names are numbers
- All cities doesn't contain pipe in it's name
- The city explodes if two or more aliens are in the same city (instead of only two), and the city explodes with all of them

## Installation

To install, just run
```bash
go mod install
```

## Usage

Before using, you have to configure a file named cities (you can change it through args) following this logic:
```text
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
Qu-ux north=Foo south=Bee west=Bar
```
After that, you can run the application with
```bash
go run cmd/game/main.go
```
If you'd like to personalize the available variables, you also can, with the following
```bash
go run cmd/game/main.go -input <input_file> -n <amount_of_aliens>
```

## Tests

To run the tests, you can just execute them all with
```bash
go test ./...
```