from __future__ import annotations
import utils
from tqdm import tqdm
from PIL import Image, ImageDraw
import av
import multiprocessing
import os

DAY = 9

class Square:
    max_x: int
    min_x: int
    max_y: int
    min_y: int
    color: str
    width: int

    def __init__(self, corners: list[tuple[int,int]], color: str = 'lightblue', width: int = 1):
        self.max_x = max([t[0] for t in corners])
        self.min_x = min([t[0] for t in corners])
        self.max_y = max([t[1] for t in corners])
        self.min_y = min([t[1] for t in corners])
        self.color = color
        self.width = width

class FrameData:
    max_x: int
    max_y: int
    frame_id: str
    squares: list[Square]
    edges: list[tuple[tuple[int, int], tuple[int, int]]]
    def __init__(self, squares: list[Square], edges: list[tuple[tuple[int, int], tuple[int, int]]], frame_id: str, max_x: int, max_y: int):
        self.squares = squares
        self.edges = edges
        self.frame_id = frame_id
        self.max_x = max_x
        self.max_y = max_y

def make_frame(frame_data: FrameData):
    SCALE = 10

    img = Image.new('RGB', (frame_data.max_x, frame_data.max_y), color = 'black')
    draw = ImageDraw.Draw(img)

    for a,b in frame_data.edges:
        x1 = a[0] // SCALE
        y1 = a[1] // SCALE
        draw.circle((x1, y1),5, fill='red')
        x2 = b[0] // SCALE
        y2 = b[1] // SCALE
        draw.circle((x2, y2),5, fill='red')

        draw.line((x1, y1, x2, y2), fill='yellow')

    for square in frame_data.squares:
        draw.rectangle((square.min_x // SCALE, square.min_y // SCALE, square.max_x // SCALE, square.max_y // SCALE), outline=square.color, width=square.width)

    img = img.resize((4096, 4096), Image.LANCZOS)
    img.save(os.path.join("__test", f"D09_{frame_data.frame_id}.png"))


if __name__ == "__main__":

    raw = utils.get_input_lines(DAY)
    #raw = utils.get_input_lines_test(DAY)

    red_tiles: list[tuple[int,int]] = []
    for line in raw:
        t = line.split(",")
        a = int(t[0])
        b = int(t[1])
        red_tiles.append((a,b))

    cache = {}
    def pos_in_zone(x,y) -> bool:
        if (x,y) in cache:
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


    edges: list[tuple[tuple[int, int], tuple[int, int]]] = []
    for x,y in red_tiles:
        for x_, y_ in red_tiles:
            if x == x_ and y == y_:
                continue
            if x == x_ and y_ > y:
                edges.append( ((x, y), (x_, y_)) )
            if y == y_ and x_ > x:
                edges.append( ((x, y), (x_, y_)) )


    SCALE = 10

    max_x = (max([t[0] for t in red_tiles])+2) // SCALE
    max_y = (max([t[1] for t in red_tiles])+2) // SCALE

    biggest = 0
    c = []

    biggest_square = None

    frame_data = []
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

            square_edges = [
                ( (tile1[0], tile1[1]), (tile1[0], tile2[1]) ),
                ( (tile1[0], tile2[1]), (tile2[0], tile2[1]) ),
                ( (tile2[0], tile2[1]), (tile2[0], tile1[1]) ),
                ( (tile2[0], tile1[1]), (tile1[0], tile1[1]) ),
            ]
            # check if any edge crosses existing edges
            valid = all([pos_in_zone(x,y) for x,y in corners])
            if not valid:
                continue

            valid = True
            for edge1 in square_edges:
                for edge2 in edges:
                    # check if edge1 crosses edge2
                    (x1, y1), (x2, y2) = edge1
                    (x3, y3), (x4, y4) = edge2

                    denom = (y4 - y3) * (x2 - x1) - (x4 - x3) * (y2 - y1)
                    if denom == 0:
                        continue  # parallel lines

                    ua = ((x4 - x3) * (y1 - y3) - (y4 - y3) * (x1 - x3)) / denom
                    ub = ((x2 - x1) * (y1 - y3) - (y2 - y1) * (x1 - x3)) / denom

                    if 0 < ua < 1 and 0 < ub < 1:
                        valid = False
                        break
                if not valid:
                    break

            square = Square(corners + [tile1, tile2], color='lightblue', width=2)
            fd = FrameData([square], edges, str(len(frame_data)).zfill(4), max_x, max_y)
            frame_data.append(fd)
            if not valid:
                if biggest_square:
                    fd.squares.append(biggest_square)
                continue
            square.color = 'blue'
            square.width = 3
            biggest_square = square
            biggest = size
            c = corners + [tile1, tile2]

    print(biggest)
    print(c)

    with multiprocessing.Pool(processes=os.cpu_count()) as pool:
        for _ in tqdm(pool.imap_unordered(make_frame, frame_data), total=len(frame_data)):
            pass


    file_name = f"D{str(DAY).zfill(2)}.mp4"
    output = av.open(file_name, 'w')
    stream = output.add_stream('h264', 60)
    stream.width = 4096
    stream.height = 4096
    for img_file in tqdm(os.listdir("__test")):
        if not img_file.startswith(f"D09_") or not img_file.endswith(".png"):
            continue
        img = Image.open(os.path.join("__test", img_file))
        img = img.resize((4096, 4096), Image.NEAREST)
        frame = av.VideoFrame.from_image(img)
        packet = stream.encode(frame)
        output.mux(packet)

    img = Image.new('RGB', (max_x, max_y), color = 'black')
    draw = ImageDraw.Draw(img)
    for x,y in red_tiles:
        x1 = x // SCALE
        y1 = y // SCALE
        draw.circle((x1, y1),5, fill='red')

    for a,b in edges:
        x1 = a[0] // SCALE
        y1 = a[1] // SCALE
        x2 = b[0] // SCALE
        y2 = b[1] // SCALE
        draw.line((x1, y1, x2, y2), fill='yellow')

    square_max_x = max([t[0] for t in c])
    square_min_x = min([t[0] for t in c])
    square_max_y = max([t[1] for t in c])
    square_min_y = min([t[1] for t in c])

    draw = ImageDraw.Draw(img)
    draw.rectangle((square_min_x // SCALE, square_min_y // SCALE, square_max_x // SCALE, square_max_y // SCALE), outline='green', width=5)
    img = img.resize((4096, 4096), Image.LANCZOS)
    for _ in tqdm(range(120)):
        frame = av.VideoFrame.from_image(img)
        packet = stream.encode(frame)
        output.mux(packet)

    packet = stream.encode(None)
    output.mux(packet)
    output.close()
    print(f"Video saved to {file_name}")
