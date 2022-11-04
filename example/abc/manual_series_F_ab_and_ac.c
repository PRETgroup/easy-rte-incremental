
//This file should be called series_F_ab_and_ac.c
//This is autogenerated code. Edit by hand at your peril!
#include "manual_series_F_ab_and_ac.h"

inputs_ab_and_ac_t possibleInputs[INPUT_OPTIONS] = {
	{0},
	{1},
};
outputs_ab_and_ac_t possibleOutputs[OUTPUT_OPTIONS] = {
	{0,0},
	{0,1},
	{1,0},
	{1,1},
};

const inputs_ab_and_ac_t unacceptableInput = {BAD};
const outputs_ab_and_ac_t unacceptableOutput = {BAD, BAD};

void ab_and_ac_init_all_vars(enforcervars_ab_and_ac_t* me, inputs_ab_and_ac_t* inputs, outputs_ab_and_ac_t* outputs) {
	//set any input vars with default values
	inputs->A = 0;
	//set any output vars with default values
	outputs->B = 0;
	outputs->C = 0;
	
	me->_policy_ab_state = POLICY_STATE_ab_and_ac_ab_b0;
	//input policy internal vars
	
	
	me->_policy_ac_state = POLICY_STATE_ab_and_ac_ac_c0;
	//input policy internal vars
	
	
}

void ab_and_ac_run_via_enforcer(enforcervars_ab_and_ac_t* me, inputs_ab_and_ac_t* inputs, outputs_ab_and_ac_t* outputs) {
	// Prepare set of all possible inputs (as all inputs are acceptable until enforcers reduce this set)
	inputs_ab_and_ac_t acceptableInputs[INPUT_OPTIONS];
	memcpy(acceptableInputs, possibleInputs, INPUT_OPTIONS*BYTES_PER_INPUT);
	// Same for all possible outputs..
	outputs_ab_and_ac_t acceptableOutputs[OUTPUT_OPTIONS];
	memcpy(acceptableOutputs, possibleOutputs, OUTPUT_OPTIONS*BYTES_PER_OUTPUT);

	// Run input enforcers
	// 	acceptableInputs is the set that gets reduced by each enforcer
	ab_and_ac_run_input_enforcer_ac(me, inputs, acceptableInputs); 
	ab_and_ac_run_input_enforcer_ab(me, inputs, acceptableInputs);
	
	// Select input
	select_input(inputs, acceptableInputs);

	ab_and_ac_run(inputs, outputs);

	// Run output enforcers
	ab_and_ac_run_output_enforcer_ab(me, inputs, outputs, acceptableOutputs);
	ab_and_ac_run_output_enforcer_ac(me, inputs, outputs, acceptableOutputs);

	// Select output
	select_output(outputs, acceptableOutputs);

	// Update output enforcer locations
	ab_and_ac_transition_output_enforcer_ab(me, inputs, outputs);
	ab_and_ac_transition_output_enforcer_ac(me, inputs, outputs);
	
}


bool compareInputs(const inputs_ab_and_ac_t* inp1, const inputs_ab_and_ac_t* inp2) {
	return (inp1->A == inp2->A);
}

bool isBadInput(const inputs_ab_and_ac_t* inputs) {
	const inputs_ab_and_ac_t* badInput = &unacceptableInput;
	return compareInputs(inputs, badInput);
}

bool compareOutputs(const outputs_ab_and_ac_t* out1, const outputs_ab_and_ac_t* out2) {
	return ((out1->B == out2->B) && (out1->C == out2->C));
}

bool isBadOutput(const outputs_ab_and_ac_t* outputs) {
	const outputs_ab_and_ac_t* badOutput = &unacceptableOutput;
	return compareOutputs(outputs, badOutput);
}


// Input Intersection
void input_option_intersection(const inputs_ab_and_ac_t* acceptableInputs, const uint8_t numAccept, inputs_ab_and_ac_t* inputOptions) {
	// for every option
	for (uint16_t i = 0; i < INPUT_OPTIONS; i++) {
		// check if it is in acceptable outputs
		// printf("Option: %d\n", inputOptions[i].A);

		bool accept = false;
		for (uint8_t j = 0; j < numAccept; j++) {
			// if ((inputOptions[i].A == acceptableInputs[j].A)) {
			if (compareInputs(&inputOptions[i], &acceptableInputs[j])) {
				accept = true;
			} 
		}
		if (!accept) {
			// printf("Deny: %d\n", inputOptions[i].A);
			memcpy(&inputOptions[i], &unacceptableInput, BYTES_PER_INPUT);
		//} else {
			// printf("Accept: %d\n", inputOptions[i].A);
		}
	}
}

// Output Intersection
void output_option_intersection(const outputs_ab_and_ac_t* acceptableOutputs, const uint8_t numAccept, outputs_ab_and_ac_t* outputOptions) {
	// for every option
	for (uint16_t i = 0; i < OUTPUT_OPTIONS; i++) {
		// check if it is in acceptable outputs
		// printf("Option: %d, %d\n", outputOptions[i].B, outputOptions[i].C);

		bool accept = false;
		for (uint8_t j = 0; j < numAccept; j++) {
			// printf("Acceptable Input: %d, %d\n", acceptableOutputs[j].B, acceptableOutputs[j].C);
			// if ((outputOptions[i].B == acceptableOutputs[j].B) && (outputOptions[i].C == acceptableOutputs[j].C)) {
			if (compareOutputs(&outputOptions[i], &acceptableOutputs[j])) {
				accept = true;
			} 
		}
		if (!accept) {
			// printf("Deny: %d, %d\n", outputOptions[i].B, outputOptions[i].C);
			memcpy(&outputOptions[i], &unacceptableOutput, BYTES_PER_OUTPUT);
		// } else {
			// printf("Accept: %d, %d\n", outputOptions[i].B, outputOptions[i].C);
		}
	}
}

// Select Input
void select_input(inputs_ab_and_ac_t* inputs, inputs_ab_and_ac_t* inputOptions) {
	// TODO: replace with min-edit/min-select/rand-select
	// printf("Current %d, %d\n", inputs->x_accel, inputs->x_hold);
	for (uint16_t i = 0; i < INPUT_OPTIONS; i++) {
		if (
			// (inputOptions[i].A != BAD)
			!isBadInput(&inputOptions[i])
			) {
    		// printf("Acceptable: x_accel:%d, x_hold:%d, y_accel:%d, y_hold:%d, z_accel:%d, z_hold:%d, rpm_up:%d, rpm_hold:%d\r\n", inputOptions[i].x_accel, inputOptions[i].x_hold, inputOptions[i].y_accel, inputOptions[i].y_hold, inputOptions[i].z_accel, inputOptions[i].z_hold, inputOptions[i].rpm_up, inputOptions[i].rpm_hold);
			if (
				compareInputs(&inputOptions[i], inputs)
				// (inputOptions[i].A == inputs->A)
				) {
				// printf("No input edit needed\n");
				return;
			}
		}
	}

	for (uint16_t i = 0; i < INPUT_OPTIONS; i++) {
		// printf("i: %d\n", i);
		if (
			// (inputOptions[i].A != BAD)
			!isBadInput(&inputOptions[i])
			) {
			// printf("Good option %d, %d, %d\n", inputOptions[i].rpm_up, inputOptions[i].x_accel, inputOptions[i].x_hold);
			printf("Input edited\n");
			memcpy(inputs, &inputOptions[i], BYTES_PER_INPUT);
			return;
		}
	}
}

// Select Output
void select_output(outputs_ab_and_ac_t* outputs, outputs_ab_and_ac_t* outputOptions) {
	// TODO: replace with min-edit/min-select/rand-select
	// printf("Current %d, %d\n", outputs->x_accel, outputs->x_hold);
	for (uint16_t i = 0; i < OUTPUT_OPTIONS; i++) {
		if (
			// (outputOptions[i].B != BAD) &&
			// (outputOptions[i].C != BAD)
			!isBadOutput(&outputOptions[i])
			) {
    		// printf("Acceptable: B:%d, C:%d\r\n", outputOptions[i].B, outputOptions[i].C);
			if (
				// (outputOptions[i].B == outputs->B) && (outputOptions[i].C == outputs->C)
				compareOutputs(&outputOptions[i], outputs)
				) {
				// printf("No output edit needed\n");
    			// printf("Selected: B:%d, C:%d\r\n", outputs->B, outputs->C);
				return;
			}
		}
	}

	for (uint16_t i = 0; i < OUTPUT_OPTIONS; i++) {
		// printf("i: %d\n", i);
		if (
			// (outputOptions[i].B != BAD) &&
			// (outputOptions[i].C != BAD)
			!isBadOutput(&outputOptions[i])
			) {
			// printf("Good option %d, %d\n", outputOptions[i].B, outputOptions[i].C);
			printf("Output Edited\n");
			memcpy(outputs, &outputOptions[i], BYTES_PER_OUTPUT);

			return;
		}
	}
}


//Input policies

//INPUT POLICY ab BEGIN
//This will run the input enforcer for ab_and_ac's policy ab
void ab_and_ac_run_input_enforcer_ab(enforcervars_ab_and_ac_t* me, inputs_ab_and_ac_t* inputs, inputs_ab_and_ac_t* inputOptions) {
	switch(me->_policy_ab_state) {
		case POLICY_STATE_ab_and_ac_ab_b0:
			
			if( !(inputs->A) || inputs->A) {
				//transition b0 -> violation on not (A) or A
				
			} 
			
			break;

		case POLICY_STATE_ab_and_ac_ab_b1:
			
			if( !(inputs->A) || inputs->A) {
				//transition b1 -> violation on not (A) or A
				
				// Previously this was a defined recovery
				// inputs->A = 0;

				// Set of acceptable inputs
				const uint8_t numAccept = 1;
				const inputs_ab_and_ac_t acceptableInputs[1] = {{0}};
				
				input_option_intersection(acceptableInputs, numAccept, inputOptions);
				
			} 
			
			break;

		
	}
}

//INPUT POLICY ab END

//INPUT POLICY ac BEGIN
//This will run the input enforcer for ab_and_ac's policy ac
void ab_and_ac_run_input_enforcer_ac(enforcervars_ab_and_ac_t* me, inputs_ab_and_ac_t* inputs, inputs_ab_and_ac_t* inputOptions) {
	switch(me->_policy_ac_state) {
		case POLICY_STATE_ab_and_ac_ac_c0:
			
			if( !(inputs->A) || inputs->A) {
				//transition c0 -> violation on not (A) or A
				//select a transition to solve the problem
				
				//Recovery instructions manually provided.
				
			} 
			
			break;

		case POLICY_STATE_ab_and_ac_ac_c1:
			
			if( !(inputs->A) || inputs->A) {
				//transition c1 -> violation on not (A) or A
				//select a transition to solve the problem
				
				// Previously this was a defined recovery
				// inputs->A = 0;

				// Set of acceptable inputs
				const uint8_t numAccept = 1;
				const inputs_ab_and_ac_t acceptableInputs[1] = {{0}};
				
				input_option_intersection(acceptableInputs, numAccept, inputOptions);
				
			} 
			
			break;

		
	}
}

//INPUT POLICY ac END



//output policies

//OUTPUT POLICY ab BEGIN
//This will run the input enforcer for ab_and_ac's policy ab
void ab_and_ac_run_output_enforcer_ab(enforcervars_ab_and_ac_t* me, inputs_ab_and_ac_t* inputs, outputs_ab_and_ac_t* outputs, outputs_ab_and_ac_t* outputOptions) {
	//advance timers
	// printf("AB Output Running - loc: A:%d, B:%d, C:%d, State:%d\r\n", inputs->A, outputs->B, outputs->C, me->_policy_ab_state);

	//run enforcer
	switch(me->_policy_ab_state) {
		case POLICY_STATE_ab_and_ac_ab_b0:
			
			if(( !(inputs->A) && outputs->B) || (inputs->A && outputs->B)) {
				//transition b0 -> violation on ( ( !A and B ) or ( A and B ) )
				//select a transition to solve the problem
				// Previously 
				// outputs->B = 0;

				// Now provide a set of acceptable outputs
				const uint8_t numAccept = 2;
				const outputs_ab_and_ac_t acceptableOutputs[2] = {
					{0,0},
					{0,1}
				};
				// Reduce the set of outputOptions 
				output_option_intersection(acceptableOutputs, numAccept, outputOptions);
				
			} 

			break;

		case POLICY_STATE_ab_and_ac_ab_b1:
			
			if(( !(inputs->A) && !(outputs->B)) || inputs->A) {
				//transition b1 -> violation on ( ( !A and !B ) or ( A ) )
				//select a transition to solve the problem
				
				//Recovery instructions manually provided.
				
			} 
			if( !(outputs->B)) {
				//transition b1 -> violation on ( !B )
				//select a transition to solve the problem
				
				// Previously 
					//Recovery instructions manually provided.
				// outputs->B = 1;

				// Now provide a set of acceptable outputs
				const uint8_t numAccept = 2;
				const outputs_ab_and_ac_t acceptableOutputs[2] = {
					{1,0},
					{1,1}
				};
				// Reduce the set of outputOptions 
				output_option_intersection(acceptableOutputs, numAccept, outputOptions);
				
			} 

			break;
	}
}

void ab_and_ac_transition_output_enforcer_ab(enforcervars_ab_and_ac_t* me, inputs_ab_and_ac_t* inputs, outputs_ab_and_ac_t* outputs){
	// printf("AB Transition: A:%d, B:%d, C:%d, State:%d\r\n", inputs->A, outputs->B, outputs->C, me->_policy_ab_state);

	//select transition to advance state
	switch(me->_policy_ab_state) {
		case POLICY_STATE_ab_and_ac_ab_b0:
			
			if( !(inputs->A) && !(outputs->B)) {
				//transition b0 -> b0 on ( !A and !B )
				me->_policy_ab_state = POLICY_STATE_ab_and_ac_ab_b0;
				//set expressions
				
				break;
			} 
			if(inputs->A && !(outputs->B)) {
				//transition b0 -> b1 on ( A and !B )
				me->_policy_ab_state = POLICY_STATE_ab_and_ac_ab_b1;
				//set expressions
				
				break;
			} 
			if(( !(inputs->A) && outputs->B) || (inputs->A && outputs->B)) {
				//transition b0 -> violation on ( ( !A and B ) or ( A and B ) )
				me->_policy_ab_state = POLICY_STATE_ab_and_ac_ab_violation;
				//set expressions
				
				break;
			} 
			
			//ensure a transition was taken in this state
			assert(false && "ab_and_ac_ab_b0 must take a transition"); //if we are still here, then no transition was taken and we are no longer satisfying liveness

			break;

		case POLICY_STATE_ab_and_ac_ab_b1:
			
			if( !(inputs->A) && outputs->B) {
				//transition b1 -> b0 on ( !A and B )
				me->_policy_ab_state = POLICY_STATE_ab_and_ac_ab_b0;
				//set expressions
				
				break;
			} 
			if(( !(inputs->A) && !(outputs->B)) || inputs->A) {
				//transition b1 -> violation on ( ( !A and !B ) or ( A ) )
				me->_policy_ab_state = POLICY_STATE_ab_and_ac_ab_violation;
				//set expressions
				
				break;
			} 
			if( !(outputs->B)) {
				//transition b1 -> violation on ( !B )
				me->_policy_ab_state = POLICY_STATE_ab_and_ac_ab_violation;
				//set expressions
				
				break;
			} 
			
			//ensure a transition was taken in this state
			assert(false && "ab_and_ac_ab_b1 must take a transition"); //if we are still here, then no transition was taken and we are no longer satisfying liveness

			break;

		
	}
	// printf("AB New Loc: State:%d\r\n", me->_policy_ab_state);

	//ensure we did not violate (i.e. we did not enter violation state)
	assert(me->_policy_ab_state != POLICY_STATE_ab_and_ac_ab_violation);
}

//OUTPUT POLICY ab END

//OUTPUT POLICY ac BEGIN
//This will run the input enforcer for ab_and_ac's policy ac
void ab_and_ac_run_output_enforcer_ac(enforcervars_ab_and_ac_t* me, inputs_ab_and_ac_t* inputs, outputs_ab_and_ac_t* outputs, outputs_ab_and_ac_t* outputOptions) {
	//advance timers
	
	
	//run enforcer
	switch(me->_policy_ac_state) {
		case POLICY_STATE_ab_and_ac_ac_c0:
			
			if(( !(inputs->A) && outputs->C) || (inputs->A && outputs->C)) {
				//transition c0 -> violation on ( ( !A and C ) or ( A and C ) )
				//select a transition to solve the problem
				
				// Previously
				// outputs->C = 0;
				
				// Now provide a set of acceptable outputs
				const uint8_t numAccept = 2;
				const outputs_ab_and_ac_t acceptableOutputs[2] = {
					{0,0},
					{1,0}
				};
				// Reduce the set of outputOptions 
				output_option_intersection(acceptableOutputs, numAccept, outputOptions);
				
			} 

			break;

		case POLICY_STATE_ab_and_ac_ac_c1:
			
			if(( !(inputs->A) && !(outputs->C)) || inputs->A) {
				//transition c1 -> violation on ( ( !A and !C ) or ( A ) )
				//select a transition to solve the problem
				
				//Recovery instructions manually provided.
				
			} 
			if( !(outputs->C)) {
				//transition c1 -> violation on ( !C )
				//select a transition to solve the problem
				
				// Previously
				// outputs->C = 1;
								
				// Now provide a set of acceptable outputs
				const uint8_t numAccept = 2;
				const outputs_ab_and_ac_t acceptableOutputs[2] = {
					{0,1},
					{1,1}
				};
				// Reduce the set of outputOptions 
				output_option_intersection(acceptableOutputs, numAccept, outputOptions);
				
			} 

			break;

		
	}
}

void ab_and_ac_transition_output_enforcer_ac(enforcervars_ab_and_ac_t* me, inputs_ab_and_ac_t* inputs, outputs_ab_and_ac_t* outputs){
	//select transition to advance state
	switch(me->_policy_ac_state) {
		case POLICY_STATE_ab_and_ac_ac_c0:
			
			if( !(inputs->A) && !(outputs->C)) {
				//transition c0 -> c0 on ( !A and !C )
				me->_policy_ac_state = POLICY_STATE_ab_and_ac_ac_c0;
				//set expressions
				
				break;
			} 
			if(inputs->A && !(outputs->C)) {
				//transition c0 -> c1 on ( A and !C )
				me->_policy_ac_state = POLICY_STATE_ab_and_ac_ac_c1;
				//set expressions
				
				break;
			} 
			if(( !(inputs->A) && outputs->C) || (inputs->A && outputs->C)) {
				//transition c0 -> violation on ( ( !A and C ) or ( A and C ) )
				me->_policy_ac_state = POLICY_STATE_ab_and_ac_ac_violation;
				//set expressions
				
				break;
			} 
			
			//ensure a transition was taken in this state
			assert(false && "ab_and_ac_ac_c0 must take a transition"); //if we are still here, then no transition was taken and we are no longer satisfying liveness

			break;

		case POLICY_STATE_ab_and_ac_ac_c1:
			
			if( !(inputs->A) && outputs->C) {
				//transition c1 -> c0 on ( !A and C )
				me->_policy_ac_state = POLICY_STATE_ab_and_ac_ac_c0;
				//set expressions
				
				break;
			} 
			if(( !(inputs->A) && !(outputs->C)) || inputs->A) {
				//transition c1 -> violation on ( ( !A and !C ) or ( A ) )
				me->_policy_ac_state = POLICY_STATE_ab_and_ac_ac_violation;
				//set expressions
				
				break;
			} 
			if( !(outputs->C)) {
				//transition c1 -> violation on ( !C )
				me->_policy_ac_state = POLICY_STATE_ab_and_ac_ac_violation;
				//set expressions
				
				break;
			} 
			
			//ensure a transition was taken in this state
			assert(false && "ab_and_ac_ac_c1 must take a transition"); //if we are still here, then no transition was taken and we are no longer satisfying liveness

			break;

		
	}

	//ensure we did not violate (i.e. we did not enter violation state)
	assert(me->_policy_ac_state != POLICY_STATE_ab_and_ac_ac_violation);
}

//OUTPUT POLICY ac END



//This function is provided in "series_F_ab_and_ac.c"
//It will check the state of the enforcer monitor code
//It returns one of the following:
//0: currently true (safe)
//1: always true (safe)
//-1: currently false (unsafe)
//-2: always false (unsafe)
//It will need to do some reachability analysis to achieve this

