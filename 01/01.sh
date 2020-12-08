#!/usr/bin/python

import sys

# commas removed from the input file
def readfile(filename):
    file = open(filename, "r")
    arr = []
    for line in file:
        for ch in line:
            arr.append(ch)
    return arr
    
def parse(arr):

    floor = 0
    idx = 0
    basement = 0

    for ch in arr:

        if ch == "(":
            floor += 1
        elif ch == ")":
            floor -= 1

        idx += 1
        if basement == 0 and floor == -1:
            basement = idx
            

    print "floor " + str(floor) + " - basement " + str(basement) 

if __name__ == "__main__":
    filename = "input"
    if len(sys.argv) > 1 :
        filename = sys.argv[1]
    arr = readfile(filename)
    parse(arr)
 
