### 22/06/23
I'm thinking about adopting Go's approach of the lexer adding semicolons to the end of lines and then the parser interprerting these as the ends of statements. The semicolon needs to be inserted under special criteria because often a newline does not necessarily denote the end of a statement.

If everything in tim is an expression (e.g. ("hello world"), (1, 2, 3), (1 + 2)), an expression must be able to have methods like ("hello world").print()

I'm going to skip expression functions (e.g. .print()) for the time being until I have variables and callables in place.

I should probably update the lexical definition so that statements that begin with "(" are considered a list. And then have "list" expression. I need to be able to differentiate between parentheses used for qualifying expressions and lists. Does an expression need to be wrapped in parentheses even?

### 23/06/23
How do you distinguish between a contained expression e.g. "hello mr " + (3 + 6) (the 3 + 6 bit) and a list? When does LEFT_PAREN not indicate a list? When there is only one value in the list?

(2, 3) + (4, 5) = (2, 3, 4, 5)

(2, 3) + ("4") = (2, 3, "4")

(2 + 3) + "hello" = (5, "hello")

So perhaps using + for concatenating strings should not be allowed. And I don't really like the + syntax.

I would prefer:
(2, 3).concat(4, 5) // 2, 3, 4, 5
(3, 3).concat("four") // 2, 3, "four"

How to evaluate the following statement?
("hello", "world").join(" ").print()

Does join return a list to print?

Or are all the methods properties of the list and then get evaluated later? Yes, probably.

Will a user defined function only be callable as an argument to a native function?

There will probably never be a callable on its own. It will always be a list function, so I should write it as such.