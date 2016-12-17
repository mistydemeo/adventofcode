from collections import defaultdict

present_map = defaultdict(lambda: defaultdict(lambda: 0))
present_map[0][0] += 2
house_count = 1

X = 'x'
Y = 'y'

santa_position = {X: 0, Y: 0}
robo_position = {X: 0, Y: 0}
MAP = {
    '^': (Y, 1),
    'v': (Y, -1),
    '>': (X, 1),
    '<': (X, -1),
}

with open('input.txt') as input:
    for i, c in enumerate(input.read()):
        position = santa_position if i % 2 == 0 else robo_position
        axis, delta = MAP[c]
        position[axis] += delta

        present_map[position[X]][position[Y]] += 1
        if present_map[position[X]][position[Y]] == 1:
            house_count += 1

print(house_count)
