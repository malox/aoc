#!/usr/bin/env python3

import sys
path = sys.argv[1]

class Fish:
    def __init__(self, val = 8):
        self.val = val
    def __repr__(self):
        return "%d" % self.val

    def incr(self, newFish):
        self.val -= 1
        if self.val < 0:
            self.val = 6
            newFish.append(Fish())

# parse commands
fishes = [ Fish(int(x)) for x in open(path,'r').read()[:-1].split(',') ]
print(fishes)

# --------------------------- part 1

for day in range(82):
    newFishes = []
    for fish in fishes:
        fish.incr(newFishes)
    fishes += newFishes
    #print("{} {}".format(str(day+1),str(fishes)))
    if day == 17 or day == 79 or day == 255 :
        print("day {} - len {}".format(str(day+1), str(len(fishes))))

# --------------------------- part 2

vals = [ int(x) for x in open(path,'r').read()[:-1].split(',') ]

# copied from https://gist.github.com/suntriber/6c3638be973b14b7bab5cc2248227b3b
def solve(n, vals):
    from collections import Counter
    fish_dict = Counter(vals)

    for _ in range(n):
        # decrease keys by one each iteration
        fish_dict = {k-1: v for k, v in fish_dict.items()}

        # add new fishes at key 8 from fish all that has reached -1
        new_fishes = fish_dict.get(-1, 0)
        fish_dict[8] = new_fishes

        # reset all -1 fish to 6
        reset_fish = fish_dict.pop(-1, 0)
        six_fishes = fish_dict.get(6, 0)
        fish_dict[6] = reset_fish + six_fishes
    
    return sum(fish_dict.values())

print("80 %d" % solve(80, vals))
print("256 %d" % solve(256, vals))

# --------------------------- 



