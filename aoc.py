#!/usr/bin/env python3
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


def get_default_outfile():
    p = (
        Path(os.getcwd())
        / str(datetime.datetime.now().year)
        / f"day{parser.parse_args().day}"
    )

    p.mkdir(parents=True, exist_ok=True)
    return p / "in.txt"


parser.add_argument(
    "-o",
    "--out",
    type=str,
    help="The output file to write the data to",
    default=get_default_outfile(),
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
