# Day 3 - 3

There is not third puzzle, but as we're learning, this exercise seemed a good
opportunity to have a go at concurrency. We can group each set of three bags
and have a separate routine calculate the score.

## Goroutines

A goroutine is a function that can run concurrently with other functions.
Channels allow different goroutines to communicate and synchronise their
execution.

> Do not communicate by sharing memory; instead, share memory by communicating.

What we'll do is create a function that reads the input file. When it has a
slice of three bags it'll dump it into a channel and read the next set.
We'll also have a number of worker functions that process these groups and
determine the score for each group and send it into a channel.
We'll take all those scores and add them up in a final routine.
When all routines are done we can terun the value and exit.

# Result

As was expected, there wasn't a perfomance boost from using multiple elves
to sort through the bags (I'm really embracing this metaphor). With 300 bags
there just wasn't enough data to offset the overhead of running things in
parallel. The applicaiton funcitoned correctly on the first try though, so not bad!

```
  2023/07/26 19:12:14 Advent of Code 2022 - 3.3
  2023/07/26 19:12:14 Reading input from 'input.txt'.
  2023/07/26 19:12:14 badgeFinder elf #4, reporting for duty!
  2023/07/26 19:12:14 badgeFinder elf #3, reporting for duty!
  2023/07/26 19:12:14 badgeFinder elf #2, reporting for duty!
  2023/07/26 19:12:14 badgeFinder elf #0, reporting for duty!
  2023/07/26 19:12:14 badgeFinder elf #6, reporting for duty!
  2023/07/26 19:12:14 badgeFinder elf #1, reporting for duty!
  2023/07/26 19:12:14 badgeFinder elf #5, reporting for duty!
  2023/07/26 19:12:14 badgeFinder elf #7, reporting for duty!
  2023/07/26 19:12:14 badgeFinder elf #7, signing off after processing 10 group's bags!
  2023/07/26 19:12:14 badgeFinder elf #3, signing off after processing 16 group's bags!
  2023/07/26 19:12:14 badgeFinder elf #1, signing off after processing 12 group's bags!
  2023/07/26 19:12:14 badgeFinder elf #6, signing off after processing 12 group's bags!
  2023/07/26 19:12:14 badgeFinder elf #2, signing off after processing 13 group's bags!
  2023/07/26 19:12:14 badgeFinder elf #0, signing off after processing 11 group's bags!
  2023/07/26 19:12:14 badgeFinder elf #4, signing off after processing 16 group's bags!
  2023/07/26 19:12:14 badgeFinder elf #5, signing off after processing 10 group's bags!
  2023/07/26 19:12:14 =========
  2023/07/26 19:12:14 Score for input file 'input.txt' is: 2805.

  real    0m0.135s
  user    0m0.146s
  sys     0m0.115s
```