#!/usr/bin/env python3

import sys
path = sys.argv[1]

# parse commands
def getSquids(path):
    squids = []
    for line in [ x[:-1] for x in open(path,'r').readlines() ]:
        squids.append([int(y) for y in line])
    return squids

squids = getSquids(path)
max_row, max_col = len(squids), len(squids[0])
flashes = 0

# --------------------------- common

def dump(squids):
    print("-"*len(squids[0]))
    for row in squids:
        print("".join([str(x) for x in row]))
    print("-"*len(squids[0]))

dump(squids)

def incr(squids, x, y, max_row, max_col):
    global flashes
    if x < 0 or x >= max_row: return
    if y < 0 or y >= max_col: return
    squids[x][y] += 1
    if squids[x][y] == 10:
        flashes += 1
        for idx in range (-1,2):
            for jdx in range (-1,2):
                if idx == 0 and jdx == 0: continue
                incr(squids, x+idx, y+jdx, max_row, max_col)

def normalise(squids, max_row, max_col):
    for x in range (max_row):
        for y in range (max_col):
            if squids[x][y] > 9:
                squids[x][y] = 0

def bigFlash(squids):
    for x in squids:
        for y in x:
            if y != 0:
                return False
    return True

# --------------------------- part 1

#
#iterations = int(sys.argv[2])
#for it in range(iterations):
#    for x in range (max_row):
#        for y in range (max_col):
#            incr(squids, x, y, max_row, max_col)
#    normalise(squids, max_row, max_col)
#    if it < 10 or it%10 == 0:
#        print("Iteration {}".format(it+1))
#        dump(squids)
#dump(squids)
#print("flashes = {}".format(flashes))

# --------------------------- part 2

#squids = getSquids(path)
#flashes = 0
rounds = 0
while not bigFlash(squids):
    for x in range (max_row):
        for y in range (max_col):
            incr(squids, x, y, max_row, max_col)
    normalise(squids, max_row, max_col)
    rounds += 1
    if rounds == 100:
        print("flashes = {}".format(flashes))
print("rounds = {}".format(rounds))

# --------------------------- 


