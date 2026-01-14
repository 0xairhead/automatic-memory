#include "parser.h"
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

// Format: [4 bytes MAGIC] [1 byte Width] [1 byte Height] [Data...]
// Magic must be "IMG!"

int parse_image_header(const uint8_t *data, size_t size) {
    if (size < 6) return -1;

    // Check Magic
    if (data[0] == 'I' && data[1] == 'M' && data[2] == 'G' && data[3] == '!') {
        uint8_t width = data[4];
        uint8_t height = data[5];

        // Bug: If width=0, we allocate 0 bytes but might write to it?
        // Actually lets do a classic heap overflow.
        
        // We expect 'width * height' bytes of data to follow.
        int expected_size = width * height;
        
        // VULNERABILITY: Integer overflow if width*height > 255?
        // No, let's make it simpler for this demo.
        
        if (width == 100 && height == 100) {
             // Simulate a "special path" that processes data
             // Let's assume we copy the rest of the data into a buffer
             // allocated based on width/height, but we mess up the math.
             
             char *buffer = (char*)malloc(width * height);
             
             // If the input provides more data than we expect, we overflow
             if (size - 6 > (size_t)(width * height)) {
                 // Classic overflow: memcpy(dest, src, size_from_input)
                 memcpy(buffer, data + 6, size - 6);
                 
                 // This will crash with AddressSanitizer!
                 free(buffer);
                 return 1;
             }
             free(buffer);
        }
    }
    return 0;
}
