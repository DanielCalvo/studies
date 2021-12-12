#include <stdio.h>

int main() {
    float fahr, celsius;
    int lower, upper, step;
    lower = 0;
    upper = 300;
    step = 20;

    fahr = lower;

    while (fahr <= upper) {
        celsius = 5 * (fahr-32) / 9;
        //Print the first number of each line in a field 3 digits wide, with 0 and 1 decimal digits respectively
        printf("%3.0f %6.1f\n", fahr, celsius); //Oof, "%df" is not very readable
        fahr = fahr + step;
    }


}