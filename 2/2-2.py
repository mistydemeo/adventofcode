class Present(object):
    def __init__(self, width, height, length):
        self.width = int(width)
        self.height = int(height)
        self.length = int(length)

    def smallest_face(self):
        side1 = (self.length, self.width)
        side2 = (self.width, self.height)
        side3 = (self.height, self.length)
        return sorted([side1, side2, side3], key=lambda s: sum(s))[0]

    def smallest_perimeter(self):
        face = self.smallest_face()
        return face[0] * 2 + face[1] * 2

    def calculate_volume(self):
        return self.width * self.height * self.length

with open("input.txt") as input:
    print(sum(map(lambda p: p.smallest_perimeter() + p.calculate_volume(), [Present(*present.split('x')) for present in input.read().splitlines()])))
