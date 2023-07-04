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

I need to implement variables and scope before thinking about native functions.

### 26/06/23

I'm struggling with assignment and global variables. Top level list items aren't being created as globals that can be used in other parts of the program, only within the same list. I think this is because the first thing we do is enter into a list, which creates a new environment.

Maybe this has nothing to do with how lists are executed, but how variables are defined.

Creating a new environment with the same enclosing environment is still new and therefore a different pointer, which means it is different.

Later... I think I have sorted it. I should probably write some tests now.

I think a callable will have to be a statement, not an expression. I don't think there's any instance of a function that will not be attached to a list in some way.

### 04/07/23

I am completely dumbfounded as to how to print statements and expressions. A list is a special kind of statement, whose values (other statements and expressions) we want to print.

The approach I have taken thus far is to create a "PrintVisitor" and then how to handle the printing of each type of statement & expression is handled by the struct itself. However, this is a problem because what I actually want to do is _evaluate_ the value first with the interpreter, then print the values in their own kind of way. But I don't have access to the interpreter in this current implementation.

^^ I think this works now.

I need to figure out how to chain methods. It would mean that the initialiser for a call statement could be another call statement, not a list.

None of that print logic worked. The solution was much simpler.

I really need to write tests. Everything that worked before is now broken.