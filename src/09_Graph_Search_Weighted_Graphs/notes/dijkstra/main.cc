#include <iostream>
#include <map>
#include <vector>
#include <string>
#include <limits>
#include <queue>
#include <utility>

using namespace std;

map<string, int> calculateDistances(map<string, vector<pair<string, int> > > graph, string starting_vertex)
{
	map<string, int> distances;
	for (pair<string, vector<pair<string, int> > > elem : graph) {
		string key = elem.first;
		distances[key] = numeric_limits<int>::max();
	}
	distances[starting_vertex] = 0;

	priority_queue<pair<int, string> > pq;
	pq.push(make_pair(0, starting_vertex));

	while (pq.size() > 0) {
		pair<int, string> top = pq.top();
		pq.pop();
		int current_distance = top.first;
		string current_vertex = top.second;

		if (current_distance > distances[current_vertex]) {
			continue;
		}

		for (pair<string, int> elem : graph[current_vertex]) {
			string neighbor = elem.first;
			int weight = elem.second;
			int distance = current_distance + weight;

			// Only consider this path if it's better than previous
			if (distance < distances[neighbor]) {
				distances[neighbor] = distance;
				pq.push(make_pair(distance, neighbor));
			}
		}
	}

	return distances;
}

int main()
{
	map<string, vector<pair<string, int> > > simple_graph = {
		{ "U", { { "V", 2 }, { "W", 5 }, { "X", 1 } } },
		{ "V", { { "U", 2 }, { "X", 2 }, { "W", 3 } } },
		{ "W", { { "V", 3 }, { "U", 5 }, { "X", 3 }, { "Y", 1 }, { "Z", 5 } } },
		{ "X", { { "U", 1 }, { "V", 2 }, { "W", 3 }, { "Y", 1 } } },
		{ "Y",
		  {
			  { "X", 1 },
			  { "W", 1 },
			  { "Z", 1 },
		  } },
		{ "Z", { { "W", 5 }, { "Y", 1 } } },
	};

	map<string, int> distances = calculateDistances(simple_graph, "X");

	for (pair<string, int> elem : distances) {
		string key = elem.first;
		int distance = elem.second;

		cout << "key: " << key << ", distance: " << distance << std::endl;
	}
}
