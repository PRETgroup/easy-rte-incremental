#include "series_F_ab_and_ac.h"
#include <stdio.h>
#include <stdint.h>

void print_data(uint32_t count, inputs_ab_and_ac_t inputs, outputs_ab_and_ac_t outputs) {
    printf("Tick %7d: A:%d, B:%d, C:%d\r\n", count, inputs.A, outputs.B, outputs.C);
}

int main() {
    enforcervars_ab_and_ac_t enf;
    inputs_ab_and_ac_t inputs;
    outputs_ab_and_ac_t outputs;
    
    ab_and_ac_init_all_vars(&enf, &inputs, &outputs);

    uint32_t count = 0;
    while(count++ < 30) {
        if(count % 5 == 0) {
            inputs.A = true;
        } else {
            inputs.A = false;
        }
        print_data(count, inputs, outputs);

        ab_and_ac_run_via_enforcer(&enf, &inputs, &outputs);

        print_data(count, inputs, outputs);
    }
}

void ab_and_ac_run(inputs_ab_and_ac_t *inputs, outputs_ab_and_ac_t *outputs) {
    //do nothing
}

