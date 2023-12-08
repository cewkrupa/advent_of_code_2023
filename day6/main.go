package main

import (
	"aoc-2023/internal/util/v1"
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	v1.Init(one, two)
}

type Race struct {
	Duration int
	Record   int
}

func one(f *os.File) {
	races := parseRaces(f)

	product := 1
	for _, race := range races {
		ways := solveRace(race)
		product = product * len(ways)
	}
	fmt.Printf("product: %v\n", product)
}

func two(f *os.File) {
	race := parseRace(f)

	ways := solveRace(race)

	fmt.Printf("ways: %v\n", len(ways))
}

func solveRace(race Race) []int {
	// y = -1x^2 + duration*x
	// y := (-1 * (x ^ 2)) + (race.Duration * x)
	// https://amsi.org.au/ESA_Senior_Years/SeniorTopic2/2a/2a_2content_9.html
	// the 'record' line is a straight line across distance (perpendicular to Y-axis).
	// We're looking for the intersection of the record line
	// and the quadratic formula for acceleration,
	// and then finding all whole numbers between those two points on the x axis
	// line:
	// y = mx + d
	// y = 0(x) + race.Record

	// 1.
	//   a. construct formula (coefficients) for race

	// ax^2+bx+c=mx+d => ax^2+(b−m)x+(c−d)=0
	a := float64(-1) // this is related to accel, somehow
	b := float64(race.Duration)
	c := float64(0)

	m := float64(0)
	d := float64(race.Record)

	//   b. intersect with record line
	// ax^2+(b−m)x+(c−d)=0
	b2 := b - m
	c2 := c - d

	// 2. find discriminant for formula
	// quadratic: x = (-b ±√(b^2 - 4ac)) / 2a
	discriminant := (b2 * b2) - (4 * a * c2)
	sqr := math.Sqrt(discriminant)
	// 3. Solve formula for points

	p1 := (-b2 + sqr) / 2 * a
	p2 := (-b2 - sqr) / 2 * a

	// 4. find whole numbers between points

	// if the points are whole numbers, they're _at_ the record
	// and if they're at the record, we didn't beat it 🥈
	// so let's add a little padding in one direction or the other.
	if math.Mod(p1, 1.0) == 0 {
		p1 = p1 + 0.1
	}
	if math.Mod(p2, 1.0) == 0 {
		p2 = p2 - 0.1
	}

	p1NextInt := int(math.Ceil(p1))
	p2PrevInt := int(math.Floor(p2))

	var waysToBeat []int
	for i := p1NextInt; i <= p2PrevInt; i++ {
		waysToBeat = append(waysToBeat, i)
	}
	return waysToBeat
}

func parseRace(f *os.File) Race {
	scanner := bufio.NewScanner(f)

	r := regexp.MustCompile(`(\d+)`)

	scanner.Scan()
	timeLine := scanner.Text()

	timeStrings := r.FindAllString(timeLine, -1)
	timeJoin := strings.Join(timeStrings, "")
	dur, err := strconv.Atoi(timeJoin)
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	recordLine := scanner.Text()

	recordStrings := r.FindAllString(recordLine, -1)
	recJoin := strings.Join(recordStrings, "")
	record, err := strconv.Atoi(recJoin)
	if err != nil {
		panic(err)
	}

	return Race{
		Duration: dur,
		Record:   record,
	}
}

func parseRaces(f *os.File) []Race {
	scanner := bufio.NewScanner(f)

	var races []Race
	r := regexp.MustCompile(`(\d+)`)

	scanner.Scan()
	timeLine := scanner.Text()

	timeStrings := r.FindAllString(timeLine, -1)
	for _, t := range timeStrings {
		d, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		race := Race{
			Duration: d,
		}
		races = append(races, race)
	}

	scanner.Scan()
	recordLine := scanner.Text()

	recordStrings := r.FindAllString(recordLine, -1)
	for i, rec := range recordStrings {
		d, err := strconv.Atoi(rec)
		if err != nil {
			panic(err)
		}
		races[i].Record = d
	}

	return races
}
