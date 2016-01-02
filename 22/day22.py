#!/bin/python

class Wizard:
    def __init__(self, name, damage, mana, hp):
        self.name = name
        self.damage = damage

        self.mana = mana
        self.total_mana = mana
        self.total_hp = hp
        self.hp = hp

        self.armor = 0

        self.shield_effect = 0
        self.poison_effect = 0
        self.recharge_effect = 0

    def take_effect(self):
        if self.shield_effect:
            self.shield_effect -= 1
            print "Sheild's timer is now %d" % self.shield_effect
            if self.shield_effect <= 0:
                self.armor = 0
                print "Shield wears off"

        if self.poison_effect:
            self.poison_effect -= 1
            self.hp -= 3
            print "Poison timer is now %d" % self.poison_effect
            if self.poison_effect <= 0:
                print "Posion wears off"

        if self.recharge_effect:
            self.recharge_effect -= 1
            print "Recharge timer is now %d" % self.recharge_effect
            self.mana += 101
            if self.recharge_effect <= 0:
                print "Recharge wears off"

        return self.hp <= 0

    def recover(self):
        self.hp = self.total_hp
        self.mana = self.total_mana

    def fight(self, other_player):
        did_you_win = False
        while True:
            print "-- Player turn --"
            status_check()
            if self.take_effect():
                did_you_win = False
                break

            if other_player.take_effect():
                did_you_win = True
                break

            if self.attack(other_player):
                did_you_win = True
                break

            print "-- Boss turn --"
            status_check()
            if self.take_effect():
                did_you_win = False
                break

            if other_player.take_effect():
                did_you_win = True
                break

            if other_player.attack(self):
                did_you_win = False
                break

        self.recover()
        other_player.recover()
        return did_you_win

    def attack(self, other_player, spell=""):
        damage = self.damage
        if spell is "Magic Missile":
            self.mana -= 53
            damage = 4

        if spell is "Drain":
            self.mana -= 73
            damage = 2
            self.hp += 2

        if spell is "Shield":
            self.mana -= 113
            self.shield_effect = 6
            self.armor += 7

        if spell is "Poison":
            self.mana -= 173
            other_player.poison_effect = 6

        if spell is "Recharge":
            self.mana -= 229
            self.recharge_effect = 5

        other_player.hp -= damage - other_player.armor
        return other_player.hp <= 0

def test():
    you = Wizard(name="player", hp=10, damage=0, mana=250)
    boss = Wizard(name="boss", hp=13, damage=8, mana=0)
    assert(not you.attack(boss, spell="Poison"))
    assert(not boss.attack(you))
    assert(not you.attack(boss, spell="Magic Missile"))
    you.recover()
    boss.recover()
    assert(you.fight(boss))
    print "all tests pass!"

def partone():
    pass

def parttwo():
    pass

if __name__ == '__main__':
    test()
    # partone()
    # parttwo()
