#include "F_px_py.h"
#include <stdio.h>

#include <time.h>

#define RUNS 10
#define TICKS_PER_RUN 100000000

void print_data(uint32_t count, inputs_px_py_t inputs, outputs_px_py_t outputs) {
    printf("Tick %7d: x:%d, x_accel:%d, x_hold:%d, y:%d, y_accel:%d, y_hold:%d\r\n", count, inputs.x, outputs.x_accel, outputs.x_hold, inputs.y, outputs.y_accel, outputs.y_hold);
}

int main() {
    double cpu_time_used[RUNS] = {0};
    for (uint16_t run = 0; run < RUNS; run++) {
        clock_t start, end;

        start = clock();
        enforcervars_px_py_t enf;
        inputs_px_py_t inputs;
        outputs_px_py_t outputs;
        
        px_py_init_all_vars(&enf, &inputs, &outputs);

        uint32_t count = 0;
        inputs.x = 0;
        inputs.y = 0;
        while(count++ < TICKS_PER_RUN) {
            if(count < 10 == 1) {
                inputs.x = inputs.x + 1;
                outputs.x_accel = true;
                outputs.x_hold = true;
                inputs.y = inputs.y + 1;
                outputs.y_accel = true;
                outputs.y_hold = true;
            } else {
                inputs.x = 10;
                outputs.x_accel = true;
                outputs.x_hold = true;
                inputs.y = 10;
                outputs.y_accel = true;
                outputs.y_hold = true;
            }

            px_py_run_via_enforcer(&enf, &inputs, &outputs);

            // print_data(count, inputs, outputs);
        }
        
        end = clock();
        cpu_time_used[run] = ((double) (end - start)) / CLOCKS_PER_SEC;

        printf("%f\n", cpu_time_used[run]);
    }   
}

void px_py_run(inputs_px_py_t *inputs, outputs_px_py_t *outputs) {
    //do nothing
}

