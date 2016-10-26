import sys


class Wire(object):
    def __init__(self, input=None):
        self.input = input
        self.value = None

    def evaluate(self):
        globals = {
            'WIRES': WIRES,
            '__builtins__': {},
        }

        # lol
        return eval(self.input, globals, {})

    def get_value(self):
        if self.value is not None:
            return self.value

        self.value = self.evaluate()
        return self.value

WIRES = {}

OPERATORS = {
    'AND': '&',
    'OR': '|',
    'LSHIFT': '<<',
    'RSHIFT': '>>',
    'NOT': '~',
}


def is_numeric(input):
    try:
        int(input)
    except ValueError:
        return False
    else:
        return True


def is_lowercase(input):
    return input.lower() == input


def transform_expression(expression):
    tokens = []

    for value in expression.split(' '):
        if is_numeric(value):
            tokens.append(value)
        elif is_lowercase(value):
            tokens.append("WIRES['{}'].get_value()".format(value))
        else:
            tokens.append(OPERATORS[value])

    return ' '.join(tokens)


def handle_command(input):
    expression, target = [s.strip() for s in input.split('->')]
    try:
        wire = WIRES[target]
    except KeyError:
        wire = Wire(input=transform_expression(expression))
        WIRES[target] = wire

if __name__ == '__main__':
    try:
        wire = sys.argv[1]
    except IndexError:
        wire = 'a'

    code = sys.stdin.read()
    for line in code.splitlines():
        handle_command(line)

    final_wire = WIRES[wire]
    value = final_wire.get_value()

    for wire in WIRES.values():
        # Reset all wires
        wire.value = None

    # Wire b takes the value previously assigned to a
    WIRES['b'].value = value

    print(final_wire.get_value())
