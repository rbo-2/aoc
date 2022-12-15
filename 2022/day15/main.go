package main

import (
	"aoc2022/common"
	"fmt"
	"sort"
	"strings"
)

type coord struct {
	x, y int
}

type rng [2]int

type sensor struct {
	coordi         coord
	radius         int
	closestBeacons coord
}

type space struct {
	sensors []sensor
	beacons []coord
}

func (r rng) limit(l rng) (rng, error) {
	if !r.intersectWith(l) {
		return rng{}, fmt.Errorf("cannot subtract ranges")
	}
	dx0, dx1 := l[0]-r[0], l[1]-r[1]
	if dx0 < 0 && dx1 > 0 {
		return r, nil
	}
	if dx0 > 0 && dx1 < 0 {
		return l, nil
	}
	if dx0 < 0 && dx1 < 0 {
		return rng{r[0], l[1]}, nil
	}
	if dx0 > 0 && dx1 > 0 {
		return rng{l[0], r[1]}, nil
	}
	return r, nil
}

func (r rng) intersectWith(r2 rng) bool {
	if r[1] < r2[0] || r2[1] < r[0] {
		return false
	}
	return true
}

func (r rng) mergeWith(r2 rng) rng {
	x1 := common.Min(r[0], r2[0])
	x2 := common.Max(r[1], r2[1])
	return rng{x1, x2}
}
func computeRadius(center, beacon coord) int {
	return common.Abs(center.x-beacon.x) + common.Abs(center.y-beacon.y)
}

func (s sensor) scanRow(row int) (rng, error) {
	if row > s.coordi.y+s.radius || row < s.coordi.y-s.radius {
		return rng{}, fmt.Errorf("row isn't included in sensor's range")
	}
	dx := s.radius - common.Abs(s.coordi.y-row)
	return rng{
		s.coordi.x - dx,
		s.coordi.x + dx,
	}, nil
}

func parseSpace(input []string) space {
	sensors := make([]sensor, 0)
	beacons := make([]coord, 0)
	for _, i := range input {
		sb := strings.Split(i, ":")
		sxy := strings.Split(sb[0], ",")
		sx := common.Int(sxy[0])
		sy := common.Int(sxy[1])
		bxy := strings.Split(sb[1], ",")
		bx := common.Int(bxy[0])
		by := common.Int(bxy[1])
		newBeacon := coord{bx, by}
		newSensor := sensor{coordi: coord{sx, sy}, closestBeacons: newBeacon}
		newSensor.radius = computeRadius(newSensor.coordi, newBeacon)
		sensors = append(sensors, newSensor)
		beacons = append(beacons, newBeacon)
	}
	return space{
		sensors,
		beacons,
	}
}

func mergeRanges(ranges []rng) []rng {

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	for i := 0; i < len(ranges)-1; {
		if ranges[i].intersectWith(ranges[i+1]) {
			ranges[i] = ranges[i].mergeWith(ranges[i+1])
			j := i + 1
			for j < len(ranges)-1 {
				temp := ranges[j]
				ranges[j] = ranges[j+1]
				ranges[j+1] = temp
				j++
			}
			ranges = ranges[:len(ranges)-1]
			continue
		}
		i++
	}

	return ranges
}

func part1(spc space) int {
	row := 10
	ranges := make([]rng, 0)
	for _, sens := range spc.sensors {
		newRange, err := sens.scanRow(row)
		if err != nil {
			continue
		}
		ranges = append(ranges, newRange)
	}
	ranges = mergeRanges(ranges)
	count := 0
	for _, r := range ranges {
		count += common.Abs(r[0] - r[1])
	}
	return count
}

func part2(spc space) int {
	min := 0
	max := 4000000
	for i := min; i <= max; i++ {
		ranges := make([]rng, 0)
		for _, sens := range spc.sensors {
			newRange, err := sens.scanRow(i)
			if err != nil {
				continue
			}
			newRange, err = newRange.limit(rng{min, max})
			if err != nil {
				continue
			}
			ranges = append(ranges, newRange)
		}
		ranges = mergeRanges(ranges)
		count := 0
		for _, r := range ranges {
			count += common.Abs(r[0]-r[1]) + 1
		}
		if count != max-min+1 {
			return i + 4000000*(ranges[0][1]+1)
		}
	}

	count := 0
	return count
}

func main() {
	input := common.Open(common.Args(1)).Lines()
	newSpace := parseSpace(input)
	//fmt.Printf("part1: %v\n", part1(newSpace))
	fmt.Printf("part1: %v\n", part2(newSpace))
}
