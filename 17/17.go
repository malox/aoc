package main

import (
	"fmt"
	"helper"
	"strings"
)

// -----------------------------------------------------------------

type layer []string
type space []layer
type hyperspace []space

// -----------------------------------------------------------------

func emptyRow(colsize int) string {
	return strings.Repeat(".", colsize)
}

func emptyLayer(rowsize, colsize int) layer {
	var emptylayer layer
	for idx := 0; idx < rowsize; idx++ {
		emptylayer = append(emptylayer, emptyRow(colsize))
	}
	return emptylayer
}

func emptySpace(zsize, rowsize, colsize int) space {
	var emptySpace space
	for idx := 0; idx < zsize; idx++ {
		emptySpace = append(emptySpace, emptyLayer(rowsize, colsize))
	}
	return emptySpace
}

// -----------------------------------------------------------------

func expandLayer(oldLayer *layer) layer {
	var newlayer layer
	ymax := len((*oldLayer)[0])
	newlayer = append(newlayer, emptyRow(ymax+2))
	for _, row := range *oldLayer {
		newlayer = append(newlayer, "."+row+".")
	}
	newlayer = append(newlayer, emptyRow(ymax+2))
	return newlayer
}

func expandSpace(oldSpace *space) space {
	var newspace space
	xmax := len((*oldSpace)[0])
	ymax := len((*oldSpace)[0][0])
	newspace = append(newspace, emptyLayer(xmax+2, ymax+2))
	for idx := range *oldSpace {
		newspace = append(newspace, expandLayer(&((*oldSpace)[idx])))
	}
	newspace = append(newspace, emptyLayer(xmax+2, ymax+2))
	return newspace
}

func expandHyperspace(oldHyperspace *hyperspace) hyperspace {
	var newHyperspace hyperspace
	zmax := len((*oldHyperspace)[0])
	xmax := len((*oldHyperspace)[0][0])
	ymax := len((*oldHyperspace)[0][0][0])
	newHyperspace = append(newHyperspace, emptySpace(zmax+2, xmax+2, ymax+2))
	for idx := range *oldHyperspace {
		newHyperspace = append(newHyperspace, expandSpace(&((*oldHyperspace)[idx])))
	}
	newHyperspace = append(newHyperspace, emptySpace(zmax+2, xmax+2, ymax+2))
	return newHyperspace
}

// -----------------------------------------------------------------

func activeNeighbors(zdx, xdx, ydx int, sp *space) int {
	actives := 0
	zmax := len(*sp)
	xmax := len((*sp)[0])
	ymax := len((*sp)[0][0])

	for z := zdx - 1; z <= zdx+1; z++ {
		if z < 0 || z >= zmax {
			continue
		}
		for x := xdx - 1; x <= xdx+1; x++ {
			if x < 0 || x >= xmax {
				continue
			}
			for y := ydx - 1; y <= ydx+1; y++ {
				if y < 0 || y >= ymax {
					continue
				}
				if z == zdx && x == xdx && y == ydx {
					continue
				}
				if (*sp)[z][x][y] == '#' {
					actives++
				}
			}
		}
	}
	return actives
}

func activeHyperNeighbors(wdx, zdx, xdx, ydx int, hypersp *hyperspace) int {
	actives := 0
	wmax := len(*hypersp)
	zmax := len((*hypersp)[0])
	xmax := len((*hypersp)[0][0])
	ymax := len((*hypersp)[0][0][0])

	for w := wdx - 1; w <= wdx+1; w++ {
		if w < 0 || w >= wmax {
			continue
		}
		for z := zdx - 1; z <= zdx+1; z++ {
			if z < 0 || z >= zmax {
				continue
			}
			for x := xdx - 1; x <= xdx+1; x++ {
				if x < 0 || x >= xmax {
					continue
				}
				for y := ydx - 1; y <= ydx+1; y++ {
					if y < 0 || y >= ymax {
						continue
					}
					if w == wdx && z == zdx && x == xdx && y == ydx {
						continue
					}
					if (*hypersp)[w][z][x][y] == '#' {
						actives++
					}
				}
			}
		}
	}
	return actives
}

// -----------------------------------------------------------------

func dumpSpace(sp *space) {
	fmt.Println(strings.Repeat("=", 23))

	for zdx, lay := range *sp {
		fmt.Printf("\nz=%d\n", zdx)
		for _, row := range lay {
			fmt.Printf("\t%s\n", row)
		}
		fmt.Println()
	}

	fmt.Println(strings.Repeat("=", 23))
}

func dumpHyperspace(hs *hyperspace) {
	fmt.Println(strings.Repeat("=", 23))

	for wdx, sp := range *hs {
		for zdx, lay := range sp {
			fmt.Printf("\nz=%d, w=%d\n", zdx, wdx)
			for _, row := range lay {
				fmt.Printf("\t%s\n", row)
			}
			fmt.Println()
		}
	}
	fmt.Println(strings.Repeat("=", 23))
}

// -----------------------------------------------------------------

func parse(lines []string, cycles int) {
	var sp space
	var lay layer
	for _, line := range lines {
		lay = append(lay, line)
	}
	sp = append(sp, lay)
	dumpSpace(&sp)

	for i := 0; i < cycles; i++ {
		newsp := expandSpace(&sp)
		oldsp := expandSpace(&sp)
		for zdx, oldlayer := range oldsp {
			for xdx, row := range oldlayer {
				for ydx := range row {
					actives := activeNeighbors(zdx, xdx, ydx, &oldsp)
					// fmt.Printf("act=%d for cube=%c z=%d x=%d y=%d - newrow=%s\n", actives, row[ydx], zdx, xdx, ydx, newsp[zdx][xdx])
					if row[ydx] == '.' && actives == 3 {
						newrow := newsp[zdx][xdx]
						newsp[zdx][xdx] = newrow[:ydx] + "#" + newrow[ydx+1:] // become active
					} else if row[ydx] == '#' && !(actives == 2 || actives == 3) {
						newrow := newsp[zdx][xdx]
						newsp[zdx][xdx] = newrow[:ydx] + "." + newrow[ydx+1:] // become inactive
					}
				}
			}
		}

		sp = newsp
		// dumpSpace(&sp)
	}

	count := 0
	for _, mlayer := range sp {
		for _, row := range mlayer {
			count += strings.Count(row, "#")
		}
	}
	fmt.Printf("\n\t Count: %d\n\n", count)
}

// -----------------------------------------------------------------

func hyperparse(lines []string, cycles int) {
	var hypersp hyperspace
	var sp space
	var lay layer
	for _, line := range lines {
		lay = append(lay, line)
	}
	sp = append(sp, lay)
	hypersp = append(hypersp, sp)
	dumpHyperspace(&hypersp)

	for i := 0; i < cycles; i++ {
		newhypersp := expandHyperspace(&hypersp)
		oldhypersp := expandHyperspace(&hypersp)
		for wdx, oldspace := range oldhypersp {
			for zdx, oldlayer := range oldspace {
				for xdx, row := range oldlayer {
					for ydx := range row {
						actives := activeHyperNeighbors(wdx, zdx, xdx, ydx, &oldhypersp)
						// fmt.Printf("act=%d for hypercube=%c w=%d z=%d x=%d y=%d - newrow=%s\n", actives, row[ydx], wdz, zdx, xdx, ydx, newhypersp[wdx][zdx][xdx])
						if row[ydx] == '.' && actives == 3 {
							newrow := newhypersp[wdx][zdx][xdx]
							newhypersp[wdx][zdx][xdx] = newrow[:ydx] + "#" + newrow[ydx+1:] // become active
						} else if row[ydx] == '#' && !(actives == 2 || actives == 3) {
							newrow := newhypersp[wdx][zdx][xdx]
							newhypersp[wdx][zdx][xdx] = newrow[:ydx] + "." + newrow[ydx+1:] // become inactive
						}
					}
				}
			}
		}

		hypersp = newhypersp
		// dumpHyperspace(&hypersp)
	}

	count := 0
	for _, mspace := range hypersp {
		for _, mlayer := range mspace {
			for _, row := range mlayer {
				count += strings.Count(row, "#")
			}
		}
	}
	fmt.Printf("\n\t HyperCount: %d\n\n", count)
}

// -----------------------------------------------------------------

func main() {
	lines := helper.FileAsStringList()
	parse(lines, 6)
	hyperparse(lines, 6)
}
