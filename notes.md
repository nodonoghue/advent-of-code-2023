# Day 12 Notes
- Part one took longer than expected, expanding the rows by 5+ will have an almost exponential increase in the number of permutations.  This is not going to brute force well at all.  Need to study up on dynamic programming, in particular memoization.  This seems to be the generally accepted "correct" method to get this solution.

# Day 13 Notes
- to find symmetry look at the rows and cols from the outside and look at the opposite edge, and move the points inward one at a time until an equality is found.  then move inward on both ends and confirm more equality until the inner two most reflections are found.

- Column symmetry search can also be done by rotating the grid and performing the row symmetry check again.

- Possible to start at index 1 instead of 0 if the rows are odd since outter rows and cols that don't exactly match are ignored, skipping the first of a reflection won't change the point of reflection.  The point of reflection is the important find.

- Part two will change the condition to find mirroring sections that have exactly one error.