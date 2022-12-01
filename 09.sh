#!/usr/bin/env python3

import sys
path = sys.argv[1]

# parse commands
points = []
for line in [ x[:-1] for x in open(path,'r').readlines() ]:
    points.append([int(y) for y in line])
#print("points {} ".format(points))

max_row, max_col = len(points), len(points[0])

# --------------------------- part 1

def check(points, x, y, max_row, max_col):
    val = 11
    for j in range(x-1, x+2):
        if j >= 0 and j < max_row:
            for k in range(y-1, y+2):
                if j == x and k == y:
                    continue
                if k >= 0 and k < max_col :
                    val = min(val, points[j][k])
    return points[x][y] < val

lows = []

for x in range(max_row):
    for y in range(max_col):
        if check(points, x, y, max_row, max_col):
            lows.append((x,y))
            
tot = 0
for low in lows:
    tot += points[low[0]][low[1]] + 1

print("part one %s" % str(tot))

# --------------------------- part 2

tot = []

def mapBasin(points, max_row, max_col, low):
    bassin = {low}
    explored = set()
    while bassin != explored:
        for pp in set(bassin):
            if pp in explored: continue
            x = pp[0]
            while x >= 0 and points[x][pp[1]] != 9:
                bassin.add((x, pp[1]))
                x -= 1
            x = pp[0]
            while x < max_row and points[x][pp[1]] != 9:
                bassin.add((x, pp[1]))
                x += 1
            y = pp[1]
            while y >= 0 and points[pp[0]][y] != 9:
                bassin.add((pp[0], y))
                y -= 1
            y = pp[1]
            while y < max_col and points[pp[0]][y] != 9:
                bassin.add((pp[0], y))
                y += 1
            explored.add(pp)
    #print("low {} - bassin size {} => {}".format(low, len(bassin), bassin))
    return len(bassin)

tot = sorted([mapBasin(points, max_row, max_col, x) for x in lows], reverse=True)[:3]
print("part two {}".format(tot[0]*tot[1]*tot[2]))

# --------------------------- 


