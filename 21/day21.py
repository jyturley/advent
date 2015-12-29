#!/bin/python
from possible_equipment import equipment

class Player:
    def __init__(self, name, damage, armor, hp):
        self.name = name
        self.damage = damage
        self.armor = armor
        self.total_hp = hp
        self.hp = hp

    def equip(self, item):
        self.damage += item[0]
        self.armor += item[1]
        if item[2]:
            if item[2] > 0:
                self.damage += item[2]
            else:
                self.armor += -item[2]

        if item[3]:
            if item[3] > 0:
                self.damage += item[3]
            else:
                self.armor += -item[3]

    def unequip(self, item):
        self.damage -= item[0]
        self.armor -= item[1]
        if item[2]:
            if item[2] > 0:
                self.damage -= item[2]
            else:
                self.armor -= -item[2]

        if item[3]:
            if item[3] > 0:
                self.damage -= item[3]
            else:
                self.armor -= -item[3]

    def recover(self):
        self.hp = self.total_hp

    def fight(self, other_player):
        did_you_win = False
        still_fighting = True
        while still_fighting:
            if self.attack(other_player):
                # print "%s is the winner!" % (self.name)
                still_fighting = False
                did_you_win = True
                break

            if other_player.attack(self):
                # print "%s is the winner!" % (other_player.name)
                still_fighting = False
                did_you_win = False
                break

        self.recover()
        other_player.recover()
        return did_you_win

    def attack(self, other_player):
        net_damage = self.damage - other_player.armor
        other_player.hp -= net_damage
        
        # print "%s attacks %s and deals %d-%d=%d damage. %s goes down to %d hit points." \
        #     % (self.name, other_player.name, self.damage, other_player.armor, net_damage, other_player.name, other_player.hp)
        return other_player.hp <= 0

def calculate_cost(item):
    weapons = {
        4:8,
        5:10,
        6:25,
        7:40,
        8:74
    }

    armor = {
        0:0,
        1:13,
        2:31,
        3:53,
        4:75,
        5:102
    }

    ring = {
        0:0,
        1:25,
        2:50,
        3:100,
        -1:20,
        -2:40,
        -3:80
    }

    return weapons[item[0]] + armor[item[1]] + ring[item[2]] + ring[item[3]]

def find_all_possible_item_combos():
    equipments_with_cost = []
    for items in equipment:
        cost = calculate_cost(items)
        equipments_with_cost.append((cost, items))

    return equipments_with_cost


def test():
    you  = Player(name="you", hp=8, damage=5, armor=5)
    boss = Player(name="boss", hp=12, damage=7, armor=2)
    assert(you.fight(boss))

    print "all tests pass!"

def partone():
    you  = Player(name="you", hp=100, damage=0, armor=0)
    boss = Player(name="boss", hp=104, damage=8, armor=1)

    for cost, item in sorted(find_all_possible_item_combos(), key=lambda item:item[0]):
        you.equip(item)
        if you.fight(boss):
            print cost, item
            return cost
        you.unequip(item)

def parttwo():
    you  = Player(name="you", hp=100, damage=0, armor=0)
    boss = Player(name="boss", hp=104, damage=8, armor=1)

    for cost, item in sorted(find_all_possible_item_combos(), reverse=True, key=lambda item:item[0]):
        you.equip(item)
        if not you.fight(boss):
            print cost, item
            return cost
        you.unequip(item)

if __name__ == '__main__':
    # test()
    # partone()
    parttwo()
