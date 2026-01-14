#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int main() {
    char buffer[10];
    printf("Enter input: ");
    if (fgets(buffer, sizeof(buffer), stdin)) {
        if (strcmp(buffer, "crash\n") == 0) {
            printf("Boom!\n");
            abort(); // Simulate a crash
        } else {
            printf("Safe.\n");
        }
    }
    return 0;
}
