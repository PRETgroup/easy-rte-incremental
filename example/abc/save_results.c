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
    fprintf(fp, "run, cpu_time\n");

    fclose(fp);
}

void save_time(uint16_t run, double cpu_time) {
    FILE *fp;
    fp = fopen(SAVE_FILE_NAME, "a");
    if(fp == NULL) {
        printf("File can't be opened %f\n", cpu_time);
        return;
    }

    fprintf(fp, "%d, %f\n", run, cpu_time);

    fclose(fp);
}
