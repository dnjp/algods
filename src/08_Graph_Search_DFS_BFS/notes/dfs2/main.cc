#include <iostream>
#include <string>
#include <vector>
#include <utility>
#include <set>
#include <map>
#include <functional>

using namespace std;

class Vertex {
    public:
	Vertex()
	{
	}

	int discovery;
	int finish;
};

map<string, Vertex> dfs(map<string, vector<string> > graph, string starting_vertex)
{
	set<string> visited;
	vector<int> counter{ 0 };
	map<string, Vertex> traversal_times;

	function<void(string)> traverse;
	traverse = [&visited, &counter, &traversal_times, &graph, &traverse](string vertex) {
		visited.insert(vertex);
		counter[0] += 1;
		traversal_times[vertex].discovery = counter[0];

		for (string next_vertex : graph[vertex]) {
			const bool in_visited = visited.find(next_vertex) != visited.end();
			if (!in_visited) {
				traverse(next_vertex);
			}
		}

		counter[0] += 1;
		traversal_times[vertex].finish = counter[0];
	};

	traverse(starting_vertex);
	return traversal_times;
}

int main()
{
	map<string, vector<string> > simple_graph = {
		{ "A", { "B", "D" } }, { "B", { "C", "D" } }, { "C", {} },
		{ "D", { "E" } },      { "E", { "B", "F" } }, { "F", { "C" } },
	};

	map<string, Vertex> traversal_times = dfs(simple_graph, "A");

	for (pair<string, Vertex> element : traversal_times) {
		cout << "key: " << element.first << std::endl;
		cout << "  {" << std::endl;
		cout << "    discovery: " << element.second.discovery << std::endl;
		cout << "    finish: " << element.second.finish << std::endl;
		cout << "  }" << std::endl;
	}
}
