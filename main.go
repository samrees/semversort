package main

import (
	"bufio"
	"flag"
	"fmt"
	semver "github.com/Masterminds/semver/v3"
	"io"
	"os"
	"sort"
)

func main() {
	fgreatest := flag.Bool("greatest", false, "display the greatest version for a given list")
	fleast := flag.Bool("least", false, "display the least version for a given list")
	constraint := flag.String("constraint", "", "list versions greatest to least, if versions pass given constraint.")

	flag.Parse()
	if !(*fgreatest == true || *fleast == true || *constraint != "") {
		flag.Usage()
		os.Exit(0)
	}

	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("you must pipe in a list of semvers to sort.")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	var rawvers []string

	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		rawvers = append(rawvers, string(input))
	}

	vers := make([]*semver.Version, len(rawvers))

	for i, r := range rawvers {
		v, err := semver.NewVersion(r)
		if err != nil {
			printedErr := fmt.Errorf("Error parsing version: %s", err)
			fmt.Println(printedErr)
			os.Exit(2)
		}

		vers[i] = v
	}

	// greatest to least
	sort.Sort(sort.Reverse(semver.Collection(vers)))

	if *fgreatest == true {
		fmt.Println(vers[0])
		os.Exit(0)
	}
	if *fleast == true {
		fmt.Println(vers[len(vers)-1])
		os.Exit(0)
	}

	if *constraint != "" {
		c, err := semver.NewConstraint(*constraint)
		if err != nil {
			printedErr := fmt.Errorf("Error parsing constraint: %s", err)
			fmt.Println(printedErr)
			os.Exit(3)
		}
		for _, v := range vers {
			status, _ := c.Validate(v)
			if status == true {
				fmt.Println(v)
			}
		}
		os.Exit(0)
	}

}
