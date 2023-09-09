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

I really need to write tests. Everything that worked before is now broken (which I know is insane.)

### 05/07

In the parser, how might I distinguish between a list and function args? e.g. (1, 2) Could I say that a function's args are a list? They are parsed differently. Perhaps now is the time to make args a ListStmt? I had thought of args as simple expressions, not statements like lists.

Perhaps the way we parse function args and lists should be similar. Effectively a while (for) loop with a kind of condition. The presence of the => determines whether it is a function or a list and anything else should trickle through the layers. What's the collective term for list and function args? Iterable?

### 07/07

How should variables be printed? For example:

(myVar: "hello", "world")

and

(helloWorld: (name) => >> "hello" + name, helloWorld)

I'm thinking:
- If the variable is an expression statement (i.e it has a primitive value) print the value.
- If it's a function, print "<function>". This seems a bit sucky though! PHP for example prints "Closure Object ()"
    - Would "function" be a usable type that one could filter out of a list? Or considered a null value.

### 10/07

Sometimes native functions need to know what the underlying ast type is in order to determine what to return e.g. the case of "get" with a string. I'm wondering if the interpreter should also do this.

A list can be either an array or a dictionary - there is currently no distinction, but should the interpreter represent it as a `map[string]interface{}` instead of a `[]interface{}`?

Is it possible to iterate over a map in the same order? Go docs say you shouldn't expect the order to be maintained. And we can't order by type if all the keys are interfaces. So perhaps a map with string keys is the way to go. How do other interpreted languages represent maps/dicts?

[Python](https://morepypy.blogspot.com/2015/01/faster-more-memory-efficient-and-more.html) appears to hold a separate, ordered list of mixed indices.

In PHP, an array can be a mixture of indexed and "associative" elements. The docs say "An array in PHP is actually an ordered map". 

How do we know we're on the last item in a map?

So the real question is how to order keys in a map?: The order in which they're defined, which is the order in which they're read by the parser. Though the parser shouldn't be responsible for storing the position because it might change during an interpreter operation.

Writing a programming language is like planting a garden.

### 11/07

The interpreter is converting all ints to float64. That's a problem because when we want to pass an int argument to a function (e.g. .Get(1)), it's actually looking for a float64(1.00) key, which is wrong.

If it's an int, evaluate it as an int. If float, evaluate as float. And so on.

Why the hell did I decide to use floats?

If we're say, performing "2 - 1.2" we know that the left side is an int and the right a float. But Go won't allow us to compare 2 different types. They must both be floats.

However, if performing 2 - 1, we don't want to do 2.00 - 1.00 (converting the ints to floats), we just want to keep the same types and return them.

### 16/07

Everything's broken.

The float issue was actually due to the lexer parsing every number as a float.

I'll still need to fix every comparison (+,-,/ etc) of numbers in the interpreter.

Replacing slices of interfaces for list expressions with an ordered map has broken everything - I borrowed an OrderedMap implementation for the time being which I'm in the process of rolling out.

Later...

There are lots of errors in simple operations we could perform before. My test coverage isn't good enough.

### 17/07

vscode-go has a feature where you can visualise test coverage for a piece of code. It's very useful for seeing which parts of the interpreter we know work. I've found how impossible any of this would be without tests.

I think alot of the type checking I wrote yesterday with reflection will be slow. I should try to refactor it with a type switch.

I was actually feeling quite dispirited by all of the Go type stuff. Writing about it here made the problem seem smaller and more manageable. Hopefully I can get back on track by implementing new language features soon.

The lexer doesn't appear to be adding semicolons at the end of lists. I think I previously implemented this wrong.

### 18/07

Given the statement:
```
(isTrue: (5 * 10) == 50)
```
How would we evaluate the expression 5 * 10 without parentheses? The parser will think this is a list with a single value.

If it _is_ a list, we could write this as `(5 * 10) == (50)`, in which case we would still need to compare the equality of two lists. This would likely be much slower (unless we write an extra condition for when a list contains a single expression).

Or should I use a separate token for expressions?

### 05/09
#### First Light
This has been too hard. Twice I've put down timlang because implementing my original vision has been too challenging.

I think the problem in its simplest form has been that telling the parser to start looking backwards _and_ forwards when it encounters a token is a lot more complicated than using keywords. So from here on, by using keywords it will only need to search forwards. This will likely also make it much faster.

I'm still figuring out the syntax. I'll also keep the old code in the repo. I intend to change "timlang" to something else - it's a bit of a joke gone too far.

### 06/09

What problem/s am I trying to solve?
- It can be hard to perceive data structures (like key/value pairs) when the structure of the code does not mirror the data
- key/value structures and indexed can't usually be interacted with in the same way, which requires knowing their type ahead of time
    - a data structure usually contains a mixture of things that we know and things we don't
    - we should be able to debug data in code
- Function libraries can be inconsistent i.e. order of args

### 09/09

I think I've figured out a rough syntax I like. The implementation will influence the design as much as the paradigms in play