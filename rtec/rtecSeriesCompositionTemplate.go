package rtec

import (
	"text/template"

	"github.com/PRETgroup/goFB/goFB/stconverter"
)

const rtecSeriesCompositionTemplate = `{{define "_policyIn"}}{{$block := .}}
// Input policies
{{range $polI, $pol := $block.Policies}}{{$pfbEnf := getPolicyEnfInfo $block $polI}}
{{if not $pfbEnf}}//{{$pol.Name}} is broken!
{{else}}{{/* this is where the policy comes in */}}//INPUT POLICY {{$pol.Name}} BEGIN
//This will run the input enforcer for {{$block.Name}}'s policy {{$pol.Name}}
void {{$block.Name}}_run_input_enforcer_{{$pol.Name}}(enforcervars_{{$block.Name}}_t* me, inputs_{{$block.Name}}_t* inputs, uint8_t* newAcceptableInputsBool) {
	switch(me->_policy_{{$pol.Name}}_state) {
		{{range $sti, $st := $pol.States}}case POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_{{$st.Name}}:
			{{range $tri, $tr := $pfbEnf.InputPolicy.GetViolationTransitions}}{{if eq $tr.Source $st.Name}}{{/*
			*/}}
			if({{$cond := getCECCTransitionCondition $block (compileExpression $tr.STGuard)}}{{$cond.IfCond}}) {
				//transition {{$tr.Source}} -> {{$tr.Destination}} on {{$tr.Condition}}
				//select a transition to solve the problem
				{{$solution := $pfbEnf.SolveViolationTransition $tr true}}
				{{if $solution.Comment}}//{{$solution.Comment}}{{end}}
				// {{range $soleI, $sole := $solution.Expressions}}{{$sol := getCECCTransitionCondition $block (compileExpression $sole)}}{{$sol.IfCond}};
				
				// Set of acceptable inputs
				// {{getAcceptableOptions (compileExpression $sole) $block.InputVars $block.Name "inputs"}}
				// input_option_intersection(acceptableOptions, numAccept, inputOptions);
				
				{{getUnacceptableOptions (compileExpression $sole) $block.InputVars $block.Name "inputs"}}
				unacceptable_input_option_intersection(newAcceptableInputsBool, numUnaccept, unacceptableOptions);
				{{end}}
			} {{end}}{{end}}
			
			break;

		{{end}}
	}
}
{{end}}
//INPUT POLICY {{$pol.Name}} END
{{end}}{{end}}

{{define "_policyOut"}}{{$block := .}}
//output policies
{{range $polI, $pol := $block.Policies}}{{$pfbEnf := getPolicyEnfInfo $block $polI}}
{{if not $pfbEnf}}//{{$pol.Name}} is broken!
{{else}}{{/* this is where the policy comes in */}}//OUTPUT POLICY {{$pol.Name}} BEGIN
//This will run the input enforcer for {{$block.Name}}'s policy {{$pol.Name}}
void {{$block.Name}}_run_output_enforcer_{{$pol.Name}}(enforcervars_{{$block.Name}}_t* me, inputs_{{$block.Name}}_t* inputs, outputs_{{$block.Name}}_t* outputs, uint8_t* newAcceptableOutputsBool) {
	//advance timers
	{{range $varI, $var := $pfbEnf.OutputPolicy.GetDTimers}}
	me->{{$var.Name}}++;{{end}}
	
	//run enforcer
	switch(me->_policy_{{$pol.Name}}_state) {
		{{range $sti, $st := $pol.States}}case POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_{{$st.Name}}:
			{{range $tri, $tr := $pfbEnf.OutputPolicy.GetViolationTransitions}}{{if eq $tr.Source $st.Name}}{{/*
			*/}}
			if({{$cond := getCECCTransitionCondition $block (compileExpression $tr.STGuard)}}{{$cond.IfCond}}) {
				//transition {{$tr.Source}} -> {{$tr.Destination}} on {{$tr.Condition}}
				//select a transition to solve the problem
				{{$solution := $pfbEnf.SolveViolationTransition $tr false}}
				{{if $solution.Comment}}//{{$solution.Comment}}{{end}}
				// {{range $soleI, $sole := $solution.Expressions}}{{$sol := getCECCTransitionCondition $block (compileExpression $sole)}}{{$sol.IfCond}};

				// Previously
				// {{getAcceptableOptions (compileExpression $sole) $block.OutputVars $block.Name "outputs"}}
				// output_option_intersection(acceptableOptions, numAccept, outputOptions);

				// Now provide a set of unacceptable outputs
				{{getUnacceptableOptions (compileExpression $sole) $block.OutputVars $block.Name "outputs"}}
				// Reduce the set of outputOptions 
				unacceptable_output_option_intersection(newAcceptableOutputsBool, numUnaccept, unacceptableOptions);
				{{end}}
			} {{end}}{{end}}

			break;

		{{end}}
	}
}

void {{$block.Name}}_transition_output_enforcer_{{$pol.Name}}(enforcervars_{{$block.Name}}_t* me, inputs_{{$block.Name}}_t* inputs, outputs_{{$block.Name}}_t* outputs){
	//select transition to advance state
	switch(me->_policy_{{$pol.Name}}_state) {
		{{range $sti, $st := $pol.States}}case POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_{{$st.Name}}:
			{{range $tri, $tr := $pfbEnf.OutputPolicy.Transitions}}{{if eq $tr.Source $st.Name}}{{/*
			*/}}
			if({{$cond := getCECCTransitionCondition $block (compileExpression $tr.STGuard)}}{{$cond.IfCond}}) {
				//transition {{$tr.Source}} -> {{$tr.Destination}} on {{$tr.Condition}}
				me->_policy_{{$pol.Name}}_state = POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_{{$tr.Destination}};
				//set expressions
				{{range $exi, $ex := $tr.Expressions}}
				me->{{$ex.VarName}} = {{$ex.Value}};{{end}}
				break;
			} {{end}}{{end}}
			
			//ensure a transition was taken in this state
			assert(false && "{{$block.Name}}_{{$pol.Name}}_{{$st.Name}} must take a transition"); //if we are still here, then no transition was taken and we are no longer satisfying liveness

			break;

		{{end}}
	}

	//ensure we did not violate (i.e. we did not enter violation state)
	assert(me->_policy_{{$pol.Name}}_state != POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_violation);
}
{{end}}
//OUTPUT POLICY {{$pol.Name}} END
{{end}}
{{end}}

{{define "functionH"}}{{$block := index .Functions .FunctionIndex}}{{$blocks := .Functions}}
//This file should be called series_F_{{$block.Name}}.h
//This is autogenerated code. Edit by hand at your peril!

#include <stdint.h>
#include <stdbool.h>
#include <assert.h>
#include <string.h>
#include <stdio.h>

#define true 1
#define false 0
#define BAD 3

#define INPUT_OPTIONS {{twoToThePower (len $block.InputVars)}} // number possible output combinations
#define OUTPUT_OPTIONS {{twoToThePower (len $block.OutputVars)}} // number possible output combinations

//the dtimer_t type
typedef uint64_t dtimer_t;

//For each policy, we need an enum type for the state machine
{{range $polI, $pol := $block.Policies}}
enum {{$block.Name}}_policy_{{$pol.Name}}_states { {{if len $pol.States}}{{range $index, $state := $pol.States}}{{if $index}}, {{end}}
	POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_{{$state}}{{end}}{{else}}POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_unknown{{end}},
	POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_violation 
};
{{end}}

//Inputs to the function {{$block.Name}}
typedef struct {
	{{range $index, $var := $block.InputVars}}{{$var.Type}} {{$var.Name}}{{if $var.ArraySize}}[{{$var.ArraySize}}]{{end}};
	{{end}}
} inputs_{{$block.Name}}_t;

#define BYTES_PER_INPUT {{(len $block.InputVars)}} // bool stored as bytes, so num outputs x 1

//Outputs from the function {{$block.Name}}
typedef struct {
	{{range $index, $var := $block.OutputVars}}{{$var.Type}} {{$var.Name}}{{if $var.ArraySize}}[{{$var.ArraySize}}]{{end}};
	{{end}}
} outputs_{{$block.Name}}_t;

#define BYTES_PER_OUTPUT {{(len $block.OutputVars)}} // bool stored as bytes, so num outputs x 1

//enforcer state and vars:
typedef struct {
	{{range $polI, $pol := $block.Policies}}enum {{$block.Name}}_policy_{{$pol.Name}}_states _policy_{{$pol.Name}}_state;
	{{$pfbEnf := getPolicyEnfInfo $block $polI}}{{if not $pfbEnf}}//Policy is broken!{{else}}//internal vars
	{{range $vari, $var := $pfbEnf.OutputPolicy.InternalVars}}{{if not $var.Constant}}{{$var.Type}} {{$var.Name}}{{if $var.ArraySize}}[{{$var.ArraySize}}]{{end}};
	{{end}}{{end}}
	{{end}}
	{{end}}
} enforcervars_{{$block.Name}}_t;

{{range $polI, $pol := $block.Policies -}}
{{range $varI, $var := $pol.InternalVars -}}
{{if $var.Constant}}#define CONST_{{$pol.Name}}_{{$var.Name}} {{$var.InitialValue}}{{end}}
{{end}}{{end}}

//This function is provided in "series_F_{{$block.Name}}.c"
//It sets up the variable structures to their initial values
void {{$block.Name}}_init_all_vars(enforcervars_{{$block.Name}}_t* me, inputs_{{$block.Name}}_t* inputs, outputs_{{$block.Name}}_t* outputs);

//This function is provided in "series_F_{{$block.Name}}.c"
//It will run the synthesised enforcer and call the controller function
void {{$block.Name}}_run_via_enforcer(enforcervars_{{$block.Name}}_t* me, inputs_{{$block.Name}}_t* inputs, outputs_{{$block.Name}}_t* outputs);

//This function is provided from the user
//It is the controller function
extern void {{$block.Name}}_run(inputs_{{$block.Name}}_t* inputs, outputs_{{$block.Name}}_t* outputs);

//enforcer functions

{{range $polI, $pol := $block.Policies}}
//This function is provided in "series_F_{{$block.Name}}.c"
//It will run the input enforcer for {{$block.Name}}'s policy {{$pol.Name}}
void {{$block.Name}}_run_input_enforcer_{{$pol.Name}}(enforcervars_{{$block.Name}}_t* me, inputs_{{$block.Name}}_t* inputs, uint8_t* newAcceptableInputsBool);

//This function is provided in "series_F_{{$block.Name}}.c"
//It will run the output enforcer for {{$block.Name}}'s policy {{$pol.Name}}
void {{$block.Name}}_run_output_enforcer_{{$pol.Name}}(enforcervars_{{$block.Name}}_t* me, inputs_{{$block.Name}}_t* inputs, outputs_{{$block.Name}}_t* outputs, uint8_t* newAcceptableOutputsBool);


//It will run update the location of output enforcer for {{$block.Name}}'s policy {{$pol.Name}}
void {{$block.Name}}_transition_output_enforcer_{{$pol.Name}}(enforcervars_{{$block.Name}}_t* me, inputs_{{$block.Name}}_t* inputs, outputs_{{$block.Name}}_t* outputs);


//This function is provided in "series_F_{{$block.Name}}.c"
//It will check the state of the enforcer monitor code
//It returns one of the following:
//0: currently true (safe)
//1: always true (safe)
//-1: currently false (unsafe)
//-2: always false (unsafe)
//It will need to do some reachability analysis to achieve this
int {{$block.Name}}_check_rv_status_{{$pol.Name}}(enforcervars_{{$block.Name}}_t* me);

{{end}}
// Select function will select the FIRST acceptable input contained in inputOptions, while respecting transparency (will not edit if not required)
void select_input(inputs_{{$block.Name}}_t* inputs, inputs_{{$block.Name}}_t* inputOptions);
void select_input_new(inputs_{{$block.Name}}_t* uneditedInputs, uint8_t* newAcceptableInputsBool, inputs_{{$block.Name}}_t* possibleInputs);
// Select function will select the FIRST acceptable output contained in outputOptions, while respecting transparency (will not edit if not required)
void select_output(outputs_{{$block.Name}}_t* outputs, outputs_{{$block.Name}}_t* outputOptions);
void select_output_new(outputs_{{$block.Name}}_t* uneditedOutputs, uint8_t* newAcceptableOutputsBool, outputs_{{$block.Name}}_t* possibleOutputs);

{{end}}

{{define "functionC"}}{{$block := index .Functions .FunctionIndex}}{{$blocks := .Functions}}
//This file should be called series_F_{{$block.Name}}.c
//This is autogenerated code. Edit by hand at your peril!
#include "series_F_{{$block.Name}}.h"

inputs_{{$block.Name}}_t possibleInputs[INPUT_OPTIONS] = { {{getBinaryCombinations (len $block.InputVars)}}
};
uint8_t inputAcceptable[INPUT_OPTIONS] = { {{getTrueFor (len $block.InputVars)}}
};

outputs_{{$block.Name}}_t possibleOutputs[OUTPUT_OPTIONS] = { {{getBinaryCombinations (len $block.OutputVars)}}
};
uint8_t outputAcceptable[OUTPUT_OPTIONS] = { {{getTrueFor (len $block.OutputVars)}}
};

const inputs_{{$block.Name}}_t unacceptableInput = {{getBADString (len $block.InputVars) }}
const outputs_{{$block.Name}}_t unacceptableOutput = {{getBADString (len $block.OutputVars) }}

void {{$block.Name}}_init_all_vars(enforcervars_{{$block.Name}}_t* me, inputs_{{$block.Name}}_t* inputs, outputs_{{$block.Name}}_t* outputs) {
	//set any input vars with default values
	{{range $index, $var := $block.InputVars}}{{if $var.InitialValue}}{{$initialArray := $var.GetInitialArray}}{{if $initialArray}}{{range $initialIndex, $initialValue := $initialArray}}inputs->{{$var.Name}}[{{$initialIndex}}] = {{$initialValue}};
	{{end}}{{else}}inputs->{{$var.Name}} = {{$var.InitialValue}};
	{{end}}{{end}}{{end}}
	//set any output vars with default values
	{{range $index, $var := $block.OutputVars}}{{if $var.InitialValue}}{{$initialArray := $var.GetInitialArray}}{{if $initialArray}}{{range $initialIndex, $initialValue := $initialArray}}outputs->{{$var.Name}}[{{$initialIndex}}] = {{$initialValue}};
	{{end}}{{else}}outputs->{{$var.Name}} = {{$var.InitialValue}};
	{{end}}{{end}}{{end}}

	{{if $block.Policies}}{{range $polI, $pol := $block.Policies}}
	me->_policy_{{$pol.Name}}_state = {{if $pol.States}}POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_{{(index $pol.States 0).Name}}{{else}}POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_unknown{{end}};
	{{$pfbEnf := getPolicyEnfInfo $block $polI}}{{if not $pfbEnf}}//Policy is broken!{{else}}//input policy internal vars
	{{range $vari, $var := $pfbEnf.OutputPolicy.InternalVars}}{{if not $var.Constant}}
	{{$initialArray := $var.GetInitialArray}}{{if $initialArray}}{{range $initialIndex, $initialValue := $initialArray}}me->{{$var.Name}}[{{$initialIndex}}] = {{$initialValue}};
	{{end}}{{else}}me->{{$var.Name}} = {{if $var.InitialValue}}{{$var.InitialValue}}{{else}}0{{end}};
	{{end}}{{end}}{{end}}{{end}}
	{{end}}{{end}}
}

void {{$block.Name}}_run_via_enforcer(enforcervars_{{$block.Name}}_t* me, inputs_{{$block.Name}}_t* inputs, outputs_{{$block.Name}}_t* outputs) {
	// Prepare set of all possible inputs (as all inputs are acceptable until enforcers reduce this set)
	// inputs_{{$block.Name}}_t acceptableInputs[INPUT_OPTIONS];
	// memcpy(acceptableInputs, possibleInputs, INPUT_OPTIONS*BYTES_PER_INPUT);
	// Same for all possible outputs..
	// outputs_{{$block.Name}}_t acceptableOutputs[OUTPUT_OPTIONS];
	// memcpy(acceptableOutputs, possibleOutputs, OUTPUT_OPTIONS*BYTES_PER_OUTPUT);

	// New approach, 1xOPTIONS array holding boolean of if this i/o combo is acceptable or not.
	uint8_t newAcceptableInputsBool[INPUT_OPTIONS];
	memcpy(newAcceptableInputsBool, inputAcceptable, INPUT_OPTIONS*sizeof(uint8_t));

	uint8_t newAcceptableOutputsBool[OUTPUT_OPTIONS];
	memcpy(newAcceptableOutputsBool, outputAcceptable, OUTPUT_OPTIONS*sizeof(uint8_t));

	// Run input enforcers
	// 	acceptableInputs is the set that gets reduced by each enforcer	{{$lpol := len $block.Policies}}
	{{range $polI, $pol_unused := $block.Policies}}{{$acti := sub (sub $lpol 1) $polI}}{{$pol := index $block.Policies $acti}} {{$block.Name}}_run_input_enforcer_{{$pol.Name}}(me, inputs, newAcceptableInputsBool);
	{{end}}
	// Select input
	// select_input(inputs, acceptableInputs);
	select_input_new(inputs, newAcceptableInputsBool, possibleInputs);
	
	{{$block.Name}}_run(inputs, outputs);
	
	// Run output enforcers
	{{range $polI, $pol := $block.Policies}}{{$block.Name}}_run_output_enforcer_{{$pol.Name}}(me, inputs, outputs, newAcceptableOutputsBool);
	{{end}}
	// Select output
	// select_output(outputs, acceptableOutputs);
	select_output_new(outputs, newAcceptableOutputsBool, possibleOutputs);

	// Update output enforcer locations
	{{range $polI, $pol := $block.Policies}}{{$block.Name}}_transition_output_enforcer_{{$pol.Name}}(me, inputs, outputs);
	{{end}}
}


bool compareInputs(const inputs_{{$block.Name}}_t* inp1, const inputs_{{$block.Name}}_t* inp2) {
	return ({{range $index, $var := $block.InputVars}}(inp1->{{$var.Name}} == inp2->{{$var.Name}}){{if $var.ArraySize}}[{{$var.ArraySize}}]{{end}} {{if isNotLastElement $index $block.InputVars}}&&{{end}} {{end}});
}

bool isBadInput(const inputs_{{$block.Name}}_t* inputs) {
	const inputs_{{$block.Name}}_t* badInput = &unacceptableInput;
	return compareInputs(inputs, badInput);
}

bool compareOutputs(const outputs_{{$block.Name}}_t* out1, const outputs_{{$block.Name}}_t* out2) {
	return ({{range $index, $var := $block.OutputVars}}(out1->{{$var.Name}} == out2->{{$var.Name}}){{if $var.ArraySize}}[{{$var.ArraySize}}]{{end}} {{if isNotLastElement $index $block.OutputVars}}&&{{end}} {{end}});
}

bool isBadOutput(const outputs_{{$block.Name}}_t* outputs) {
	const outputs_{{$block.Name}}_t* badOutput = &unacceptableOutput;
	return compareOutputs(outputs, badOutput);
}

// Input Intersection
void unacceptable_input_option_intersection(uint8_t* newAcceptableInputsBool, const uint16_t numNewUncccept, const inputs_{{$block.Name}}_t* uneditedInputs) {
	// for every new unacceptable option
	for (uint16_t j = 0; j < numNewUncccept; j++) {
		uint16_t unAccptOptnIndex ={{getInputIndex $block.InputVars true}};
		newAcceptableInputsBool[unAccptOptnIndex] = false;
	}
}

// Output Intersection
void unacceptable_output_option_intersection(uint8_t* newAcceptableOutputsBool, const uint16_t numNewUncccept, const outputs_{{$block.Name}}_t* uneditedOutputs) {
	// for every new unacceptable option
	for (uint16_t j = 0; j < numNewUncccept; j++) {
		uint16_t unAccptOptnIndex ={{getOutputIndex $block.OutputVars true}};
		newAcceptableOutputsBool[unAccptOptnIndex] = false;
	}
}

// Select Input
void select_input(inputs_{{$block.Name}}_t* inputs, inputs_{{$block.Name}}_t* inputOptions) {
	// TODO: replace with min-edit/min-select/rand-select
	for (uint16_t i = 0; i < INPUT_OPTIONS; i++) {
		if (
			// (inputOptions[i].A != BAD)
			!isBadInput(&inputOptions[i])
			) {
			if (
				compareInputs(&inputOptions[i], inputs)
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
			// printf("Input edited\n");
			memcpy(inputs, &inputOptions[i], BYTES_PER_INPUT);
			return;
		}
	}
}

// Select Input
void select_input_new(inputs_{{$block.Name}}_t* uneditedInputs, uint8_t* newAcceptableInputsBool, inputs_{{$block.Name}}_t* possibleInputs) {
	// Check if current input is accepting
	uint16_t uneditedInputsIndex ={{getInputIndex $block.InputVars false}};
	if (newAcceptableInputsBool[uneditedInputsIndex] == true) {
		// If accepting return unchanged
		// printf("NEW APPROACH - No input edit needed\n");
		return;
	}

	// If we get here we know we need to edit
	// for each in newAcceptableInputsBool
	for (uint16_t i = 0; i < OUTPUT_OPTIONS; i++) {
		// if accepting
		if (newAcceptableInputsBool[i] == true){
			// look up what it should be in possibleInputs
			// set inputs to possibleInputs
			memcpy(uneditedInputs, &possibleInputs[i], BYTES_PER_INPUT);
			return;
		}
	}

	// assert that we dont get here 
	// (as this would mean we didnt find accepting or edit to an accepting and something is wrong)
	assert(false && "We must either perform no edit (transparency) or find a suitable edit."); 
}

// Select Output
void select_output(outputs_{{$block.Name}}_t* outputs, outputs_{{$block.Name}}_t* outputOptions) {
	// TODO: replace with min-edit/min-select/rand-select
	for (uint16_t i = 0; i < OUTPUT_OPTIONS; i++) {
		if (
			!isBadOutput(&outputOptions[i])
			) {
			if (
				compareOutputs(&outputOptions[i], outputs)
				) {
				// printf("No output edit needed\n");
				return;
			}
		}
	}

	for (uint16_t i = 0; i < OUTPUT_OPTIONS; i++) {
		// printf("i: %d\n", i);
		if (
			!isBadOutput(&outputOptions[i])
			) {
			// printf("Output Edited\n");
			memcpy(outputs, &outputOptions[i], BYTES_PER_OUTPUT);
			return;
		}
	}
}

// Select Output
void select_output_new(outputs_{{$block.Name}}_t* uneditedOutputs, uint8_t* newAcceptableOutputsBool, outputs_{{$block.Name}}_t* possibleOutputs) {
	// Check if current output is accepting
	uint16_t uneditedOutputsIndex ={{getOutputIndex $block.OutputVars false}};
	if (newAcceptableOutputsBool[uneditedOutputsIndex] == true) {
		// If accepting return unchanged
		// printf("NEW APPROACH - No output edit needed\n");
		return;
	}

	// If we get here we know we need to edit
	// for each in newAcceptableOutputsBool
	for (uint16_t i = 0; i < OUTPUT_OPTIONS; i++) {
		// if accepting
		if (newAcceptableOutputsBool[i] == true){
			// look up what it should be in possibleOutputs
			// set outputs to possibleOutputs
			memcpy(uneditedOutputs, &possibleOutputs[i], BYTES_PER_OUTPUT);

			return;
		}
	}

	// assert that we dont get here 
	// (as this would mean we didnt find accepting or edit to an accepting and something is wrong)
	assert(false && "We must either perform no edit (transparency) or find a suitable edit."); 
}
{{if $block.Policies}}{{template "_policyIn" $block}}{{end}}

{{if $block.Policies}}{{template "_policyOut" $block}}{{end}}

//This function is provided in "series_F_{{$block.Name}}.c"
//It will check the state of the enforcer monitor code
//It returns one of the following:
//0: currently true (safe)
//1: always true (safe)
//-1: currently false (unsafe)
//-2: always false (unsafe)
//It will need to do some reachability analysis to achieve this
{{/*int {{$block.Name}}_check_rv_status_{{$pol.Name}}(enforcervars_{{$block.Name}}_t* me) {
	//TODO: this function
	return -2;
}*/}}
{{end}}
{{define "mainCBMCC"}}{{$block := index .Functions .FunctionIndex}}{{$blocks := .Functions}}
//This file should be called cbmc_main_{{$block.Name}}.c
//This is autogenerated code. Edit by hand at your peril!

//It can be used with the cbmc model checker
//Call it using the following command: 
//$ cbmc cbmc_main_{{$block.Name}}.c{{range $blockI, $block := $blocks}} series_F_{{$block.Name}}.c{{end}}

{{range $blockI, $block := $blocks}}
#include "series_F_{{$block.Name}}.h"{{end}}
#include <stdio.h>
#include <stdint.h>

int main() {

{{range $blockI, $block := $blocks}}
	//I/O and state for {{$block.Name}}
	enforcervars_{{$block.Name}}_t enf_{{$block.Name}};
    inputs_{{$block.Name}}_t inputs_{{$block.Name}};
    outputs_{{$block.Name}}_t outputs_{{$block.Name}};

	//set values to known state
    {{$block.Name}}_init_all_vars(&enf_{{$block.Name}}, &inputs_{{$block.Name}}, &outputs_{{$block.Name}});

	//introduce nondeterminism
    //a nondet_xxxxx function name tells cbmc that it could be anything, but must be unique
    //randomise inputs
	{{range $inputI, $input := $block.InputVars -}}
	inputs_{{$block.Name}}.{{$input.Name}} = nondet_{{$block.Name}}_input_{{$inputI}}();
	{{end}}

	//randomise enforcer state, i.e. clock values and position (excepting violation state)
	{{range $polI, $pol := $block.Policies -}}
	{{range $varI, $var := $pol.InternalVars -}}
	{{if not $var.Constant}}enf_{{$block.Name}}.{{$var.Name}} = nondet_{{$block.Name}}_enf_{{$pol.Name}}_{{$varI}}();{{end}}
	{{end}}
	enf_{{$block.Name}}._policy_{{$pol.Name}}_state = nondet_{{$block.Name}}_enf_{{$pol.Name}}_state() % {{len $pol.States}};
	{{end}}

    //run the enforcer (i.e. tell CBMC to check this out)
	{{$block.Name}}_run_via_enforcer(&enf_{{$block.Name}}, &inputs_{{$block.Name}}, &outputs_{{$block.Name}});
{{end}}
}

{{range $blockI, $block := $blocks}}
void {{$block.Name}}_run(inputs_{{$block.Name}}_t *inputs, outputs_{{$block.Name}}_t *outputs) {
    //randomise controller

    {{range $outputI, $output := $block.OutputVars -}}
	outputs->{{$output.Name}} = nondet_{{$block.Name}}_output_{{$outputI}}();
	{{end}} 
}
{{end}}
{{end}}
`

var cSeriesCompositionTemplateFuncMap = template.FuncMap{
	"getCECCTransitionCondition": getCECCTransitionCondition,

	"getPolicyEnfInfo": getPolicyEnfInfo,

	"compileExpression": stconverter.CCompileExpression,

	"sub":   sub,
	"times": times,

	"twoToThePower":          twoToThePower,
	"getBinaryCombinations":  getBinaryCombinations,
	"getBADString":           getBADString,
	"getTrueFor":             getTrueFor,
	"getAcceptableOptions":   getAcceptableOptions,
	"getUnacceptableOptions": getUnacceptableOptions,
	"isNotLastElement":       isNotLastElement,
	"getInputIndex":          getInputIndex,
	"getOutputIndex":         getOutputIndex,
}

var cSeriesCompositionTemplates = template.Must(template.New("").Funcs(cSeriesCompositionTemplateFuncMap).Parse(rtecSeriesCompositionTemplate))
