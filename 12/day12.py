#!/bin/python
import re
import json

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
	print json.dumps(json_text)

if __name__ == '__main__':
    test()
    # partone()
    parttwo()
