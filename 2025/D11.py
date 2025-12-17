from __future__ import annotations
import utils

DAY = 11

raw = utils.get_input_lines(DAY)
#raw = utils.get_input_lines_test(DAY)

class Device:
    name: str
    connections: list[Device]

    def __init__(self, name: str) -> None:
        self.name = name
        self.connections = []

    def __repr__(self) -> str:
        return self.name

devices: dict[str, Device] = {}
for line in raw:
    t = line.split(": ")
    device_name = t[0]
    connection_names = t[1].split(" ")
    if device_name not in devices:
        device = Device(device_name)
        devices[device_name] = device
    else:
        device = devices[device_name]

    for cn in connection_names:
        if cn not in devices:
            connection = Device(cn)
            devices[cn] = connection
        else:
            connection = devices[cn]
        device.connections.append(connection)

# for device in devices.values():
#     print(device.name, device.connections)

def traverse_devices(current_device: Device, destination: Device) -> int:
    cache = {}
    def traverse(current_device: Device, destination: Device, ) -> int:
        if current_device in cache:
            return cache[current_device]

        if current_device == destination:
            return 1

        count = 0
        for device in current_device.connections:
            count += traverse(device, destination)
        cache[current_device] = count
        return count
    return traverse(current_device, destination)

print(traverse_devices(devices["you"], devices["out"]))

dac_fft = traverse_devices(devices["dac"], devices["fft"])
fft_dac = traverse_devices(devices["fft"], devices["dac"])

steps = ["svr"]
if dac_fft == 0:
    steps.append("fft")
    steps.append("dac")
else:
    steps.append("dac")
    steps.append("fft")
steps.append("out")

result = 1
for i in range(len(steps)-1):
    a = steps[i]
    b = steps[i+1]
    count = traverse_devices(devices[a], devices[b])
    result *= count

print(result)

