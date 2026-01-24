#include <stdio.h>
#include <string.h>
int main(int argc, char *argv[]) {
    if (argc < 2) return 1;
    if (strcmp(argv[1], "Flamingo!23") == 0) {
        printf("Access Granted\n");
    } else {
        printf("Access Denied\n");
    }
    return 0;
}

