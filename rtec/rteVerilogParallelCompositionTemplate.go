package rtec

import (
	"text/template"

	"github.com/PRETgroup/goFB/goFB/stconverter"
)

const rteVerilogParallelCompositionTemplate = `
{{define "functionVerilog"}}{{$block := index .Functions .FunctionIndex}}{{$blocks := .Functions}}{{$pbfPolicies := getAllPolicyEnfInfo $block}}
//This file should be called F_{{$block.Name}}.sv
//This is autogenerated code. Edit by hand at your peril!!!

//Warning: This is experimental parallel composition code.

{{range $polI, $pol := $block.Policies}}
	module F_combinatorialVerilog_{{$block.Name}}_policy_{{$pol.Name}} (
		input wire clk,

		//inputs (plant to controller){{range $index, $var := $block.InputVars}}
		input wire {{getVerilogWidthArrayForType $var.Type}} {{$var.Name}}_ptc_in,
		output reg {{getVerilogWidthArrayForType $var.Type}} {{$var.Name}}_ptc_out,
		output reg {{getVerilogWidthArrayForType $var.Type}} {{$var.Name}}_dont_care,
		{{end}}
		//outputs (controller to plant){{range $index, $var := $block.OutputVars}}
		input wire {{getVerilogWidthArrayForType $var.Type}} {{$var.Name}}_ctp_in,
		output reg {{getVerilogWidthArrayForType $var.Type}} {{$var.Name}}_ctp_out,
		output reg {{getVerilogWidthArrayForType $var.Type}} {{$var.Name}}_dont_care,
		{{end}}

		//helpful internal variable outputs {{range $vari, $var := $pol.InternalVars}}{{if not $var.Constant}}
		output wire {{getVerilogWidthArrayForType $var.Type}} {{$var.Name}}_out,
		{{end}}{{end}}
		
		//helpful state output input var{{if $polI}},
		{{end}}
		output reg {{getVerilogWidthArray (add (len $pol.States) 1)}} {{$block.Name}}_policy_{{$pol.Name}}_state_out

		);


		//For each policy, we need define types for the state machines
		localparam
			{{if len $pol.States}}{{range $index, $state := $pol.States}}` + `POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_{{$state}} = {{$index}},
			{{end}}{{else}}POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_unknown 0 {{end}}` + `POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_violation = {{if len $pol.States}}{{len $pol.States}};{{else}}1;{{end}}

		//For each policy, we need a reg for the state machine
		reg {{getVerilogWidthArray (add (len $pol.States) 1)}} {{$block.Name}}_policy_{{$pol.Name}}_c_state = 0;
		reg {{getVerilogWidthArray (add (len $pol.States) 1)}} {{$block.Name}}_policy_{{$pol.Name}}_n_state = 0;

		{{range $index, $var := $block.InputVars}}
		{{getVerilogType $var.Type}}{{$var.Name}} {{if $var.InitialValue}}/* = {{$var.InitialValue}}*/{{else}}= 0{{end}};
		{{end}}{{range $index, $var := $block.OutputVars}}
		{{getVerilogType $var.Type}}{{$var.Name}} {{if $var.InitialValue}}/* = {{$var.InitialValue}}*/{{else}}= 0{{end}};
		{{end}}
		{{$pfbEnf := index $pbfPolicies $polI}}{{if not $pfbEnf}}//Policy is broken!{{else}}//internal vars
		{{range $vari, $var := $pfbEnf.OutputPolicy.InternalVars}}{{if $var.Constant}}wire {{getVerilogWidthArrayForType $var.Type}} {{$var.Name}} = {{$var.InitialValue}}{{else}}{{getVerilogType $var.Type}} {{$var.Name}}{{if $var.InitialValue}}/* = {{$var.InitialValue}}*/{{end}}{{end}};
		{{end}}{{end}}

		initial begin{{range $index, $var := $block.InputVars}}
			{{getVerilogWidthArrayForType $var.Type}} {{$var.Name}}_ptc_out = 0;
			{{getVerilogWidthArrayForType $var.Type}} {{$var.Name}}_dont_care = 0;
		{{end}}{{range $index, $var := $block.OutputVars}}
			{{getVerilogWidthArrayForType $var.Type}} {{$var.Name}}_ctp_out = 0;
			{{getVerilogWidthArrayForType $var.Type}} {{$var.Name}}_dont_care = 0;
		{{end}}
		end

		always @(posedge clk)
		begin
			{{$block.Name}}_policy_{{$pol.Name}}_c_state = {{$block.Name}}_policy_{{$pol.Name}}_n_state;
		end

		always @({{range $index, $var := $block.InputVars}}{{$var.Name}}_ptc_in, {{end}}{{range $index, $var := $block.OutputVars}}{{$var.Name}}_ctp_in {{if equal $index (subtract (len $block.OutputVars) 1)}}{{else}},{{end}}{{end}}) begin
			// Default no location change
			{{$block.Name}}_policy_{{$pol.Name}}_n_state = {{$block.Name}}_policy_{{$pol.Name}}_c_state;
			
			// Default no change to inputs/outputs (transparency) {{range $index, $var := $block.InputVars}}
			{{$var.Name}} = {{$var.Name}}_ptc_in;
			{{end}}
			{{range $index, $var := $block.OutputVars}}{{$var.Name}} = {{$var.Name}}_ctp_in;
			{{end}}

			//capture time inputs
			{{$pfbEnf := index $pbfPolicies $polI}}{{if not $pfbEnf}}//Policy is broken!{{else}}//internal vars
			{{range $vari, $var := $pfbEnf.OutputPolicy.InternalVars}}{{if not $var.Constant}}{{$var.Name}} = {{$var.Name}}_in{{if $var.IsDTimer}} + 1{{end}};{{end}}
			{{end}}{{end}}

			{{if $block.Policies}}//input policies
			{{$pfbEnf := index $pbfPolicies $polI}}
			{{if not $pfbEnf}}//{{$pol.Name}} is broken!
			{{else}}{{/* this is where the policy comes in */}}	//INPUT POLICY {{$pol.Name}} BEGIN 
				case({{$block.Name}}_policy_{{$pol.Name}}_c_state)
					{{range $sti, $st := $pol.States}}` + `POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_{{$st.Name}}: begin
						{{range $tri, $tr := $pfbEnf.InputPolicy.GetViolationTransitions}}{{if eq $tr.Source $st.Name}}{{/*
						*/}}
						if ({{$cond := getVerilogECCTransitionCondition $block (compileExpression $tr.STGuard)}}{{$cond.IfCond}}) begin
							//transition {{$tr.Source}} -> {{$tr.Destination}} on {{$tr.Condition}}
							//select a transition to solve the problem
							{{$solution := $pfbEnf.SolveViolationTransition $tr true}}
							{{if $solution.Comment}}//{{$solution.Comment}}{{end}}
							{{range $soleI, $sole := $solution.Expressions}}{{$sol := getVerilogECCTransitionCondition $block (compileExpression $sole)}}{{$sol.IfCond}};
							{{end}}
						end{{end}}{{end}}
					end
					{{end}}
				endcase
			{{end}}
			//INPUT POLICY {{$pol.Name}} END{{end}}

			//output policies
			{{if not $pfbEnf}}//{{$pol.Name}} is broken!
			{{else}}{{/* this is where the policy comes in */}}	//OUTPUT POLICY {{$pol.Name}} BEGIN 
				
				case({{$block.Name}}_policy_{{$pol.Name}}_c_state)
					{{range $sti, $st := $pol.States}}` + `POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_{{$st.Name}}: begin
						{{range $tri, $tr := $pfbEnf.OutputPolicy.GetViolationTransitions}}{{if eq $tr.Source $st.Name}}{{/*
						*/}}
						if ({{$cond := getVerilogECCTransitionCondition $block (compileExpression $tr.STGuard)}}{{$cond.IfCond}}) begin
							//transition {{$tr.Source}} -> {{$tr.Destination}} on {{$tr.Condition}}
							//select a transition to solve the problem
							{{$solution := $pfbEnf.SolveViolationTransition $tr false}}
							{{if $solution.Comment}}//{{$solution.Comment}}{{end}}
							{{range $soleI, $sole := $solution.Expressions}}{{$sol := getVerilogECCTransitionCondition $block (compileExpression $sole)}}{{$sol.IfCond}};
							{{end}}
						end {{end}}{{end}}
					end
					{{end}}
				endcase

				//transTaken_{{$block.Name}}_policy_{{$pol.Name}} = 0;
				//select transition to advance state
				case({{$block.Name}}_policy_{{$pol.Name}}_c_state)
					{{range $sti, $st := $pol.States}}` + `POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_{{$st.Name}}: begin
						{{range $tri, $tr := $pfbEnf.OutputPolicy.GetTransitionsForSource $st.Name}}{{/*
						*/}}
						{{if $tri}}else {{end}}if ({{$cond := getVerilogECCTransitionCondition $block (compileExpression $tr.STGuard)}}{{$cond.IfCond}}) begin
							//transition {{$tr.Source}} -> {{$tr.Destination}} on {{$tr.Condition}}
							{{$block.Name}}_policy_{{$pol.Name}}_n_state = ` + `POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_{{$tr.Destination}};
							//set expressions
							{{range $exi, $ex := $tr.Expressions}}
							{{$ex.VarName}} = {{$ex.Value}};{{end}}
							//transTaken_{{$block.Name}}_policy_{{$pol.Name}} = 1;
						end {{end}} else begin
							//only possible in a violation
							{{$block.Name}}_policy_{{$pol.Name}}_n_state = ` + `POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_violation;
							//transTaken_{{$block.Name}}_policy_{{$pol.Name}} = 1;
						end
					end
					{{end}}
					default begin
						//if we are here, we're in the violation state
						//the violation state permanently stays in violation
						{{$block.Name}}_policy_{{$pol.Name}}_n_state = ` + `POLICY_STATE_{{$block.Name}}_{{$pol.Name}}_violation;
						//transTaken_{{$block.Name}}_policy_{{$pol.Name}} = 1;
					end
				endcase
			{{end}}
			//OUTPUT POLICY {{$pol.Name}} END

			// Ground all signals the enf doesnt care about
			// Inputs{{range $index, $var := $block.InputVars}}
			if ({{$var.Name}} === {{$var.Name}}_ptc_in) begin
				{{$var.Name}}_dont_care <= 1;
				{{$var.Name}} = 0;
			end else begin
				{{$var.Name}}_dont_care <= 0;
			end
			{{end}}
			// Outputs{{range $index, $var := $block.OutputVars}}
			if ({{$var.Name}} === {{$var.Name}}_ctp_in) begin
				{{$var.Name}}_dont_care <= 1;
				{{$var.Name}} = 0;
			end else begin
				{{$var.Name}}_dont_care <= 0;
			end
			{{end}}
			// Post input enforced 
			{{range $index, $var := $block.InputVars}}{{$var.Name}}_ptc_out = {{$var.Name}};
			{{end}}
			// Post output enforced 
			{{range $index, $var := $block.OutputVars}}{{$var.Name}}_ctp_out = {{$var.Name}};
			{{end}}
		end

		//emit state/time inputs
		assign {{$block.Name}}_policy_{{$pol.Name}}_state_out =  {{$block.Name}}_policy_{{$pol.Name}}_c_state;
		{{$pfbEnf := index $pbfPolicies $polI}}{{if not $pfbEnf}}//Policy is broken!{{else}}//internal vars
		{{range $vari, $var := $pfbEnf.OutputPolicy.InternalVars}}{{if not $var.Constant}}assign {{$var.Name}}_out = {{$var.Name}};{{end}}
		{{end}}{{end}}

	endmodule

	{{end}}
{{end}}
`

var verilogParallelCompositionTemplateFuncMap = template.FuncMap{
	"getVerilogECCTransitionCondition": getVerilogECCTransitionCondition,
	"getVerilogType":                   getVerilogType,
	"getPolicyEnfInfo":                 getPolicyEnfInfo,
	"getAllPolicyEnfInfo":              getAllPolicyEnfInfo,
	"getVerilogWidthArray":             getVerilogWidthArray,
	"getVerilogWidthArrayForType":      getVerilogWidthArrayForType,
	"add1IfClock":                      add1IfClock,

	"compileExpression": stconverter.VerilogCompileExpression,

	"add":      add,
	"subtract": subtract,
	"equal":    equal,
}

var verilogParallelCompositionTemplates = template.Must(template.New("").Funcs(verilogParallelCompositionTemplateFuncMap).Parse(rteVerilogParallelCompositionTemplate))
