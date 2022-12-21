#include <stdio.h>
#include <stdint.h>

#define SAVE_FILE_NAME "cpu_time_results.csv"
#define RUNS 10
#define TICKS_PER_RUN 1000

void save_setup(char run_reference[]);
void save_time(uint16_t run, double cpu_time_s, double tick_time_ms);
