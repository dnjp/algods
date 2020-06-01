/*
 * Problem:
 *   A pangram is a phrase which contains every letter at least once,
 *   such as "the quick brown fox jumps over the lazy dog". Write a
 *   function which determines if a given string is a pangram.
 *
 * The Unknown:
 *   Whether or not the string contains all letters of the alphabet
 *
 * The Data:
 *   The given string
 *
 * The Condition:
 *   The string is a pangram if every letter of the alphabet is present
 */

#include <ctype.h>
#include <stdlib.h>
#include <string.h>
#include <unordered_set>

bool contains(char *word, char c);
bool is_alphabetic(char c);

/*
 * Solution A:
 *   Check for the existence of each letter of the alphabet in the phrase
 */
bool is_pangram_a(char *word)
{
	const char *alphabet = "abcdefghijklmnopqrstuvwxyz";
	for (int i = 0; i < sizeof(alphabet); i++) {
		if (!contains(word, alphabet[i])) {
			return false;
		}
	}
	return true;
}

/*
 * Solution B:
 *   Create new string with distinct alphabetical characters from input string
 *   and test if the length of the new string is 26
 */
bool is_pangram_b(char *word)
{
	char distinct_alpha[27] = "";
	for (int i = 0; i < strlen(word); i++) {
		if (strlen(distinct_alpha) == 26) {
			return true;
		}

		char ch = tolower(word[i]);
		if (!is_alphabetic(ch)) {
			continue;
		}

		int alpha_index = ch - 97;
		bool should_insert = (int)distinct_alpha[alpha_index] == 0;

		if (should_insert) {
			distinct_alpha[alpha_index] = ch;
		}
	}
	if (strlen(distinct_alpha) == 26) {
		return true;
	}
	return false;
}

/*
 * Solution C:
 *   Create a set containing the alphabet and remove each letter in the set when
 *   encountered in the given string. If the set is empty at the end, the string
 *   is a pangram.
 */
bool is_pangram_c(char *word)
{
	std::unordered_set<std::string> alpha_set({ "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
						    "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z" });

	for (int i = 0; i < strlen(word); i++) {
		char ch = tolower(word[i]);
		if (is_alphabetic(ch)) {
			std::string ch_str(1, ch);
			alpha_set.erase(ch_str);
		}
	}
	return alpha_set.empty();
}

bool contains(char *word, char c)
{
	for (int j = 0; j < strlen(word); j++) {
		if (tolower(word[j]) == tolower(c)) {
			return true;
		}
	}
	return false;
}

bool is_alphabetic(char c)
{
	if (c >= 'a' && c <= 'z') {
		return true;
	}
	return false;
}

class TestCase {
    public:
	TestCase(const char *w, bool sp)
	{
		// const char *input = w.c_str();
		word = strdup(w);
		should_pass = sp;
	}
	char *word;
	bool should_pass;
};

void print_result(const char *ver, char *word, bool passes, bool should_pass)
{
	if (passes == should_pass) {
		printf("%s:\t%s => passed\n", ver, word);
	} else {
		printf("%s:\t%s => failed\n", ver, word);
	}
}

void run_tests()
{
	TestCase cases[11] = {

		/* true cases */
		TestCase("Two driven jocks help fax my big quiz", true),
		TestCase("Pack my box with five dozen liquor jugs", true),
		TestCase("The five boxing wizards jump quickly", true),
		TestCase("Bright vixens jump; dozy fowl quack", true),
		TestCase("Jackdaws love my big sphinx of quartz", true),
		TestCase("J.Q. Schwartz flung V.D. Pike my box", true),

		/* false cases */
		TestCase("hello", false),
		TestCase("You can only be afraid of what you think you know", false),
		TestCase("C is quirky, flawed, and an enormous success", false),
		TestCase("When in doubt, use brute force.", false),
		TestCase("", false),
	};

	bool a_failed = false;
	bool b_failed = false;
	bool c_failed = false;

	int size = sizeof(cases) / sizeof(cases[0]);
	for (int i = 0; i < size; i++) {
		TestCase tc = cases[i];
		char *word = tc.word;
		bool should_pass = tc.should_pass;

		bool passes_a = is_pangram_a(word);
		bool passes_b = is_pangram_b(word);
		bool passes_c = is_pangram_c(word);

		if (passes_a != should_pass) {
			a_failed = true;
		}
		if (passes_b != should_pass) {
			b_failed = true;
		}
		if (passes_c != should_pass) {
			c_failed = true;
		}

		printf("----------------------\n");
		print_result("a", word, passes_a, should_pass);
		print_result("b", word, passes_b, should_pass);
		print_result("c", word, passes_c, should_pass);
	}

	printf("---------------------\n");
	printf("      RESULTS\n");
	printf("---------------------\n");
	printf(" Solution A: %s\n", a_failed == true ? "failed" : "passed");
	printf(" Solution B: %s\n", b_failed == true ? "failed" : "passed");
	printf(" Solution C: %s\n", c_failed == true ? "failed" : "passed");
	printf("---------------------\n");
}

int main(void)
{
	run_tests();
	return 0;
}
