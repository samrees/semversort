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
		reverseFlag  bool
	)

	flag.BoolVar(&greatestFlag, "greatest", false, "display the greatest version for a given list")
	flag.BoolVar(&leastFlag, "least", false, "display the least version for a given list")
	flag.StringVar(&constraint, "constraint", "", "list versions only if versions pass given constraint")
	flag.BoolVar(&reverseFlag, "reverse", false, "lists verions greatest to least")

	flag.Parse()

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

	// least to greatest
	sort.Sort(semver.Collection(versions))

	if greatestFlag {
		fmt.Println(versions[len(versions)-1])
		os.Exit(0)
	}

	if leastFlag {
		fmt.Println(versions[0])
		os.Exit(0)
	}

	//greatest to least
	if reverseFlag {
		sort.Sort(sort.Reverse(semver.Collection(versions)))
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
	} else {
		for _, v := range versions {
			fmt.Println(v)
		}
	}
	os.Exit(0)
}
