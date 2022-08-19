
import os
import shutil
import datetime

#
# Expects directory structure as follows
# easy-rte-composition/example/pacemaker/monolithic/p1/synthesis_F_p1.sv 
#                                       /monolithic/p1_and_p2
#                                       /monolithic/p1_and_p2...
#                                       /parallel/p1
#                                       /parallel/p1_and_p2..
#
# You will need to compile the Verilog using easy-rte in the pacemaker directory, then shift these into the folders above
#   make verilog_enf PROJECT=pacemaker FILE=p1_and_p2 PARSEARGS=-product COMPILEARGS=-synthesis
#
# You should run this script from the easy-rte-composition directory
#

easyrte_directory = "/easy-rte-composition/" # eg C:/easy-rte-composition
quartus_directory = "L:/intelFPGA_21.1/quartus/bin64/" # eg L:/intelFPGA_21.1/quartus/bin64/ 

def generate_tcl_mon(fname):
  fp = open(fname + ".tcl", "w")

  string = """
	# Load Quartus Prime Tcl Project package
	package require ::quartus::project

	set need_to_close_project 0
	set make_assignments 1

	# Check that the right project is open
	if {[is_project_open]} {
		if {[string compare $quartus(project) """ + "/" + fname + "/" + """ ]} {
			puts "Project  """+ fname +""" is not open"
			set make_assignments 0
		}
	} else {
		# Only open if not already open
		if {[project_exists  """+ fname +"""]} {
			project_open -revision """+ fname.replace("_", "") +""" """+ fname +"""
		} else {
			project_new -revision """+ fname.replace("_", "") +""" """+ fname +"""
		}
		set need_to_close_project 1
	}

	# Make assignments
	#if {$make_assignments} {
		set_global_assignment -name FAMILY "Cyclone IV E"
		set_global_assignment -name DEVICE EP4CE115F29C7
		set_global_assignment -name TOP_LEVEL_ENTITY """ + fname + """
		set_global_assignment -name ORIGINAL_QUARTUS_VERSION 21.1.0
		set_global_assignment -name PROJECT_CREATION_TIME_DATE " """ + datetime.datetime.now().strftime("%H:%M:%S %B %d, %Y ") + """ "
		set_global_assignment -name LAST_QUARTUS_VERSION "21.1.0 Lite Edition"
		set_global_assignment -name PROJECT_OUTPUT_DIRECTORY output_files
		set_global_assignment -name MIN_CORE_JUNCTION_TEMP 0
		set_global_assignment -name MAX_CORE_JUNCTION_TEMP 85
		set_global_assignment -name ERROR_CHECK_FREQUENCY_DIVISOR 1
		set_global_assignment -name NOMINAL_CORE_SUPPLY_VOLTAGE 1.2V
		set_global_assignment -name POWER_PRESET_COOLING_SOLUTION "23 MM HEAT SINK WITH 200 LFPM AIRFLOW"
		set_global_assignment -name POWER_BOARD_THERMAL_MODEL "NONE (CONSERVATIVE)"
		set_global_assignment -name EDA_SIMULATION_TOOL "ModelSim (Verilog)"
		set_global_assignment -name EDA_TIME_SCALE "1 ps" -section_id eda_simulation
		set_global_assignment -name EDA_RUN_TOOL_AUTOMATICALLY OFF -section_id eda_simulation
		set_global_assignment -name EDA_OUTPUT_DATA_FORMAT "VERILOG HDL" -section_id eda_simulation
		set_global_assignment -name EDA_GENERATE_FUNCTIONAL_NETLIST OFF -section_id eda_board_design_timing
		set_global_assignment -name EDA_GENERATE_FUNCTIONAL_NETLIST OFF -section_id eda_board_design_symbol
		set_global_assignment -name EDA_GENERATE_FUNCTIONAL_NETLIST OFF -section_id eda_board_design_signal_integrity
		set_global_assignment -name EDA_GENERATE_FUNCTIONAL_NETLIST OFF -section_id eda_board_design_boundary_scan
		set_global_assignment -name EDA_DESIGN_ENTRY_SYNTHESIS_TOOL "Precision Synthesis"
		set_global_assignment -name EDA_LMF_FILE mentor.lmf -section_id eda_design_synthesis
		set_global_assignment -name EDA_INPUT_DATA_FORMAT VQM -section_id eda_design_synthesis
		set_global_assignment -name ALLOW_REGISTER_RETIMING ON
		set_global_assignment -name PARTITION_NETLIST_TYPE SOURCE -section_id Top
		set_global_assignment -name PARTITION_FITTER_PRESERVATION_LEVEL PLACEMENT_AND_ROUTING -section_id Top
		set_global_assignment -name PARTITION_COLOR 16764057 -section_id Top

		# Commit assignments
		export_assignments

		# Close project
		if {$need_to_close_project} {
			project_close
		}
	#}
	"""
  fp.write(string)
  fp.close()

if __name__ == "__main__":  

	os.chdir("example/pacemaker")
	basedir = "example" 

	composition_methods = ["monolithic", "parallel"]
	policies = ["p1", "p1_and_p2", "p1_and_p2_and_p3", "p1_and_p2_and_p3_and_p4", "p1_and_p2_and_p3_and_p4_and_p5"]
	
	# For debugging purposes
	# composition_methods = ["monolithic"]
	# policies = ["p1"]

	for c in composition_methods:
		for p in policies:
			#direc = basedir +"/pacemaker/"
			#print(direc)
			direc = c + "/" + p
			print(direc)
			os.chdir(direc)
			if c == "monolithic":
				fname = "synthesis_F_" + p
			else:
				fname = "parallel_F_" + p

			print("Running for Monolithic: "+ direc + fname)

			generate_tcl_mon(fname)
			tcl_file = easyrte_directory + "example/pacemaker/" + direc + '/' + fname + ".tcl"
			print(tcl_file)
			top_level_entity = fname
			os.system(quartus_directory + "quartus_sh.exe -t " + tcl_file)
			os.system(quartus_directory + "quartus_map.exe " + top_level_entity)
			try:
				os.system(quartus_directory + "quartus_fit.exe " + top_level_entity)
				os.system(quartus_directory + "quartus_sta.exe " + top_level_entity + " --model=slow")
			except:
				pass
			
			os.chdir(easyrte_directory+ basedir + "/" + "pacemaker")
	os.chdir(easyrte_directory)
	
        
