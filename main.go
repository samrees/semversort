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

var (
	ignoreParseError bool
	quiet            bool
	constraint       string
	reverseFlag      bool
	num              int
	latest           bool
	oldest           bool
)

func main() {
	flag.BoolVar(&ignoreParseError, "ignore", false, "Output the error and ignore if a version can't be parsed")
	flag.BoolVar(&quiet, "quiet", false, "Suppress error output")
	flag.StringVar(&constraint, "constraint", "", "Filter versions by constraints if given")
	flag.BoolVar(&reverseFlag, "reverse", false, "lists versions latest to oldest")
	flag.IntVar(&num, "num", 0, "Number of versions to display")
	flag.BoolVar(&latest, "latest", false, "Display the latest version")
	flag.BoolVar(&oldest, "oldest", false, "Display the oldest version")

	flag.Parse()

	if oldest {
		num = 1
	} else if latest {
		reverseFlag = true
		num = 1
	}

	versions := semver.Collection{}

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			outputToStderr("Error reading input: %s\n", err.Error())
			os.Exit(1)
		}

		line := string(input)
		ver, err := semver.NewVersion(line)
		if err != nil {
			outputToStderr("Invalid version: %q\n", line)
			if ignoreParseError {
				continue
			} else {
				os.Exit(2)
			}
		}

		versions = append(versions, ver)
	}

	// oldest to latest
	sort.Sort(versions)

	if constraint != "" {
		c, err := semver.NewConstraint(constraint)
		if err != nil {
			outputToStderr("Invalid constraint %q: %+v\n", constraint, err)
			os.Exit(3)
		}
		versions = filterByConstraint(versions, c)
	}

	if reverseFlag {
		sort.Sort(sort.Reverse(versions))
	}

	if num > 0 {
		if num > len(versions) {
			num = len(versions)
		}
		versions = versions[:num]
	}

	for _, v := range versions {
		fmt.Println(v)
	}
}

func filterByConstraint(versions semver.Collection, c *semver.Constraints) semver.Collection {
	r := semver.Collection{}
	for _, v := range versions {
		ok, _ := c.Validate(v)
		if ok {
			r = append(r, v)
		}
	}
	return r
}

func outputToStderr(format string, args ...interface{}) {
	if quiet {
		return
	}
	fmt.Fprintf(os.Stderr, format, args...)
}
