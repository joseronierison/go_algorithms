# Go Algorithms: Search

## Binary search

How does it works, ex.:
```
search for 3 in the array [1, 2, 3, 5, 5, 6, 7, 8, 9]

step 1 -> [1, 2, 3, 5, 5, 6, 7, 8, 9]
           ^           ^           ^   -> 3 != array[4] and 3 < array[4] then it continues ..
        start(0)   middle(4)      end(8)            ^
                                                array[4] == 5

step 2 -> [1, 2, 3, 5, 5, 6, 7, 8, 9]
           ^  ^     ^                  -> 3 != array[1] and 3 > array[2] then it continues ..
     start(0) |     |
              |    end(3)
            middle(1)

step 2 -> [1, 2, 3, 5, 5, 6, 7, 8, 9]
                 ^  ^                  -> 3 == array[2], THEN RETURN 2
                 |  |
                 | end(3)
            start(2)
            middle(2)
```

## Complexity
The binary search has some difference when compared to a linear search, for instance:
- it requires an sorted array
- it does not require to iterate over all the items of the array.

Leaving aside the sorting complexy and assuming we are receiving a ordered array, lest compare the execution complexity for an array with 10000 (10ˆ3) where the item is on the 9998 position of the array.

Using a the **linear search** the loop would go over all the itens in the array to find the last iten of the array. That is the worst case, so the function complexity would behave this way:
- For an array with 10000 items it will iterate 10000 times in the worst case.
- For an array with 100 items it will iterate 1000 times in the worst case
- For an array with 19 items it will iterate 19 times in the worst case.

`O(n)`

Using a **binary search** the loop will iterate much less time as the algorithm reduces the search size in half each iteration. Eg.:
- For an array with with 10000 items it will for the worst case (inexistent value in the array):
```
#1   10000 -> check the item 5000
#2   5000  -> check the items 2500
#3   2500  -> chek the item 1250
#4   1250  -> check the item 625
#5   625   -> check for the item 312
#6   312   -> check for the item 156
#7   156   -> check fo the item 78
#8   78    -> check for the item 39
#9   39    -> check for the item 19
#10  19    -> check for the item 9
#11  9     -> check for the item 4
#12  4     -> check for the item 2
#13  2     -> check for the item 2
#14  1     -> check for the item 1
END
```

As you can see for the example with 10000 itens, instead of iterating 10ˆ3 times on the array, with a binary search is possible to itarete 14 in the worst case.
`O(log n)`

## Implementation

There are two implementation for binary search in the samples. 
- One of the use a for loop that is interrupted once the condition is found. 
- The other one use a recursion to check the array with reduced array as the algorithm does not find the item in the current step.

Take a look on the code to go further.