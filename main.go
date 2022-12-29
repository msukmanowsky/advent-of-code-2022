package main

import (
	"flag"
	"log"
	"os"

	"github.com/msukmanowsky/advent-of-code-2022/day1"
	"github.com/msukmanowsky/advent-of-code-2022/day2"
	"github.com/msukmanowsky/advent-of-code-2022/day3"
	"github.com/msukmanowsky/advent-of-code-2022/day4"
	"github.com/msukmanowsky/advent-of-code-2022/day5"
	"github.com/msukmanowsky/advent-of-code-2022/day6"
	"github.com/msukmanowsky/advent-of-code-2022/day7"
	"github.com/msukmanowsky/advent-of-code-2022/day8"
	"github.com/msukmanowsky/advent-of-code-2022/day9"
)

func main() {
	day1Cmd := flag.NewFlagSet("day1", flag.ExitOnError)
	day2Cmd := flag.NewFlagSet("day2.1", flag.ExitOnError)
	day2_1Cmd := flag.NewFlagSet("day2.2", flag.ExitOnError)
	day3_1Cmd := flag.NewFlagSet("day3.1", flag.ExitOnError)
	day3_2Cmd := flag.NewFlagSet("day3.2", flag.ExitOnError)
	day4_1Cmd := flag.NewFlagSet("day4.1", flag.ExitOnError)
	day4_2Cmd := flag.NewFlagSet("day4.2", flag.ExitOnError)
	day5_1Cmd := flag.NewFlagSet("day5.1", flag.ExitOnError)
	day5_2Cmd := flag.NewFlagSet("day5.2", flag.ExitOnError)
	day6_1Cmd := flag.NewFlagSet("day6.1", flag.ExitOnError)
	day6_2Cmd := flag.NewFlagSet("day6.2", flag.ExitOnError)
	day7_1Cmd := flag.NewFlagSet("day7.1", flag.ExitOnError)
	day7_2Cmd := flag.NewFlagSet("day7.2", flag.ExitOnError)
	day8_1Cmd := flag.NewFlagSet("day8.1", flag.ExitOnError)
	day8_2Cmd := flag.NewFlagSet("day8.2", flag.ExitOnError)
	day9_1Cmd := flag.NewFlagSet("day9.1", flag.ExitOnError)
	day9_2Cmd := flag.NewFlagSet("day9.2", flag.ExitOnError)

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
	case "day3.1":
		day3_1Cmd.Parse(os.Args[2:])
		day3.Day3_1()
	case "day3.2":
		day3_2Cmd.Parse(os.Args[2:])
		day3.Day3_2()
	case "day4.1":
		day4_1Cmd.Parse(os.Args[2:])
		day4.Day4_1()
	case "day4.2":
		day4_2Cmd.Parse(os.Args[2:])
		day4.Day4_2()
	case "day5.1":
		day5_1Cmd.Parse(os.Args[2:])
		day5.Day5_1()
	case "day5.2":
		day5_2Cmd.Parse(os.Args[2:])
		day5.Day5_2()
	case "day6.1":
		day6_1Cmd.Parse(os.Args[2:])
		day6.Day6_1()
	case "day6.2":
		day6_2Cmd.Parse(os.Args[2:])
		day6.Day6_2()
	case "day7.1":
		day7_1Cmd.Parse(os.Args[2:])
		day7.Day7_1()
	case "day7.2":
		day7_2Cmd.Parse(os.Args[2:])
		day7.Day7_2()
	case "day8.1":
		day8_1Cmd.Parse(os.Args[2:])
		day8.Day8_1()
	case "day8.2":
		day8_2Cmd.Parse(os.Args[2:])
		day8.Day8_2()
	case "day9.1":
		day9_1Cmd.Parse(os.Args[2:])
		day9.Day9_1()
	case "day9.2":
		day9_2Cmd.Parse(os.Args[2:])
		day9.Day9_2()
	default:
		log.Fatalln("expected a valid subcommand")
	}
}
