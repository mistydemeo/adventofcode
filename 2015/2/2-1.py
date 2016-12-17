class Present(object):
    def __init__(self, width, height, length):
        self.width = int(width)
        self.height = int(height)
        self.length = int(length)

    def surface_area(self):
        return (2 * self.length * self.width) + \
               (2 * self.width * self.height) + \
               (2 * self.height * self.length)

    def calculate_slack(self):
        side1 = self.length * self.width
        side2 = self.width * self.height
        side3 = self.height * self.length
        return min([side1, side2, side3])

with open("input.txt") as input:
    print(sum(map(lambda p: p.surface_area() + p.calculate_slack(), [Present(*present.split('x')) for present in input.read().splitlines()])))
