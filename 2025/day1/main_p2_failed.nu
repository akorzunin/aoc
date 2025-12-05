#!/usr/bin/env nu
def main [file: path = "in.txt"] {
    let input = open $file
    | lines
    | split chars
    | each {|e| {
        dir: ($e | get 0)
        steps: ($e | skip 1 | str join | into int)
    }}

    mut position = 50
    mut count = 0
    for i in $input {
        let curr = $position
        match $i.dir {
            "R" => {
                let add = ($position + $i.steps) // 100
                $count += $add
                $position = ($position + $i.steps) mod 100
                if $position == 0 and $add == 0 {
                    $count += 1
                }
                {pos: $curr, steps: $i.steps, dir: $i.dir, count: $count, end: $position, add: $add} | to tsv | print
                print ""
            },
            "L" => {
                let add = (($position - $i.steps) | math abs) // 100
                $count += $add
                if ($position - $i.steps) < 0 and $position != 0 {
                    print "a"
                    $count += 1
                }
                $position = ($position - $i.steps) mod 100
                if $position < 0 {
                    $position = $position + 100
                }
                if ($position == 0 or $position == 99) and $add == 0 and $i.steps > 1 {
                    print "b"
                    $count += 1
                }
                {pos: $curr, steps: $i.steps, dir: $i.dir, count: $count, end: $position, add: $add} | to tsv | print
                print ""
            }
        }
    }
    print $count
}
