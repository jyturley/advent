#!/bin/python
import re
import json

def find_target(json, target, targets):
	if isinstance(json, dict):
		if target not in json.values():
			for item in json.values():
				find_target(item, target, targets)
		else:
			targets.append(json)

	if isinstance(json, list):
		for item in json:
			find_target(item, target, targets)


def sum_numbers(input):
	 return reduce(lambda x,y: x+ y, [int(x) if x else 0 for x in re.sub(r'[^0-9\-]', ' ', input).split()])

def sum_numbers_from_file(text_file):
	sum = 0
	with open(text_file, 'r') as f:
		data = f.readlines()

	for line in data:
		sum += sum_numbers(line)

	return sum

def test():
	print "all tests pass!"

def partone():
	print sum_numbers_from_file('input.json')

def parttwo():
	with open('input.json', 'r') as f:
		data = f.read()

	json_text = json.loads(data)

	red_statements = []
	find_target(json_text, 'red', red_statements)
	json_red = json.dumps(red_statements)

	print sum_numbers_from_file('input.json') - sum_numbers(json_red)


if __name__ == '__main__':
    test()
    # partone()
    parttwo()
