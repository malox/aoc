#!/usr/bin/python

import sys

filename = sys.argv[1] if len(sys.argv) > 1 else "test"
lines = [(y[0], y[2], int(y[4])) for y in [x.split() for x in open(filename, "r").readlines()]]
print(lines)

nodes = dict()
name = "name"
visited = "visited"
friends = "friends"

def newnode(nodename):
  node = dict()
  node[name] = nodename
  node[visited] = False
  node[friends] = dict()
  return node

def linknodes(source, dest, weight):
  source[friends][dest[name]]=weight
  dest[friends][source[name]]=weight

def setupnodes(nodes, source, dest, weight):
  if not source in nodes:
    nodes[source] = newnode(source)
  if not dest in nodes:
    nodes[dest] = newnode(dest)
  linknodes(nodes[source], nodes[dest], weight)

for step in lines:
  source, dest, weight = step[0], step[1], step[2]
  setupnodes(nodes, source, dest, weight)

#print(nodes)  

def visits(nodes, shortest, current, start, weight):
  one = nodes[start]
  one[visited] = True
  current.append(start)
  if len(nodes) == len(current):
    shortest[str(current)] = weight
  else:
    for next in one[friends]:
      if nodes[next][visited]: continue
      visits(nodes, shortest, current, next, weight +one[friends][next])
  current.remove(start)  
  one[visited] = False


shortest = dict()
current = []

for start in nodes:
  visits(nodes, shortest, current, start, 0)

#print(current)
#print(shortest)

print("part one {}".format(min(shortest.values())))
print("part two {}".format(max(shortest.values())))
