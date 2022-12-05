#!/usr/bin/python3

import sys
from itertools import permutations

persons = dict()
parse = lambda x, d : d.update( {(x[0], x[10]) : int(x[3]) if x[2] == "gain" else -1*int(x[3])} )
[ parse(x[:len(x)-2].split(), persons) for x in open(sys.argv[1], "r").readlines() ]
#print(persons)

def behappy(persons):
  happiness = dict()
  for seats in permutations(set([x[0] for x in persons])):
    happy = 0
    for it in range(len(seats)):
      right = 0 if it == len(seats)-1 else it+1
      happy += persons[(seats[it], seats[right])] 
      left = len(seats)-1 if it == 0 else it-1
      happy += persons[(seats[it], seats[left])]
    happiness[str(seats)] = happy
  return max(happiness.values())  

print("partone {}".format(behappy(persons)))

for dude in set([x[0] for x in persons]):
  persons[("myself", dude)] = 0
  persons[(dude, "myself")] = 0

print("parttwo {}".format(behappy(persons)))

