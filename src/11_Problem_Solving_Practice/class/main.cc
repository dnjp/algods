#include <iostream>
#include <vector>
#include <map>

using namespace std;

/*
 * Triangle Problem:
 * 
 * 	Given a triangle. In each node in the triangle there is a number.
 * 	Goal: Find a path from top of the triangle to the bottom that maximizes sum of values
 * 
 * 	Input: [[1], [5, 2], [3, 3, 10], [1, 1, 6, 4]]
 * 	Expected: 1 -> 2 -> 10 -> 6
 * 	Output: 19
 * 
 * 
 * 	f(i,j)
 * 		max sum  from (i,j) to the bottom
 * 		max(f(i+1, j), f(i+1, j+1))
 * 
 */

int triangleMax(vector<vector<int> > triangle)
{
	// could use the 2d triangle array and modify in place
	// stores max of level, index
	// map<pair<int,int>> cache;
	int numLevels = triangle.size();

	int total = 0;

	for (int i = numLevels; i > 0; i--) {
		int maxNum = 0;
		for (int j = 0; j < triangle[i].size(); j++) {
			if (triangle[i][j] > maxNum) {
				maxNum = triangle[i][j];
			}
		}
		total += maxNum;
	}

	return total;
}

/*
 * Duplicate within d
 * 
 * Given an array of numbers and a distance (d), return boolean 
 * representing whether there are two matching integers within 
 * distance d
 * 
 *     [2, 3, 6, 3, 2]
 *     d = 1 // FALSE
 *     d = 2 // tRUE
 */

// TODO watch out for array overflow
bool duplicateInD(vector<int> nums, int distance)
{
	// could use map to track things? maybe a set?
	// value -> index

	// numMap
	for (int i = 0; i < nums.size(); i++) {
		/*
		 * if nums[i] in numMap and i[k]
		 */
		for (int j = i + 1; j < i + distance; j++) {
			if (nums[i] == nums[j]) {
				return true;
			}
		}
	}

	return false;
}

/*
 * We have 
 */

int main()
{
	// vector<vector<int> > tri = { { 1 }, { 5, 2 }, { 3, 3, 10 }, { 1, 1, 6, 4 } };
	// cout << "triangle sum: " << triangleMax(tri) << endl;

	// vector<int> nums = { 2, 3, 3, 7, 2 };
	// cout << "duplicate in d: " << duplicateInD(nums, 2);
	return 0;
}