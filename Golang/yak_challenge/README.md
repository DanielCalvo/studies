

### Dani's program that sorts lists

- To complete this challenge, I used external sorting.

	//This program uses external sorting to sort large amounts of data that will not fit into memory.
	//It has 3 phases:

	//1. Sorting
	//Slices of integers with a certain number of lines are read from a file containing an unsorted list.
	//These slices containing parts of the original list are sorted and saved on temporary files.
	//This website explains the first part well:
	//https://www.geeksforgeeks.org/external-sorting/

	//2. Merging
	//The program loops over the first element of every sorted slice, looking for the one with the highest value.
	//Whichever slice has the highest element, has this element written to the file containing the final list.
	//We then move over to the next element on this slice.
	//Wikipedia explains this part way better than me:
	//https://en.wikipedia.org/wiki/Merge_algorithm#K-way_merging
	//(Thank you Wikipedia, very cool!)

	//3. Printing
	//We look for how many values the user asked (up to 30000000) and scan the final file for these values
	//Then we print them!

	//Oh there's also a number generator at number_generator.go
	//Check it out, it can create an unsorted list as a starting point for this program


### Improvements
- No tests
- Feedback would be nice
- Practice!