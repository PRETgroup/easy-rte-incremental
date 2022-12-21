#include "save_results.h"

void save_setup(char run_reference[]) {
    FILE *fp;
    fp = fopen(SAVE_FILE_NAME, "a");
    if(fp == NULL) {
        printf("File can't be opened\n");
        return;
    }

    fprintf(fp, "%s", run_reference);
    fprintf(fp, "\n");
    fprintf(fp, "run, cpu_time_s, tick_time_ms\n");

    fclose(fp);
}

void save_time(uint16_t run, double cpu_time_s, double tick_time_ms) {
    FILE *fp;
    fp = fopen(SAVE_FILE_NAME, "a");
    if(fp == NULL) {
        printf("File can't be opened %f\n", cpu_time_s);
        return;
    }

    fprintf(fp, "%d, %f, %f\n", run, cpu_time_s, tick_time_ms);

    fclose(fp);
}
