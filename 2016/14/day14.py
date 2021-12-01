#!/bin/python
import re

class Reindeer:
	def __init__(self, name, speed, duration, rest):
		self.name = name
		self.speed = speed
		self.duration = duration
		self.rest = rest

	def distance_at_time(self, time):
		max_fly_distance = self.speed * self.duration
		time_unit = self.duration + self.rest
		distance_after_repeated_full_fly_rests = max_fly_distance * (time // time_unit)
		num_seconds_left = time % time_unit
		if num_seconds_left < self.duration:
			return distance_after_repeated_full_fly_rests + self.speed * num_seconds_left
		else:
			return distance_after_repeated_full_fly_rests + max_fly_distance

def calculate_points_at_time(time, reindeers):
	scores = {}
	for reindeer in reindeers:
		scores[reindeer] = 0

	for t in xrange(1, time+1):
		for reindeer in winning_reindeers_at_time(t, reindeers):
			scores[reindeer] += 1

	return scores


def winning_reindeers_at_time(time, reindeers):
	winning_so_far = []
	max_distance_so_far = 0
	for reindeer in reindeers:
		d = reindeer.distance_at_time(time)
		if d > max_distance_so_far:
			max_distance_so_far = d
			winning_so_far = [reindeer]
		elif d == max_distance_so_far:
			winning_so_far.append(reindeer)

	return winning_so_far

def parse(text_file):
	with open(text_file, 'r') as f:
		data = f.readlines()

	reindeers = []

	r = re.compile(r'(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.')
	for reindeer_stat in data:
		match = re.match(r, reindeer_stat)
		if match:
			reindeer = Reindeer(match.group(1), int(match.group(2)), int(match.group(3)), int(match.group(4)))
			reindeers.append(reindeer)

	return reindeers

def test():
	comet = Reindeer("Comet", speed=14, duration=10, rest=127)
	dancer = Reindeer("Dancer", speed=16, duration=11, rest=162)
	assert(comet.distance_at_time(1000) == 1120)
	assert(dancer.distance_at_time(1000) == 1056)

	assert(len(winning_reindeers_at_time(2503, [comet, dancer])) == 1)
	assert(len(winning_reindeers_at_time(2503, [comet, comet])) == 2)

	scores = calculate_points_at_time(1000, [comet, dancer])
	assert(scores[comet] == 312)
	assert(scores[dancer] == 689)

	print "all tests pass!"

def partone():
	t = 2503
	reindeers = parse('input.txt')
	distances_at_time = [r.distance_at_time(t) for r in reindeers]
	print max(distances_at_time)

def parttwo():
	t = 2503
	reindeers = parse('input.txt')
	scores = calculate_points_at_time(t, reindeers)
	print max(scores.values())


if __name__ == '__main__':
    test()
    # partone()
    parttwo()
