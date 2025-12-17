import os
import requests
from dotenv import load_dotenv
load_dotenv()

def download_input(day: int):
    path = f"D{str(day).zfill(2)}Input.txt"
    if os.path.isfile(path):
        return # file already exists
    print(f"Downloading Day {day} input")
    seassion_cookie = os.getenv("SESSION_COOKIE")
    cookie = {"session": seassion_cookie}
    url = f"https://adventofcode.com/2025/day/{day}/input"
    response = requests.get(url, cookies=cookie)
    if response.status_code != 200:
        print(f"Error getting input, status: {response.status_code}")
    with open(path, "w") as f:
        f.write(response.text)
    print(f"Input saved to {path}")

def get_input(day: int) -> str:
    path = f"D{str(day).zfill(2)}Input.txt"
    if not os.path.isfile(path):
        download_input(day)
    with open(path, "r") as f:
        return f.read()

def get_input_lines(day: int) -> list[str]:
    input_str = get_input(day)
    return input_str.splitlines()

def get_input_test(day: int) -> str:
    path = f"D{str(day).zfill(2)}InputTest.txt"
    if not os.path.isfile(path):
        raise FileNotFoundError(f"Test Values file for Day {day} does not exist")
    with open(path, "r") as f:
        return f.read()

def get_input_lines_test(day: int) -> list[str]:
    input_str = get_input_test(day)
    return input_str.splitlines()
