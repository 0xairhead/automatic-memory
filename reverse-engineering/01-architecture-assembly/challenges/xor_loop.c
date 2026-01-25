#include <stdio.h>

int main() {
    int data[] = {10, 20, 30, 40, 50};
    int key = 0x55;
    int length = 5;

    printf("Original data: ");
    for (int i = 0; i < length; i++) {
        printf("%d ", data[i]);
    }
    printf("\n");

    // The loop to reverse engineer
    for (int i = 0; i < length; i++) {
        data[i] = data[i] ^ key;
    }

    printf("XORed data: ");
    for (int i = 0; i < length; i++) {
        printf("%d ", data[i]);
    }
    printf("\n");

    return 0;
}
