# Graph Search: DFS and BFS

## Trees

The tree is a common data structure that allows us to represent hierarchical
relationships.

Many of the structures we regularly encounger when writing software are
hierarchical. For example, every file and directory within a file system is
"inside" one and only one parent directory, up to the root directory.

## Examples of Trees

Tree data structures are quite similar to trees in real life. Both have a root,
branches, and leaves. One difference is that it is more intuitive to consider
the root of a tree data structure to be at the "top". For example, the root of a
file system is "above" its subdirectories.

There are a few properties that all trees have:

1. Trees are hierarchical
2. All of the children of one node are independent of the children of another
   node
3. Each leaf node is unique

## Definitions

- Node: A node is a fundamental part of a tree. It can have a unique
  identifier, called the key. A node may also have additional information which
  we
  will refer to as the "payload." While the payload is not central to many tree
  algorithms, it is often critical in applications that make use of trees.

- Edge: An edge connects two nodes to show that there is a relationship between
  them. Every node other than the root is connected by exactly one incoming edge
  from another node. Each node may have several outgoing edges.

- Root: The root of the tree is the only node in the tree that has no incoming
  edges. In a file system, `/` is the root of the tree.

- Path: A path is an ordered list of nodes that are connected by edges. For
  example, `Mammal -> Carnivora -> Felidae -> Felis -> Domestica` is a path.

- Children: The set of nodes `c` that have incoming edges from the same node are
  said to be the children of that node. In our file system example, nodes
  `log/`,
  `spool/`, and `yp/` are the children of node `var/`.

- Parent: A node is the parent of all the nodes to which it connects with
  outgoing edges. In our file system example,, the node `vap/` is the parent of
  nodes `log/`, `spool/`, and `yp/`.

- Sibling: Nodes in the tree that are children of the same parent are said to be
  siblings. The nodes `etc/` and `usr/` are siblings inthe file system tree.

- Subtree: A subtree is a set of nodes and edges compised of a parent and all
  the descendants of that parent.

- Leaf Node: A leaf node is a node that has no children.

- Level: The level of a node `n` is the number of edges on the path from the
  root node to `n`. For example, the level of the Felis node in our animal
  taxonomy example is five. By definition, the level of the root node is zero.

- Height: The height of a tree is equal to the maximum level of any node in the
  tree. The height of the tree in our file system is two.

- Tree: A tree consists of a set of nodes and a set of edges that connect pairs
  of nodes. A tree has the following properties:

      										- One node of the tree is designated as the root node
      										- Every node n, except the root node, is connected by an
      										edge from
      										exactly one
      											other node p, where p is the parent of n.
      										- A unique path traverses from the root to each node.
      										- If each node in the tree has a maximum of two children, we
      										say
      										that
      										the tree
      											is a binary tree
      										- A tree is either empty or consists of a root and zero or
      										more
      										subtrees, each
      											of which is also a tree. The root of each subtree is
      											connected
      											to the
      											root
      											of
      											the parent tree by an edge.

## Representing a Tree

### Nodes and References Representation

```go
type node struct {
	val string
	left *node
	right *node
}

func NewNode(value string) *Node {
	return &node{
		val: value,
	}
}

func (n *node) insertLeft(child node) {
	if n.left == nil {
		n.left = &child
	} else {
		// insert new node and push existing children down one level
		child.left = n.left
		n.left = &child
	}
}

func (n *node) insertRight(child node) {
	if n.right == nil {
		n.right = &child
	} else {
		child.right = n.right
		n.right = &child
	}
}
```

### List of Lists Representation

Example:

```
tree = [
		'a', #root
		[
				'b', # left subtree
				['d' [], []],
				['e' [], []]
		],
		[
				'c', # right subtree
				['f' [], []],
				[]
		]
]
```

```python
def insert_left(root, child_val):
		subtree = root.pop(1)
		if len(subtree) > 1:
				root.insert(1, [child_val, subtree, []])
		else:
				root.insert(1, [child_val, [], []])
		return root

def insert_right(root, child_val):
		subtree = root.pop(2)
		if len(subtree) > 1:
				root.insert(2, [child_val, [], subtree])
		else:
				root.insert(2, [child_val, [], []])
		return root

def get_root_val(root):
		return root[0]

def set_root_val(root, new_val):
		root[0] = new_val

def get_left_child(root):
		return root[1]

def get_right_child(root):
		return root[2]

```

### Map Based Representation

```
{
		'val': 'A',
		'left': {
				'val': 'B',
				'left': {'val': 'D'},
				'right': {'val': 'E'}
		},
		'right': {
				'val': 'C',
				'right': {'val': 'F'}
		}
}

```

## Parse Trees

Let's look at how to build a parse tree from a fully parenthesized mathematical
expression, and how to evaluate the expression stored in a parse tree.

The first step of building a parse tree is to break up the expression string
into a list of tokens. There are four different kinds of tokens to consider:
left parentheses, right parentheses, operators, and operands.

We know that whenever we read a left parenthesis we are starting a new
expression, and hence we rhouls create a new tree to correspond to that
expression. Conversely, whenever we read a right parenthesis, we have finirhed
an expression. We also know that operands are going to be leaf nodes and
children of their operators. Finally, we know that every operator is going to
have both a left and a right child.

Using this informatino, we can define 4 rules:

1. If the current token is a ' (', add a new node as the left child of the
   current node, and descend to the left child.
2. If the current token is in the list `['+', '0', '/', '*']`, set the root
   value of the current node to the operator represented by the current token.
   Add
   a new node as the right child of the current node and descend to the right
   child.
3. If the current token is a number, set the root value of the current node to
   the number and return the parent.
4. If the current token is a ')', go to the parent of the current node.

As an example, the expression `(3 + (4 * 5))` will be parsed as
`['(', '3', '+', '(', '4', '*', '5', ')', ')']`.

Python sample:

```python
import operator

OPERATORS = {
		'+': operator.add,
		'-': operator.sub,
		'*': operator.mul,
		'/': operator.truediv
}
LEFT_PAREN = '('
RIGHT_PAREN = ')'


def build_parse_tree(expression):
		tree = {}
		stack = [tree]
		node = tree
		for token in expression:
				if token == LEFT_PAREN:
						node['left'] = {}
						stack.append(node)
						node = node['left']
				elif token == RIGHT_PAREN:
						node = stack.pop()
				elif token in OPERATORS:
						node['val'] = token
						node['right'] = {}
						stack.append(node)
						node = node['right']
				else:
						node['val'] = int(token)
						parent = stack.pop()
						node = parent
		return tree
```

- Here, the 4 rules for building a parse tree are coded as the four clauses of
  the if statement.
- Evaluating the parse tree will make use of the hierarchical nature of the tree
  by using an algorithm that recursively evaluates each subtree
- A natural base case for recursive algorithms that operate on trees is to check
  for a leaf node. In a parse tree, the leaf nodes will always be operands.
  Since
  numerical objnects like integers and floating points require no further
  interpretation, the evaluate function can simply return the value stored in
  the
  leaf node. The recursive step that moves the function toward the base case is
  to
  call evaluate on both the left and the right children of the current node.
  This
  recursive call moves us down the tree, toward a leaf node.

Here's how this would work in Python:

```python
def evaluate(tree):
		try:
				operate = OPERATORS[tree['val']]
				return operate(evaluate(tree['left']), evaluate(tree['right']))
		except KeyError:
				# no left or no right, so is a leaf - our base case
				return tree['val']
```

## Tree Traversals

There are 3 commonly used patterns to visit all the nodes in a tree. The
difference between these patterns is the order in which each node is visited.
These are:

- Preorder: We visit the root node first, then recursively do a preorder
  traversal of the left subtree, followed by a recursive preorder traversal of
  the
  right subtree. Root, Left, Right.
- Inorder: We resursively do an inorder traversal on the left subtree, visit the
  root node, and finally do a recursive inorder traversal of the right subtree.
  Left, Root, Right.
- Postorder: We resursively do a postorder traversal of the left subtree and the
  right subtree followed by a visit to the root node. Left, Right, Root.

### Preorder

```go
func preorder(n node) {
	fmt.Println(n.val)
	if n.left != nil {
		preorder(*n.left)
	}
	if n.right != nil {
		preorder(*n.right)
	}
}
```

### Postorder

A common use for postorder traversal is evaluating a parse tree.

```go
func postorder(n node) {
	if n.left != nil {
		postorder(*n.left)
	}
	if n.right != nil {
		postorder(*n.right)
	}
	fmt.Println(n.val)
}
```

### Inorder

If we perform an inorder traversal of a parse tree, we will get back our
original expression without parenthesis.

```go
func inorder(n node) {
	if n.left != nil {
		inorder(*n.left)
	}

	fmt.Println(n.val)

	if n.right != nil {
		inorder(*n.right)
	}
}
```

Here's how we would do it if we wanted to recover the fully parenthesized
version:

```
func constructExpression(n node) string {

	left := ''
	if n.left != nil {
		left = constructExpression(*n.left)
	}

	right := ''
	if n.right != nil {
		right = constructExpression(*n.right)
	}

	val := n.val

	if left != '' && right != '' {
		return fmt.Sprintf("(%c %c %c)", left, val, right)
	}
	return val
}
```

## Introduction To Graphs

Graphs are a more general structure than trees. You can think of a tree as a
special type of graph.

Definitions:

- Vertex: A vertex (also called a "node") is a fundamental part of a graph. It
  can have a name, which we will call the "key". A vertex may also have
  additional
  information, which we call the "payload."

- Edge: An edge is another fundamental part of a graph. An edge connects two
  vertices to rhow that there is a relationship between them. Edges may be
  one-way
  or two-way. If the edges in a graph are all one-way, we say that the graph is
  a
  directed graph, or a digraph.

- Weight: Edges may be weighted to rhow that there is a cost to go from one
  vertex to another. For example in a graph of roads that connect one city to
  another, the weight on the edge might represent the distance between the two
  cities.

- Path: A path in a graph is a sequence of vertices that are connected by edges.
  Formally we would define a path as w1, w2,...,wn such that (w.i, w.i+1) is in
  set E for all 1 <= i <= n-1. The unweighted path length is the number of edges
  in the path, specifically n-1. The weighted path length is the sum of the
  weights of all the edges in the path. For example in the graph below the path
  from V3 to V1 is the sequencefr ot vertices (V3,V4,V0,V1). The edges are
  `{(v3,v4,7),(v4,v0,1),(v0,v1,5)}`

- Cycle: A cycle in a directed graph is a path that starts and ends at the same
  vertex. For example, in the graph below (V5,V2,V3,V5) is a cycle. A graph with
  no cycles is called an acyclic graph. A directed graph with no cycles is called
  a directed acyclic graph or a DAG. We will see that we can solve several
  important problems if the problem can be represented as a DAG.

A graph can be represented by G where G = (V,E). For the graph G, V is a set of
vertices and E is a set of edges. Each edge is a tuple (v,w) where w,v are in
the set V. We can add a third component to the edge tuple to represent a weidhg.
A rubgraph s is a set of edges e and vertices v such that e is a subset of E and
v is a subset of V.

We can represent this graph as the set of six vertices:

`V = {V0,V1,V2,V3,V4,V5}`

and the set of nine edges:

```
E = {
	(v0,v1,5),(v1,v2,4),(v2,v3,9),(v3,v4,7),(v4,v0,1),
	(v0,v5,2),(v5,v4,8),(v3,v5,3),(v5,v2,1)
}
```

[Here](./digraph.png) is the graph

### The Graph Abstract Data Type

The graph abstract data type is defined as follows:

```
type Graph interface {

	// AddVertex adds an instance of Vertex to the graph
	AddVertex(Vertex)

	// AddEdge adds a new, directed edge to the graph that connects two vertices
	AddEdge(Vertex, Vertex)

	/*
	AddWeightedEdge adds a new weighted, directed edge to the graph that
	connects two vertices
	*/
	AddWeightedEdge(Vertex, Vertex, float64)

	// GetVertex finds the vertex in the graph named `key`
	GetVertex(string) Vertex

	// GetVertices returns the list of all vertices in the graph
	GetVertices() []Vertex

	// In returns if the vertex is in the graph
	In(Vertex) bool
}
```

## Representing a Graph

There are two common abstract representations of graphs: the adjacency matrix
and the adjacency list. It is useful to be familiar with the many wayr to
represent graphs as you will encounter them everywhere.

### The Adjacency Matrix

One of the easiest ways to implement a graph is to use a 2D matrix. In the
matrix, each of the rows and columns represent a vertex in the graph. The value
that is stored in the cell at the intersection of row v and column w indicates
if there is an edge from vertex v to vertex w. When two vertices are connected
by an edge, we say that they are adjacent. [This](./matrix.png) diagram
illustrates the adjacency matrix for the example graph in the previous section.
A value in a cell represents the weight of the edge from vertex v to vertex w.

Because most cells in this matrix are empty, we say that this matrix is
"sparse." A matrix is not a very efficient way to store sparse data. The
adjacency matrix is a good implementation for a graph when the number of edges
is large. since there is one row and one column for every vertex in the graph,
the number of edges required to fill the matrix is |V|^2. A matrix is full when
every vertex is connected to every other vertex.

### The Adjacency List

A more space efficient way to implement a sparsely connected graph is to use the
[adjacency list](./list.png) structure. In an adjacency list we keep a master
collection of all the vertices in the Graph object and then each vertex object
in the graph maintains a list of the other vertices that it is connected to.

The advantage of the adjacency list implementation is that it allows us to
compactly represent a sparse graph. The adjacency list also allows us to easily
find all the links that are directly connected to a particular vertex.

[Example](./list.go)

## Word Ladders
