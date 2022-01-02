# The Collatz Conjecture (in every language)
For those who don't know what this is, [here's a video that explains it perfectly](https://www.youtube.com/watch?v=094y1Z2wpJg). It explains how, through this algorithm, every number will boil down to `1`. And if you continue the process, you'll get stuck in a loop of `4, 2, 1, 4, 2, 1, 4, 2, 1`.

Here's how it works:
1. Pick a random number (for this example, we'll use `10`)
2. If it's even, divide it by `2`. If it's odd, multiply it by `3` and add `1`.
*Our number is now `5`*
3. Repeat step 2

Eventually, it'll lead you down the same path. Every time.

10 → 5 → 16 → 8 → **4 → 2 → 1**

13 → 40 → 20 → 10 → 5 →16 → 8 → **4 → 2 → 1**

45 → 136 → 68 → 43 → 17 → 52 → 26 → 13 → 40 → 20 → 10 → 5 → 16 → 8 → **4 → 2 → 1**

I made this for fun when I was bored. It didn't alleviate my boredom, but it added a new thing to my profile so here it is!
