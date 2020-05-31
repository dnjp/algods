/*
 * Problem:
 *   A pangram is a phrase which contains every letter at least once,
 *   such as "the quick brown fox jumps over the lazy dog". Write a
 *   function which determines if a given string is a pangram.
 *
 * The Unknown:
 *   Whether or not the string is a pangram
 *
 * The Data:
 *   The given string
 *
 * The Condition:
 *   The string is a pangram if every letter of the alphabet is present
 */

#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
#include <stdbool.h>
#include <ctype.h>
#include <string.h>

typedef struct {
	char *word;
	bool should_pass;
} test_case;

const char *ALPHABET = "abcdefghijklmnopqrstuvwxyz";
bool contains(char *word, char c);
bool is_alphabetic(char c);

/*
 * Solution A:
 *   Check for the existence of each letter of the alphabet in the phrase
 */
bool is_pangram_a(char *word)
{
	for (int i = 0; i < sizeof(ALPHABET); i++) {
		if (!contains(word, ALPHABET[i])) {
			return false;
		}
	}
	return true;
}

/*
 * Solution B:
 *   Remove duplicate and nonalphabetic characters from the phrase and test
 *   if the length of the new phrase is 26
 */
bool is_pangram_b(char *word)
{
	char alphabet[27] = "";
	for (int i = 0; i < strlen(word); i++) {
		char ch = tolower(word[i]);
		if (!is_alphabetic(ch)) {
			continue;
		}

		int alpha_index = ch - 97;
		bool should_insert = (int)alphabet[alpha_index] == 0;

		if (should_insert) {
			alphabet[alpha_index] = ch;
		}
	}
	if (strlen(alphabet) == 26) {
		return true;
	}
	return false;
}

/*
 * Solution C:
 *   Check for the existence of each letter in the word
 */
bool is_pangram_c(char *word)
{
	return false;
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

test_case *create_case(char *word, bool should_pass)
{
	test_case *tc = malloc(sizeof(test_case));
	tc->word = strdup(word);
	tc->should_pass = should_pass;
	return tc;
}

void free_case(test_case *tc)
{
	free(tc->word);
	free(tc);
}

void print_result(char *version, char *word, bool passes, bool should_pass)
{
	if (passes == should_pass) {
		printf("%s:\t%s => passed\n", version, word);
	} else {
		printf("%s:\t%s => failed\n", version, word);
	}
}

void run_tests()
{
	test_case *cases[10] = {

		/* true cases */
		create_case("Two driven jocks help fax my big quiz", true),
		create_case("Pack my box with five dozen liquor jugs", true),
		create_case("The five boxing wizards jump quickly", true),
		create_case("Bright vixens jump; dozy fowl quack", true),
		create_case("Jackdaws love my big sphinx of quartz", true),
		create_case("J.Q. Schwartz flung V.D. Pike my box", true),

		/* false cases */
		create_case("hello", false),
		create_case("You can only be afraid of what you think you know",
			    false),
		create_case("C is quirky, flawed, and an enormous success",
			    false),
		create_case("When in doubt, use brute force.", false),
	};

	bool a_failed = false;
	bool b_failed = false;
	bool c_failed = false;

	int size = sizeof(cases) / sizeof(cases[0]);
	for (int i = 0; i < size; i++) {
		test_case *tc = cases[i];
		char *word = tc->word;
		bool should_pass = tc->should_pass;

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

		print_result("a", word, passes_a, should_pass);
		print_result("b", word, passes_b, should_pass);
		print_result("c", word, passes_c, should_pass);

		free_case(cases[i]);

		printf("----------------------\n");
	}

	printf("---------------------\n");
	printf("     RESULTS\n");
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
