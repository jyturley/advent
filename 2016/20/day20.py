#!/bin/python
import re
import time
from collections import defaultdict

def total_presents(house_num):
    total_presents = 0
    possible_elves = (x for x in xrange(1, house_num + 1))
    divisible_numbers = filter(lambda x: house_num % x == 0, possible_elves)
    return sum(divisible_numbers) * 10

def test():
    assert(total_presents(1) == 10)
    assert(total_presents(2) == 30)
    assert(total_presents(3) == 40)
    assert(total_presents(4) == 70)
    assert(total_presents(5) == 60)
    assert(total_presents(6) == 120)
    assert(total_presents(7) == 80)
    assert(total_presents(8) == 150)
    assert(total_presents(9) == 130)
    print total_presents(803881)
    print total_presents(710640)
    print total_presents(662000)

    t = time.time()
    # total_presents(33100000)
    print time.time() - t

    print "all tests pass!"

# 776160
# 803881
# 1000081
# 1010017
# 1030081
# 1050000
# 1080001
# 1090081 too high
def partone():
    presents = defaultdict(int)
    upper_bound = 803881
    for elf in xrange(1, upper_bound):
        for house in xrange(elf, upper_bound, elf):
            presents[house] += 10 * elf

        if presents[elf] > 33100000:
            print total_presents(house)
            print house
            break

def parttwo():
    presents = defaultdict(int)
    upper_bound = 803881
    for elf in xrange(1, upper_bound):
        for house in xrange(elf, elf * 51, elf):
            presents[house] += 11 * elf

        if presents[elf] > 33100000:
            print elf, presents[elf]
            break

if __name__ == '__main__':
    # test()
    # partone()
    parttwo()
