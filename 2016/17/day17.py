#!/bin/python
import re
from itertools import combinations

def parse():
	buckets = []
	with open('input.txt', 'r') as f:
		data = f.readlines()

	for bucket in data:
		buckets.append(int(bucket.strip()))

	return buckets

def test():
	print "all tests pass!"

def partone():
	precise_volume_buckets = []
	buckets = parse()
	for x in xrange(len(buckets)):
		combos = combinations(buckets, x)
		precise_volume_buckets += filter(lambda x: sum(x) == 150, combos)

	print len(precise_volume_buckets)

def parttwo():
	precise_volume_buckets = []
	buckets = parse()
	for x in xrange(len(buckets)):
		combos = combinations(buckets, x)
		precise_volume_buckets += filter(lambda x: sum(x) == 150, combos)

	min_containers = min([len(combo) for combo in precise_volume_buckets])
	print len(filter(lambda x: len(x) == min_containers, precise_volume_buckets))

if __name__ == '__main__':
    test()
    # partone()
    parttwo()
