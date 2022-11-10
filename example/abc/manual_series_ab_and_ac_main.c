#include "manual_series_F_ab_and_ac.h"
#include <stdio.h>
#include <stdint.h>

#include "save_results.h"
#include <time.h>
#define RUNS 10
#define TICKS_PER_RUN 10000000
#define RUN_REFERENCE "ab_and_ac"

void print_data(uint32_t count, inputs_ab_and_ac_t inputs, outputs_ab_and_ac_t outputs) {
    printf("Tick %7d: A:%d, B:%d, C:%d\r\n", count, inputs.A, outputs.B, outputs.C);
}

int main() {
    char run_reference[] = RUN_REFERENCE;
    save_setup(run_reference);
    double cpu_time_used[RUNS] = {0};
    for (uint16_t run = 0; run < RUNS; run++) {

        clock_t start, end;

        start = clock();

        enforcervars_ab_and_ac_t enf;
        inputs_ab_and_ac_t inputs;
        outputs_ab_and_ac_t outputs;
        
        ab_and_ac_init_all_vars(&enf, &inputs, &outputs);

        uint32_t count = 0;
        while(count++ < TICKS_PER_RUN) {
            if(count % 5 == 0) {
                inputs.A = true;
            } else {
                inputs.A = false;
            }
            // print_data(count, inputs, outputs);

            ab_and_ac_run_via_enforcer(&enf, &inputs, &outputs);

            // print_data(count, inputs, outputs);
        }

        end = clock();
        cpu_time_used[run] = ((double) (end - start)) / CLOCKS_PER_SEC;

        save_time(run+1, cpu_time_used[run]);
        printf("%f\n", cpu_time_used[run]);
    }
}

void ab_and_ac_run(inputs_ab_and_ac_t *inputs, outputs_ab_and_ac_t *outputs) {
    //do nothing
}

