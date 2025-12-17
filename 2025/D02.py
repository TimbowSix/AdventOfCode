import utils

raw = utils.get_input(2)
#raw = utils.get_input_test(2)

invalid_ids = []

for id_range in raw.split(","):
    start, end = id_range.split("-")
    start = int(start)
    end = int(end)

    curr = start
    while curr <= end:
        curr_str = str(curr)
        if len(curr_str) % 2 == 0:
            if curr_str[:len(curr_str)//2] == curr_str[len(curr_str)//2:]:
                invalid_ids.append(curr)

        curr += 1

print(sum(invalid_ids))

invalid_ids_p2 = []
for id_range in raw.split(","):
    start, end = id_range.split("-")
    start = int(start)
    end = int(end)

    curr = start
    while curr <= end:
        curr_str = str(curr)
        for i in range(1, len(curr_str)//2+1):
            if len(curr_str) % i != 0:
                continue
            part = curr_str[:i]
            repeated = part * (len(curr_str)//i)
            if repeated == curr_str:
                invalid_ids_p2.append(curr)
                break

        curr += 1
print(sum(invalid_ids_p2))
