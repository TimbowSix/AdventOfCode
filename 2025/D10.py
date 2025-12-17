from __future__ import annotations
import utils
from tqdm import tqdm

DAY = 10

raw = utils.get_input_lines(DAY)
#raw = utils.get_input_lines_test(DAY)

class Machine:
    target_lamps: str
    buttons: list[Button]
    joltages: list[int]
    def __init__(self, lamps, buttons, joltages) -> None:
        self.target_lamps = lamps
        self.lamps = ["." for _ in lamps]
        self.buttons = buttons
        self.joltages = joltages

class Button:
    toggles: list[int]
    def __init__(self, toggles: list[int]) -> None:
        self.toggles = toggles

    def __repr__(self) -> str:
        return f"{tuple(self.toggles)}"

machines: list[Machine] = []

for line in raw:
    s = line.split()
    lamps = s[0][1:-1]
    joltages = s[-1][1:-1].split(",")
    joltages = [int(j) for j in joltages]
    buttons = []
    for elem in s[1:-1]:
        elem = elem[1:-1].split(",")
        c = [int(i) for i in elem]
        buttons.append(Button(c))

    machine = Machine(
        lamps,
        buttons,
        joltages
    )
    machines.append(machine)
    #print(lamps, buttons, joltages)

class Node:
    lamps: str
    previous_count: int
    next_nodes: dict[Button, Node]

    def __init__(self, lamps: str, previous_count: int) -> None:
        self.lamps = lamps
        self.next_nodes = {}
        self.previous_count = previous_count

    def get_next_node(self, button: Button, known_states: dict[str, Node]) -> Node:
        if button in self.next_nodes:
            return self.next_nodes[button]

        new_lamps = list(self.lamps)
        for toggle in button.toggles:
            if new_lamps[toggle] == ".":
                new_lamps[toggle] = "#"
            else:
                new_lamps[toggle] = "."

        lamps_str = "".join(new_lamps)
        if lamps_str in known_states:
            next_node = known_states[lamps_str]
            if next_node.previous_count > self.previous_count + 1:
                next_node.previous_count = self.previous_count + 1
        else:
            next_node = Node(lamps_str, self.previous_count + 1)
            known_states[lamps_str] = next_node
        self.next_nodes[button] = next_node
        return next_node

result = 0
for machine in tqdm(machines):
    break
    print(machine.target_lamps)
    known_states: dict[str, Node] = {}

    # build tree
    #machine = machines[0]
    initial_node = Node("".join(machine.lamps), 0)
    known_states[initial_node.lamps] = initial_node
    queue: list[Node] = [initial_node]

    while queue:
        current_node = queue.pop(0)
        #print(" ".join(current_node.lamps), current_node.previous_count)
        #print(len(known_states))
        for button in machine.buttons:
            next_node = current_node.get_next_node(button, known_states)
            if next_node.previous_count == current_node.previous_count + 1:
                queue.append(next_node)
        #print(current_node.lamps, current_node.lamps == "".join(machine.target_lamps))
        if current_node.lamps == "".join(machine.target_lamps):
            #print(f"done: {current_node.previous_count}")
            result += current_node.previous_count
            break

print(result)

# a*(3)
# b*(1,3)
# c*(2)
# d*(2,3)
# e*(0,2)
# f*(0,1)
# ={3,5,4,7}

# e+f = 3
# b+f = 5
# c+d+e = 4
# a+b+d = 7
# a+b+c+d+e+f < 19

# a = 7 - b - d
# b = 5 - f
# c = 4 - d
# d = 4 - c
# e = 3 - f
# f = 3 - e


# Configuring the first machine's counters requires a minimum of 10 button presses.
# One way to do this is by pressing (3) once, (1,3) three times, (2,3) three times, (0,2) once, and (0,1) twice.
# 1 3 0 3 1 2
# a = 1
# b = 3
# c = 0
# d = 3
# e = 1
# f = 2

# 1+2 = 3 | a+f = 3
# 3+2 = 5 | b+f = 5
# 0+3 = 4 | c+d = 4
# 1+3+3 = 7 | a+b+d = 7

# e+f = 3
# b+f = 5
# c+d = 4
# a+b+d = 7

max_a = 7
max_b = 5
max_c = 4
max_d = 4
max_e = 3
max_f = 3

# for a in range(0, max_a + 1):
#     for b in range(0, max_b + 1):
#         for c in range(0, max_c + 1):
#             for d in range(0, max_d + 1):
#                 for e in range(0, max_e + 1):
#                     for f in range(0, max_f + 1):
#                         if a == 1 and b == 3 and c == 0 and d == 3 and e == 1 and f == 2:
#                             print("Example")
#                         if e+f == 3 and b+f == 5 and c+d+e == 4 and a+b+d == 7:
#                             print(a,b,c,d,e,f, a+b+c+d+e+f)

# result = 0
# for machine in tqdm(machines):
#     print(machine.joltages)
#     max_presses_per_button: dict[Button, int] = {}
#     # find max presses per button
#     for button in machine.buttons:
#         max_presses = float('inf')
#         for idx in button.toggles:
#             target_joltage = machine.joltages[idx]
#             if target_joltage < max_presses:
#                 max_presses = target_joltage
#         max_presses_per_button[button] = max_presses
#     #print(max_presses_per_button)
#     # all combinations of button presses from 0 to max presses
#     from itertools import product
#     button_press_ranges = [range(0, max_presses_per_button[button] + 1) for button in machine.buttons]
#     min_presses = float("inf")
#     for button_presses in tqdm(product(*button_press_ranges)):
#         #print(button_presses)
#         valid = True
#         for idx, target_joltage in enumerate(machine.joltages):
#             joltage = 0
#             for button, presses in zip(machine.buttons, button_presses):
#                 if idx in button.toggles:
#                     joltage += presses
#             if joltage != target_joltage:
#                 valid = False
#                 break
#         if valid and sum(button_presses) < min_presses:
#             min_presses = sum(button_presses)
#     result += min_presses
#     #print(min_presses)

# print(result)

# a*(0,5)
# b*(1,2,3,4,5)
# c*(1,3,4,5)
# d*(3,4)
# e*(2,3,5)
# f*(0,1,2,5)
# ={29,40,23,42,39,52}

# a = 29 - f
# b = 40 - c - d - e - f
# c = 23 - b - d - e - f
# d = 42 - b - c - e - f
# e = 39 - b - c - d - f
# f = 52 - a - b - c - d - e

