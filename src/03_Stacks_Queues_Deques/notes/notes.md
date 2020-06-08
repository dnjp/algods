# Stacks, Queues, and Deques

Stacks, queues, deques, and lists are data collections with items ordered
according to how they are added or removed.

- Once an item is added, it stays in the same position relative to its
  neighbors.
- This is why these types of collections are are called _Linear Data
  Structures_.
- They have two ends - "left" and "right", top" and "bottom", or "front" and
  "rear".
- The difference between these data structures is based on where additions and
  removals occur.

## Stacks

A stack is an ordered collection of items where the addition of new items and
the removal of existing items always takes place at the same end.

- This end is typically called the "top" where the opposite end is known as the
  "base".
- This is a LIFO (last in, first out) data structure
- An example of a stack in everyday life would be a stack of plates on a table,
  or books on a desk.
- In a web browser, your browsing history is a stack. As you visit more pages,
  they are added to the stack. When you click the "Back" button, you are popping
  that page off the stack for viewing.

### The Stack Abstract Data Type

An abstract data type (ADT), is a logical description of how we view the data
and the allowed operations without regard to how they'll be implemented.

- We're only concerned with what the data represents, and not with how it will
  be constructed.
- This level of abstraction encapsulates the data and hides implementation
  details from the user (information hiding).

A data structure is an implementation of an ADT and requires a physical view of
the data using some collection of primitive data types and other programming
constructs.

This separation allows us to define complex data models for our problems without
giving any details about how the model will actually be built.

- Implementation-independent view of the data

### Example

My Go implementation can be seen [here](./examples/cmd/stack/stack.go)

The stack abstract data type is an ordered collection of items
where items are added to and removed from the top. The interface for a stack is:

```go
type Stack interface {
	// Creates a new, empty stack
	New() *Stack

	// Adds given item to top of stack
	Push(interface{})

	// Removes and returns the top item from the stack
	Pop() *interface{}

	// Returns the top item from the stack but doesn't remove it
	Peek() *interface{}

	// Returns whether the stack is empty
	IsEmpty() bool

	// Returns the number of items on the stack
	Size() int
}
```

### Balanced Parentheses

In Lisp, you might see a function like

```
(defun square(n)
  (* n n))
```

which defines a function `square` that returns the square of its argument `n`.
In Lisp, parentheses must be balanced. Balanced parentheses means that each
opening symbol has a corresponding closing symbol and the pairs of parentheses
are properly nested. Here are correct examples:

```
(()()()())

(((())))

(()((())()))
```

Here are incorrect examples:

```
((((((())

()))

(()()(()
```

The challenge is to write an algorithm that will read a string of parentheses
from left to right and decide whether the symbols are balanced.

- As you process symbols from left to right, the most recent opening parenthesis
  must match the next closing symbol.
- The first opening symbol processed may have to wait until the very last symbol
  for its match.
- Closing symbols match opening symbols in the reverse order of their
  appearance; they match from the inside out.

Starting with an empty stack, process the parenthesis strings from left to
right.

- If a symbol is an opening parenthesis, push it on the stack as a signal that a
  corresponding closing symbol needs to appear later.
- If, on the other hand, a symbol is a closing parenthesis, pop the stack.
- As long as it's possible to pop the stack to match every closing symbol, the
  parentheses remain balanced.
- If at any time there's no opening symbol on the stack to match a closing
  symbol, the string is not balanced properly. At the end of the string, when
  all symbols have been processed, the stack should be empty.

[Here](./examples/cmd/balanced/balanced.go) is my Go implementation

### Decimal to Binary

[Here](./examples/cmd/divideBy2/divideBy2.go) is my Go implementation

### Infix, Prefix, and Postfix Expressions

We are accustomed to writing mathematical expressions like B _ C. In this case,
we know that the variable B is being multiplied by the variable C since the
multiplication operator _ appears between them in the expression. This is known
as infix notation since the operator is in between the two operands that it's
working on.

If we consider A + B _ C, how do we know which operands the operator is working
on? We know that each operator has a precedence level. Operators of higher
precedence are used before operators of a lower precedence level. In this case,
B and C are multiplied first, and A is then added to the result. (A + B) _ C
would force the addition of A and B to be done first before the multiplication.

Computers need to know exactly what operations to perform and in what order. One
way to write an expression that guarantees there will not be confusion is to use
a fully parenthesized expression. This type of expression uses one pair of
parentheses for each operator. The expression A + B _ C + D can be rewritten as
((A + (B _ C)) + D) to show that the multiplication happens first, followed by
the leftmost addition.

Languages like Lisp use prefix notation. So instead of a statement like (B _ C),
we would write (_ B C) where the operator (_) is applied to all operands in the
expression. Outside of programming, we would write an expression like A + B _ C
as + A \* B C.

In postfix notation, the expression would be A B C _ +. The order of operations
is preserved since the _ appears immediately after the B and the C.

Now consider the infix expression (A + B) _ C. Recall that in this case, infix
requires the parentheses to force the performance of the addition before the
multiplication. In prefix notation, the multiplication operator is moved in
front of the entire expression, giving us _ + A B C. Likewise, in postfix the
expression would be A B + C \*.

// TODO continue

## Queues

A queue is an ordered collection of items where the addition of new items
happens at one end, called the "rear", and removal of existing items occurs at
the other end, called the "front". As an element enters the queue it starts at
the rear and makes its way toward the front, waiting until that time when it is
the next element to be removed.

- Queues are FIFO (first-in first-out) data structures
  - The most recently added item in the queue must wait at the end of the
    collection.
  - The item that has been in the collection the longest is at the front.

The interface for a Queue is as follows:

```go
type Queue interface {
	// constructs a new empty queue
	New() *Queue

	// Adds a new item to the rear of the queue
	Enqueue(interface{})

	// Removes the front item from the queue
	Dequeue() *interface{}

	// Returns whether or not the queue is empty
	IsEmpty() bool

	// Returns the number of items in the queue
	Size() int
}
```

[Here](./examples/queue/queue.go) is my Go implementation

### Example

[Here](./examples/cmd/hot_potato/hot_potato.go) is an example of the game Hot
Potato

## Deques

A deque, also known as a double-ended queue, is an ordered collection of items
similar to the queue. It has two ends, a front and a rear, and the items remain
positioned in the collection. What makes a dequeue different is the
unrestricted nature of adding and removing items. New items can be added at
either the front of the rear. Likewise, existing items can be removed from
either end. In a sense, this hybrid linear structure provides all the
capabilities of stacks and queues in a single data structure.

The deque operations are given below:

```go
type Deque interface {
	// Creqtes an empty deque
	New() *Deque

	// Adds a new item to the front of the deque
	AddFront(interface{})

	// Adds a new item to the rear of the deque
	AddRear(interface{})

	// Removes the front item from the deque
	RemoveFront() *interface{}

	// Removes the rear item from the deque
	RemoveRear() *interface{}

	// Returns whether the deque is empty
	IsEmpty() bool

	// Returns the number of items in the deque
	Size() int
}
```
