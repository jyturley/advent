#!/bin/python
import re

def look_and_say(input):
	streak = 1
	previous = input[0]
	say = ""

	if len(input) == 1:
		return "1%s" % input[0]

	for i in xrange(1, len(input)):
		if input[i] is not previous:
			say += str(streak) + previous
			previous = input[i]
			streak = 1
		else:
			streak += 1

	say += str(streak) + previous

	return say

def test():
	assert(look_and_say("1") == "11")
	assert(look_and_say("11") == "21")
	assert(look_and_say("21") == "1211")
	assert(look_and_say("1211") == "111221")
	assert(look_and_say("111221") == "312211")
	print "all tests pass!"

def partone():
	input = "1113122113"
	iterations = 40

	for x in xrange(iterations):
		input = look_and_say(input)

	print len(input)

def parttwo():
	input = "1113122113"
	iterations = 50

	for x in xrange(iterations):
		input = look_and_say(input)

	print len(input)

if __name__ == '__main__':
    test()
    partone()
    parttwo()
