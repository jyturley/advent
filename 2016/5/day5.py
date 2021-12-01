#!/bin/python

def pair_of_letters_constraint(string):
    for first_letter_index in xrange(len(string) - 1):
        second_letter_index = first_letter_index + 1
        pair = string[first_letter_index] + string[second_letter_index]
        if pair in string[second_letter_index+1:]:
            return True

    return False


def sandwich_letter_constraint(string):
    for i in xrange(len(string) - 2):
        if string[i] == string[i + 2]:
            return True

    return False

def is_nice_2(string):
    return pair_of_letters_constraint(string) and sandwich_letter_constraint(string)

def is_nice_1(string):
    vowels = set(["a", "e", "i", "o", "u"])
    vowel_count = 0
    double_letter_pass = False
    previous_char = ""

    if "ab" in string: return False
    if "cd" in string: return False
    if "pq" in string: return False
    if "xy" in string: return False

    for char in string:
        if vowel_count < 3 and char in vowels:
            vowel_count += 1
        if double_letter_pass == False and char == previous_char:
            double_letter_pass = True

        previous_char = char

    return double_letter_pass and vowel_count >= 3

def test():
    assert(is_nice_1("ugknbfddgicrmopn") == True)
    assert(is_nice_1("aaa") == True)
    assert(is_nice_1("jchzalrnumimnmhp") == False)
    assert(is_nice_1("haegwjzuvuyypxyu") == False)
    assert(is_nice_1("dvszwmarrgswjxmb") == False)
    print "part one tests pass!"

    assert(pair_of_letters_constraint("aaa") == False)
    assert(sandwich_letter_constraint("xxyxx") == True)
    assert(is_nice_2("qjhvhtzxzqqjkmpb") == True)
    assert(is_nice_2("xxyxx") == True)
    assert(is_nice_2("uurcxstgmygtbstg") == False)
    assert(is_nice_2("ieodomkazucvgmuy") == False)
    print "part two tests pass!"

def partone():
    nice_count = 0
    with open('input.txt', 'r') as f:
        data = f.readlines()
        for string in data:
            nice_count += is_nice_1(string)

    print nice_count


def parttwo():
    nice_count = 0
    with open('input.txt', 'r') as f:
        data = f.readlines()
        for string in data:
            nice_count += is_nice_2(string)

    print nice_count

if __name__ == '__main__':
    test()
    # partone()
    parttwo()
