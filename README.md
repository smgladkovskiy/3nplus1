3n+1 problem
------------

Realisation of Collatz conjecture (or 3n+1 problem) written on go.

Program tries to iterate through numbers from 1 to X and run recursive cycle with conditions:

- if number is odd it multiplied by 3 and add 1
- if number is even it divided by 2
- if resulted number is 1 cycle ends
- else resulted number passed on the entry of the same cycle

As a result, the program displays:

1. number it reached in the iteration
2. time it took to iterate
3. the maximum number of steps for iteration
4. the largest number that was reached during the iteration