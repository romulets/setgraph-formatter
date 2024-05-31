package main

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestFull(t *testing.T) {
	conf := args{
		inputFile: "./testdata/in.txt",
		sort:      true,
		sortFile:  "./testdata/format.txt",
	}

	expected, err := os.ReadFile("./testdata/out.txt")
	if err != nil {
		t.Error(err)
		return
	}

	transformed, err := readAndConvertSessions(conf)
	if err != nil {
		t.Error(err)
		return
	}

	if string(expected) != transformed {
		t.Errorf("out.txt doesn't match result, result:\n %s", transformed)
		return
	}
}

func TestClean(t *testing.T) {
	input := `Squat (machine)	45/50/55kg 3*15/12/10
Standing leg curl	15kg 3*18/15/12
Leg press + calf raises	120kg 3*25/20/15
Hip Thrusts	60kg 3*18/16/14
Leg Extension	30kg 3*20
Reverse grip lat pulldown	35kg 3*15
Single arm iliac pulldown	5/7.5/7.5kg 3*15/16/16
Abs rolling wheel	0kg 3*15

Strength Training • 1 hr, 3 min

Tracked on Setgraph






Squat (machine)	45/50/55kg 3*15/12/10
Standing leg curl	15kg 3*18/15/12
Leg press + calf raises	120kg 3*25/20/15
Hip Thrusts	60kg 3*18/16/14
Leg Extension	30kg 3*20
Reverse grip lat pulldown	35kg 3*15
Single arm iliac pulldown	5/7.5/7.5kg 3*15/16/16
Abs rolling wheel	0kg 3*15

Strength Training • 1 hr, 3 min

Tracked on Setgraph



`

	expected := []string{
		"Squat (machine)	45/50/55kg 3*15/12/10",
		"Standing leg curl	15kg 3*18/15/12",
		"Leg press + calf raises	120kg 3*25/20/15",
		"Hip Thrusts	60kg 3*18/16/14",
		"Leg Extension	30kg 3*20",
		"Reverse grip lat pulldown	35kg 3*15",
		"Single arm iliac pulldown	5/7.5/7.5kg 3*15/16/16",
		"Abs rolling wheel	0kg 3*15",
		"",
		"Squat (machine)	45/50/55kg 3*15/12/10",
		"Standing leg curl	15kg 3*18/15/12",
		"Leg press + calf raises	120kg 3*25/20/15",
		"Hip Thrusts	60kg 3*18/16/14",
		"Leg Extension	30kg 3*20",
		"Reverse grip lat pulldown	35kg 3*15",
		"Single arm iliac pulldown	5/7.5/7.5kg 3*15/16/16",
		"Abs rolling wheel	0kg 3*15",
	}

	output := cleanInput(input)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Clean hasn't return expected output. \nExpected:\n%s\n\nGot:\n%s\n", strings.Join(expected, "\n"), strings.Join(output, "\n"))
	}

}
