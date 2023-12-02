# Go Advent Template
This is an overengineered template to simplify managing inputs/solutions for Advent of Code.

## Usage
You can run the following for each day's new problem:
```
go run ./new
```
This will append a directory to `solutions` named after the day number (e.g. `solutions/1`). The new directory will contain a Go file in which to implement the solution (e.g. `2.go`).

It will also attempt to fetch the input automatically, but only if you have set your Advent of Code session cookie to the env var `AOC_SESSION`. You'll need to set this manually, generally by copying it from your browser after logging in. If it fails to do so, it will instead create an empty `input.txt` for you to manually populate.

When solutions are run, input files are split by newlines and passed in as string slices to the `Solution` function.

When you are ready to run your solution, you can use:
```
go run ./generate
go run ./solutions [which_day's_solution_to_run]
```

You only need to generate (i.e. `go run ./generate`) a new main file when a new solution is added. This will likely mean once a day.
