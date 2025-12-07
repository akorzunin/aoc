#!/usr/bin/env nu
def main [file: path = "in.txt"] {
    let input = open $file
        | lines
        | str join
        | split row ","
        | each {|e| $e
            | split row "-"
            | into int
            | zip [start, end]
            | each { reverse }
            | into record
        }
    let invalid_ids = $input | par-each {|i|
        for id in ($i.start..$i.end) {
            if not (is_id_valid_p2 ($id | into string)) {
                return $id
            }
        }
    }
    print -n "sum: " ($invalid_ids | math sum) "\n"
}

def is_id_valid [id: string] {
    let half: int = ($id | str length) / 2 | into int
    if $half < 1 {
        return true
    }
    let first_part = $id | str substring 0..($half - 1)
    let last_part = $id | str substring ($half)..(($in | str length) - 1)
    if $first_part == $last_part {
        return false
    }
    return true
}

def is_id_valid_p2 [id: string] {
    let half: int = ($id | str length) / 2 | into int
    if $half < 1 {
        return true
    }
    for i in (0..($half - 1)) {
        let is_invalid = $id
            | split row ($in | str substring 0..$i)
            | all {|e| $e | is-empty }
        if $is_invalid {
            return false
        }
    }
    return true
}
