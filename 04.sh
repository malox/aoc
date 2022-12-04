#!/usr/bin/env python3

import re
import sys
path = sys.argv[1]
conv = lambda nums : [(int(nums[0]), int(nums[1])), (int(nums[2]), int(nums[3]))]
lines = [conv(x) for x in (re.split(",|-", x[:len(x)-1]) for x in open(path,'r').readlines())]
#print(lines)

total = 0
for task in lines:
    one, two = task[0], task[1]
    if two[0] >= one[0] and two[1] <= one[1] :
        total += 1
    elif one[0] >= two[0] and one[1] <= two[1] :
        total += 1

print("part one: {}".format(total))



total = 0
for task in lines:
    one, two = task[0], task[1]
    if two[0] >= one[0] and one[1] >= two[0] :
        total += 1
    elif one[0] >= two[0] and two[1] >= one[0] :
        total += 1

print("part two: {}".format(total))
