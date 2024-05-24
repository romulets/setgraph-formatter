package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const setNameSep = " • "

type setParser func(string) []rep

var (
	pattern1 = regexp.MustCompile("^[0-9]+ sets: [0-9]+ rep( [0-9.]+ kg)?")
	pattern2 = regexp.MustCompile("^([0-9]+×[0-9.]+ kg,?)+")
	pattern3 = regexp.MustCompile("^[0-9.]+ kg: ([0-9]+,? )+rep$")
	pattern4 = regexp.MustCompile("^([0-9]+,? )+rep$")
	pattern5 = regexp.MustCompile("^[0-9]+ rep: ([0-9.]+,? )+kg")

	patternMap = map[*regexp.Regexp]setParser{
		pattern1: parsePattern1,
		pattern2: parsePattern2,
		pattern3: parsePattern3,
		pattern4: parsePattern4,
		pattern5: parsePattern5,
	}
)

type liftSession struct {
	sets []liftSet
}

type liftSet struct {
	name string
	reps []rep
}

type rep struct {
	count  int
	weight float64
}

func main() {
	input, err := getInput()
	if err != nil {
		fmt.Println("error: ", err.Error())
		os.Exit(1)
	}

	sessions := parseInput(input)
	transformed := convertToString(sessions)

	if err := saveSession(transformed); err != nil {
		fmt.Println("error: ", err.Error())
		os.Exit(1)
	}
}

func parseInput(input []string) []liftSession {
	sessions := make([]liftSession, 0, 31) // arbitrary number of a month max

	curSess := newSession()
	for _, l := range input {
		if l == "" {
			sessions = append(sessions, curSess)
			curSess = newSession()
			continue
		}

		parts := strings.Split(l, setNameSep)
		name := parts[0]
		unparsedRep := strings.Trim(parts[1], " ")
		curSess.sets = append(curSess.sets, liftSet{name: name, reps: parseRep(unparsedRep)})
	}

	sessions = append(sessions, curSess)

	return sessions
}

func parseRep(unparsedRep string) []rep {
	for p, f := range patternMap {
		if p.MatchString(unparsedRep) {
			return f(unparsedRep)
		}
	}

	fmt.Println("[WARN] UNKOWN PATTERN: ", unparsedRep)
	return nil
}

func parsePattern1(l string) []rep {
	parts := strings.Split(l, " sets: ")
	setCount, _ := strconv.Atoi(parts[0])

	repParts := strings.Split(parts[1], " rep")
	repCount, _ := strconv.Atoi(repParts[0])
	repWeight := float64(0)
	if len(repParts) > 1 {
		cleanWeight := strings.TrimSuffix(strings.TrimPrefix(repParts[1], " "), " kg")
		f, _ := strconv.ParseFloat(cleanWeight, 64)
		repWeight = float64(f)
	}

	reps := make([]rep, setCount)

	for i := 0; i < setCount; i++ {
		reps[i].count = repCount
		reps[i].weight = repWeight
	}

	return reps
}

func parsePattern2(l string) []rep {
	parts := strings.Split(l, ", ")

	reps := make([]rep, len(parts))
	for i, p := range parts {
		countWeight := strings.Split(p, "×")
		count, _ := strconv.Atoi(countWeight[0])
		weight, _ := strconv.ParseFloat(strings.TrimSuffix(countWeight[1], " kg"), 64)
		reps[i].count = count
		reps[i].weight = float64(weight)
	}

	return reps
}

func parsePattern3(l string) []rep {
	parts := strings.Split(l, " kg: ")
	weight, _ := strconv.ParseFloat(parts[0], 64)
	counts := strings.Split(strings.TrimSuffix(parts[1], " rep"), ", ")

	reps := make([]rep, len(counts))
	for i, count := range counts {
		reps[i].count, _ = strconv.Atoi(count)
		reps[i].weight = float64(weight)
	}

	return reps
}

func parsePattern4(l string) []rep {
	counts := strings.Split(strings.TrimSuffix(l, " rep"), ", ")

	reps := make([]rep, len(counts))
	for i, count := range counts {
		reps[i].count, _ = strconv.Atoi(count)
		reps[i].weight = float64(0)
	}

	return reps
}

func parsePattern5(l string) []rep {
	parts := strings.Split(l, " rep: ")
	count, _ := strconv.Atoi(parts[0])
	weights := strings.Split(strings.TrimSuffix(parts[1], " kg"), ", ")

	reps := make([]rep, len(weights))
	for i, rWeight := range weights {
		reps[i].count = count
		weight, _ := strconv.ParseFloat(rWeight, 64)
		reps[i].weight = float64(weight)
	}
	return reps
}

func getInput() ([]string, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("No input file provided")
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n"), nil
}

func saveSession(session string) error {
	now := time.Now()
	dateF := now.Format("20060102")
	return os.WriteFile(dateF+".out", []byte(session), 0664)
}

func convertToString(sessions []liftSession) string {
	b := strings.Builder{}
	for _, s := range sessions {
		b.WriteString(s.string())
		b.WriteRune('\n')
	}

	return b.String()
}

func (l liftSession) string() string {
	b := strings.Builder{}
	for _, set := range l.sets {
		b.WriteString(set.string())
	}

	return b.String()
}

func (l liftSet) string() string {
	b := strings.Builder{}
	b.WriteString(l.name)
	b.WriteRune('\t')

	b.WriteString(l.repWeightsString())
	b.WriteString("kg ")

	b.WriteString(strconv.Itoa(len(l.reps)))
	b.WriteRune('*')
	b.WriteString(l.repCountsString())

	b.WriteRune('\n')
	return b.String()
}

func newSession() liftSession {
	return liftSession{
		sets: make([]liftSet, 0, 20), // arbitrary numer of 20 being a lot of sets
	}
}

func (l liftSet) repCountsString() string {
	if l.allRepsSameCount() {
		return strconv.Itoa(l.reps[0].count)
	}

	counts := make([]string, len(l.reps))
	for i, rep := range l.reps {
		counts[i] = strconv.Itoa(rep.count)
	}

	return strings.Join(counts, "/")
}

func (l liftSet) repWeightsString() string {
	if l.allRepsSameWeight() {
		return strconv.FormatFloat(l.reps[0].weight, 'g', -1, 32)
	}

	counts := make([]string, len(l.reps))
	for i, rep := range l.reps {
		counts[i] = strconv.FormatFloat(rep.weight, 'g', -1, 32)
	}

	return strings.Join(counts, "/")
}

func (l liftSet) allRepsSameWeight() bool {
	for i := 1; i < len(l.reps); i++ {
		if l.reps[0].weight != l.reps[i].weight {
			return false
		}
	}

	return true
}

func (l liftSet) allRepsSameCount() bool {
	for i := 1; i < len(l.reps); i++ {
		if l.reps[0].count != l.reps[i].count {
			return false
		}
	}

	return true
}
