from __future__ import annotations
import utils
from tqdm import tqdm

DAY = 9

raw = utils.get_input_lines(DAY)
raw = utils.get_input_lines_test(DAY)

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

ranges_x = {} # y : (min_x, max_x)
ranges_y = {} # x : (min_y, max_y)
for tile1 in red_tiles:
    for tile2 in red_tiles:
        if tile1[1] == tile2[1]:
            y = tile1[1]
            min_x = min(tile1[0], tile2[0])
            max_x = max(tile1[0], tile2[0])
            if y not in ranges_x:
                ranges_x[y] = (min_x, max_x)
            else:
                cur_min, cur_max = ranges_x[y]
                ranges_x[y] = (min(min_x, cur_min), max(max_x, cur_max))
        if tile1[0] == tile2[0]:
            x = tile1[0]
            min_y = min(tile1[1], tile2[1])
            max_y = max(tile1[1], tile2[1])
            if x not in ranges_y:
                ranges_y[x] = (min_y, max_y)
            else:
                cur_min, cur_max = ranges_y[x]
                ranges_y[x] = (min(min_y, cur_min), max(max_y, cur_max))

# print(ranges_x)
# print(ranges_y)
cache = {}
def pos_in_zone(x,y) -> bool:
    if (x,y) in cache:
        return cache[(x,y)]
    if y in ranges_x:
        min_x, max_x = ranges_x[y]
        if min_x <= x <= max_x:
            cache[(x,y)] = True
            return True
    if x in ranges_y:
        min_y, max_y = ranges_y[x]
        if min_y <= y <= max_y:
            cache[(x,y)] = True
            return True

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
        corners = [
            (tile1[0], tile2[1]),
            (tile2[0], tile1[1]),
        ]

        #print(tile1, tile2, corners)
        valid = all([pos_in_zone(x,y) for x,y in corners])
        if not valid:
            continue
        biggest = size
        c = corners + [tile1, tile2]


print(biggest)
print(c)
#4602673662

max_x = max([t[0] for t in red_tiles])+2
max_y = max([t[1] for t in red_tiles])+2

grid = [["." for _ in range(max_x)] for _ in range(max_y)]
for x,y in red_tiles:
    grid[y][x] = "#"
for y in range(len(grid)):
    for x in range(len(grid[y])):
        symbol = grid[y][x]
        if symbol == ".":
            if pos_in_zone(x,y):
                grid[y][x] = "X"
for x,y in c:
    grid[y][x] = "O"
for i in grid:
    print(" ".join(i))
