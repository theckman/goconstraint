// Copyright (c) 2017 Tim Heckman
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file at the root of this repository.

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
)

const (
	maxMajor = uint64(1)

	minMinor = uint64(1)
	maxMinor = uint64(10)
)

var versionsToCheck = []string{
	"go1.1", "go1.2", "go1.3", "go1.4",
	"go1.5", "go1.6", "go1.7", "go1.8",
	"go1.9", "go1.10",
}

var verRegexp = regexp.MustCompile(`go(\d+)\.(\d+)(?:\.(\d+))?`)

func main() {
	toTest, toFail, err := versionsToTest()

	if err != nil {
		log.Fatalf("failed to get versions to test: %s", err)
	}

	for _, test := range toTest {
		if err := runTest(test, false); err != nil {
			log.Fatalf("last test suite (%q) unexpectedly failed: %s", test, err)
		}
	}

	for _, test := range toFail {
		if err := runTest(test, true); err != nil {
			log.Fatalf("last test suite (%q) unexpectedly failed: %s", test, err)
		}
	}
}

func versionsToTest() (pass []string, fail []string, err error) {
	major, minor, _, err := parseVersion(runtime.Version())

	if err != nil {
		log.Fatalf("failed to getVersion(): %s", err)
	}

	if err := validateVersions(major, minor); err != nil {
		log.Fatal(err)
	}

	var toTest []string
	var toFail []string

	for _, ver := range versionsToCheck {
		verMajor, verMinor, _, err := parseVersion(ver)

		if err != nil {
			return nil, nil, fmt.Errorf("unexpected parse error for %q: %s", ver, err)
		}

		if verMajor != major {
			toFail = append(toFail, ver)
		} else {
			if verMinor <= minor {
				toTest = append(toTest, ver)
			} else {
				toFail = append(toFail, ver)
			}
		}
	}

	return toTest, toFail, nil
}

func parseVersion(version string) (major, minor, patch uint64, err error) {
	match := verRegexp.FindAllStringSubmatch(version, -1)

	if len(match) < 1 {
		return 0, 0, 0, fmt.Errorf("unknown runtime version: %s", version)
	}

	if major, err = strconv.ParseUint(match[0][1], 10, 64); err != nil {
		return
	}

	if minor, err = strconv.ParseUint(match[0][2], 10, 64); err != nil {
		return
	}

	if match[0][3] != "" {
		if patch, err = strconv.ParseUint(match[0][3], 10, 64); err != nil {
			return
		}
	}

	return
}

func validateVersions(major, minor uint64) error {
	if major != maxMajor {
		return fmt.Errorf("unknown Go major version: %d", major)
	}

	if minor < minMinor || minor > maxMinor {
		return fmt.Errorf("unknown Go minor version: %d; valid range %d - %d", minor, minMinor, maxMinor)
	}

	return nil
}

func runTest(test string, shouldFail bool) error {
	dir := fmt.Sprintf("./%s/...", test)

	fmt.Printf("go test -v %s # should fail: %t\n", dir, shouldFail)

	cmd := exec.Command("go", "test", "-v", dir)

	if !shouldFail {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout
	}

	err := cmd.Run()

	if shouldFail {
		if err != nil {
			fmt.Printf("(%q test failed as expected; continuing)\n", test)
			return nil
		}
		log.Fatalf("expected %q to fail", test)
	}

	return err
}
