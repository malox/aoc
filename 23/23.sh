#!/usr/bin/python

import sys

# commas removed from the input file
def readfile(filename):
    file = open(filename, "r")
    arr = []
    for line in file:
        if len(line) > 0:
            arr.append(line.replace(',', ''))
    return arr
    
def parse(arr, start = 0):

    a = start
    b = 0
    idx = 0

    while idx < len(arr):
        #print " idx " + str(idx) + " - a " + str(a) + " b " + str(b) + " - instr " + arr[idx]
        cmds = arr[idx].split()
        if cmds[0] == "hlf":
            if cmds[1] == "a" :
                a /= 2
            else:
                b /= 2
            idx += 1
        elif cmds[0] == "tpl":
            if cmds[1] == "a" :
                a *= 3
            else:
                b *= 3
            idx += 1
        elif cmds[0] == "inc":
            if cmds[1] == "a" :
                a += 1
            else:
                b += 1
            idx += 1
        elif cmds[0] == "jmp":
            idx += int(cmds[1])
        elif cmds[0] == "jie":
            c = 0
            if cmds[1] == "a" :
                c = a%2
            else:
                c = b%2
            if c == 0 :
                idx += int(cmds[2])
            else:
                idx += 1
        elif cmds[0] == "jio":
            c = 0
            if cmds[1] == "a" :
                c = a
            else:
                c = b
            if c == 1 :
                idx += int(cmds[2])
            else:
                idx += 1
        else:
            print "unknown command : " + line

    print "a " + str(a) + " - b " + str(b)

if __name__ == "__main__":
    filename = "test"
    if len(sys.argv) > 1 :
        filename = sys.argv[1]
    arr = readfile(filename)
    parse(arr)
    parse(arr, 1)

