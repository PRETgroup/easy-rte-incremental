
//This file should be called series_F_ab_and_ac.h
//This is autogenerated code. Edit by hand at your peril!

#include <stdint.h>
#include <stdbool.h>
#include <assert.h>
#include <string.h>
#include <stdio.h>

#define true 1
#define false 0
#define BAD 3

#define INPUT_OPTIONS 2 // number possible output combinations
#define OUTPUT_OPTIONS 4 // number possible output combinations

//the dtimer_t type
typedef uint64_t dtimer_t;

//For each policy, we need an enum type for the state machine

enum ab_and_ac_policy_ab_states { 
	POLICY_STATE_ab_and_ac_ab_b0, 
	POLICY_STATE_ab_and_ac_ab_b1,
	POLICY_STATE_ab_and_ac_ab_violation 
};

enum ab_and_ac_policy_ac_states { 
	POLICY_STATE_ab_and_ac_ac_c0, 
	POLICY_STATE_ab_and_ac_ac_c1,
	POLICY_STATE_ab_and_ac_ac_violation 
};


//Inputs to the function ab_and_ac
typedef struct {
	uint8_t A;
	
} inputs_ab_and_ac_t;

#define BYTES_PER_INPUT 1 // Bool = 1 Byte, A is only, therefore 1 Byte

//Outputs from the function ab_and_ac
typedef struct {
	uint8_t B;
	uint8_t C;
	
} outputs_ab_and_ac_t;

#define BYTES_PER_OUTPUT 2 // Boolx2 = 2 Bytes

//enforcer state and vars:
typedef struct {
	enum ab_and_ac_policy_ab_states _policy_ab_state;
	//internal vars
	
	
	enum ab_and_ac_policy_ac_states _policy_ac_state;
	//internal vars
	
	
	
} enforcervars_ab_and_ac_t;



//This function is provided in "series_F_ab_and_ac.c"
//It sets up the variable structures to their initial values
void ab_and_ac_init_all_vars(enforcervars_ab_and_ac_t* me, inputs_ab_and_ac_t* inputs, outputs_ab_and_ac_t* outputs);

//This function is provided in "series_F_ab_and_ac.c"
//It will run the synthesised enforcer and call the controller function
void ab_and_ac_run_via_enforcer(enforcervars_ab_and_ac_t* me, inputs_ab_and_ac_t* inputs, outputs_ab_and_ac_t* outputs);

//This function is provided from the user
//It is the controller function
extern void ab_and_ac_run(inputs_ab_and_ac_t* inputs, outputs_ab_and_ac_t* outputs);

//enforcer functions


//This function is provided in "series_F_ab_and_ac.c"
//It will run the input enforcer for ab_and_ac's policy ab
void ab_and_ac_run_input_enforcer_ab(enforcervars_ab_and_ac_t* me, inputs_ab_and_ac_t* inputs, uint8_t* newAcceptableInputsBool);

//This function is provided in "series_F_ab_and_ac.c"
//It will run the output enforcer for ab_and_ac's policy ab
void ab_and_ac_run_output_enforcer_ab(enforcervars_ab_and_ac_t* me, inputs_ab_and_ac_t* inputs, outputs_ab_and_ac_t* outputs, uint8_t* newAcceptableOutputsBool);


//It will run update the location of output enforcer for ab_and_ac's policy ab
void ab_and_ac_transition_output_enforcer_ab(enforcervars_ab_and_ac_t* me, inputs_ab_and_ac_t* inputs, outputs_ab_and_ac_t* outputs);


//This function is provided in "series_F_ab_and_ac.c"
//It will check the state of the enforcer monitor code
//It returns one of the following:
//0: currently true (safe)
//1: always true (safe)
//-1: currently false (unsafe)
//-2: always false (unsafe)
//It will need to do some reachability analysis to achieve this
int ab_and_ac_check_rv_status_ab(enforcervars_ab_and_ac_t* me);


//This function is provided in "series_F_ab_and_ac.c"
//It will run the input enforcer for ab_and_ac's policy ac
void ab_and_ac_run_input_enforcer_ac(enforcervars_ab_and_ac_t* me, inputs_ab_and_ac_t* inputs, uint8_t* newAcceptableInputsBool);

//This function is provided in "series_F_ab_and_ac.c"
//It will run the output enforcer for ab_and_ac's policy ac
void ab_and_ac_run_output_enforcer_ac(enforcervars_ab_and_ac_t* me, inputs_ab_and_ac_t* inputs, outputs_ab_and_ac_t* outputs, uint8_t* newAcceptableOutputsBool);


//It will run update the location of output enforcer for ab_and_ac's policy ac
void ab_and_ac_transition_output_enforcer_ac(enforcervars_ab_and_ac_t* me, inputs_ab_and_ac_t* inputs, outputs_ab_and_ac_t* outputs);


//This function is provided in "series_F_ab_and_ac.c"
//It will check the state of the enforcer monitor code
//It returns one of the following:
//0: currently true (safe)
//1: always true (safe)
//-1: currently false (unsafe)
//-2: always false (unsafe)
//It will need to do some reachability analysis to achieve this
int ab_and_ac_check_rv_status_ac(enforcervars_ab_and_ac_t* me);


void select_input_new(inputs_ab_and_ac_t* uneditedInputs, uint8_t* newAcceptableInputsBool, inputs_ab_and_ac_t* possibleInputs);
void select_output_new(outputs_ab_and_ac_t* uneditedOutputs, uint8_t* newAcceptableOutputsBool, outputs_ab_and_ac_t* possibleOutputs);