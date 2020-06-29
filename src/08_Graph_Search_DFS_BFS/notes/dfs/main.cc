#include <iostream>
#include <list>

using namespace std;

class Graph {
    private:
	int numOfVertices;
	list<int> *adj;
	void DFS_helper(int s, bool *visited);

    public:
	Graph(int v);
	void addEdge(int v, int w);
	void DFS(int s);
};

Graph::Graph(int v)
{
	numOfVertices = v;
	adj = new list<int>[v];
}

void Graph::addEdge(int v, int w)
{
	adj[v].push_back(w);
}

void Graph::DFS_helper(int s, bool *visited)
{
	cout << "visiting node " << s << endl;
	visited[s] = true;

	for (auto i = adj[s].begin(); i != adj[s].end(); i++) {
		if (!visited[*i]) {
			cout << "going to vertex " << *i << " from vertex " << s << endl;
			DFS_helper(*i, visited);
		}
	}
}

void Graph::DFS(int s)
{
	bool *visited = new bool[numOfVertices];
	for (int i = 0; i < numOfVertices; i++) {
		visited[i] = false;
	}

	DFS_helper(s, visited);
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

	g.DFS(0);

	return 0;
}
