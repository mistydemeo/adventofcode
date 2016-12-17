#include <inttypes.h>
#include "stdio.h"
#include "stdlib.h"
#include "string.h"

#include <CommonCrypto/CommonDigest.h>

const char* KEY = "bgvyzdsv";

int guess(int i, int prefix_size)
{
    char test_string[100];
    unsigned char result[CC_MD5_DIGEST_LENGTH];

    snprintf(test_string, sizeof(test_string), "%s%i", KEY, i);
    CC_MD5(test_string, strnlen(test_string, 100), result);
    char hex[7];
    snprintf(hex, sizeof(hex), "%02x%02x%02x", result[0], result[1], result[2]);

    return (strncmp("000000", hex, prefix_size) == 0);
}

int get_prefix_size(char *s)
{
    int prefix_size;
    char* endptr;
    prefix_size = (int)strtol(s, &endptr, 10);
    if (*endptr != 0) {
        fprintf(stderr, "Argument was not an integer: %s\n", s);
        exit(1);
    } else {
        return prefix_size;
    }
}

int main(int argc, char *argv[])
{
    int prefix_size;
    if (argc < 2) {
        prefix_size = 5; // default
    } else {
        prefix_size = get_prefix_size(argv[1]);
    }

    int i = 0;
    while(1) {
        if (guess(i, prefix_size)) {
            printf("%i\n", i);
            exit(0);
        } else {
            i++;
        }
    }
}
