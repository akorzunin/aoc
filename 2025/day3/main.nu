#!/usr/bin/env nu
def main [file: path = "in.txt"] {
    let input = open $file
        | lines
        | par-each {|e| $e
            mut el = $e
            mut res = null
            while ($res | is-empty) {
                $res = ($el | parse_row)
                $el = $el | split chars | slice 0..-2 | str join
            }
            return {
                value: $res
                row: ($e | find ...($e |split chars | sort | last 2))
            }
        }
    print $input
    print -n "sum: " ($input | get value | into int | math sum) "\n"
}

def parse_row [] {
    let e = $in
    let ec = $in | split chars
    let max_el = $ec | math max
    let max_el_idx = $ec | enumerate | find -c [item] $max_el | first | get index

    let after_max_el = $ec | slice ($max_el_idx + 1)..-1
    if ($after_max_el | length) >= 1 {
        # get max el from after max el
        let after_max_el_max = $after_max_el | math max
        return ($max_el + $after_max_el_max)
    }
    if ($max_el == ($ec | last)) {
        return (($ec | slice 0..-2 | math max) + $max_el)
    }
    return null
}
