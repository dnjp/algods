#include <iostream>
#include <map>
#include <vector>
#include <string>

using namespace std;

int fibMemo(int n, map<int, int> m)
{
	if (n < 2) {
		return n;
	}

	if (m.find(n) != m.end()) {
		return m[n];
	}
	int result = fibMemo(n - 1, m) + fibMemo(n - 2, m);
	m.insert(make_pair(n, result));
	return result;
}

int fibTab(int n)
{
	vector<int> f = { 0, 1 };
	while (f.size() <= n) {
		int s = f.size();
		f.push_back(f[s - 1] + f[s - 2]);
	}
	return f[n];
}

int fibonacci(int n)
{
	map<int, int> m;
	// return fibRecur(n, m);
	return fibTab(n);
}

/*
 * min cost to jump from 0 to i
 * 	base case is 0
 * 	build it up
 * 
 * 	cost = [0] * n
 * 	cost [0] = 0
 * 	cost [1] = abs(h[1] - h[0])
 * 
 * 	for i in range(2,n):
 * 		cost[i] = min(
 * 			abs(h[i] - h[i-1]) + cost[i-1],
 * 			abs(h[i] - h[i-1]) + cost[i-2],
 *		)
 * 
 * ------------------------------------------------------------
 * 
 * 
 * min cost to jump from i to the end
 * 
 * 	cost = [0] * n   # initializes array with 0 values
 * 	cost[n-1] = 0
 * 	cost[n-2] = abs(h[n-1] - h[n-2])
 * 	for (i=n-3; i >= 0; i--) {
 *		cost[i] = min(
 *	 		# 1 hop jump # 
 *	 		abs(h[i+1]-h[i]) + cost[i+1],
 *	 		abs(h[i+2]-h[i]) + cost[i+2],
 *		)
 * 	}
 * 
 * 	return cost[0] # tells cost of getting from 0 to the end
 * 
 * ------------------------------------------------------------
 * 
 * min cost from start 
 * 
 * 	def min_hop(n, heights, target):
 * 		if target == 0: return 0
 * 		if target == 1: return abs(heights[1] - h[0])
 * 
 * 		small_jump = min_hop(n, h, i-1) + abs(h[i] - h[i-1])
 * 		big_jump = min_hop(n, h, i-2) + abs(h[i] - h[i-2])
 * 
 * 		return min(small_jump, big_jump)
 * 
 */

int main()
{
	// cout << fibonacci(5) << endl;
	return 0;
}