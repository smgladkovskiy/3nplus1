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

## Current results

As for now, algorithm counts numbers on 6 cores Intel Core i7

* from 1 to 10e6 in about 40-50 seconds
* from 1 to 10e7 in about 8 minutes
* from 10e7 to 10e8 already more than hour...

Trying to optimize...