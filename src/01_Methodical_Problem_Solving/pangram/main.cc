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
#include <vector>

bool contains(char *phrase, char c);
bool is_alphabetic(char c);

/*
 * Solution A:
 *   The simplest solution: check for the existence of each letter of the alphabet in the phrase
 */
bool is_pangram_a(char *phrase)
{
	const char *alphabet = "abcdefghijklmnopqrstuvwxyz";
	for (int i = 0; i < sizeof(alphabet); i++) {
		if (!contains(phrase, alphabet[i])) {
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
bool is_pangram_b(char *phrase)
{
	char distinct_alpha[27] = "";
	for (int i = 0; i < strlen(phrase); i++) {
		if (strlen(distinct_alpha) == 26) {
			return true;
		}

		char ch = tolower(phrase[i]);
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
bool is_pangram_c(char *phrase)
{
	std::unordered_set<std::string> alpha_set({ "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
						    "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z" });

	for (int i = 0; i < strlen(phrase); i++) {
		char ch = tolower(phrase[i]);
		if (is_alphabetic(ch)) {
			std::string ch_str(1, ch);
			alpha_set.erase(ch_str);
		}
	}
	return alpha_set.empty();
}

bool contains(char *phrase, char c)
{
	for (int j = 0; j < strlen(phrase); j++) {
		if (tolower(phrase[j]) == tolower(c)) {
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
	TestCase(const char *input, bool sp)
	{
		phrase = strdup(input);
		should_pass = sp;
	}
	char *phrase;
	bool should_pass;
};

class Test {
    public:
	Test(std::vector<TestCase> c)
	{
		cases = c;
	}
	void Run();

    private:
	void PrintResult(const char *ver, char *phrase, bool passes, bool should_pass);

	std::vector<TestCase> cases;
};

void Test::PrintResult(const char *ver, char *phrase, bool passes, bool should_pass)
{
	if (passes == should_pass) {
		printf("%s:\t%s => passed\n", ver, phrase);
	} else {
		printf("%s:\t%s => failed\n", ver, phrase);
	}
}

void Test::Run()
{
	bool a_failed = false;
	bool b_failed = false;
	bool c_failed = false;

	for (TestCase tc : cases) {
		char *phrase = tc.phrase;
		bool should_pass = tc.should_pass;

		bool passes_a = is_pangram_a(phrase);
		bool passes_b = is_pangram_b(phrase);
		bool passes_c = is_pangram_c(phrase);

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
		this->PrintResult("a", phrase, passes_a, should_pass);
		this->PrintResult("b", phrase, passes_b, should_pass);
		this->PrintResult("c", phrase, passes_c, should_pass);
	}

	printf("---------------------\n");
	printf("	  RESULTS\n");
	printf("---------------------\n");
	printf(" Solution A: %s\n", a_failed == true ? "failed" : "passed");
	printf(" Solution B: %s\n", b_failed == true ? "failed" : "passed");
	printf(" Solution C: %s\n", c_failed == true ? "failed" : "passed");
	printf("---------------------\n");
}

int main(void)
{
	Test t(std::vector<TestCase>({
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
	}));
	t.Run();
	return 0;
}
