# Day 6: Tuning Trouble

The preparations are finally complete; you and the Elves leave camp on foot and
begin to make your way toward the star fruit grove.

As you move through the dense undergrowth, one of the Elves gives you a handheld
device. He says that it has many fancy features, but the most important one to
set up right now is the communication system.

However, because he's heard you have significant experience dealing with
signal-based systems, he convinced the other Elves that it would be okay to
give you their one malfunctioning device - surely you'll have no problem fixing
it.

As if inspired by comedic timing, the device emits a few colorful sparks.

To be able to communicate with the Elves, the device needs to lock on to their
signal. The signal is a series of seemingly-random characters that the device
receives one at a time.

To fix the communication system, you need to add a subroutine to the device
that detects a start-of-packet marker in the datastream. In the protocol being
used by the Elves, the start of a packet is indicated by a sequence of four
characters that are all different.

The device will send your subroutine a datastream buffer (your puzzle input);
your subroutine needs to identify the first position where the four most
recently received characters were all different. Specifically, it needs to
report the number of characters from the beginning of the buffer to the end of
the first such four-character marker.

For example, suppose you receive the following datastream buffer:

`mjqjpqmgbljsphdztnvjfqwrcgsmlb``
After the first three characters (mjq) have been received, there haven't been
enough characters received yet to find the marker. The first time a marker
could occur is after the fourth character is received, making the most recent
four characters mjqj. Because j is repeated, this isn't a marker.

The first time a marker appears is after the seventh character arrives. Once it
does, the last four characters received are jpqm, which are all different. In
this case, your subroutine should report the value 7, because the first
start-of-packet marker is complete after 7 characters have been processed.

Here are a few more examples:

`bvwbjplbgvbhsrlpgdmjqwftvncz`: first marker after character 5
`nppdvjthqldpwncqszvftbrmjlhg`: first marker after character 6
`nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`: first marker after character 10
`zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`: first marker after character 11
How many characters need to be processed before the first start-of-packet
marker is detected?

# Solution

For this puzzle, we'll move a window of n characters (in this case, 4) over the
input string. Whenever all 4 characters are different, we have found out marker.

We could compare all 4 characters for every step, but that's hardly efficient,
nor does it allow us to easily change the number of characters required for a
marker.

What we can do is maintain a directory of characters that are currently in the
window. As we shift over the string, we remove the oldest character, and then
add the new character. When we go and add the new character, and it is already
in the directory, we'll know we have hit our index. This means that, for any
size of input, and any size of key, we will only use a fixed amount of memory
(determined by the number of possible characters in our string) and compute.

The data-structure we use to track this is dependent on a trade-of. The fastest
solution is to recognise that we're just reading bytes limited to a speciofic
subset of ascii. In this case the lower-case letters spanning a at 97 to z
at 122. As such, we can have a fixed-length array of bools that we
address by `read byte - 97` and set to `true` and `false`. This'll only use
26 bytes of memory and be very quick to set for each character read.

The downside to this approach is that it is less easy to change should our
search-space change. If we expected a possible change we could either use a map
instead, or define an interface that'd allow us to easily switch out methods.
If the compiler does a good job, the latter should provide us with flexibility
without causing performance overhead.

Alongside this structure we'll keep an in-order buffer (a list would do well)
of the characters in our moving window to allow us to update our character map,
and an index to keep track of where we are in the file.

Finally, we'll read the file byte-by-byte so that we can process a large file.
We'll only be limited in file size we can process by the maximum size of the
index. If we us a uint64 that is massive. The odds of this running on a 32-bit
system are not wiorth consideraiton given the application.

## Assumptions

We'll assume we can only encounter lower-case characters.

