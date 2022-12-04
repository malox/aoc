#!/usr/bin/env python3

import re
import sys
path = sys.argv[1]

lines = open(path,'r').read().split(", ")
#print(lines)

# 1=N 2=E 3=S 4=W
dir = 1
x,y = 0,0

calc = lambda x : int(x[1:])

found = False
path = set()
path.add((0,0))

def pathAdd(a,b):
    global path, found
    #print("pos ({},{})".format(a,b))
    if not found and (a,b) in path:
        print("part two {}".format(abs(a)+abs(b)))
        found = True
    path.add((a,b))

def move(s):
    global dir, x, y, path, found
    xold,yold = x,y
    #print("================================================")
    #print("{} dir={} xold={} yold={}".format(s, dir, x, y))
    
    turn = s[0]
    if turn == 'R':
        dir += 1
        if dir > 4:
            dir = 1
    if turn == 'L':
        dir -= 1
        if dir < 1:
            dir = 4

    if dir == 1:
        y += calc(s)
        for tt in range(yold, y+1):
            if (xold,yold) != (x,tt) : pathAdd(x,tt)
        
    elif dir == 2:
        x += calc(s)
        for tt in range(xold, x+1):
            if (xold,yold) != (tt,y) : pathAdd(tt,y)

    elif dir == 3:
        y -= calc(s)
        for tt in range(yold, y-1, -1):
            if (xold,yold) != (x,tt) : pathAdd(x,tt)

    elif dir == 4:
        x -= calc(s)
        for tt in range(xold, x-1, -1):
            if (xold,yold) != (tt,y) : pathAdd(tt,y)

    #print("{} dir={} x={} y={}".format(s, dir, x, y))
    #print("================================================")

for step in lines:
    move(step)
    
print("part one {}".format(abs(x)+abs(y)))