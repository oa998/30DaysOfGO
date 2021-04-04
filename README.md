## 30DaysOfGO

#### I know nothing about _golang_, but I want to learn.

I made the following list of 30 projects that I think will be reasonable challenges to help me learn GO. Some of these are very easy, others are more difficult. My goal is to do at least one a day until they are all done. I'm not going to do them in any particular order. I'm going to use any standard libraries where it makes sense unless the challenge is to implement a standard library function - such as a sorting algorithm.

Learning GO with 30 small projects:

1. **Rock Paper Scissors**

- create a main function that generates two random throws of rock-paper-scissors and announces the winning throw (or a draw)

2. **Binary Search algorithm**

- implement any sort algorithm (except bubble) and then implement a binary search to find the index of a particular entry.

3. **Fibonacci check**

- given a number, return N or (-1). If the number is in the fibonacci sequence, return the index of the location Or (-1) if it's not part of the sequence

4. **Ceasar Encode/Decode**

- Two functions. Function one takes a string and returns an encoded string and a number (positive or negative). Function two takes a string and a number and returns the decoded string.

5. **Pig Latin** ✅ (2. duration ~15 minutes)

- given a statement of words, take the first letter of every word, moving it to the end of the word and adding ‘ay’

6. **Coordinate Library**

- Two functions. Each takes two (x, y) coordinates. One function returns the midpoint of the two coordinates. The other function returns the distance between the two coordinates.

7. **Slice Library**

- Two functions. Function one takes a single slice of elements and returns it randomized. Function two takes 2 slices of the the same types and returns one list where elements of list two are "weaved" (every-other) into list one.

8. **Permissions Function** ✅ (3. duration ~20 minutes)

- Create permissions assigned to 1-hot "permissions" vectors (integers). Given any integer, return a list of strings that identify what permissions that integer "has". (1101 & 0011 = 0001)

9. **tr -s** ✅ (4. duration ~8 minutes)

- given a string, return a string where all adjacent, repeated characters are reduced to just 1 instance of that character. ("aaabbcdd" -> "abcd")

10. **Cubing**

- Function tries to find the best way to fit small boxes into a bigger box. Write a function that takes 2 structs each with 3 integers (length, width, height). Struct 1 is the big box. Struct 2 is the little box. You may rotate the smaller box 90 degrees on any axis. Find the maximum number of little boxes you can fit into the bigger box.

11. **Odds or Evens** ✅ (5. duration ~10 minutes)

- Function that takes a positive integer. Without converting the integer to a string, determine which is larger: the sum of all the odd digits or the sum of all the even digits. Return the string "even" or "odd" accordingly.

12. **Coprime**

- Function that takes 2 positive integers. Return True if the two numbers are coprime. Else return false.

13. **Largest subsequence** ✅ (6. duration ~20 minutes)

- Function that takes 2 positive integers. The first number should be greater than 9999. The second number should be less than 10. Search the first number for the largest number with as many digits as stated by the second parameter. (378289324, 5 => 89324)

14. **Hash Table**

- Implement a very simple hash table data structure. The structure should have a fixed size and store that many values. Hash strings using any preferred approach and store the string using the hash. Collisions should log the original entry before over-writting it.

15. **Encode/Decode**

- Create a function that takes in a string and returns both the string with characters shuffled around and a slice of tuples representing instructions to unshuffle the output. Create a second function that takes in the shuffled string and a slice of tuples, returning the string unshuffled. Tuples are indexes in the shuffled string and a number of movements in a direction for that letter.
  ("abcxyz" => "acbxzy" & [(1, 1), (4, 1)])

16. **Newton's Method for square root**

- implement Newton's method to estimate the square root of a number. Handle invalid inputs so your code doesn't get into an infinite loop. https://en.wikipedia.org/wiki/Newton%27s_method#Square_root

17. **Base 64**

- implement a function to encode a string into base 64. implement a function to decode a base 64 string into a string. Refer: http://fm4dd.com/programming/base64/base64_algorithm.htm

18. **Hex to RGB**

- implement a function that takes a hex string ("#44ffaa") and outputs the rgb triplet (68, 255, 170)

19. **Balance the parentheses**

- implement a function that takes a string containing any number of open and close parentheses ("()))(") an outputs a string where additional parentheses were added to close all pairs ("((()))()")

20. **Strong password**

- implement a function that takes in a string and returns "Valid" if the password is good. If the password doesn't contain at least two of the following, it is invalid: uppercase letter, lowercase letter, number, special character. Use regular expressions in your implementation. Return a string explaining the error(s) if the password is invalid ("Must contain two special characters")

21. **Random Number Generator**

- implement your own random number generating function. Find any approach on the internet that you want to implement and do so.

22. **Contact Tracing**

- Implement two functions. Function one takes in a slice of tuples. The tuples consist of the names of individuals. These two individuals had direct contact with one another. Function two takes a single tuple of two names. Function two returns a number (or list of names) representing the minimum distance between those two people given all the tuples stored from function one.

23. **Parallel routines**

- Implement two functions that take the same inputs and have the same outputs. Each function will take a slice of positive integers and return a slice of integers. Within the function, iterate over the list of floats, and for each float count the number of values between 1 that value that have a perfect square root. (use the math package). Return this count. (for example, the number 12 exceeds 3 perfect square roots: 1, 4, 9). In one of the functions, use parallel processing to improve speed of execution.

24. **Consuming API**

- Implement a method to call a GET API that responds with JSON. Log that data to the console in any way you'd like. (API: https://randomuser.me/api/ Reference Material: https://randomuser.me/)

25. **Blackjack**

- Implement a main method that deals enough cards to both players until they each have a hand that is worth more than 16. If any player has a hand worth more than 21, that player loses. If both players have hands worth 21 or less, the one with the highest hand wins.

26. **Repeated Phrases** ✅ (1. duration ~2 hours)

- Given a text file of prose (sentences, like exerpts from a book), separate the text into sentences then compose a map of phrases where the key is the phrase and the value is an integer. Phrases are any word combinations of 2 or more words. Ignore capitalization and punctionation. A sentence is delimited by a period (.) or exclamation point (!).

27.
