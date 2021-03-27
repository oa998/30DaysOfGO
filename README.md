## 30DaysOfGO

Learning GO with 30 small projects:

1. Rock Paper Scissors

- create a main function that generates two random throws of rock-paper-scissors and announces the winning throw (or a draw)

2. Binary Search algorithm

- implement any sort algorithm (except bubble) and then implement a binary search to find the index of a particular entry.

3. Fibonacci check

- given a number, return N or (-1). If the number is in the fibonacci sequence, return the index of the location Or (-1) if it's not part of the sequence

4. Ceasar Encode/Decode

- Two functions. Function one takes a string and returns an encoded string and a number (positive or negative). Function two takes a string and a number and returns the decoded string.

5. Pig Latin

- given a statement of words, take the first letter of every word, moving it to the end of the word and adding ‘ay’

6. Coordinate Library

- Two functions. Each takes two (x, y) coordinates. One function returns the midpoint of the two coordinates. The other function returns the distance between the two coordinates.

7. Slice Library

- Two functions. Function one takes a single slice of elements and returns it randomized. Function two takes 2 slices of the the same types and returns one list where elements of list two are "weaved" (every-other) into list one.

8. Permissions Library

- Create permissions assigned to 1-hot "permissions" vectors (integers). Given any integer, return a list of strings that identify what permissions that integer "has". (1101 & 0011 = 0001)

9. tr -s

- given a string, return a string where all adjacent, repeated characters are reduced to just 1 instance of that character. ("aaabbcdd" -> "abcd")

10. Cubing

- Function tries to find the best way to fit small boxes into a bigger box. Write a function that takes 2 structs each with 3 integers (length, width, height). Struct 1 is the big box. Struct 2 is the little box. You may rotate the smaller box 90 degrees on any axis. Find the maximum number of little boxes you can fit into the bigger box.

11. Odds or Evens

- Function that takes a positive integer. Without converting the integer to a string, determine which is larger: the sum of all the odd digits or the sum of all the even digits. Return the string "even" or "odd" accordingly.

12. Coprime

- Function that takes 2 positive integers. Return True if the two numbers are coprime. Else return false.

13. Largest subsequence

- Function that takes 2 positive integers. The first number should be greater than 9999. The second number should be less than 10. Search the first number for the largest number with as many digits as stated by the second parameter. (378289324, 5 => 89324)

14. Hash Table

- Implement a very simple hash table data structure. The structure should have a fixed size and store that many values. Hash strings using any preferred approach and store the string using the hash. Collisions should log the original entry before over-writting it.

15. Encode/Decode

- Create a function that takes in a string and returns both the string with characters shuffled around and a slice of tuples representing instructions to unshuffle the output. Create a second function that takes in the shuffled string and a slice of tuples, returning the string unshuffled. Tuples are indexes in the shuffled string and a number of movements in a direction for that letter.
  ("abcxyz" => "acbxzy" & [(1, 1), (4, 1)])

16. Newton's Method for square root

- implement Newton's method to estimate the square root of a number. Handle invalid inputs so your code doesn't get into an infinite loop. https://en.wikipedia.org/wiki/Newton%27s_method#Square_root

17. Base 64

- implement a function to encode a string into base 64. implement a function to decode a base 64 string into a string. Refer: http://fm4dd.com/programming/base64/base64_algorithm.htm

18. Hex to RGB

- implement a function that takes a hex string ("#44ffaa") and outputs the rgb triplet (68, 255, 170)

19. Balance the parentheses

- implement a function that takes a string containing any number of open and close parentheses ("()))(") an outputs a string where additional parentheses were added to close all pairs ("((()))()")

20. Strong password

- implement a function that takes in a string and returns "Valid" if the password is good. If the password doesn't contain at least two of the following, it is invalid: uppercase letter, lowercase letter, number, special character. Use regular expressions in your implementation. Return a string explaining the error(s) if the password is invalid ("Must contain two special characters")

21. Random Number Generator

- implement your own random number generating function. Find any approach on the internet that you want to implement and do so.

22.
