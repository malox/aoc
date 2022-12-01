#!/usr/bin/env python3

import sys, copy
path = sys.argv[1]

# --------------------------- common

cave = []
for line in open(path,'r').readlines():
    cave.append([int(c) for c in line if c.isdigit()])

def visit(cave, x, y, risk, max_row, max_col):
    val = cave[x][y]
    minval = risk[x][y]
    if x :
        minval = min(minval, risk[x-1][y])
    if x+1 < max_row :
        minval = min(minval, risk[x+1][y])
    if y :
        minval = min(minval, risk[x][y-1])
    if y+1 < max_col :
        minval = min(minval, risk[x][y+1])
    risk[x][y] = minval + val
    

def findPath(cave, risk, max_row, max_col):
    risk[0][0] = 0 # start position comes with a zero - always!
    for x in range(max_row):
        for y in range(max_col):
            visit(cave, x, y, risk, max_row, max_col)

def doit(cave, max_row, max_col):
    risk = []
    for idx in range(max_row): 
        risk.append([pow(max_row,max_col)]*max_col)
    result = lambda risk, max_row, max_col : risk[max_row-1][max_col-1]-risk[0][0]
    while True:
        findPath(cave, risk, max_row, max_col)
        one = result(risk, max_row, max_col)
        findPath(cave, risk, max_row, max_col)
        two = result(risk, max_row, max_col)
        if one == two:
            return one

# --------------------------- part 1

print("part one {}".format(doit(cave, len(cave), len(cave[0]))))

# --------------------------- part 2

def expand(cave):
    newcave = [] 
    incr = lambda x,y : x+y if x+y <= 9 else x+y-9
    for row in cave:
        newrow = []
        for idx in range(5):
            newrow += [ incr(val, idx) for val in row ]
        newcave.append(newrow)
    import copy
    tmpcave = copy.deepcopy(newcave)
    for idx in range(1,5):
        for row in tmpcave:
            newcave.append([ incr(val, idx) for val in row ])
    return newcave

newcave = expand(cave)
print("part two {}".format(doit(newcave, len(newcave), len(newcave[0]))))

# --------------------------- 

