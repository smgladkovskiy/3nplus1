package main

import (
	"flag"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile = flag.String("memprofile", "", "write memory profile to `file`")
	minIPower  = flag.Float64("minIPower", 0, "starting number power of 10")
	maxIPower  = flag.Float64("maxIPower", 7, "ending number power of 10")
)
