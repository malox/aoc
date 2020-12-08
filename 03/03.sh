#!/usr/bin/python

import sys

def readfile(filename):
    file = open(filename, "r")
    str = ""
    for line in file:
        str += line.replace('\n', '')
    return str

def movePos(row, col, ch):
    if ch == '^' :
        row -= 1
    elif ch == 'v' :
        row += 1
    elif ch == '<' :
        col -= 1
    elif ch == '>' :
        col += 1
    else :
        print "unknown char " + ch
    return row, col

def visitHouses(input, houses):
    row = 0
    col = 0
    for ch in input:
        row, col = movePos(row, col, ch)
        pos = (row, col)
        x = houses.get(pos)
        if x == None:
            houses.update({pos : 1})
        else:
           houses[pos] += 1

def parseOne(input):
    print "input[" + input[:20] + str("..........]" if len(input) > 20 else "]")

    houses = {(0,0) : 1}
    visitHouses(input, houses)

    #print " - Houses : " + str(houses)
    print " - Houses count : " + str(len(houses))

def parseTwo(input):
    santa = ""
    robot = ""
    for idx in range(len(input)) :
        if idx % 2 == 0 :
            santa += input[idx]
        else:
            robot += input[idx]
    
    houses = {(0,0) : 1}
    visitHouses(santa, houses)
    visitHouses(robot, houses)

    #print " - Houses : " + str(houses)
    print " - Houses count : " + str(len(houses))

def parse(input):
    parseOne(input)
    parseTwo(input)

if __name__ == "__main__":
    filename = "input"
    if len(sys.argv) > 1 :
        filename = sys.argv[1]

    parse(">")
    parse("^>v<")
    parse("^v^v^v^v^v")
    parse(readfile(filename))
 
