#!/usr/bin/python

import sys

def readfile(filename):
    file = open(filename, "r")
    arr = []
    for line in file:
        if len(line) > 0:
            boxstr = line.replace('\n', '').split('x')
            boxint = []
            for face in boxstr :
                boxint.append(int(face))
                boxint.sort()
            arr.append(boxint)
    return arr

def parse(arr):

    paper = 0
    ribbon = 0
    for box in arr:
        paper += 2*box[0]*box[1] + 2*box[1]*box[2] + 2*box[0]*box[2] + box[0]*box[1]
        ribbon += 2*box[0] + 2*box[1] + box[0]*box[1]*box[2]

    print "paper " + str(paper) + " - ribbon " + str(ribbon)

if __name__ == "__main__":
    filename = "input"
    if len(sys.argv) > 1 :
        filename = sys.argv[1]
    arr = readfile(filename)
    parse(arr)
 
