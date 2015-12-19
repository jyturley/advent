#!/bin/python
import re
import sqlite3

def parse_to_db(text_file, db):
	with open(text_file, 'r') as f:
		data = f.readlines()

	conn = sqlite3.connect(db)
	c = conn.cursor()
	c.execute("SELECT 1 FROM Sues LIMIT 2")
	print c.fetchall()
	if c.fetchone():
		print "Table exists"
	else:
		print "Table does not exist, creating new one..."
		c.execute("""CREATE TABLE Sues (
			children integer,
			cats integer,
			samoyeds integer,
			pomeranians integer,
			akitas integer,
			viszslas integer,
			goldfish integer,
			trees integer,
			cars integer,
			perfumes integer)""")
		db.commit()

	c.execute("SELECT * FROM Sues")
	print c.fetchall()


	r = re.compile(r'Sue (\d+): (\w+: \d+), (\w+: \d+), (\w+: \d+)')
	for sue in data:
		m = re.match(r, sue)
		if m:
			id = m.group(1)
			clues = [clue.split(": ") for clue in m.group(2, 3, 4)] # ([cats, 7], [dogs, 4], [cows, 5])


def test():
	print "all tests pass!"

def partone():
	pass

def parttwo():
	parse_to_db('input.txt', 'day16.db')
	# conn = sqlite3.connect('day16.db')
	# c = conn.cursor()

if __name__ == '__main__':
    # test()
    # partone()
    parttwo()
