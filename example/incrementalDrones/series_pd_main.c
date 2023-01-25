#include "series_F_pd.h"
#include <stdio.h>

#include "save_results.h"
#include <time.h>
#define RUN_REFERENCE "series_pd"

void print_data(uint32_t count, inputs_pd_t inputs, outputs_pd_t outputs) {
    printf("Tick %7d: x:%d, x_up:%d, x_down:%d, y:%d, y_up:%d, y_down:%d, rpm:%d, rpm_up:%d, rpm_down:%d, x2:%d, x2_up:%d, x2_down:%d\r\n", count, inputs.x, outputs.x_up, outputs.x_down, inputs.y, outputs.y_up, outputs.y_down, inputs.rpm, outputs.rpm_up, outputs.rpm_down, inputs.x2, outputs.x2_up, outputs.x2_down);
}

int main() {
    char run_reference[] = RUN_REFERENCE;
    save_setup(run_reference);

    double cpu_time_used[RUNS] = {0};
    double cpu_time_ms_per_tick[RUNS] = {0};

    for (uint16_t run = 0; run < RUNS; run++) {
        clock_t start, end;

        start = clock();
        enforcervars_pd_t enf;
        inputs_pd_t inputs;
        outputs_pd_t outputs;
        
        pd_init_all_vars(&enf, &inputs, &outputs);

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
        inputs.x = 2;
        inputs.y = 200;
        inputs.rpm = 7825;

        inputs.x2 = 3;
        inputs.y2 = 200;
        inputs.rpm2 = 7825;
        // print_data(0, inputs, outputs);
        while(count++ < TICKS_PER_RUN) {
            // inputs.x = inputs.x + 1;
            outputs.x_up = false;
            outputs.x_down = false;

            outputs.x2_up = false;
            outputs.x2_down = false;

            pd_run_via_enforcer(&enf, &inputs, &outputs);

            print_data(count, inputs, outputs);
        }
        
        end = clock();
        cpu_time_used[run] = ((double) (end - start)) / CLOCKS_PER_SEC;
        cpu_time_ms_per_tick[run] = (cpu_time_used[run] * 1000.0) / ((double) TICKS_PER_RUN);

        save_time(run+1, cpu_time_used[run], cpu_time_ms_per_tick[run]);

        printf("Run: %d, CPU: %.0f ms, Per Tick: %f ms\n", run+1, cpu_time_used[run]*1000, cpu_time_ms_per_tick[run]);
    }   
}

void pd_run(inputs_pd_t *inputs, outputs_pd_t *outputs) {
    //do nothing
}

