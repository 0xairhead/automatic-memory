#include <stdio.h>

// A simple function to demonstrate cdecl/stdcall
// This specific function will likely be cdecl by default on Linux
int add_numbers(int a, int b) {
    int sum = a + b;
    return sum;
}

int main() {
    int x = 10;
    int y = 20;
    int result = 0;

    printf("Starting the addition...\n");
    
    // Call the function
    result = add_numbers(x, y);

    printf("The result is: %d\n", result);
    
    return 0;
}
