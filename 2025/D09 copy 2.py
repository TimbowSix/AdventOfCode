from __future__ import annotations
import utils
from tqdm import tqdm

DAY = 9

raw = utils.get_input_lines(DAY)
#raw = utils.get_input_lines_test(DAY)

red_tiles: list[tuple[int,int]] = []
for line in raw:
    t = line.split(",")
    a = int(t[0])
    b = int(t[1])
    red_tiles.append((a,b))

#print(red_tiles)

biggest = 0
for tile1 in red_tiles:
    for tile2 in red_tiles:
        size = (abs(tile1[0]-tile2[0])+1)*(abs(tile1[1]-tile2[1])+1)
        if size > biggest:
            biggest = size

print(biggest)

cache = {}
def pos_in_zone(x,y) -> bool:
    if (x,y) in cache:
        #print((x,y), cache[(x,y)])
        return cache[(x,y)]
    n = len(red_tiles)
    inside = False
    px, py = x, y
    for i in range(n):
        x1, y1 = red_tiles[i][0], red_tiles[i][1]
        x2, y2 = red_tiles[(i + 1) % n][0], red_tiles[(i + 1) % n][1]

        # check if cross
        if ((y1 > py) != (y2 > py)) and (px < (x2 - x1) * (py - y1) / (y2 - y1) + x1):
            inside = not inside
    cache[(x,y)] = inside
    return inside

biggest = 0
c = []
for tile1 in tqdm(red_tiles):
    for tile2 in red_tiles:
        if tile1 == tile2:
            continue
        size = (abs(tile1[0]-tile2[0])+1)*(abs(tile1[1]-tile2[1])+1)
        if size <= biggest:
            continue
        max_x = max(tile1[0], tile2[0])
        min_x = min(tile1[0], tile2[0])
        max_y = max(tile1[1], tile2[1])
        min_y = min(tile1[1], tile2[1])

        valid = True
        #print(tile1, tile2, size)
        for x in range(min_x, max_x+1):
            if not pos_in_zone(x, min_y):
                valid = False
                break
        if not valid:
            continue
        for y in range(min_y, max_y+1):
            if not pos_in_zone(min_x, y):
                valid = False
                break
        if not valid:
            continue

        biggest = size
        c = [tile1, tile2]


print(biggest)
# print(c)
# #4602673662

# max_x = max([t[0] for t in red_tiles])+2
# max_y = max([t[1] for t in red_tiles])+2

# grid = [["." for _ in range(max_x)] for _ in range(max_y)]
# for x,y in red_tiles:
#     grid[y][x] = "#"
# for y in range(len(grid)):
#     for x in range(len(grid[y])):
#         symbol = grid[y][x]
#         if symbol == ".":
#             if pos_in_zone(x,y):
#                 grid[y][x] = "X"
# for x,y in c:
#     grid[y][x] = "O"
# for i in grid:
#     print(" ".join(i))
