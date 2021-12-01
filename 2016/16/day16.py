#!/bin/python
import re
import sqlite3
import sys

def find_correct_sue(connection):
	c = connection.cursor()
	# query = 'SELECT rowid FROM Sues WHERE children=3'
	for row in c.execute('''SELECT rowid, children, cats, samoyeds,
		                           pomeranians, akitas, vizslas, goldfish,
		                           trees, cars, perfumes
		                    FROM Sues
		                    WHERE (children = 3  OR children = 0) 
		                      AND (cats > 7 OR cats = 0)
		                      AND (pomeranians < 3 OR pomeranians = 0)
		                      AND akitas = 0
		                      AND vizslas = 0
		                      AND samoyeds = 2
		                      AND (trees > 3 OR trees = 0)
		                      AND (goldfish < 5  OR goldfish = 0)'''):

		print row

def cleanup(connection):
	try:
		c = connection.cursor()
		c.execute('DROP TABLE Sues')
		connection.commit()
	except:
		pass

def create_db(db):
	conn = sqlite3.connect(db)
	c = conn.cursor()
	c.execute("""CREATE TABLE Sues (
		children integer,
		cats integer,
		samoyeds integer,
		pomeranians integer,
		akitas integer,
		vizslas integer,
		goldfish integer,
		trees integer,
		cars integer,
		perfumes integer)""")
	conn.commit()

	for x in xrange(500):
		c.execute('INSERT INTO Sues Values (0, 0, 0, 0, 0, 0, 0, 0, 0, 0)')
		
	conn.commit()
	conn.close()

def parse_to_db(text_file, connection):
	with open(text_file, 'r') as f:
		data = f.readlines()

	r = re.compile(r'Sue (\d+): (\w+: \d+), (\w+: \d+), (\w+: \d+)')
	c = connection.cursor()
	for sue in data:
		m = re.match(r, sue)
		if m:
			id = m.group(1)
			clues = [clue.split(": ") for clue in m.group(2, 3, 4)] # ([cats, 7], [dogs, 4], [cows, 5])
			for name, value in clues:
				# statement = 'INSERT INTO Sues ({}) VALUES (?)'.format(name)
				statement = 'UPDATE Sues SET {column} = ? WHERE rowid = {id}'.format(column=name, id=id)
				c.execute(statement, tuple([int(value)]))
			connection.commit()
	for row in c.execute('''SELECT rowid, children, cats, samoyeds,
		                           pomeranians, akitas, vizslas, goldfish,
		                           trees, cars, perfumes
		                    FROM Sues
		                    WHERE rowid < 20 '''):
		# print row
		pass

def test():
	print "all tests pass!"

def partone():
	pass

def parttwo():
	db = 'day16.db'
	conn = sqlite3.connect(db)
	cleanup(conn)
	create_db(db)

	parse_to_db('input.txt', conn)
	find_correct_sue(conn)


	cleanup(conn)
	conn.close()

if __name__ == '__main__':
    # test()
    # partone()
    parttwo()
