import re
from collections import defaultdict


class QuakeLogParser:
    def __init__(self, log_file):
        self.log_file = log_file
        self.matches = {}
        self.parse_log()

    def parse_log(self):
        with open(self.log_file, 'r') as file:
            lines = file.readlines()
            current_match = None
            for line in lines:
                if 'InitGame:' in line:
                    current_match = f'game_{len(self.matches) + 1}'
                    self.matches[current_match] = {
                        'total_kills': 0,
                        'players': set(),
                        'kills': defaultdict(int)
                    }
                elif 'Kill:' in line:
                    self.process_kill(current_match, line)

    def process_kill(self, current_match, line):
        pattern = r'(\d+):(\d+) Kill: (\d+) (\d+) (\d+): (.*) killed (.*) by (.*)'
        match = re.match(pattern, line)
        if match:
            killer = match.group(6)
            victim = match.group(7)
            mod = match.group(8)

            if killer != '<world>':
                self.matches[current_match]['kills'][killer] += 1
            else:
                self.matches[current_match]['kills'][victim] -= 1

            self.matches[current_match]['players'].add(killer)
            self.matches[current_match]['players'].add(victim)
            self.matches[current_match]['total_kills'] += 1

    def get_match_summary(self):
        for match, data in self.matches.items():
            print(f"{match}:")
            print(f"  Total Kills: {data['total_kills']}")
            print("  Players: [")
            for player in data['players']:
                print(f"    '{player}',")
            print("  ]")
            print("  Kills: {")
            for player, kills in data['kills'].items():
                print(f"    '{player}': {kills},")
            print("  }")
            print("}")
