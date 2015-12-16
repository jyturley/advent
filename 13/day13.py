#!/bin/python
import re

def calculate_table(table, preferences):
	total_happiness = 0
	table = table + table[0:2]
	for x in xrange(len(table) - 2):
		left = table[x]
		person = table[x + 1]
		right = table[x + 2]
		total_happiness += preferences[person][left] + preferences[person][right]

	return total_happiness

def parse(input):
	with open(input, 'r') as f:
		data = f.readlines()

	preferences = {}
	family = set()
	r = re.compile(r'(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+).')
	for stat in data:
		match = re.match(r, stat)
		if match:
			main_person = match.group(1)
			happiness_amount = int(match.group(3)) if "gain" in stat else -int(match.group(3))
			adjacent_person = match.group(4)

			family.add(main_person)

			if main_person not in preferences:
				preferences[main_person] = {}
			preferences[main_person][adjacent_person] = happiness_amount

	return family, preferences


def test():
	family, preferences = parse('test_input.txt')
	assert(calculate_table(["David", "Alice", "Bob", "Carol"], preferences) == 330)

	print "all tests pass!"

def partone():
	family, preferences = parse('input.txt')

def parttwo():
	pass

if __name__ == '__main__':
    test()
    partone()
    # parttwo()
