#!/bin/python

class Wizard:
    def __init__(self, name, damage, mana, hp):
        self.name = name
        self.damage = damage

        self.mana = mana
        self.total_mana_used = 0
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
            print "Poison deals 3 damage. Timer is now %d" % self.poison_effect
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
        self.total_mana_used = 0

    def fight(self, other_player):
        did_you_win = False
        while True:
            print "-- Player turn --"
            # status_check()
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
            # status_check()
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
        if spell == "Magic Missile":
            print "Magic Missle is casted for 4 damage"
            self.mana -= 53
            self.total_mana_used += 53
            damage = 4

        if spell == "Drain":
            print "Drain is casted for 2 damage"
            self.mana -= 73
            self.total_mana_used += 73
            damage = 2
            self.hp += 2

        if spell == "Shield":
            print "Shield is casted"
            self.mana -= 113
            self.total_mana_used += 113
            self.shield_effect = 6
            self.armor += 7

        if spell == "Poison":
            print "Poison is casted"
            self.mana -= 173
            self.total_mana_used += 173
            other_player.poison_effect = 6

        if spell == "Recharge":
            print "Recharge is casted"
            self.mana -= 229
            self.total_mana_used += 229
            self.recharge_effect = 5

        other_player.hp -= damage - other_player.armor
        return other_player.hp <= 0

    def atttack_sequence(self, moves, other_player):
        self.recover()
        other_player.recover()
        for move in moves:
            print "-- Player Turn --"
            fight_status(self, other_player)
            if self.take_effect():
                return -self.total_mana_used

            if other_player.take_effect():
                return self.total_mana_used

            print "Player casts %s" % move
            if self.attack(other_player, spell=move):
                return self.total_mana_used

            print "-- Boss Turn --"
            fight_status(self, other_player)
            if self.take_effect():
                return -self.total_mana_used

            if other_player.take_effect():
                return self.total_mana_used

            print "Boss attacks for %d damage" % other_player.damage
            if other_player.attack(self):
                return -self.total_mana_used

        # raise Exception('Both players are alive. Fight is not over')
        return -self.total_mana_used

def fight_status(player, boss):
    print "- %s has %d hp, %d armor, %d mana" % (player.name, player.hp, player.armor, player.mana)
    print "- Boss has %s hp" % boss.hp

def generate_attack_combos(max_num_moves):
    pass

def test():
    you = Wizard(name="Player", hp=10, damage=0, mana=250)
    boss = Wizard(name="Boss", hp=13, damage=8, mana=0)
    # assert(not you.attack(boss, spell="Poison"))
    # assert(not boss.attack(you))
    # assert(not you.attack(boss, spell="Magic Missile"))
    assert(you.atttack_sequence(["Poison", "Magic Missile"], boss) > 0)
    print "all tests pass!"

def partone():
    you = Wizard(name="Player", hp=10, damage=0, mana=250)
    boss = Wizard(name="Boss", hp=13, damage=8, mana=0)

    combos = generate_attack_combos(max_num_moves=50)
    winning_combos_scores = filter(lambda x: x > 0, map(lambda x: you.atttack_sequence(x, boss), combos))

    return min(winning_combos_scores)


def parttwo():
    pass

if __name__ == '__main__':
    test()
    # partone()
    # parttwo()
