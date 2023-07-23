# Day 3: Rucksack Reorganization
## Problem
One Elf has the important job of loading all of the rucksacks with supplies for
the jungle journey. Unfortunately, that Elf didn't quite follow the packing
instructions, and so a few items now need to be rearranged.

Each rucksack has two large compartments. All items of a given type are meant to
go into exactly one of the two compartments. The Elf that did the packing failed
to follow this rule for exactly one item type per rucksack.

The Elves have made a list of all of the items currently in each rucksack (your
puzzle input), but they need your help finding the errors. Every item type is
identified by a single lowercase or uppercase letter (that is, a and A refer to
different types of items).

The list of items for each rucksack is given as characters all on a single line.
A given rucksack always has the same number of items in each of its two
compartments, so the first half of the characters represent items in the first
compartment, while the second half of the characters represent items in the
second compartment.

For example, suppose you have the following list of contents from six rucksacks:

```
    vJrwpWtwJgWrhcsFMMfFFhFp
    jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
    PmmdzqPrVvPwwTWBwg
    wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
    ttgJtRGJQctTZtZT
    CrZsJsPPZsGzwwsLwLmpwMDw
```

The first rucksack contains the items `vJrwpWtwJgWrhcsFMMfFFhFp`, which means
its first compartment contains the items `vJrwpWtwJgWr`, while the second
compartment contains the items `hcsFMMfFFhFp`. The only item type that appears
in both compartments is lowercase `p`.
The second rucksack's compartments contain `jqHRNqRjqzjGDLGL` and
`rsFMfFZSrLrFZsSL`. The only item type that appears in both compartments is
uppercase `L`.
The third rucksack's compartments contain `PmmdzqPrV` and `vPwwTWBwg`; the only
common item type is uppercase `P`.
The fourth rucksack's compartments only share item type `v`.
The fifth rucksack's compartments only share item type `t`.
The sixth rucksack's compartments only share item type `s`.
To help prioritize item rearrangement, every item type can be converted to a
priority:

Lowercase item types `a` through `z` have priorities 1 through 26.
Uppercase item types `A` through `Z` have priorities 27 through 52.
In the above example, the priority of the item type that appears in both
compartments of each rucksack is 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19
(s); the sum of these is 157.

Find the item type that appears in both compartments of each rucksack. What is
the sum of the priorities of those item types?

# Solution

- Set a total score to 0.
- For each rucksack:
  - Get the number of items.
  - Calculate the midpoint.
  - Create a map[char]bool (actual contained type doesn't matter)
  - Iterate over the items in the bag. In the first half of the rucksack:
    - Insert each item into the map
  - In the second half of the rucksack, check if each item is in the map. If it
    is not, continue, if it is we have found our duplicate.
  - If a duplicate is found, look up the value and return it.
  - If no duplictae is found return 0
  - Add the returned value to the score.
- Return the score.

## Considerations

As the characters in the backpacks can be mapped to an integer we could use
an array of bools to track items we have already encountered. A map however is
easier to read, and flexible should the rules change later. If we wanted max
performance we might still use an array, as inserts into the map may cause some
performance overhead.