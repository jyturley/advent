#!/bin/python
import re
import numpy
from collections import deque

def parse_instruction(instruction, circuit):
	signal, wire = instruction.split(" -> ")

	m = re.match(r"^(\d+)$", signal)
	if m:
		circuit[wire] = numpy.uint16(m.group(1))
		return True
	
	m = re.match(r"^([a-z][a-z]?)$", signal)
	if m:
		if m.group(1) in circuit:
			circuit[wire] = circuit[m.group(1)]
		else:
			return False
		return True

	m = re.match(r"NOT (\w+)", signal)
	if m:
		if m.group(1) in circuit:
			circuit[wire] = ~ numpy.uint16(circuit[m.group(1)])
		else:
			return False
		return True
	
	m = re.match(r"([\w\d]+) AND (\w+)", signal)
	if m:
		left_and = None 
		if m.group(1) == "1":
			left_and = 1
		elif m.group(1) in circuit:
			left_and = numpy.uint16(circuit[m.group(1)])
		else:
			return False

		if m.group(2) in circuit:
			circuit[wire] = left_and & numpy.uint16(circuit[m.group(2)])
		else:
			return False
		return True
	
	m = re.match(r"(\w+) OR (\w+)", signal)
	if m:
		if m.group(1) in circuit and m.group(2) in circuit:
			circuit[wire] = circuit[m.group(1)] | circuit[m.group(2)]
		else:
			return False
		return True
	
	m = re.match(r"(\w+) LSHIFT (\d+)", signal)
	if m:
		if m.group(1) in circuit:
			circuit[wire] = circuit[m.group(1)] << numpy.uint16(m.group(2))
		else:
			return False
		return True
			
	
	m = re.match(r"(\w+) RSHIFT (\d+)", signal)
	if m:
		if m.group(1) in circuit:
			circuit[wire] = circuit[m.group(1)] >> numpy.uint16(m.group(2))
		else:
			return False
		return True
			
	return False


def test():
	circuit = {}
	parse_instruction("123 -> x", circuit)
	parse_instruction("456 -> y", circuit)
	parse_instruction("x AND y -> d", circuit)
	parse_instruction("x OR y -> e", circuit)
	parse_instruction("x LSHIFT 2 -> f", circuit)
	parse_instruction("y RSHIFT 2 -> g", circuit)
	parse_instruction("NOT x -> h", circuit)
	parse_instruction("NOT y -> i", circuit)

	print circuit

	assert(circuit['d'] == 72)
	assert(circuit['e'] == 507)
	assert(circuit['f'] == 492)
	assert(circuit['g'] == 114)
	assert(circuit['h'] == 65412)
	assert(circuit['i'] == 65079)
	assert(circuit['x'] == 123)
	assert(circuit['y'] == 456)

	parse_instruction("d -> a", circuit)
	parse_instruction("a -> ax", circuit)
	parse_instruction("ax -> k", circuit)
	assert(circuit['a'] == 72)
	assert(circuit['k'] == 72)

	print "all tests pass!"

def follow_insctructions(text_file):
	circuit = {}
	instructions = []
	with open( text_file, 'r') as f:
		instructions = f.readlines()

	instructions_q = deque(instructions)
	while instructions_q:
		instruction = instructions_q.popleft().rstrip()
		if not parse_instruction(instruction, circuit):
			instructions_q.append(instruction)

	print sorted(circuit.items())
	print circuit['a']
	return circuit

def partone():
	follow_insctructions("input_1.txt")

def parttwo():
	follow_insctructions("input_2.txt")

if __name__ == '__main__':
    test()
    partone()
    parttwo()
