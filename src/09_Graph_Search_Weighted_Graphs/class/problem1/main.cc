#include <iostream>
#include <limits>
#include <string>

string START = "s";
string S = "s";
string GOAL = "x";
string G = "x";
NORMAL = N = " " MOUNTAIN = M = "^" RIVER = R = "-"

	const map<string, double>
		COSTS = {
			{ "START", 0.0 },  { "GOAL", 1.0 },
			{ "NORMAL", 1.0 }, { "MOUNTAIN", numeric_limits<double>::infinity() },
			{ "RIVER", 2.0 },
		};

int main()
{
}
