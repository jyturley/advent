#!/bin/python
import re

def parse_instruction(instruction):
	m = re.compile(r"(\d+,\d+) through (\d+,\d+)")
	m = re.match("turn off (\d+,\d+) through (\d+,\d+)", instruction)
	if m:
		print m.group(1)
		print m.group(2)

	if instruction.startswith("turn on"):
		return

	if instruction.startswith("turn off"):
		return

	if instruction.startswith("toggle"):
		return

def test():
	parse_instruction("turn off 660,55 through 986,197")
	print "all tests pass!"


def partone():
	grid = [[0 for x in range(1000)] for x in range(1000)]


def parttwo():
	pass

if __name__ == '__main__':
    test()
    # partone()
    # parttwo()
