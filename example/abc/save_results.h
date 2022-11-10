#include <stdio.h>
#include <stdint.h>

#define SAVE_FILE_NAME "cpu_time_results.csv"

void save_setup(char run_reference[]);
void save_time(uint16_t run, double cpu_time);
