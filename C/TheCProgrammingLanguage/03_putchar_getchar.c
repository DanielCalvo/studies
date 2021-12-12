//
// Created by daniel on 12/12/21.
//

#include <stdio.h>

int main() {
    int c, n1;

//    c = getchar();
//    while (c != EOF) {
//        putchar(c);
//        c = getchar();
//    }

    n1 = 0;
    while ((c = getchar()) != EOF) {
        if (c == '\n') {
           ++n1;
        }
        printf("%d\n",n1);
    }

}
