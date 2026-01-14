#include <stdint.h>
#include <stddef.h>
#include "parser.h"

// This function is called by libFuzzer repeatedly
extern "C" int LLVMFuzzerTestOneInput(const uint8_t *Data, size_t Size) {
    // Pass the random data directly to our library function
    parse_image_header(Data, Size);
    return 0;
}
