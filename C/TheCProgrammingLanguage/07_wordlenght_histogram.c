#include <stdio.h>

#define IN 1
#define OUT 0

int main() {

    int c, state, charCount, i;
    state = OUT;
    charCount = i = 0;
    int ndigit[10];

    for (i = 0; i < 10; ++i) {
        ndigit[i] = 0;
    }

    while ((c = getchar()) != EOF) {
        if (c == ' ' || c == '\n' || c == '\t') {
            state = OUT;
            if (charCount > 0) {
                ++ndigit[charCount];
                charCount = 0;
            }
        } else if (state == OUT) {
            state = IN;
        }
        if (state == IN) {
            ++charCount;
        }
    }

    //Well, histogram-sorta, I'm not too interested in making the display pretty
    for (i = 0; i < 10; ++i) {
        printf("Words with length %d: %d\n", i, ndigit[i]);
    }
}
