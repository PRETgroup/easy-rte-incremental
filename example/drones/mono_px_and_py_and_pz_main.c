#include "F_px_and_py_and_pz.h"
#include <stdio.h>

#include <time.h>

#define RUNS 10
#define TICKS_PER_RUN 10000000

void print_data(uint32_t count, inputs_px_and_py_and_pz_t inputs, outputs_px_and_py_and_pz_t outputs) {
    printf("Tick %7d: x:%d, x_accel:%d, x_hold:%d, y:%d, y_accel:%d, y_hold:%d, z:%d, z_accel:%d, z_hold:%d, rpm:%d, rpm_up:%d, rpm_hold:%d, x2:%d, x2_accel:%d, x2_hold:%d\r\n", count, inputs.x, outputs.x_accel, outputs.x_hold, inputs.y, outputs.y_accel, outputs.y_hold, inputs.z, outputs.z_accel, outputs.z_hold, inputs.rpm, outputs.rpm_up, outputs.rpm_hold, inputs.x2, outputs.x2_accel, outputs.x2_hold);
}

int main() {
    double cpu_time_used[RUNS] = {0};
    for (uint16_t run = 0; run < RUNS; run++) {
        clock_t start, end;

        start = clock();
        enforcervars_px_and_py_and_pz_t enf;
        inputs_px_and_py_and_pz_t inputs;
        outputs_px_and_py_and_pz_t outputs;
        
        px_and_py_and_pz_init_all_vars(&enf, &inputs, &outputs);

        uint32_t count = 0;
        inputs.x = 0;
        inputs.y = 0;
        inputs.z = 0;
        inputs.rpm = 800;
        inputs.x2 = 0;
        while(count++ < TICKS_PER_RUN) {
            if(count < 10 == 1) {
                inputs.x = inputs.x + 1;
                outputs.x_accel = true;
                outputs.x_hold = true;
                inputs.y = inputs.y + 1;
                outputs.y_accel = true;
                outputs.y_hold = true;
                inputs.z = inputs.z + 1;
                outputs.z_accel = true;
                outputs.z_hold = true;
                inputs.rpm = inputs.rpm + 50;
                outputs.rpm_up = true;
                outputs.rpm_hold = true;
                inputs.x2 = inputs.x2 + 1;
                outputs.x2_accel = true;
                outputs.x2_hold = true;
            } else {
                inputs.x = 10;
                outputs.x_accel = true;
                outputs.x_hold = true;
                inputs.y = 10;
                outputs.y_accel = true;
                outputs.y_hold = true;
                inputs.z = 10;
                outputs.z_accel = true;
                outputs.z_hold = true;
                inputs.rpm = 1250;
                outputs.rpm_up = true;
                outputs.rpm_hold = true;
                inputs.x2 = 10;
                outputs.x2_accel = true;
                outputs.x2_hold = true;
            }

            px_and_py_and_pz_run_via_enforcer(&enf, &inputs, &outputs);

            // print_data(count, inputs, outputs);
        }
        
        end = clock();
        cpu_time_used[run] = ((double) (end - start)) / CLOCKS_PER_SEC;

        printf("%f\n", cpu_time_used[run]);
    }   
}

void px_and_py_and_pz_run(inputs_px_and_py_and_pz_t *inputs, outputs_px_and_py_and_pz_t *outputs) {
    //do nothing
}

