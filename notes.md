# Day 12 Notes
- For part one possible to brute force by taking the number of `?` and turning into a boolean array, sending through a method to get all permutations, putting the characters back into the source line and checking if valid.

- To check if valid, possibly use a regex for the pattern matching?

# Day 13 Notes
- to find symmetry look at the rows and cols from the outside and look at the opposite edge, and move the points inward one at a time until an equality is found.  then move inward on both ends and confirm more equality until the inner two most reflections are found.

- Column symmetry search can also be done by rotating the grid and performing the row symmetry check again.

- Possible to start at index 1 instead of 0 if the rows are odd since outter rows and cols that don't exactly match are ignored, skipping the first of a reflection won't change the point of reflection.  The point of reflection is the important find.

- Part two will change the condition to find mirroring sections that have exactly one error.

- validation regex: `#{1}/*.+/*#{5}/*.+/*#{1}/*.+/*#{1}/*.+/*#{1}`