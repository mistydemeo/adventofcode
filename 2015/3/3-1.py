from collections import defaultdict

present_map = defaultdict(lambda: defaultdict(lambda: 0))
present_map[0][0] += 1
house_count = 1

X = 'x'
Y = 'y'

position = {X: 0, Y: 0}
MAP = {
    '^': (Y, 1),
    'v': (Y, -1),
    '>': (X, 1),
    '<': (X, -1),
}

with open('input.txt') as input:
    for c in input.read():
        axis, delta = MAP[c]
        position[axis] += delta

        present_map[position[X]][position[Y]] += 1
        if present_map[position[X]][position[Y]] == 1:
            house_count += 1

print(house_count)
