# use external modules whose repo has been deleted and then recreated with the same name

This is an example of a package, the main package stored in the `src/use-external-module-repo-deleted/main` folder, that uses an external module stored in a git repo.

The original git repo of the external model defined the versions v0.0.0 and v0.0.1 and then was loaded into the public registry pkg.go.dev

Then this repo has been deleted and a new repo, with the same name has been created. This new repo defines only version v0.0.0

This external module has some packages residing below a folder names `internal`. The exported names (e.g. variables and functions) exported by the pacakges below the `internal` folder are not visible for this package even if they are exported (i.e. they start with a capital letter). These exported names are visible only within the module that defines them.

This is the [main.go](./main/main.go) file.

## build

Build the executable running the following command from the go-class project directory

`go build -o ./bin/use-external-module-repo-deleted ./src/use-external-module-repo-deleted/main`

## run

Run the example calling the following command from the go-class project directory

`./bin/use-external-module-repo-deleted`
