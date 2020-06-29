#include <iostream>
#include <list>

using namespace std;

class Graph {
    private:
	// Number of vertices
	int numVertices;
	// Pointer to adjacency list
	list<int> *adj;

    public:
	Graph(int v);
	void addEdge(int v, int w);
	void BFS(int s);
};

Graph::Graph(int v)
{
	numVertices = v;
	adj = new list<int>[v];
}

void Graph::addEdge(int v, int w)
{
	adj[v].push_back(w);
}

void Graph::BFS(int startingPoint)
{
	bool *visited = new bool[numVertices];
	for (int i = 0; i < numVertices; i++) {
		visited[i] = false;
	}

	list<int> queue;

	visited[startingPoint] = true;
	queue.push_back(startingPoint);

	while (!queue.empty()) {
		// remove front of queue
		startingPoint = queue.front();
		queue.pop_front();

		// get all adjacent vertices from that vertex
		cout << "Checking adjacent vertices for vertex " << startingPoint << endl;
		for (auto i = adj[startingPoint].begin(); i != adj[startingPoint].end(); i++) {
			if (!visited[*i]) {
				cout << "Visit and enqueue " << *i << endl;
				// mark as visited
				visited[*i] = true;

				queue.push_back(*i);
			}
		}
	}
}

int main()
{
	Graph g(6);

	// connections for vertex 0
	g.addEdge(0, 1);
	g.addEdge(0, 2);
	// connections for vertex 1
	g.addEdge(1, 0);
	g.addEdge(1, 3);
	g.addEdge(1, 4);
	// connections for vertex 2
	g.addEdge(2, 0);
	g.addEdge(2, 4);
	// connections for vertex 3
	g.addEdge(3, 1);
	g.addEdge(3, 4);
	g.addEdge(3, 5);
	// connections for vertex 4
	g.addEdge(4, 1);
	g.addEdge(4, 2);
	g.addEdge(4, 3);
	g.addEdge(4, 5);
	// connections for vertex 5
	g.addEdge(5, 3);
	g.addEdge(5, 4);

	g.BFS(4);

	return 0;
}
