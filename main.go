package main

import (
	"flag"
	"log"
	"os"

	"github.com/msukmanowsky/advent-of-code-2022/day1"
	"github.com/msukmanowsky/advent-of-code-2022/day2"
)

func main() {
	day1Cmd := flag.NewFlagSet("day1", flag.ExitOnError)
	day2Cmd := flag.NewFlagSet("day2.1", flag.ExitOnError)
	day2_1Cmd := flag.NewFlagSet("day2.2", flag.ExitOnError)

	if len(os.Args) < 2 {
		log.Fatalln("subcommand expected but not received")
	}

	switch os.Args[1] {
	case "day1":
		day1Cmd.Parse(os.Args[2:])
		day1.Day1()
	case "day2.1":
		day2Cmd.Parse(os.Args[2:])
		day2.Day2_1()
	case "day2.2":
		day2_1Cmd.Parse(os.Args[2:])
		day2.Day2_2()
	default:
		log.Fatalln("expected a valid subcommand")
	}
}
