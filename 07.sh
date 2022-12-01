#!/usr/bin/env python3

import sys
path = sys.argv[1]

# parse commands
crabs = [ int(x) for x in open(path,'r').read()[:-1].split(',') ]
#print(crabs)

# --------------------------- part 1

def part_one(crabs):
    min_fuel = 99999999999999999999999999999999999999999999999999
    for crab in crabs:
        fuel = 0
        for cc in crabs:
            fuel += (crab-cc) if crab > cc else (cc-crab)
            if fuel > min_fuel : 
                break
        min_fuel = min(min_fuel, fuel)
    
    print("part one %d" % min_fuel)

part_one(crabs)

# --------------------------- part 2

def part_two(crabs):
    min_fuel = 99999999999999999999999999999999999999999999999999
    store = dict()
    for crab in crabs:
        fuel = 0
        for cc in crabs:
            distance = (crab-cc) if crab > cc else (cc-crab)
            if distance not in store.keys() :
                 store[distance] = sum([idx for idx in range(distance+1)])
            fuel += store[distance]
            if fuel > min_fuel : 
                break
        min_fuel = min(min_fuel, fuel)
    
    print("part two %d" % min_fuel)

part_two(crabs)

# --------------------------- 



