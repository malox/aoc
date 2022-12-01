#!/usr/bin/env python3

import sys

path = sys.argv[1]
input = [int(x) for x in open(path,'r').readlines()]

def count(points):
    previous = -1
    incr = 0
    for val in points:
        if previous != -1:
            if val > previous:
                incr += 1
        previous = val
    print(incr)

measure = []
for idx in range(len(input)-2):
    measure.append(input[idx]+input[idx+1]+input[idx+2])

count(input)
count(measure)
