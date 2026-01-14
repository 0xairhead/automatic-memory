#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int main() {
    char buffer[128];
    printf("Enter input: ");
    if (fgets(buffer, sizeof(buffer), stdin)) {
        // Level 1: Magic keyword
        if (buffer[0] == 'M') {
            printf("Level 1 passed\n");
            // Level 2
            if (buffer[1] == 'a') {
                printf("Level 2 passed\n");
                // Level 3
                if (buffer[2] == 'z') {
                    printf("Level 3 passed\n");
                    // Level 4
                    if (buffer[3] == 'e') {
                        printf("Level 4 passed\n");
                         // Level 5
                         if (buffer[4] == '!') {
                            printf("CRASH! You solved the maze!\n");
                            abort();
                         }
                    }
                }
            }
        }
    }
    return 0;
}
