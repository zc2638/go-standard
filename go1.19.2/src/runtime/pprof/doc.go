// Package pprof writes runtime profiling data in the format expected
// by the pprof visualization tool.
//
// # Profiling a Go program
//
// The first step to profiling a Go program is to enable profiling.
// Support for profiling benchmarks built with the standard testing
// package is built into go test. For example, the following command
// runs benchmarks in the current directory and writes the CPU and
// memory profiles to cpu.prof and mem.prof:
//
//	go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
//
// To add equivalent profiling support to a standalone program, add
// code like the following to your main function:
//
//	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
//	var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
//
//	func main() {
//	    flag.Parse()
//	    if *cpuprofile != "" {
//	        f, err := os.Create(*cpuprofile)
//	        if err != nil {
//	            log.Fatal("could not create CPU profile: ", err)
//	        }
//	        defer f.Close() // error handling omitted for example
//	        if err := pprof.StartCPUProfile(f); err != nil {
//	            log.Fatal("could not start CPU profile: ", err)
//	        }
//	        defer pprof.StopCPUProfile()
//	    }
//
//	    // ... rest of the program ...
//
//	    if *memprofile != "" {
//	        f, err := os.Create(*memprofile)
//	        if err != nil {
//	            log.Fatal("could not create memory profile: ", err)
//	        }
//	        defer f.Close() // error handling omitted for example
//	        runtime.GC() // get up-to-date statistics
//	        if err := pprof.WriteHeapProfile(f); err != nil {
//	            log.Fatal("could not write memory profile: ", err)
//	        }
//	    }
//	}
//
// There is also a standard HTTP interface to profiling data. Adding
// the following line will install handlers under the /debug/pprof/
// URL to download live profiles:
//
//	import _ "net/http/pprof"
//
// See the net/http/pprof package for more details.
//
// Profiles can then be visualized with the pprof tool:
//
//	go tool pprof cpu.prof
//
// There are many commands available from the pprof command line.
// Commonly used commands include "top", which prints a summary of the
// top program hot-spots, and "web", which opens an interactive graph
// of hot-spots and their call graphs. Use "help" for information on
// all pprof commands.
//
// For more information about pprof, see
// https://github.com/google/pprof/blob/master/doc/README.md.
package pprof
