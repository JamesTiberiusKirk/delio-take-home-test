package main

import (
	"flag"
	"strings"
)

type Flags struct {
	Original string
	Stocks   []string
}

func (f *Flags) Set(value string) error {
	f.Stocks = strings.Split(value, ",")
	f.Original = value
	return nil
}

func (f *Flags) String() string {
	return f.Original
}

func buildFlags() Flags {
	f := Flags{}

	flag.Var(&f, "stocks", "List of stocks separated by a comma (,)")
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
	}

	return f
}
