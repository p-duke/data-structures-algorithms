# Sum of All Values

A school in HackerLand organized a sholarship exam with this interesting mathematics problem on addition.

Given a string _num_ that consists of digits (0 to 9) a "+" can be inserted between its characters. No adjacent '+' characters are allowed. The value of the expression is then evaluated. Find the sum of the values of all possible expressions after inserting "+" characters any number of times (possibly zero). Since the answer can be large, return the value modulo (10^9 + 7).

## Example
Consider _num = "123"_

All possible valid expressions are shown.
- "1 + 23", value = 24
- "12 + 3", value = 15
- "1 + 2 + 3", value = 6
- "123", value = 123

Their sum is 24 + 15 + 6 + 123 = 168. Thus, the answer is 168 modulo (10^9 + 7) which is 168.

## Function Description
Complete the function getExpressionSums in the editor below.

getExpressionSums has the following parameter:
- _string num:_ the string of numbers

## Returns
_int_: the sum of all possible expression values modulo (10^9 + 7)

## Constraints
- 1 <= |num| <= 17
- _nun_ consists of characters from 0 - 9 (both inclusive) only.
