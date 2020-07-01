#include <vector>
#include <assert.h>
#include <iostream>
#include <map>
#include <queue>
#include <utility>

using namespace std;

class Solution {
    public:
	Solution()
	{
	}

	int networkDelayTime(vector<vector<int> > &times, int N, int K)
	{
		/*
		 * Problem:
		 * 	How long will it take for all nodes to receive the signal
		 * 	sent from K? Return -1 if it is impossible.
		 *
		 * Data:
		 * 	times: list of travel times as directed edges
		 * 		times[i] = (u, v, w)
		 * 		u: source node
		 * 		v: target node
		 * 		w: time it takes for a signal to travel from u to v
		 * 	N: network nodes 1-10
		 * 	K: Starting Node
		 */

		map<int, vector<pair<int, int> > > edges;
		for (auto edge : times) {
			int source = edge[0];
			int target = edge[1];
			int weight = edge[2];
			edges[source].push_back(make_pair(target, weight));
		}

		int delay = 0;
		int node = K;
		priority_queue<pair<int, int> > frontier;
		frontier.push(make_pair(delay, node));
		map<int, int> visited;

		while (!frontier.empty()) {
			pair<int, int> item = frontier.top();
			frontier.pop();
			int delay = item.first;
			int node = item.second;

			if (visited.find(node) != visited.end()) {
				// we've already visited this node
				continue;
			}

			// add node to visited map
			visited[node] = delay;

			for (pair<int, int> n : edges[node]) {
				int neighbor = n.first;
				int edgeWeight = n.second;
				if (visited.find(neighbor) != visited.end()) {
					// we've already visited this node
					continue;
				}
				frontier.push(make_pair(delay + edgeWeight, neighbor));
			}
		}

		if (visited.size() != N) {
			return -1;
		}

		int totalDelayTime = 0;
		for (pair<int, int> item : visited) {
			int node = item.first;
			int delay = item.second;
			if (delay > totalDelayTime) {
				totalDelayTime = delay;
			}
		}

		return totalDelayTime;
	}
};

int main()
{
	Solution s;
	vector<vector<int> > times = {
		{ 2, 1, 1 },
		{ 2, 3, 1 },
		{ 3, 4, 1 },
	};
	int N = 4;
	int K = 2;

	int delayTime = s.networkDelayTime(times, N, K);
	cout << "Delay Time: " << delayTime << ", Expected: 2" << endl;
	assert(delayTime == 2);

	times = {
		{ 1, 2, 1 },
		{ 2, 3, 2 },
		{ 1, 3, 4 },
	};

	N = 3;
	K = 1;

	delayTime = s.networkDelayTime(times, N, K);
	cout << "Delay Time: " << delayTime << ", Expected: 3" << endl;
	assert(delayTime == 3);
}
