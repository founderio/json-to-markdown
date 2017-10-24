# Json to Markdown converter

This tool converts json backup files of a specific android notes application to rudimentary markdown. See example.json for an example input.
Still in beta phase, written for personal usage.

## Building from source

The supplied makefile simply runs `go build` with a custom output file. On Unix systems, simply run `make` from the source directory.

## Using the tool

Run `jsonmarkdown -h` to get this overview:

```
Usage of ./jsonmarkdown:
  -i value
        shorthand for --input
  -input value
        Input files. You can add this option multiple times: --input file1.json --input file2.json
  -o string
        Shorthand for --output (default "notes.md")
  -output string
        Output file: --output out.md (default "notes.md")
```

To convert the example json, you can run `./jsonmarkdown -i example.json` to convert to the output file `notes.md`.
