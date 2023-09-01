// console luhn generator
package main

import (
	"fmt"

	"github.com/alecthomas/kingpin/v2"

	"github.com/freepaddler/luhna"
)

var (
	count  = kingpin.Flag("count", "luhn numbers count").Default("6").Short('c').Int()
	prefix = kingpin.Flag("prefix", "luhn numbers prefix").Default("").Short('p').String()
	length = kingpin.Flag("length", "luhn numbers length").Default("16").Short('l').Int()
)

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate)
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	for i := 0; i < *count; i++ {
		fmt.Println(luhna.Generate(*prefix, *length))
	}

}
