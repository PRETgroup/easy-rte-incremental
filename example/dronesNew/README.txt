Boundary Example

1) Run easy-rte-composition
make c_enf PROJECT=dronesNew FILE=boundary  COMPILEARGS=-seriesComposition

2) Compile series_boundary_main.c, save_results.c, and series_F_boundary.C
gcc example/dronesNew/series_boundary_main.c example/dronesNew/series_F_boundary.c example/dronesNew/save_results.c -o example_dronesNew_boundary_series

3) Run example_dronesNew_boundary_series
./example_dronesNew_boundary_series

make c_enf PROJECT=dronesNew FILE=boundary  PARSEARGS=-product
gcc example/dronesNew/mono_boundary_main.c example/dronesNew/F_boundary.c example/dronesNew/save_results.c -o example_dronesNew_boundary_mono
./example_dronesNew_boundary_mono

Simultaneous
make c_enf PROJECT=dronesNew FILE=simultaneous  COMPILEARGS=-seriesComposition
gcc example/dronesNew/series_simultaneous_main.c example/dronesNew/series_F_simultaneous.c example/dronesNew/save_results.c -o example_dronesNew_simultaneous_series
./example_dronesNew_simultaneous_series

jamming
make c_enf PROJECT=dronesNew FILE=jamming  COMPILEARGS=-seriesComposition
gcc example/dronesNew/series_jamming_main.c example/dronesNew/series_F_jamming.c example/dronesNew/save_results.c -o example_dronesNew_jamming_series
./example_dronesNew_jamming_series
				printf("Ding v1:%d\n", me->v1);


--------------------------------
pb 1
--------------------------------
make c_enf PROJECT=dronesNew FILE=pb  COMPILEARGS=-seriesComposition
gcc -ftime-report example/dronesNew/series_pb_main.c example/dronesNew/series_F_pb.c example/dronesNew/save_results.c -o example_dronesNew_pb_series
./example_dronesNew_pb_series

make c_enf PROJECT=dronesNew FILE=pb PARSEARGS=-product
gcc -ftime-report example/dronesNew/mono_pb_main.c example/dronesNew/F_pb.c example/dronesNew/save_results.c -o example_dronesNew_pb_mono
./example_dronesNew_pb_mono

--------------------------------
pb ps 2
--------------------------------
make c_enf PROJECT=dronesNew FILE=pb_and_ps  COMPILEARGS=-seriesComposition
gcc -ftime-report example/dronesNew/series_pb_and_ps_main.c example/dronesNew/series_F_pb_and_ps.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_series
./example_dronesNew_pb_and_ps_series

make c_enf PROJECT=dronesNew FILE=pb_and_ps PARSEARGS=-product
gcc -ftime-report example/dronesNew/mono_pb_and_ps_main.c example/dronesNew/F_pb_and_ps.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_mono
./example_dronesNew_pb_and_ps_mono

--------------------------------
pb ps pj 3
--------------------------------
make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj  COMPILEARGS=-seriesComposition
gcc -ftime-report example/dronesNew/series_pb_and_ps_and_pj_main.c example/dronesNew/series_F_pb_and_ps_and_pj.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_series
./example_dronesNew_pb_and_ps_and_pj_series

make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj PARSEARGS=-product
gcc -ftime-report example/dronesNew/mono_pb_and_ps_and_pj_main.c example/dronesNew/F_pb_and_ps_and_pj.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_mono
./example_dronesNew_pb_and_ps_and_pj_mono

--------------------------------
pb ps pj pr 4
--------------------------------
make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj_and_pr  COMPILEARGS=-seriesComposition
gcc -ftime-report example/dronesNew/series_pb_and_ps_and_pj_and_pr_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_series
./example_dronesNew_pb_and_ps_and_pj_and_pr_series

make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj_and_pr PARSEARGS=-product
gcc -ftime-report example/dronesNew/mono_pb_and_ps_and_pj_and_pr_main.c example/dronesNew/F_pb_and_ps_and_pj_and_pr.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_mono
./example_dronesNew_pb_and_ps_and_pj_and_pr_mono

--------------------------------
pb ps pj pr pb2 5
--------------------------------
make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj_and_pr_and_pb2  COMPILEARGS=-seriesComposition
gcc -ftime-report example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_series
./example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_series

make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj_and_pr_and_pb2 PARSEARGS=-product
gcc -ftime-report example/dronesNew/mono_pb_and_ps_and_pj_and_pr_and_pb2_main.c example/dronesNew/F_pb_and_ps_and_pj_and_pr_and_pb2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_mono
./example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_mono

--------------------------------
pb ps pj pr pb2 ps2 6
--------------------------------
make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj_and_pr_and_pb2_and_ps2  COMPILEARGS=-seriesComposition
gcc -ftime-report xample/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_series
./example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_series

make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj_and_pr_and_pb2_and_ps2 PARSEARGS=-product

--------------------------------
pb ps pj pr pb2 ps2 pj2 7
--------------------------------
make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2  COMPILEARGS=-seriesComposition
gcc -ftime-report example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_series
./example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_series

make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2 PARSEARGS=-product

--------------------------------
pb ps pj pr pb2 ps2 pj2 pr2 8
--------------------------------
make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2  COMPILEARGS=-seriesComposition
gcc -ftime-report example/dronesNew/series_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2_main.c example/dronesNew/series_F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2_series
./example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2_series

make c_enf PROJECT=dronesNew FILE=pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2 PARSEARGS=-product
gcc -ftime-report example/dronesNew/mono_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2_main.c example/dronesNew/F_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2.c example/dronesNew/save_results.c -o example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2_mono
./example_dronesNew_pb_and_ps_and_pj_and_pr_and_pb2_and_ps2_and_pj2_and_pr2_mono


