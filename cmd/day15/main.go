package main

import (
	"advent_of_code_2024/pkg/aoc"
	"fmt"
	"slices"
	"strings"
)

func doHASH(s string) int {
	raw := []byte(s)
	ret := 0
	for _, c := range raw {
		ret += int(c)
		ret *= 17
		ret %= 256
	}
	return ret
}

type Pair struct {
	key string
	val int
}

type HashMap struct {
	slots [256][]Pair
}

func newHashMap() *HashMap {
	hm := &HashMap{}
	for idx := range hm.slots {
		hm.slots[idx] = []Pair{}
	}
	return hm
}

func (this *HashMap) Insert(k string, v int) {
	hash := doHASH(k)
	slot := &this.slots[hash]
	idx := slices.IndexFunc(*slot, func(p Pair) bool {
		return p.key == k
	})
	if idx != -1 {
		(*slot)[idx].val = v
	} else {
		*slot = append(*slot, Pair{k, v})
	}
}

func (this *HashMap) Remove(k string) {
	hash := doHASH(k)
	slot := &this.slots[hash]
	idx := slices.IndexFunc(*slot, func(p Pair) bool {
		return p.key == k
	})
	if idx == -1 {
		return
	}
	*slot = slices.Delete(*slot, idx, idx+1)
}

func solvePart1(tokens []string) int {
	ans := 0
	for _, token := range tokens {
		ans += doHASH(token)
	}
	return ans
}

func (this *HashMap) String() string {
	sb := strings.Builder{}
	for idx, slot := range this.slots {
		if len(slot) == 0 {
			continue
		}
		sb.WriteString(fmt.Sprintf("Box %d: %v\n", idx, slot))
	}
	return sb.String()
}

func solvePart2(tokens []string) int {
	hm := newHashMap()
	for _, token := range tokens {
		if strings.Contains(token, "=") {
			pair := strings.Split(token, "=")
			hm.Insert(pair[0], aoc.StrToInt(pair[1]))
		} else if strings.Contains(token, "-") {
			k := strings.TrimSuffix(token, "-")
			hm.Remove(k)
		}
	}

	// fmt.Println(hm)

	ans := 0
	for boxIdx, slot := range hm.slots {
		for slotIdx, label := range slot {
			// fmt.Printf("%s: %d * %d * %d = %d\n", label.key, (boxIdx + 1), (slotIdx + 1), label.val, (boxIdx+1)*(slotIdx+1)*label.val)
			ans += (boxIdx + 1) * (slotIdx + 1) * label.val
		}
	}
	return ans
}

func main() {
	input := aoc.GetInputForDay(15, false)
	tokens := strings.Split(input, ",")
	p1 := solvePart1(tokens)
	p2 := solvePart2(tokens)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
