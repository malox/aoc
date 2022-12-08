#!/usr/bin/env python3

import sys
import json

lines = [x.split() for x in open(sys.argv[1],'r').read().splitlines()]
#print(lines)

disk = dict()
current = []
size = "size"

def goto(disk, current):
  tmp = disk
  for tt in current:
    if tt not in tmp:
      tmp[tt] = dict()
      tmp[tt][size] = 0
    tmp = tmp[tt]
  return tmp

def up(disk, current, pos):
  old = current.pop(len(current)-1)
  pos = goto(disk, current)
  pos[size] += pos[old][size]
  return pos

pos = disk
for xx in lines:
  if xx[0] == "$":
    if xx[1] == "cd":
      if xx[2] == "..":
        old = current.pop(len(current)-1)
        pos = goto(disk, current)
        pos[size] += pos[old][size]
      else:
        current.append(xx[2])
        pos = goto(disk, current)
    # ls skipped
  else:
    if xx[0] != "dir":
      #print(f"xx={xx} pos={pos}")
      pos[size] += int(xx[0])
  #print(f"dir={current} disk={disk}")

while len(current) > 0:
  old = current.pop(len(current)-1)
  pos = goto(disk, current)
  if size in pos and size in pos[old]:
    pos[size] += pos[old][size] 

#print(f"{json.dumps(disk, indent=2)}")

def count(node):
  total = 0
  #print(f" >>> DBG node={node} total={total}")
  for ff in node :
    if ff == size : continue
    total += count(node[ff])
  if size in node and node[size] <= 100000:
    total += node[size]
  #print(f" <<< DBG node={node} total={total}")
  return total

print(f"part one = {count(disk)}")

freespace = 70000000 - disk["/"][size]
missing = 30000000 - freespace
print(f"freespace={freespace} missing={missing}")

minsize = disk["/"][size]
def claimSpace(node):
  global minsize, missing
  for ff in node:
    if ff == size : continue
    claimSpace(node[ff])
  if size in node and node[size] >= missing:
    minsize = min([minsize, node[size]])
claimSpace(disk)
print(f"min={minsize}")
