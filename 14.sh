#!/usr/bin/env python3

import sys, copy
path = sys.argv[1]
iterations = int(sys.argv[2])

# --------------------------- common

polymer, rawrules = open(path,'r').read().split('\n\n')
rules = dict()
for y,z in [tuple(x.split(' -> ')) for x in rawrules[:-1].split('\n')] : rules[y] = z

print("polymer : {}".format(polymer))
print("rules ({}) : {}".format(len(rules), rules))

# --------------------------- part 1 Not optimized

if iterations < 15 :
    for it in range(iterations):
        newpolymer = polymer[0]
        for idx in range(len(polymer)-1):
            if polymer[idx:idx+2] in rules :
                newpolymer += rules[polymer[idx:idx+2]] + polymer[idx+1]
                #print("  newpolymer {} - chars {} {} - rule {}".format(newpolymer, polymer[idx], polymer[idx+1], rules[polymer[idx:idx+2]]))
        polymer = newpolymer
        print("after step {} polymer : {}".format(it+1, len(polymer)))
        
    from collections import Counter
    minval, maxval = len(polymer), 0
    for val in Counter(polymer).values():
        minval, maxval = min(minval, val), max(maxval, val)
    print("Len {} - Counter {}".format(len(polymer), Counter(polymer)))
    print("part one {}".format(maxval-minval))

# --------------------------- part 2

if iterations > 15 :
    pairs = dict()
    for idx in range(len(polymer)-1):
        pairs[polymer[idx:idx+2]] = 1
    for it in range(iterations):
        new_pairs = dict()
        for key, count in pairs.items():
            rule = rules[key]
            one = key[0] + rule 
            two = rule + key[1]
            new_pairs[one] = new_pairs.get(one,0) + count
            new_pairs[two] = new_pairs.get(two,0) + count
        pairs = new_pairs
        print("after step {} pairs : {}".format(it+1, pairs))

    charcounter = dict()
    for key, count in pairs.items():
        charcounter[key[0]] = charcounter.get(key[0],0) + count
        charcounter[key[1]] = charcounter.get(key[1],0) + count
    minval, maxval = sum([val for val in charcounter.values()]), 0
    for val in charcounter.values():
        minval, maxval = min(minval, val), max(maxval, val)
    print("charcounter : {}".format(charcounter))
    result = (maxval-minval)/2
    import math
    print("part two {} or {}".format(math.floor(result), math.ceil(result)))
   
    # the return value should be rounded to lowest/highest integer, dunno why :D    

# --------------------------- 

