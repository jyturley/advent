#!/bin/python
import md5
import time

def find_hash(secret_key):
	max_duration = time.time() + 5 * 60 #5 minutes
	so_far = 0
	while time.time() < max_duration:
		try_key = secret_key + str(so_far)

		if md5.new(try_key).hexdigest()[:5] == "00000":
			return so_far
		else:
			so_far += 1

def test():
	assert(find_hash("abcdef") == 609043)
	assert(find_hash("pqrstuv") == 1048970)

def partone():
	hash = "iwrupvqb"
	print "finding answer for key: %s" % hash
	print find_hash(hash)

def parttwo():
    pass

if __name__ == '__main__':
    # test()
    partone()
