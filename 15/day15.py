#!/bin/python
import re

class Ingredient:
	def __init__(self, name, capacity, durability, flavor, texture, calories):
		self.name = name
		self.capacity = capacity
		self.durability = durability
		self.flavor = flavor
		self.texture = texture
		self.calories = calories

	def multiply(self, scalar):
		attributes = [self.capacity, self.durability, self.flavor, self.texture]
		product = map(lambda x: scalar * x, attributes)
		return product

	def multiply_calories(self, scalar):
		return self.calories * scalar

def is_recipe_target_calorie(recipe, ingredients, target_calorie):
	calorie_count = 0
	for x in xrange(len(ingredients)):
		calorie_count += ingredients[x].multiply_calories(recipe[x])

	return calorie_count == target_calorie

def find_best_recipe_with_calories(proportions, ingredients, target_calorie):
	recipes_with_target_calories = filter(lambda x: is_recipe_target_calorie(x, ingredients, target_calorie), proportions)

	return find_best_recipe(recipes_with_target_calories, ingredients)

def find_best_recipe(proportions, ingredients):
	best_recipe_so_far = ""
	best_score_so_far = 0
	for recipe in proportions:
		score = calculate_score(recipe, ingredients)
		# print recipe, score
		if score > best_score_so_far:
			best_score_so_far = score
			best_recipe_so_far = recipe

	return best_recipe_so_far, best_score_so_far

def generate_test_proportions(total):
	proportions = []
	for i in xrange(total+1):
		j = total - i
		# print i, j, i+j
		proportions.append([i, j])

	return proportions

def generate_proportions(total):
	proportions = []
	for i in xrange(total+1):
		for j in xrange(total-i+1):
			for k in xrange(total-j-i+1):
				m = total - k - j - i
				# print i, j, k, m, i+j+k+m
				proportions.append([i, j, k, m])

	return proportions

def calculate_score(proportions, ingredients):
	assert(len(proportions) == len(ingredients))
	scores = {"capacity" : 0, "durability" : 0, "flavor" : 0, "texture" : 0}
	for x in xrange(len(ingredients)):
		product = ingredients[x].multiply(proportions[x])
		scores['capacity'] += product[0]
		scores['durability'] += product[1]
		scores['flavor'] += product[2]
		scores['texture'] += product[3]

	for attr in scores:
		if scores[attr] < 0:
			scores[attr] = 0

	total = scores['capacity'] * scores['durability'] * scores['flavor'] * scores['texture']
	return total if total > 0 else 0

def parse(text_file):
	with open(text_file, 'r') as f:
		data = f.readlines()

	ingredients = []
	r = re.compile(r'(\w+): capacity ([-+]?\d+), durability ([-+]?\d+), flavor ([-+]?\d+), texture ([-+]?\d+), calories (\d+)')
	for ingredient in data:
		m = re.match(r, ingredient)
		if m:
			i = Ingredient(m.group(1), int(m.group(2)), int(m.group(3)), int(m.group(4)), int(m.group(5)), int(m.group(6)))
			ingredients.append(i)

	return ingredients

def test():
	ingredients = parse('test_input.txt')
	assert(len(ingredients) == 2)
	assert(calculate_score([44, 56], ingredients) == 62842880)

	recipes = generate_test_proportions(100)
	# proportions = generate_proportions(100)

	print find_best_recipe(recipes, ingredients)


	print "all tests pass!"

def partone():
	ingredients = parse('input.txt')
	recipes = generate_proportions(100)
	print find_best_recipe(recipes, ingredients)


def parttwo():
	ingredients = parse('input.txt')
	recipes = generate_proportions(100)
	print find_best_recipe_with_calories(recipes, ingredients, target_calorie=500)

if __name__ == '__main__':
    # test()
    # partone()
    parttwo()
