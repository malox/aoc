#!/usr/bin/env python3

import sys

def getStacks(lines):
  pos = [x for x in range(1, len(lines[0]), 4)]
  stacks = [[] for x in pos]

  idx = 0
  while lines[idx] != "":
    curr = lines[idx]
    idx += 1
    #print(curr)
    if "]" not in curr : continue
    for it in range(len(pos)):
      if curr[pos[it]] == " " : continue
      stacks[it].append(curr[pos[it]])
  idx += 1 # skip blank line
  #print(stacks)
  return idx, stacks

def parseStacks(lines, reverse):
  idx, stacks = getStacks(lines)
  for tt in range(idx, len(lines)):
    transf = lambda x : (int(x[1]), int(x[3])-1, int(x[5])-1)
    howmany, source, dest = transf(lines[tt].split())
    maybereverse = lambda flag, vect : reversed(vect) if flag else vect
    tmp = [ y for y in maybereverse(reverse, [stacks[source].pop(0) for xx in range(howmany)])] 
    stacks[dest] = tmp + stacks[dest]
    #print("{} ==> {} {} {}".format(lines[tt], howmany, source, dest))
    #print(stacks)
  print("top stacks: {}".format("".join(x[0] for x in stacks)))

lines = [x[:len(x)-1] for x in open(sys.argv[1],'r').readlines()]

parseStacks(lines, reverse=True)
parseStacks(lines, reverse=False)
