#!/usr/bin/env python3

import sys
path = sys.argv[1]
lines = [x[:len(x)-1] for x in open(path,'r').readlines()]

def score(c):
    return (ord(c) - 96) if (ord(c) > 92) else (ord(c) - 64 + 26)

def check(s1, s2):
    s = ""
    for x in set(s1) - (set(s1) - set(s2)):
        s += x
    return s


total = 0
tuples = [(x[:int(len(x)/2)], x[int(len(x)/2):len(x)]) for x in lines]
for tup in tuples:
    total += score(check(tup[0], tup[1])[0])

print("part one: {}".format(total))


tuples = []
for i in range(0, len(lines)-1, 3):
    tuples.append((lines[i], lines[i+1], lines[i+2]))

total = 0
for tup in tuples:
    total += score(check(check(tup[0], tup[1]), check(tup[0], tup[2]))[0])

print("part two: {}".format(total))
