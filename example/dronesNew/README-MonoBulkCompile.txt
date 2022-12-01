================
    Mono
================

== O0 ==
time gcc example/dronesNew/mono_pb_main.c example/dronesNew/F_pb.c example/dronesNew/save_results.c -o example_dronesNew_pb_mono
time gcc example/dronesNew/mono_pb_and_ps_main.c example/dronesNew/F_pb_and_ps.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_mono
time gcc example/dronesNew/mono_pb_and_ps_and_pj_main.c example/dronesNew/F_pb_and_ps_and_pj.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_mono
time gcc example/dronesNew/mono_pb_and_ps_and_pj_and_pr_main.c example/dronesNew/F_pb_and_ps_and_pj_and_pr.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_mono
time gcc example/dronesNew/mono_pb_and_ps_and_pj_and_pr_and_pb2_main.c example/dronesNew/F_pb_and_ps_and_pj_and_pr_and_pb2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_mono

./example_dronesNew_pb_mono
./example_dronesNew_pb_and_ps_mono
./example_dronesNew_pb_and_ps_and_pj_mono
./example_dronesNew_pb_and_ps_and_pj_and_pr_mono
./example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_mono

== O1 ==
time gcc -O1 example/dronesNew/mono_pb_main.c example/dronesNew/F_pb.c example/dronesNew/save_results.c -o example_dronesNew_pb_mono
time gcc -O1 example/dronesNew/mono_pb_and_ps_main.c example/dronesNew/F_pb_and_ps.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_mono
time gcc -O1 example/dronesNew/mono_pb_and_ps_and_pj_main.c example/dronesNew/F_pb_and_ps_and_pj.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_mono
time gcc -O1 example/dronesNew/mono_pb_and_ps_and_pj_and_pr_main.c example/dronesNew/F_pb_and_ps_and_pj_and_pr.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_mono
time gcc -O1 example/dronesNew/mono_pb_and_ps_and_pj_and_pr_and_pb2_main.c example/dronesNew/F_pb_and_ps_and_pj_and_pr_and_pb2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_mono

== O2 ==
time gcc -O2 example/dronesNew/mono_pb_main.c example/dronesNew/F_pb.c example/dronesNew/save_results.c -o example_dronesNew_pb_mono
time gcc -O2 example/dronesNew/mono_pb_and_ps_main.c example/dronesNew/F_pb_and_ps.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_mono
time gcc -O2 example/dronesNew/mono_pb_and_ps_and_pj_main.c example/dronesNew/F_pb_and_ps_and_pj.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_mono
time gcc -O2 example/dronesNew/mono_pb_and_ps_and_pj_and_pr_main.c example/dronesNew/F_pb_and_ps_and_pj_and_pr.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_mono
time gcc -O2 example/dronesNew/mono_pb_and_ps_and_pj_and_pr_and_pb2_main.c example/dronesNew/F_pb_and_ps_and_pj_and_pr_and_pb2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_mono


== O3 ==
time gcc -O3 example/dronesNew/mono_pb_main.c example/dronesNew/F_pb.c example/dronesNew/save_results.c -o example_dronesNew_pb_mono
time gcc -O3 example/dronesNew/mono_pb_and_ps_main.c example/dronesNew/F_pb_and_ps.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_mono
time gcc -O3 example/dronesNew/mono_pb_and_ps_and_pj_main.c example/dronesNew/F_pb_and_ps_and_pj.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_mono
time gcc -O3 example/dronesNew/mono_pb_and_ps_and_pj_and_pr_main.c example/dronesNew/F_pb_and_ps_and_pj_and_pr.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_mono
time gcc -O3 example/dronesNew/mono_pb_and_ps_and_pj_and_pr_and_pb2_main.c example/dronesNew/F_pb_and_ps_and_pj_and_pr_and_pb2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_mono


== Os ==
time gcc -Os example/dronesNew/mono_pb_main.c example/dronesNew/F_pb.c example/dronesNew/save_results.c -o example_dronesNew_pb_mono
time gcc -Os example/dronesNew/mono_pb_and_ps_main.c example/dronesNew/F_pb_and_ps.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_mono
time gcc -Os example/dronesNew/mono_pb_and_ps_and_pj_main.c example/dronesNew/F_pb_and_ps_and_pj.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_mono
time gcc -Os example/dronesNew/mono_pb_and_ps_and_pj_and_pr_main.c example/dronesNew/F_pb_and_ps_and_pj_and_pr.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_mono
time gcc -Os example/dronesNew/mono_pb_and_ps_and_pj_and_pr_and_pb2_main.c example/dronesNew/F_pb_and_ps_and_pj_and_pr_and_pb2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_mono
