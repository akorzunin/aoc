import os
import subprocess
import argparse
from pathlib import Path

TOKEN = os.getenv(
    "TOKEN",
    open(Path(os.getcwd()).parents[0] / ".env", "r").read().strip(),
)
# Parse the command-line arguments
parser = argparse.ArgumentParser(description="Download Advent of Code data")
parser.add_argument(
    "-y",
    "--year",
    type=int,
    help="The year for the Advent of Code event",
)
parser.add_argument(
    "-d",
    "--day",
    type=int,
    help="The day for which to download the data",
)
args = parser.parse_args()

# # Set the year and day
year = args.year or 2023
day = args.day or int(Path(os.getcwd()).stem[-1])

if __name__ == "__main__":
    res = subprocess.run(
        [
            "curl",
            f"https://adventofcode.com/2023/day/{day}/input",
            "--cookie",
            f"session={TOKEN}",
            "-o",
            "in.txt",
        ]
    )
    print(TOKEN)
