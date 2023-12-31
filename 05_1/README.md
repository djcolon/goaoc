# Day 5: Supply Stacks - part 1
The expedition can depart as soon as the final supplies have been unloaded from
the ships. Supplies are stored in stacks of marked crates, but because the
needed supplies are buried under many other crates, the crates need to be
rearranged.

The ship has a giant cargo crane capable of moving crates between stacks. To
ensure none of the crates get crushed or fall over, the crane operator will
rearrange them in a series of carefully-planned steps. After the crates are
rearranged, the desired crates will be at the top of each stack.

The Elves don't want to interrupt the crane operator during this delicate
procedure, but they forgot to ask her which crate will end up where, and they
want to be ready to unload them as soon as possible so they can embark.

They do, however, have a drawing of the starting stacks of crates and the
rearrangement procedure (your puzzle input). For example:

```
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
```
In this example, there are three stacks of crates. Stack 1 contains two crates:
crate Z is on the bottom, and crate N is on top. Stack 2 contains three crates;
from bottom to top, they are crates M, C, and D. Finally, stack 3 contains a
single crate, P.

Then, the rearrangement procedure is given. In each step of the procedure, a
quantity of crates is moved from one stack to a different stack. In the first
step of the above rearrangement procedure, one crate is moved from stack 2 to
stack 1, resulting in this configuration:

```
[D]        
[N] [C]    
[Z] [M] [P]
 1   2   3 
```
In the second step, three crates are moved from stack 1 to stack 3. Crates are
moved one at a time, so the first crate to be moved (D) ends up below the second
and third crates:

```
        [Z]
        [N]
    [C] [D]
    [M] [P]
 1   2   3
```
Then, both crates are moved from stack 2 to stack 1. Again, because crates are
moved one at a time, crate C ends up below crate M:

```
        [Z]
        [N]
[M]     [D]
[C]     [P]
 1   2   3
```

Finally, one crate is moved from stack 1 to stack 2:

```
        [Z]
        [N]
        [D]
[C] [M] [P]
 1   2   3
```

The Elves just need to know which crate will end up on top of each stack; in
this example, the top crates are C in stack 1, M in stack 2, and Z in stack 3,
so you should combine these together and give the Elves the message CMZ.

After the rearrangement procedure completes, what crate ends up on top of each
stack?

# Solution

We'll need some functionality to parse the input file. We'll define a struct for
a move (number, source, dest) and build some slice to hold the stacks of crates.

If we were naive we'd mvoe the crates between the slices one-by-one, but if we
recognise that any multiple moves will reverse the order of the subset moved
we can do it in one go and just reverse the order and move the stack in one go.
That said, that will probably be slower as it will reverse the slice in place
and then append it to another stack slice. We may just move the crates
one-by-one and be more efficient.

```
for moves
    for number in move
        crate = source.pop
        dest.append(crate)
```

And then simply return the ends of the slices in order.

## Opportunity to practice

Whilst the outcome of the crate stacking is order dependent (except for blocks
of moves that are independent of each-other) we can split the parsing of the
input file into multiple parallel processes. We'll spawn three processes
to do our work. The first will be passed the input file up to the blank
line, and parse it into the original stack layout. A second process will start
processing all the moves, and pass it into a pipe with a bit of a buffer
(to allow it to run ahead of the other pipes). Finally we'll set up a third
process to simulate our crane movements. It'll wait for the initial state, and
then start processing moves as they are passed into the pipe.
