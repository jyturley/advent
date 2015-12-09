#!/bin/python
import re
import os

def difference_in_space(text_file):
	return find_num_chars(text_file) - find_num_chars(unescape(text_file))

def find_num_chars(text_file):
	sofar = 0
	with open(text_file, 'r') as f:
		for lines in f.readlines():
			sofar += len(lines.strip())
	return sofar

def unescape(text_file):
	with open(text_file, 'r') as f1:
		orig_file_data = f1.readlines()

	new_file_name = "replaced_" + os.path.basename(text_file)
	with open(new_file_name, 'w') as f2:
		for line in orig_file_data:
			line = re.sub(r'^"|"$', '', line)
			line = line.replace("\\\"", "$")
			line = line.replace("\\\\", "^")
			line = re.sub( r"\\x[0-9a-fA-F][0-9a-fA-F]", "%", line)

			f2.write(line)

	return new_file_name

def test():
	test_file = 'test_input.txt'
	assert(find_num_chars(test_file) == 23)
	assert(find_num_chars(unescape(test_file)) == 11)
	assert(difference_in_space(test_file) == 12)
	print "all tests pass!"

def partone():
	print difference_in_space("input.txt")

def parttwo():
	pass

if __name__ == '__main__':
    test()
    partone()
    # parttwo()
