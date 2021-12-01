#!/bin/python
import re
from itertools import permutations

def longest_path(valid_routes):
	distances = [couple[0] for couple in valid_routes]
	return max(distances)
	
def shortest_path(valid_routes):
	distances = [couple[0] for couple in valid_routes]
	return min(distances)

def path_exists(from_city, to_city, graph):
	return from_city in graph and to_city in graph[from_city]

def traveling_santa(graph, cities, routes):
	valid_routes = []
	for route in routes:
		route_cities = route.split(" ")
		distance_so_far = 0
		valid_route_flag = True
		for i in xrange(len(route_cities) - 1):
			if not path_exists(route_cities[i], route_cities[i + 1], graph):
				valid_route_flag = False
				break
			distance_so_far += graph[route_cities[i]][route_cities[i + 1]]
		if valid_route_flag:
			valid_routes.append(tuple([distance_so_far, route]))

	return valid_routes

def create_routes(cities):
	output = []
	routes = [" ".join(s) for s in permutations(cities)]
	return routes

def parse(text_file):
	graph = {}
	cities = set()
	r = re.compile(r'(\w+) to (\w+) = (\d+)')
	with open(text_file, 'r') as f:
		for line in f.readlines():
			m = re.match(r, line)
			city_one, city_two, distance = m.groups()
			cities.add(city_one)
			cities.add(city_two)
			if not city_one in graph:
				graph[city_one] = {}
			graph[city_one][city_two] = int(distance)

			# because this is not a directed graph
			if not city_two in graph:
				graph[city_two] = {}
			graph[city_two][city_one] = int(distance)
	return graph, cities

def test():
	test_graph, test_cities = parse("test_input.txt")

	print test_graph
	print create_routes(test_cities)

	assert(len(test_cities) == 3)
	assert(path_exists("London", "Dublin", test_graph))
	assert(path_exists("London", "Belfast", test_graph))

	test_routes = create_routes(test_cities)
	test_valid_routes = traveling_santa(test_graph, test_cities, test_routes)
	assert(len(test_valid_routes) == 6)
	assert(shortest_path(test_valid_routes) == 605)

	print "all tests pass!"

def partone():
	graph, cities = parse("input.txt")
	valid_routes = traveling_santa(graph, cities, create_routes(cities))
	print shortest_path(valid_routes)

def parttwo():
	graph, cities = parse("input.txt")
	valid_routes = traveling_santa(graph, cities, create_routes(cities))
	print longest_path(valid_routes)

if __name__ == '__main__':
    test()
    # partone()
    parttwo()
