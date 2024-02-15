# Day 8: Treetop Tree House

The expedition comes across a peculiar patch of tall trees all planted carefully
in a grid. The Elves explain that a previous expedition planted these trees as a
reforestation effort. Now, they're curious if this would be a good location for
a tree house.

First, determine whether there is enough tree cover here to keep a tree house
hidden. To do this, you need to count the number of trees that are visible from
outside the grid when looking directly along a row or column.

The Elves have already launched a quadcopter to generate a map with the height
of each tree (your puzzle input). For example:

```
30373
25512
65332
33549
35390
```
Each tree is represented as a single digit whose value is its height, where 0
is the shortest and 9 is the tallest.

A tree is visible if all of the other trees between it and an edge of the grid
are shorter than it. Only consider trees in the same row or column; that is,
only look up, down, left, or right from any given tree.

All of the trees around the edge of the grid are visible - since they are
already on the edge, there are no trees to block the view. In this example,
that only leaves the interior nine trees to consider:

The top-left 5 is visible from the left and top. (It isn't visible from the
right or bottom since other trees of height 5 are in the way.)
The top-middle 5 is visible from the top and right.
The top-right 1 is not visible from any direction; for it to be visible, there
would need to only be trees of height 0 between it and an edge.
The left-middle 5 is visible, but only from the right.
The center 3 is not visible from any direction; for it to be visible, there
would need to be only trees of at most height 2 between it and an edge.
The right-middle 3 is visible from the right.
In the bottom row, the middle 5 is visible, but the 3 and 4 are not.
With 16 trees visible on the edge and another 5 visible in the interior, a
total of 21 trees are visible in this arrangement.

Consider your map; how many trees are visible from outside the grid?

# Solution

To find the number of trees that are visible, we will at the minimum have to
iterate over every single tree once, as we cannot assume there are no taller
or equally tall trees until we have tested them all. This means that at best,
our solution is O(n). The problem is that for each tree, whether it is visible
or not depends on every tree between the tree itself and each of the 4 edges.
This means that in an n*n grid, the visibility of any tree from any direction
is dependent on all trees to the left, right, top and bottom of it. I.e. 
(n-1)+(n-1) trees. So if we, for every tree in the n*n grid, we tested all 4
directions, our solution would be O(n^3). Not great. We can opportunistically
break when we find a tree is not visible, but this won't do much in the grand
scheme of things.

## Reverse thinking.

Rather than thinking what tree is visible from the edge, lets think

## TreeMap

To efficiently store and access the highest tree we've seen so far in each
cardinal direction, we'll pack the information into a 16 bit integer:
[
    highest tree west
    highest tree north
    highest tree east
    tree height
]
Where each is 4 bits. Note we don't hold South. We will iterate over south last,
so we don't need to store it as we'll compute the answer in the final iteration.