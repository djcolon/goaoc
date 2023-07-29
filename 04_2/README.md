# Advent of Code 2022 - Day 4 - 2

--- Part Two ---
It seems like there is still quite a bit of duplicate work planned. Instead, the Elves would like to know the number of pairs that overlap at all.

In the above example, the first two pairs (2-4,6-8 and 2-3,4-5) don't overlap, while the remaining four pairs (5-7,7-9, 2-8,3-7, 6-6,4-6, and 2-6,4-8) do overlap:

5-7,7-9 overlaps in a single section, 7.
2-8,3-7 overlaps all of the sections 3 through 7.
6-6,4-6 overlaps in a single section, 6.
2-6,4-8 overlaps in sections 4, 5, and 6.
So, in this example, the number of overlapping assignment pairs is 4.

In how many assignment pairs do the ranges overlap?

# Solution

This problem is still, simple, all we have to modify from the previous question
is the function determining whether there is overlap. This is the case if the
lower bound of one is larger than or equal to the lower bound of the other,
and lower than or equal to the upper bound, OR the same is through for the upper
bound.

Assuming the ranges are continuous, there will always be overlap if the end
of the range of the one are within the ends of the other, i.e:

lowerbound within range of other:
`one.Lower >= other.Lower && one.Lower <= other.Upper`
**OR**
upperbound within range of other:
`one.Upper >= other.Lower && one.Upper <= other.Upper`

This is not reversible so we'll need to test the reverse as well.

5 permutations:
```
# Neither in range.
one:   -----xxxx-----
other: --xx----------

# other.Upper in range
one:   -----xxxx-----
other: ----xx--------

# other.Upper and other.Lower in range
one:   ----xxxx------
other: -----xx-------

# other.Lower in range
one:   -----xxxx-----
other: --------xx----

# Neither in range.
one:   -----xxxx-----
other: ----------xx--
```