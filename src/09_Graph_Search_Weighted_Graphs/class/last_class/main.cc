#include <iostream>
#include <vector>
#include <set>
#include <utility>
#include <queue>

using namespace std;

/*
 * Jug 1 = 5 gallons
 * Jug 2 = 3 gallons
 *
 * Key idea: represent this as a matrix where the x axis
 * 	represents jug 1 and the y axis represents jug 2.
 *
 *
 *
 *
 *
 *
 *
 *
 */

/*
 * Suppose we're in the state (0, 3)
 * e.g. jug1 has 0 gallons
 * 	jug2 has 3 gallons
 */
vector<pair<int, int> > getNeighbors(int cap1, int cap2, int cur1, int cur2)
{
	// fill a jug
	vector<pair<int, int> > neighbors;
	neighbors.push_back(make_pair(cap1, cur2));
	neighbors.push_back(make_pair(cur1, cap2));

	// empty a jug
	neighbors.push_back(make_pair(0, cur2));
	neighbors.push_back(make_pair(cur1, 0));

	/*
	 * Transfer one jug to the other
	 *
	 * Limited by 2 things:
	 * - Capacity of 1
	 * - Amount in 2
	 */

	// Pour from jug2 to jug1
	int amountPoured = min(cap1 - cur1, cur2);
	neighbors.append(make_pair(cur1 + amountPoured, cur2 - amountPoured));

	// Pour from jug1 to jug2
	int amountPoured = min(cur1, cap2 - cur2);
	neighbors.append(make_pair(cur1 - amountPoured, cur2 + amountPoured));
}

// What's the shortest distance to the target state? -1 if it's not possible
int dieHard(int cap1, int cap2, int target)
{
	set<pair(int, int)> visited;
	map<int, pair<int, int> > parent;
	queue<pair<int, pair<int, int> > > q; // TODO: need to fix this

	pair<int, int> start = { 0, 0 };
	q.insert(make_pair(0, start));
	visited.push(make_pair(0, 0));

	while (!q.empty()) {
		// note: replacing queue with a stack will change this from bfs to dfs
		pair<int, pair<int, int> > item = q.front(); // TODO: make sure q.front() is correct here
		q.pop();
		int d = item.first;
		int item2 = item.second;
		int cur1 = item2.first;
		int cur2 = item2.second;

		if (cur1 == target || cur2 == target) {
			vector<pair<int, int> > path;
			pair<int, int> node = make_pair(cur1, cur2);
			while (node != nullptr) {
				path.push_back(node);
				node = parent[node]
			}

			return true;
		}

		for (auto neighbor : getNeighbors(cap1, cap2, cur1, cur2)) {
			if (visited.find(neighbor) == visited.end()) {
				q.insert(make_pair(d + 1, neighbor));
				visited.push(neighbor);
				parent[neighbors] = make_pair(cur1, cur2);
			}
		}
	}

	return false;
}

int main()
{
}
