# Daily Coding Problem: Problem #534 [Easy] 

Given a string, determine whether any permutation of it is a palindrome.

For example,
"carrace" should return true, since it can be rearranged to form "racecar",
which is a palindrome.
"daily" should return false,
since there's no rearrangement that can form a palindrome.

## Insight

A palindrome has to have even numbers of each of the letters in it,
except for possibly the middle letter:

* "aabaa" contains 4 'a', 1 'b'
* "cacac" contains 3 'c', 2 'a'
* "caac" contains 2 'a', 2 'c'

## Algorithm

1. Create counts of each character appearing in the input string.
2. Examine character counts,
counting the number of even and odd character counts.
3. If only 1 count is odd, or all counts are even, return true,
otherwise return false.

There's an optimization/simplification:
just track odd character counts.
If you get past 1, you can't possibly form a palindrome
from the input string.

## Analysis

This is surprisingly easy, *if* you have the insight about character counts.
The programming is easy,
but the interviewer has to decide if the candidate does the right thing
by counting both even and odd character counts,
which probably requires an if-then-else,
or if the candidate does the optimization/simplification.

If you're looking for a Enterprise app Java programmer,
you probably want to avoid the candidate that notices the simplification.

