# Day 3: Rucksack Reorganization - Part 2
## Problem
As you finish identifying the misplaced items, the Elves come to you with
another issue.

For safety, the Elves are divided into groups of three. Every Elf carries a
badge that identifies their group. For efficiency, within each group of three
Elves, the badge is the only item type carried by all three Elves. That is, if
a group's badge is item type B, then all three Elves will have item type B
somewhere in their rucksack, and at most two of the Elves will be carrying
any other item type.

The problem is that someone forgot to put this year's updated authenticity
sticker on the badges. All of the badges need to be pulled out of the rucksacks
so the new authenticity stickers can be attached.

Additionally, nobody wrote down which item type corresponds to each group's
badges. The only way to tell which item type is the right one is by finding the
one item type that is common between all three Elves in each group.

Every set of three lines in your list corresponds to a single group, but each
group can have a different badge item type. So, in the above example, the
first group's rucksacks are the first three lines:

```
  vJrwpWtwJgWrhcsFMMfFFhFp
  jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
  PmmdzqPrVvPwwTWBwg
```
And the second group's rucksacks are the next three lines:

```
  wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
  ttgJtRGJQctTZtZT
  CrZsJsPPZsGzwwsLwLmpwMDw
```
In the first group, the only item type that appears in all three rucksacks is
lowercase r; this must be their badges. In the second group, their badge item
type must be Z.

Priorities for these items must still be found to organize the sticker
attachment efforts: here, they are 18 (r) for the first group and 52 (Z) for
the second group. The sum of these is 70.

Find the item type that corresponds to the badges of each three-Elf group.
What is the sum of the priorities of those item types?

# Part two

Whereas in the first part of the day's puzzle we were looking at items present
in two parts of one bag, we're now looking for duplicates amongst three sets.
We could take two approaches:

1. Build a map like we did in the previous solution, update it for every bag to
   indicate whether the item index was present in all previous bags.
2. Build a map for each set.

The second solution is easier to implement, but would have serious performance
implications as we scaled to a larger number of bags (even though this is not
relevant to the puzzle). If we wanted to see whether an item is in all the other
bags, we'd have to do a retrieve from every other map to check whether the item
was in each other bag. As such, doing the final check to deduce which object
is in each bag is O(1) for solution 1, but O(n) for solution 2. We will
implement solution 2.

We can store an array of length n where n is the number of bags against each
index in the map. When an item is found, we'll mark the index in the array
corresponding to the bag true.

# Solution

- Set a total score to 0.
- Create a map[byte]*bool
- For each rucksack:
  - If the index is divisible by 3, (re)initialise our map.
  - Iterate over the items in the bag. For each item:
    - Try to retrieve it from the map.
 
      **If it exists:** 
      - **If this is the last bag:** look up the value and add it to the total
                                     score.
      - **otherwise:** mark the bag index (index % 3) in the array true.

      **If it doesn't exist:** create a new array, mark the bag index
      (index % 3) in the array true, and insert the pointer into the map.
- Return the score.

# Considerations

In the solution, we're copying every three lines into a slice, and passing that
into a function to find the duplicate between the three bags. We could have
implemented a more efficient solution by using indices as described above (%3)
to cleverly determine where we are in the group and use that to take decisions.
I started writing the solution like this, but it was neither easy to read nor
very testable. As correctness, maintainability and readability supercede
efficiency (in this case) I refactored the search across the group into a
separate function.