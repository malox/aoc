#!/usr/bin/env python3

import sys
import json

intify = lambda z : [int(y) for y in z]
lines = [intify(x) for x in open(sys.argv[1],'r').read().splitlines()]

#dump = lambda vect : print("\n".join([" ".join([str(j) for j in i]) for i in vect]))
#dump(lines)

rows, cols = len(lines), len(lines[0])
visible = [[0 for j in range(cols)] for i in range(rows)]

def check(i, j, curr):
  global lines, visible
  if lines[i][j] > curr:
    visible[i][j] = 1
    curr = lines[i][j]
  return curr

def parse(outer, inner):
  global lines, visible
  for i in outer:
    curr = -1
    for j in inner:
      curr = check(i, j, curr)
  for j in inner:
    curr = -1
    for i in outer:
      curr = check(i, j, curr)

parse(range(rows), range(cols))
parse(range(rows-1, 0, -1), range(cols-1, 0, -1))
#dump(visible)
print("partone {}".format(sum([it.count(1) for it in visible])))

def scenic(x, y):
  global rows, cols, lines
  tmp = [0, 0, 0, 0]
  self = lines[x][y]

  for i in range(x-1, -1, -1):
    tmp[0] += 1
    if lines[i][y] >= self: break
  
  for i in range(x+1, rows):
    tmp[1] += 1
    if lines[i][y] >= self: break

  for j in range(y-1, -1, -1):
    tmp[2] += 1
    if lines[x][j] >= self: break

  for j in range(y+1, cols):
    tmp[3] += 1
    if lines[x][j] >= self: break

  tot = 1
  for n in tmp:
    if n > 0 : tot *= n 

  return tot

scenary = [[scenic(i, j) for j in range(1, cols-1)] for i in range(1, rows-1)]
#dump(scenary)
print("parttwo {}".format(max([max(it) for it in scenary])))

