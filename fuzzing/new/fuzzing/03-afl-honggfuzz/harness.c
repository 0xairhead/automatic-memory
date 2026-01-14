#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>

// This is a typical target function
void target_function(char *buf, int len) {
    if (len > 0 && buf[0] == 'A') {
        if (len > 1 && buf[1] == 'F') {
            if (len > 2 && buf[2] == 'L') {
                if (len > 3 && buf[3] == '!') {
                    // Artificial crash
                    abort();
                }
            }
        }
    }
}

// __AFL_FUZZ_TESTCASE_LEN is an optimization for persistent mode, 
// but for simple cases we can just read from stdin.
int main() {
    char buf[1024];
    // Read from stdin (AFL default)
    ssize_t len = read(STDIN_FILENO, buf, sizeof(buf));
    
    if (len > 0) {
        target_function(buf, len);
    }
    return 0;
}
