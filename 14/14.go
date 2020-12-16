package main

import (
	"fmt"
	"helper"
	"strings"
)

// -----------------------------------------------------------------

func clearBit(value int64, pos uint) int64 {
	mask := int64(^(1 << pos))
	value &= mask
	return value
}

func setBit(value int64, pos uint) int64 {
	mask := int64(1 << pos)
	value |= mask
	return value
}

func maskValue(mask string, value int64, skip byte, clear byte, set byte) int64 {
	size := len(mask)
	for idx := range mask {
		if mask[idx] == skip {
			continue
		}
		if mask[idx] == clear {
			value = clearBit(value, uint(size-idx-1))
		} else if mask[idx] == set {
			value = setBit(value, uint(size-idx-1))
		}
	}
	return value
}

type maskResetter func(mask string, masks *[]string)
type maskApplier func(memidx int64, value int64, line *string, masks *[]string, mem *map[int64]int64)

// -----------------------------------------------------------------

func resetOne(mask string, masks *[]string) {
	*masks = append(*masks, mask)
}

func applyOne(memidx int64, value int64, line *string, masks *[]string, mem *map[int64]int64) {
	(*mem)[memidx] = maskValue((*masks)[0], value, 'X', '0', '1')
}

// -----------------------------------------------------------------

func computeAllMasks(masks *[]string, mask string, maskidx int) {
	size := len(mask)
	for maskidx < size {
		if mask[maskidx] == 'X' {
			computeAllMasks(masks, mask[:maskidx]+"Z"+mask[maskidx+1:], maskidx+1)
			computeAllMasks(masks, mask[:maskidx]+"1"+mask[maskidx+1:], maskidx+1)
			break
		}
		maskidx++
	}
	if maskidx == size {
		*masks = append(*masks, mask)
	}
}

func resetTwo(mask string, masks *[]string) {
	computeAllMasks(masks, mask, 0)
}

func applyTwo(memidx int64, value int64, line *string, masks *[]string, mem *map[int64]int64) {
	for mdx := range *masks {
		(*mem)[maskValue((*masks)[mdx], memidx, '0', 'Z', '1')] = value
	}
}

// -----------------------------------------------------------------

func parse(lines []string, resetter maskResetter, applier maskApplier) {
	var masks []string
	mem := map[int64]int64{}

	for _, line := range lines {
		curr := strings.Split(line, " = ")
		if curr[0] == "mask" {
			masks = []string{}
			resetter(curr[1], &masks)
			// fmt.Println("mask ", curr[1], " - masks unfolded ", masks)
		} else {
			var memidx int64
			var value int64
			fmt.Sscanf(line, "mem[%d] = %d", &memidx, &value)
			applier(memidx, value, &(line), &masks, &mem)
		}
	}

	var sum int64
	for _, val := range mem {
		sum += val
	}
	fmt.Println("sum : ", sum)
}

func main() {
	lines := helper.FileAsStringList()
	parse(lines, resetOne, applyOne)
	parse(lines, resetTwo, applyTwo)
}
