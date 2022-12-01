================
    Series
================

== XML ==

make c_enf PROJECT=dronesNew FILE=pb  COMPILEARGS=-seriesComposition
make c_enf PROJECT=dronesNew FILE=pb_and_ps  COMPILEARGS=-seriesComposition
make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj  COMPILEARGS=-seriesComposition
make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj_and_pr  COMPILEARGS=-seriesComposition
make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj_and_pr_and_pb2  COMPILEARGS=-seriesComposition
make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj_and_pr_and_pb2_and_ps2  COMPILEARGS=-seriesComposition
make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2  COMPILEARGS=-seriesComposition
make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2  COMPILEARGS=-seriesComposition

== O0 ==

time gcc example/dronesNew/series_pb_main.c example/dronesNew/series_F_pb.c example/dronesNew/save_results.c -o example_dronesNew_pb_series
time gcc example/dronesNew/series_pb_and_ps_main.c example/dronesNew/series_F_pb_and_ps.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_series
time gcc example/dronesNew/series_pb_and_ps_and_pj_main.c example/dronesNew/series_F_pb_and_ps_and_pj.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_series
time gcc example/dronesNew/series_pb_and_ps_and_pj_and_pr_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_series
time gcc example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_series
time gcc example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_series
time gcc example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_series
time gcc example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2_series


== O1 ==

time gcc -O example/dronesNew/series_pb_main.c example/dronesNew/series_F_pb.c example/dronesNew/save_results.c -o example_dronesNew_pb_series
time gcc -O example/dronesNew/series_pb_and_ps_main.c example/dronesNew/series_F_pb_and_ps.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_series
time gcc -O example/dronesNew/series_pb_and_ps_and_pj_main.c example/dronesNew/series_F_pb_and_ps_and_pj.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_series
time gcc -O example/dronesNew/series_pb_and_ps_and_pj_and_pr_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_series
time gcc -O example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_series
time gcc -O example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_series
time gcc -O example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_series
time gcc -O example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2_series

== O2 ==

time gcc -O2 example/dronesNew/series_pb_main.c example/dronesNew/series_F_pb.c example/dronesNew/save_results.c -o example_dronesNew_pb_series
time gcc -O2 example/dronesNew/series_pb_and_ps_main.c example/dronesNew/series_F_pb_and_ps.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_series
time gcc -O2 example/dronesNew/series_pb_and_ps_and_pj_main.c example/dronesNew/series_F_pb_and_ps_and_pj.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_series
time gcc -O2 example/dronesNew/series_pb_and_ps_and_pj_and_pr_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_series
time gcc -O2 example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_series
time gcc -O2 example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_series
time gcc -O2 example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_series
time gcc -O2 example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2_series

./example_dronesNew_pb_series
./example_dronesNew_pb_and_ps_series
./example_dronesNew_pb_and_ps_and_pj_series
./example_dronesNew_pb_and_ps_and_pj_and_pr_series
./example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_series
./example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_series
./example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_series
./example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2_series


== O3 ==

time gcc -O3 example/dronesNew/series_pb_main.c example/dronesNew/series_F_pb.c example/dronesNew/save_results.c -o example_dronesNew_pb_series
time gcc -O3 example/dronesNew/series_pb_and_ps_main.c example/dronesNew/series_F_pb_and_ps.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_series
time gcc -O3 example/dronesNew/series_pb_and_ps_and_pj_main.c example/dronesNew/series_F_pb_and_ps_and_pj.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_series
time gcc -O3 example/dronesNew/series_pb_and_ps_and_pj_and_pr_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_series
time gcc -O3 example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_series
time gcc -O3 example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_series
time gcc -O3 example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_series
time gcc -O3 example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2_series

== Os ==

time gcc -Os example/dronesNew/series_pb_main.c example/dronesNew/series_F_pb.c example/dronesNew/save_results.c -o example_dronesNew_pb_series
time gcc -Os example/dronesNew/series_pb_and_ps_main.c example/dronesNew/series_F_pb_and_ps.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_series
time gcc -Os example/dronesNew/series_pb_and_ps_and_pj_main.c example/dronesNew/series_F_pb_and_ps_and_pj.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_series
time gcc -Os example/dronesNew/series_pb_and_ps_and_pj_and_pr_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_series
time gcc -Os example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_series
time gcc -Os example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_series
time gcc -Os example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_series
time gcc -Os example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2_series