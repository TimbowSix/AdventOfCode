from __future__ import annotations
import utils
from tqdm import tqdm

DAY = 10

raw = utils.get_input_lines(DAY)
raw = utils.get_input_lines_test(DAY)

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
# for machine in tqdm(machines):
#     print(machine.target_lamps)
#     known_states: dict[str, Node] = {}

#     # build tree
#     #machine = machines[0]
#     initial_node = Node("".join(machine.lamps), 0)
#     known_states[initial_node.lamps] = initial_node
#     queue: list[Node] = [initial_node]

#     while queue:
#         current_node = queue.pop(0)
#         #print(" ".join(current_node.lamps), current_node.previous_count)
#         #print(len(known_states))
#         for button in machine.buttons:
#             next_node = current_node.get_next_node(button, known_states)
#             if next_node.previous_count == current_node.previous_count + 1:
#                 queue.append(next_node)
#         #print(current_node.lamps, current_node.lamps == "".join(machine.target_lamps))
#         if current_node.lamps == "".join(machine.target_lamps):
#             #print(f"done: {current_node.previous_count}")
#             result += current_node.previous_count
#             break

# print(result)

import numpy as np
from scipy.optimize import milp
from scipy.optimize import LinearConstraint
from scipy.optimize import Bounds

result = 0
for machine in machines:
    c = np.array([1 for _ in machine.buttons])
    A = np.zeros((len(machine.joltages), len(machine.buttons)))
    for j, joltage in enumerate(machine.joltages):
        for b, button in enumerate(machine.buttons):
            if j in button.toggles:
                A[j,b] = 1
    b_u = np.array(machine.joltages)
    b_l = np.array(machine.joltages)
    constraints = LinearConstraint(A, b_l, b_u)
    upper_bounds = []
    for button in machine.buttons:
        bound = min([machine.joltages[toggle] for toggle in button.toggles])
        upper_bounds.append(bound)
    #print(upper_bounds)
    bounds = Bounds([0 for _ in machine.buttons], upper_bounds)
    integrality = np.ones_like(c)
    res = milp(c=c, constraints=constraints, bounds=bounds, integrality=integrality)
    # print(res)
    #print(res.x)
    # print(sum(res.x))
    result += sum(res.x)

print(int(result))
