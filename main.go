package main

import "github.com/alecthomas/chroma/formatters/html"

type config struct {
	lineNos bool
}

func main() {
	var cfg config
	opts := getOptions(cfg)

	// Turn off line numbers
	opts = append(options, html.WithLineNumbers(false))
}

func getOptions(cfg config) []html.Option {
	var options []html.Option
	options = append(options, html.WithLineNumbers(cfg.lineNos))
	return options
}
