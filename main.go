package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/Masterminds/semver/v3"
)

func main() {
	var (
		fgreatest  bool
		fleast     bool
		constraint string
	)

	flag.BoolVar(&fgreatest, "greatest", false, "display the greatest version for a given list")
	flag.BoolVar(&fleast, "least", false, "display the least version for a given list")
	flag.StringVar(&constraint, "constraint", "", "list versions greatest to least, if versions pass given constraint.")

	flag.Parse()

	if fgreatest == false && fleast == false && constraint == "" {
		flag.Usage()
		os.Exit(0)
	}

	reader := bufio.NewReader(os.Stdin)
	var rawvers []string

	for {
		input, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading input: ", err.Error())
			os.Exit(1)
		}

		rawvers = append(rawvers, string(input))
	}

	vers := make([]*semver.Version, len(rawvers))

	for i, r := range rawvers {
		v, err := semver.NewVersion(r)
		if err != nil {
			fmt.Printf("Error parsing version '%s': %s\n", r, err.Error())
			os.Exit(2)
		}

		vers[i] = v
	}

	// greatest to least
	sort.Sort(sort.Reverse(semver.Collection(vers)))

	if fgreatest {
		fmt.Println(vers[0])
		os.Exit(0)
	}

	if fleast {
		fmt.Println(vers[len(vers)-1])
		os.Exit(0)
	}

	if constraint != "" {
		c, err := semver.NewConstraint(constraint)
		if err != nil {
			fmt.Printf("Error parsing constraint '%s': %s\n", constraint, err.Error())
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
