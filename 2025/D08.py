from __future__ import annotations
import utils
from tqdm import tqdm

DAY = 8

raw = utils.get_input_lines(DAY)
#raw = utils.get_input_lines_test(DAY)

MAX_CONNECTIONS = 1000

class Network:
    boxes: set[JunctionBox]
    def __init__(self):
        self.boxes = set()

    def has_box(self, box: JunctionBox) -> bool:
        return box in self.boxes

    def add_box(self, box: JunctionBox):
        self.boxes.add(box)
        box.network = self

    def __add__(self, other: Network) -> Network:
        new_network = Network()
        for box in self.boxes:
            new_network.add_box(box)
        for box in other.boxes:
            new_network.add_box(box)
        return new_network

class JunctionBox:
    x: int
    y: int
    z: int
    network: Network|None

    def __init__(self, x, y, z):
        self.x = x
        self.y = y
        self.z = z
        self.network = None

    def distance(self, other: JunctionBox) -> int:
        dist = ((other.x - self.x) ** 2 + (other.y - self.y) ** 2 + (other.z - self.z) ** 2)**0.5
        #dist = abs(other.x - self.x) + abs(other.y - self.y) + abs(other.z - self.z)
        return dist

    def __repr__(self) -> str:
        return f"x={self.x} y={self.y} z={self.z}"

    def are_connected(self, other: JunctionBox) -> bool:
        return self.network == other.network and self.network is not None and other.network is not None

    def connect(self, other: JunctionBox) -> Network:
        if self.are_connected(other): return self.network
        if self.network is None and other.network is None:
            network = Network()
            network.add_box(self)
            network.add_box(other)
        elif self.network is None: #other.network is not None
            other.network.add_box(self)
        elif other.network is None: # self.network is not None
            self.network.add_box(other)
        else: # self and other network is not None
            self.network + other.network

boxes: list[JunctionBox] = []
print("Parse boxes")
for line in raw:
    x, y, z = [int(n) for n in line.split(",")]
    boxes.append(JunctionBox(x, y, z))
#print(boxes)

distances_dict: dict[str, tuple[float, tuple[JunctionBox, JunctionBox]]] = {}
print("calc dists")
for box1 in tqdm(boxes):
    for box2 in boxes:
        if box1 == box2: continue
        connection_id = f"{box1}-{box2}"
        reverse_connection_id = f"{box2}-{box1}"
        if connection_id in distances_dict or reverse_connection_id in distances_dict: continue
        dist = box1.distance(box2)
        distances_dict[connection_id] = ((dist, (box1, box2)))
distances = list(distances_dict.values())
distances.sort(key=lambda x: x[0])

print("make connections")
connections = 0
for connection in distances:
    dist = connection[0]
    box1, box2 = connection[1]
    #print(f"connect {box1} - {box2} - {dist}")
    box1.connect(box2)
    connections += 1
    networks: set[Network] = set()
    networks_with_nones = set()
    for box in boxes:
        networks_with_nones.add(box.network)
        if box.network is None: continue
        networks.add(box.network)
    #print(f"Networks: {len(networks)}, Connections: {connections}")
    if connections == MAX_CONNECTIONS:
        print("get networks")
        sizes = []
        for network in networks:
            sizes.append(len(network.boxes))

        sizes.sort(reverse=True)
        #print(sizes)
        print("Part1:", sizes[0] * sizes[1] * sizes[2])
    if len(networks_with_nones) == 1:
        print(box1, box2)
        print("Part2:", box1.x * box2.x)
        break
