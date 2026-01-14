#include <stdio.h>
#include <string.h>
#include <stdlib.h>

// A fake XML parser that enforces strict structure
// <tag key="val">data</tag>

void proces_tag(char *key, char *val, char *data) {
    // VULNERABILITY: Stack buffer overflow
    // If key is "BOMB" and val is large, we overflow
    if (strcmp(key, "BOMB") == 0) {
        char buffer[10];
        // strcpy is dangerous!
        strcpy(buffer, val);
        printf("Processed BOMB tag with val: %s\n", buffer);
    } else {
        printf("Processed tag: key=%s val=%s data=%s\n", key, val, data);
    }
}

int main() {
    char input[512];
    printf("Enter XML: ");
    if (fgets(input, sizeof(input), stdin)) {
        // Strip newline
        input[strcspn(input, "\n")] = 0;

        // 1. Must start with '<'
        if (input[0] != '<') {
            printf("Error: Must start with '<'\n");
            return 1;
        }

        // 2. Parse <tag key="val">data</tag>
        // Ideally we'd use a real parser, but sscanf is easier for demo
        char key[100];
        char val[100];
        char data[100];
        char closing_tag[100];

        // This format string is very specific. 
        // It expects: <tag key="VALUE">DATA</tag>
        // Mutation fuzzers have a hard time preserving this exact structure.
        int parsed = sscanf(input, "<tag key=\"%[^\"]\">%[^<]</%[^>]>", key, val, data, closing_tag);

        if (parsed == 4) {
             if (strcmp(closing_tag, "tag") == 0) {
                 proces_tag(key, val, data);
             } else {
                 printf("Error: Mismatched closing tag\n");
             }
        } else {
            printf("Error: Invalid XML format\n");
        }
    }
    return 0;
}
