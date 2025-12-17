import utils

DAY = 5

raw = utils.get_input_lines(DAY)
#raw = utils.get_input_lines_test(DAY)

ranges = []
fresh_count = 0
for line in raw:
    if "-" in line:
        a,b = line.split("-")
        ranges.append((int(a), int(b)))
    elif line == "": continue
    else:
        i_id = int(line)
        for start, end in ranges:
            if start <= i_id <= end:
                fresh_count += 1
                break

print(fresh_count)

new_ranges = []
for a,b in sorted(ranges):
    # find biggest possible b
    if not new_ranges:
        new_ranges.append((a,b))
        continue
    last_a, last_b = new_ranges[-1]
    if a <= last_b + 1:
        # overlap
        new_ranges[-1] = (last_a, max(last_b, b))
    else:
        new_ranges.append((a,b))

result = 0
for a,b in new_ranges:
    result += b - a + 1

print(result)
