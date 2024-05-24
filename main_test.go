package main

import (
	"os"
	"strings"
	"testing"
)

func TestFull(t *testing.T) {
	in, err := os.ReadFile("./testdata/in.txt")
	if err != nil {
		t.Error(err)
		return
	}

	expected, err := os.ReadFile("./testdata/out.txt")
	if err != nil {
		t.Error(err)
		return
	}

	sessions := parseInput(strings.Split(string(in), "\n"))
	transformed := convertToString(sessions)

	if string(expected) != transformed {
		t.Errorf("out.txt doesn't match result, result:\n %s", transformed)
		return
	}
}
