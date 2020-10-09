The 8-queens problem
==============

# What is this repo?
This is a solver for the 8-queens problem that I wrote in go.

# What is the 8-queens problem?
It is a puzzle where you are asked to place 8 queens on a chess board so that they do not attack eachother.

# Why write a solver if all solutions are already known?
Because I thought it would be a nice exercise which I could solve in the go programming language as practice.

# Implementation
My approach was to do a depth-first search in the game tree where all the children of a node are seen as the same board as their parent except with a queen placed in the leftmost column where no queen is already present.

A further algorithmic optimisation could be to store which rows are already taken but I was too lazy to implement that and the code already ran in 0.007 seconds so I decided not to do it (yet), if you feel called to do so you can always send me a PR.

