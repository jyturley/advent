#!/bin/python
import re
import time
import string

modified_ascii_lowercase = string.ascii_lowercase.replace("i", "").replace("o", "").replace("l", "")

def increment_letter(letter):
	if letter is'z':
		return 'a'

	try:
		index = modified_ascii_lowercase.index(letter)
	except:
		return '$'

	# print modified_ascii_lowercase[index + 1]
	return modified_ascii_lowercase[index + 1]


def increment_string(password):
	if len(password) == 1:
		return increment_letter(password)

	if password[-1] == 'z':
		return increment_string(password[:-1]) + 'a'

	return password[:-1] + increment_letter(password[-1])

def valid_first_req(password):
	if len(password) < 3:
		return False

	streak = 1
	for i in xrange(len(password) - 1):
		if streak == 3:
			return True

		if increment_letter(password[i]) == password[i+1]:
			streak += 1
			continue

		streak = 1

	return streak >= 3

def valid_second_req(password):
	return True

def valid_third_req(password):
	pairs = []

	i = 0
	for x in xrange(len(password) - 1):
		if password[i] == password[i+1]:
			pairs.append(password[i])
			i += 1
			break
		i += 1

	if not pairs:
		return False

	for y in xrange(i, len(password) - 1):
		print y
		if password[i] == password[i+1] and password[i] != pairs[0]:
			pairs.append(password[i:i+1])
		i += 1

	print password, pairs
	return len(pairs) >= 2

def valid_pass(password):
	return valid_first_req(password) and valid_second_req(password) and valid_third_req(password)

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

	assert(valid_first_req("abc"))
	# assert(valid_first_req("hijklmn"))
	# assert(not valid_second_req("hijklmn"))

	assert(not valid_first_req("abbceffg"))
	assert(valid_third_req("abbceffg"))

	assert(not valid_third_req("abbcegjk"))

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
    # partone()
    # parttwo()
