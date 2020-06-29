#include <iostream>
#include <map>
#include <utility>
#include <set>
#include <vector>
#include <algorithm>
#include <iterator>

using namespace std;

class Graph {
    public:
	Graph();
	void addEdge(pair<int, int>, pair<int, int>);
	void build(int);

    private:
	map<pair<int, int>, set<pair<int, int> > > g;
	set<pair<int, int> > moveOffsets;

	vector<pair<int, int> > legalMovesFrom(int, int, int);
	void traverse(vector<pair<int, int> >, pair<int, int>, int);
	set<pair<int, int> > convertToSet(vector<pair<int, int> >);
	vector<pair<int, int> > convertToVector(set<pair<int, int> >);
};

Graph::Graph()
{
	moveOffsets = {
		{ -1, -2 }, { 1, -2 }, { -2, -1 }, { 2, -1 }, { -2, 1 }, { 2, 1 }, { -1, 2 }, { 1, 2 },
	};
}

void Graph::addEdge(pair<int, int> vertex_a, pair<int, int> vertex_b)
{
	g[vertex_a].insert(vertex_b);
	g[vertex_b].insert(vertex_a);
}

void Graph::build(int board_size)
{
	for (int row = 0; row < board_size; row++) {
		for (int col = 0; col < board_size; col++) {
			vector<pair<int, int> > moves = legalMovesFrom(row, col, board_size);
			for (pair<int, int> p : moves) {
				int to_row = p.first;
				int to_col = p.second;
				pair<int, int> a1 = { row, col };
				pair<int, int> a2 = { to_row, to_col };
				addEdge(a1, a2);
			}
		}
	}
}

vector<pair<int, int> > Graph::legalMovesFrom(int row, int col, int board_size)
{
	vector<pair<int, int> > moves;
	for (auto p : moveOffsets) {
		int row_offset = p.first;
		int col_offset = p.second;

		int move_row = row + row_offset;
		int move_col = col + col_offset;
		if ((0 <= move_row && move_row < board_size) && (0 <= move_col && move_col < board_size)) {
			pair<int, int> move = { move_row, move_col };
			moves.push_back(move);
		}
	}
	return moves;
}

vector<pair<int, int> > *Graph::traverse(vector<pair<int, int> > path, pair<int, int> current_vertex, int total_squares)
{
	if ((path.size() + 1) == total_squares) {
		return true;
		vector<pair<int, int> > new_path = path;
		new_path.push_back(current_vertex);
		return &new_path;
	}
	set<pair<int, int> > graph_set = g[current_vertex];
	set<pair<int, int> > pair_set = convertToSet(path);
	set<pair<int, int> > yet_to_visit;
	set_difference(graph_set.begin(), graph_set.end(), pair_set.begin(), pair_set.end(),
		       inserter(yet_to_visit, yet_to_visit.end()));

	if (yet_to_visit.size() == 0) {
		// no unvisited neighbors, so dead end
		return null_ptr;
	}

	vector<pair<int, int> > next_vertices = convertToVector(yet_to_visit);
	sort(next_vertices.begin(), next_vertices.end(), [](pair<int, int> a, pair<int, int> b) {
		// TODO: fill this in
		// use warnsdorffs heuristic:
		return g[a].size() - graph[b].size();
	});
}

set<pair<int, int> > Graph::convertToSet(vector<pair<int, int> > input)
{
	set<pair<int, int> > s(input.begin(), input.end());
	return s;
}

vector<pair<int, int> > Graph::convertToVector(set<pair<int, int> > input)
{
	vector<pair<int, int> > v(input.begin(), input.end());
	return v;
}

int main()
{
	Graph gf;
	pair<int, int> vertex_a(0, 1);
	pair<int, int> vertex_b(1, 2);
	gf.addEdge(vertex_a, vertex_b);

	return 0;
}
