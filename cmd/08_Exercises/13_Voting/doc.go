package main

// "aaabbbcddd" a -> 3 b -> 3 c -1 d -> 3

// => Winner is : a  (max votes, single person)
// => loser is : b (min votes, single person)
// => no winner
// every letter is a vote.
// if maximum number votes results twice, no winner

// Length of the string is N (i.e. n different votes)
// "aaaabbbcddd" -> Winner is a, Loser is c
// "aaabbbcddd" -> No winner
// "aaabcd" -> Winner is a
// "abcd" -> No Winner
