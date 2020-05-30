#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
#include <stdbool.h>
#include <string.h>

/*
 * A pangram is a phrase which contains every letter at least once, such as "the quick brown fox jumps over the lazy dog". Write a function which determines if a given string is a pangram.
 */

typedef struct {
	char *word;
	bool should_pass;
} test_case;

bool is_pangram_a(char *word)
{
	return false;
}

test_case *create_case(char *word, bool should_pass)
{
	test_case *case = malloc(sizeof(test_case));
	case->word = strdup(word);
	case->should_pass = should_pass;
	return case;
}

void run_tests()
{
	test_case *cases[10] = {
		create_case("Two driven jocks help fax my big quiz", true),
		create_case("Pack my box with five dozen liquor jugs", true),
		create_case("The five boxing wizards jump quickly", true),
		create_case("Bright vixens jump; dozy fowl quack", true),
		create_case("New job: fix Mr Gluck's hazy TV, PD", true),
		create_case("J.Q. Schwartz flung V.D. Pike my box", true),
		create_case("hello", false),
		create_case("You can only be afraid of what you think you know",
			    false),
		create_case("C is quirky, flawed, and an enormous success",
			    false),
		create_case("When in doubt, use brute force.", false),
	};

	int size = sizeof(cases) / sizeof(cases[0]);
	for (int i = 0; i < size; i++) {
		test_case *case = cases[i];
		char *word = case->word;
		bool should_pass = case->should_pass;
		bool passes = is_pangram_a(word);
		printf("phrase: %s\n", word);
		printf("is pangram: %d, expecting: %d\n", passes,
		       case->should_pass);

		assert(passes == should_pass);
	}
}

int main(void)
{
	run_tests();
	return 0;
}
