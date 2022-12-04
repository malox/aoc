#!/usr/bin/env python3

import sys
path = sys.argv[1]

elves = []
current = 0
for y in open(path,'r').readlines():
    if y != "\n":
        current += int(y)
    else:
        elves.append(current)
        current = 0

elves.sort()
print(elves[-1:])
print(sum(elves[-3:]))
