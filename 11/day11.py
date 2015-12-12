#!/bin/python
import re
import time
import string

modified_ascii_lowercase = string.ascii_lowercase.replace("i", "").replace("o", "").replace("l", "")

def increment_letter(letter):
	if letter is'z':
		return 'a'

	index = modified_ascii_lowercase.index(letter)
	return modified_ascii_lowercase[index + 1]


def increment_string(password):
	if len(password) == 1:
		return increment_letter(password)

	if password[-1] == 'z':
		return increment_string(password[:-1]) + 'a'

	return password[:-1] + increment_letter(password[-1])

def valid_pass():
	pass

def find_next_valid_pass(password):
	max_duration = time.time() + 5 * 60
	while time.time() < max_duration:
		if valid_pass(password):
			print password
			return password
		password = increment_string(password)

	print "no password found in time"
	return ""

def test():
	assert(increment_letter("a") == "b")
	assert(increment_letter("z") == "a")
	assert(increment_string("a") == "b")
	assert(increment_string("z") == "a")
	assert(increment_string("xz") == "ya")
	assert(increment_string("zz") == "aa")
	assert(valid_pass("abcdffaa"))
	assert(valid_pass("ghjaabcc"))
	assert(find_next_valid_pass("abcdefgh") == "abcdffaa")
	assert(find_next_valid_pass("ghijklmn") == "ghjaabcc")

	print "all tests pass!"

def partone():
	password = "hxbxwxba"
	find_next_valid_pass(password)

def parttwo():
	input = "hxbxwxba"
	pass

if __name__ == '__main__':
    test()
    partone()
    # parttwo()
