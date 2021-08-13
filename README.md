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
3. an amount of numbers in all iterations
4. the largest number that was reached during the iteration
5. ~~ten most used while iterating numbers~~ (too much time to wait this data)

## How to run

To run iterations to the number 10 with defined power, run:

    make run_iterator MAXIPOWER=6

or with `go run`:

    go run main.go iterator --max-power 6

To make run with profiling, use make command `run_iterator_profiled`.

To investigate profiling data run:

    make profile_cpu

or

    make profile_mem

for CPU or Memory profiling.

## Current results

As for now, algorithm counts numbers on 6 cores Intel Core i7 macbook pro:

* from 1 to 10e6 in about 0.9-2 seconds
* from 1 to 10e7 in about 10-13 seconds
* from 1 to 10e8 in about 1m50s-2m30s
* from 1 to 10e9 in about an hour o so
* from 1 to 10e10 already more than I can wait :)...

it optimized to run from the beginning (1) to the end. So, iterating from 10e7 to 10e8 takes more than hour instead of couple minutes.

Trying to optimize it more...