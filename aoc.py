import datetime
import os
import subprocess
import argparse
from pathlib import Path

TOKEN = os.getenv(
    "TOKEN",
    open(Path(os.getcwd()) / ".env", "r").read().strip().split("=")[-1],
)
# Parse the command-line arguments
parser = argparse.ArgumentParser(description="Download Advent of Code data")
parser.add_argument(
    "-y",
    "--year",
    type=int,
    help="The year for the Advent of Code event",
    default=datetime.datetime.now().year,
)
parser.add_argument(
    "-d",
    "--day",
    type=int,
    help="The day for which to download the data",
    default=datetime.datetime.now().day,
)
parser.add_argument(
    "-o",
    "--out",
    type=str,
    help="The output file to write the data to",
    default="in.txt",
)
args = parser.parse_args()

if __name__ == "__main__":
    res = subprocess.run(
        [
            "curl",
            f"https://adventofcode.com/{args.year}/day/{args.day}/input",
            "-s",
            "--cookie",
            f"session={TOKEN}",
            "-o",
            args.out,
        ]
    )
