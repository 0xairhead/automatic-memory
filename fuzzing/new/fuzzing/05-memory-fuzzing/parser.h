#ifndef PARSER_H
#define PARSER_H

#include <stdint.h>
#include <stddef.h>

// Imitates a library function that parses some binary format
int parse_image_header(const uint8_t *data, size_t size);

#endif
