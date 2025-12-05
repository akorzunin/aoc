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
        match $i.dir {
            "R" => {
                $position = ($position + $i.steps) mod 100
            },
            "L" => {
                $position = ($position - $i.steps) mod 100
                if $position < 0 {
                    $position = $position + 100
                }
            }
        }
        if $position == 0 {
            $count += 1
        }
    }
    print $count
}
