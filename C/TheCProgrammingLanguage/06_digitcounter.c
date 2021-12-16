#include <stdio.h>

int main() {
    int c, i, nwhite, nother;
    int ndigit[10];
    nwhite = nother = 0;

    for (i = 0; i < 10; ++i) {
        ndigit[i] = 0;
    }

    while ((c = getchar()) != EOF) {
        if (c >= '0' && c <= '9') {
            ++ndigit[c-'0']; //There's some ASCII integer subtraction going on here that only works for this example. Really obscure!
        } else if (c == ' ' || c == '\n' || c == '\t') {
            ++nwhite;
        } else {
            ++nother;
        }
    }

    for (i = 0; i < 10; ++i) {
        printf(" %d", ndigit[i]);
    }
}