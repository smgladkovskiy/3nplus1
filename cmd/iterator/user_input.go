package main

import (
	"flag"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile = flag.String("memprofile", "", "write memory profile to `file`")
	maxIPower  = flag.Int("maxIPower", 8, "ending number power of 10")
)
