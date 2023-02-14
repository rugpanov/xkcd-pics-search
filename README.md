# XKCD Comics Search

This is a Go program that allows users to search for *XKCD* comics based on keywords from their transcripts. If the
`xkcd.json` file doesn't exist, it will be fetched from the *XKCD* API and saved to disk.

## Usage

1. Clone this repository and navigate to the root directory.
2. Run `go run .`.
3. Provide whitespace separated keywords to search for *XKCD* comics based on their transcripts.
4. Type `exit` to leave the program.

## Dependencies

* Go 1.16 or higher

## License

This project is licensed under the MIT License.