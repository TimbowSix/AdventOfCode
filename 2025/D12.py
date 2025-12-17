from __future__ import annotations
import utils

DAY = 12

raw = utils.get_input_lines(DAY)
#raw = utils.get_input_lines_test(DAY)

class Shape:
    shape_id: int
    grid: list[list[str]]
    def __init__(self, shape_id: int) -> None:
        self.shape_id = shape_id
        self.grid = []

    def __repr__(self) -> str:
        return f"{self.shape_id}"

class Region:
    width: int
    length: int
    shapes: dict[Shape, int] # shape: amount

    def __init__(self, width: int, length: int) -> None:
        self.width = width
        self.length = length
        self.shapes = {}

    def __repr__(self) -> str:
        return f"{self.width}x{self.length} {self.shapes}"

shapes: list[Shape] = []
curr_shape = None
regions: list[Region] = []
for line in raw:
    if not line:
        continue
    if line.endswith(":") and line[:-1].isdigit():
        curr_shape = Shape(int(line[:-1]))
        shapes.append(curr_shape)
    elif "x" in line:
        end = line.index(":")
        w, l = line[:end].split("x")
        width = int(w)
        length = int(l)
        region = Region(width, length)
        for idx, num in enumerate(line[end+1:].split()):
            n = int(num)
            if n == 0: continue
            shape = shapes[idx]
            region.shapes[shape] = n
        regions.append(region)
    else:
        curr_shape.grid.append(list(line))

# for shape in shapes:
#     print(shape.shape_id)
#     for line in shape.grid:
#         print(line)
# print(regions)

fit_count = 0
for region in regions:
    area = region.width * region.length
    needed_area = 0
    for shape, count in region.shapes.items():
        needed_area += count * 9
    if area >= needed_area:
        fit_count += 1

print(fit_count)
