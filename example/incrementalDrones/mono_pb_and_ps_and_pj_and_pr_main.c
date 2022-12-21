#include "F_pb_and_ps_and_pj_and_pr.h"
#include <stdio.h>

#include "save_results.h"
#include <time.h>
#define RUN_REFERENCE "monolithic_pb_and_ps_and_pj_and_pr"

void print_data(uint32_t count, inputs_pb_and_ps_and_pj_and_pr_t inputs, outputs_pb_and_ps_and_pj_and_pr_t outputs) {
    // printf("Tick %7d: x:%d, x_up:%d, x_down:%d, y:%d, y_up:%d, y_down:%d, z:%d, z_up:%d, z_down:%d, rpm:%d, rpm_up:%d, rpm_down:%d, x2:%d, x2_up:%d, x2_down:%d\r\n", count, inputs.x, outputs.x_up, outputs.x_down, inputs.y, outputs.y_up, outputs.y_down, inputs.z, outputs.z_up, outputs.z_down, inputs.rpm, outputs.rpm_up, outputs.rpm_down, inputs.x2, outputs.x2_up, outputs.x2_down);
}

int main() {
    char run_reference[] = RUN_REFERENCE;
    save_setup(run_reference);

    double cpu_time_used[RUNS] = {0};
    double cpu_time_ms_per_tick[RUNS] = {0};

    for (uint16_t run = 0; run < RUNS; run++) {
        clock_t start, end;

        start = clock();
        enforcervars_pb_and_ps_and_pj_and_pr_t enf;
        inputs_pb_and_ps_and_pj_and_pr_t inputs;
        outputs_pb_and_ps_and_pj_and_pr_t outputs;
        
        pb_and_ps_and_pj_and_pr_init_all_vars(&enf, &inputs, &outputs);

        uint32_t count = 0;

        // Drone 1 
        //  min x = 0
        //  max x = 2
        //  min y = 200
        //  max y = 202
        //  rpm max = 7825

        // Drone 2 
        //  min x = 3
        //  max x = 5
        //  min y = 200
        //  max y = 202
        //  rpm max = 7825

        // Initalise to maximums 
        // controller will attempt to continue to increase beyond limits
        // incremental enforcers will prevent limitation breaches
        inputs.x = 0;
        inputs.y = 200;
        inputs.rpm = 7825;
        inputs.x2 = 3;
        inputs.y2 = 200;
        inputs.rpm2 = 7825;

        // print_data(0, inputs, outputs);
        while(count++ < TICKS_PER_RUN) {
            // inputs.x = inputs.x + 1;
            outputs.x_up = true;
            outputs.x_down = true;
            // inputs.y = inputs.y + 1;
            outputs.y_up = true;
            outputs.y_down = true;
            // inputs.rpm = inputs.rpm + 100;
            outputs.rpm_up = true;
            outputs.rpm_down = true;
            // inputs.x2 = inputs.x2 + 1;
            outputs.x2_up = true;
            outputs.x2_down = true;
            // inputs.y2 = inputs.y2 + 1;
            outputs.y2_up = true;
            outputs.y2_down = true;
            // inputs.rpm2 = inputs.rpm2 + 100;
            outputs.rpm2_up = true;
            outputs.rpm2_down = true;

            pb_and_ps_and_pj_and_pr_run_via_enforcer(&enf, &inputs, &outputs);

            // print_data(count, inputs, outputs);
        }
        
        end = clock();
        cpu_time_used[run] = ((double) (end - start)) / CLOCKS_PER_SEC;
        cpu_time_ms_per_tick[run] = (cpu_time_used[run] * 1000.0) / ((double) TICKS_PER_RUN);

        save_time(run+1, cpu_time_used[run], cpu_time_ms_per_tick[run]);

        printf("Run: %d, CPU: %.0f ms, Per Tick: %f ms\n", run+1, cpu_time_used[run]*1000, cpu_time_ms_per_tick[run]);
    }   
}

void pb_and_ps_and_pj_and_pr_run(inputs_pb_and_ps_and_pj_and_pr_t *inputs, outputs_pb_and_ps_and_pj_and_pr_t *outputs) {
    //do nothing
}

