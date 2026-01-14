#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void bug_one() {
    printf("Triggering BUG ONE (Null Pointer Dereference)\n");
    int *p = NULL;
    *p = 1;
}

void bug_two() {
    printf("Triggering BUG TWO (Buffer Overflow)\n");
    char buf[10];
    memset(buf, 'A', 100);
}

void bug_three() {
    printf("Triggering BUG THREE (Divide by Zero)\n");
    int a = 1;
    int b = 0;
    int c = a / b;
}

int main() {
    char buffer[100];
    if (fgets(buffer, sizeof(buffer), stdin)) {
        if (strncmp(buffer, "BUG1", 4) == 0) {
            bug_one();
        } else if (strncmp(buffer, "BUG2", 4) == 0) {
            bug_two();
        } else if (strncmp(buffer, "BUG3", 4) == 0) {
            bug_three();
        } else {
            printf("Normal execution.\n");
        }
    }
    return 0;
}
