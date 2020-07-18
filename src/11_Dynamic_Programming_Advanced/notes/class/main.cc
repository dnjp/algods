#include <iostream>
#include <vector>
#include <climits>

using namespace std;

/*
 *  3 states:
 * 	- has stock
 * 	- no stock, cannot buy
 * 	- no stock, can buy
 * 
 *  Technique:
 * 	- Keep separate memo tables for each of these states
 * 
 *  Notes:
 * 	- Most difficult part is identifying what things _should_ mean
 * 	- Consider what state you end up in at the end of the day
 * 	- Consider how much money you end up with at the end of the day
 */
int maxProfit(vector<int> &prices)
{
	int buy = -1;
	int total = 0;

	// represents state AFTER decision has been made


	/*
	 * ith elem is the most money you can have if you 
	 * end the day in state hasStock
	 */
	vector<int> hasStock(prices.size(), 0);

	/*
	 * ith elem is the most money you can have if you 
	 * end the day in state noStockNoBuy
	 */
	vector<int> noStockNoBuy(prices.size(), 0);
	/*
	 * ith elem is the most money you can have if you 
	 * end the day in state noStockCanBuy
	 */
	vector<int> noStockCanBuy(prices.size(), 0); 

	/*
	 * noStockCanBuy:
	 * 	         0, 
	 * 
	 * hasStock:
	 * 	-prices[i], 
	 * 
	 * noStockNoBuy:
	 * 	      -Inf, 
	 * 
	 * answer at end:
	 *  	max(hasStock[-1], noStockCanBuy[-1], noStockNoBuy[-1])
	 */

	noStockCanBuy[0] = 0;
	hasStock[0] = 0-prices[0];
	noStockNoBuy[0] = INT_MIN;

	for (int i = 1; i < prices.size(); i++) {

		// no stock can buy
		// started from hasStock on day i-1
		int didNothing = hasStock[i-1];
		// started from noStockCanBuy on day i-1
		int boughtStock = noStockCanBuy[i-1];
		hasStock[i] = max(didNothing, boughtStock);

		// no stock can buy
		int cooled = noStockNoBuy[i-1];
		int waited = noStockCanBuy[i-1];
		noStockCanBuy[i] = max(cooled, waited);

		// no stock no buy
		noStockNoBuy[i] = hasStock[i-1] + prices[i];
	}
	int last = prices.size() - 1;
	return max(noStockCanBuy[last], noStockNoBuy[last]);
}

int main()
{
	cout << "hello world" << std::endl;
	return 0;
}