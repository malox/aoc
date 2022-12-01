#!/usr/bin/env python3

import sys, copy
path = sys.argv[1]

# --------------------------- common

class Cave:
    def __init__(self, name):
        self.name = name
        self.small = name.islower()
        self.visits = 1
        self.next = []
    def link(self, other):
        self.next.append(other.name)
        other.next.append(self.name)
    def __repr__(self):
        return "{}({})[{}]".format(self.name, self.small, self.next)

def getCave(caves, name):
    if name not in caves:
        caves[name] = Cave(name)
    return caves[name]

# parse commands
caves = dict()
for line in [ x[:-1].split('-') for x in open(path,'r').readlines() ]:
    getCave(caves, line[0]).link(getCave(caves, line[1]))

#print("{}".format(caves))

# --------------------------- part 1

def find(caves, start, end, currentpath, allpaths):
    if start == end :
        currentpath.append(end)
        allpaths.add("".join(currentpath))
        currentpath.pop()
        return
    cave = caves[start]
    if cave.small and cave.visits <= 0 :
        return
    currentpath.append(start)
    cave.visits -= 1
    for next in cave.next:
        find(caves, next, end, currentpath, allpaths)
    cave.visits += 1
    currentpath.pop()

allpaths = set()
find(caves, 'start', 'end', [], allpaths)

#for x in allpaths: print(x)
print("part one {}".format(len(allpaths)))

# --------------------------- part 2

allpaths = set()
for _, cave in caves.items():
    if not cave.small : continue
    if cave.name == 'start' or cave.name == 'end' : continue
    cave.visits += 1
    find(caves, 'start', 'end', [], allpaths)
    cave.visits -= 1

#for x in allpaths: print(x)
print("part two {}".format(len(allpaths)))

# --------------------------- 

