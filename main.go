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
		greatestFlag bool
		leastFlag    bool
		constraint   string
	)

	flag.BoolVar(&greatestFlag, "greatest", false, "display the greatest version for a given list")
	flag.BoolVar(&leastFlag, "least", false, "display the least version for a given list")
	flag.StringVar(&constraint, "constraint", "", "list versions greatest to least, if versions pass given constraint.")

	flag.Parse()

	if !greatestFlag && !leastFlag && constraint == "" {
		flag.Usage()
		os.Exit(0)
	}

	reader := bufio.NewReader(os.Stdin)
	var rawVersions []string

	for {
		input, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading input: ", err.Error())
			os.Exit(1)
		}

		rawVersions = append(rawVersions, string(input))
	}

	versions := make([]*semver.Version, len(rawVersions))

	for i, r := range rawVersions {
		v, err := semver.NewVersion(r)
		if err != nil {
			fmt.Printf("Error parsing version '%s': %s\n", r, err.Error())
			os.Exit(2)
		}

		versions[i] = v
	}

	// greatest to least
	sort.Sort(sort.Reverse(semver.Collection(versions)))

	if greatestFlag {
		fmt.Println(versions[0])
		os.Exit(0)
	}

	if leastFlag {
		fmt.Println(versions[len(versions)-1])
		os.Exit(0)
	}

	if constraint != "" {
		c, err := semver.NewConstraint(constraint)
		if err != nil {
			fmt.Printf("Error parsing constraint '%s': %s\n", constraint, err.Error())
			os.Exit(3)
		}
		for _, v := range versions {
			status, _ := c.Validate(v)
			if status {
				fmt.Println(v)
			}
		}
		os.Exit(0)
	}

}
