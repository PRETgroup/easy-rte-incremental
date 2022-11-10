-----------------------------
PARALLEL HARDWARE COMPOSITION
-----------------------------

Altering Verilog template
rtec/rteVerilogTemplate.go 

1) Make edits
2) Build the compiler (make local in easy-rte directory)
3) Make the example (I'm using abc5, so make verilog_enf PROJECT=abc5 COMPILEARGS=-parallelComposition)

Added parallel composition COMPILEARGS
make verilog_enf PROJECT=abc5 COMPILEARGS=-parallelComposition
-parallelComposition is passed to the go compiler.

Added "make local" to avoid checking for goFB latest downloads

-------------------------------
SERIES SOFTWARE (C) COMPOSITION
-------------------------------

Example ABC
    Monolithic
        make c_enf PROJECT=abc FILE=ab_and_ac PARSEARGS=-product
        gcc example/abc/ab_and_ac_main.c example/abc/F_ab_and_ac.c -o example_abc_monolithic
        ./example_abc_monolithic

Build compiler
    make local

Example ABC
    Series
        make c_enf PROJECT=abc FILE=ab_and_ac  COMPILEARGS=-seriesComposition
        gcc example/abc/series_ab_and_ac_main.c example/abc/series_F_ab_and_ac.c -o example_abc_series
        ./example_abc_series

        Notes
            Limited to BOOLEANS only
            Must be defined as uint8_t in erte file

    Manual Series ABC
        gcc example/abc/manual_series_ab_and_ac_main.c example/abc/manual_series_F_ab_and_ac.c -o example_abc_manual_series
        ./example_abc_manual_series

Example Drones Series

// Delete all xml files before running these (or else you may try compile the monolithic XML)
make c_enf PROJECT=drones FILE=px_and_py  COMPILEARGS=-seriesComposition
make c_enf PROJECT=drones FILE=px_and_py_and_pz  COMPILEARGS=-seriesComposition
make c_enf PROJECT=drones FILE=px_and_py_and_pz_and_pr  COMPILEARGS=-seriesComposition
make c_enf PROJECT=drones FILE=px_and_py_and_pz_and_pr_and_px2  COMPILEARGS=-seriesComposition

gcc example/drones/series_px_and_py_main.c example/drones/series_F_px_and_py.c -o example_series_px_and_py
gcc example/drones/series_px_and_py_and_pz_main.c example/drones/series_F_px_and_py_and_pz.c -o example_series_px_and_py_and_pz
gcc example/drones/series_px_and_py_and_pz_and_pr_main.c example/drones/series_F_px_and_py_and_pz_and_pr.c -o example_series_px_and_py_and_pz_and_pr
gcc example/drones/series_px_and_py_and_pz_and_pr_and_px2_main.c example/drones/series_F_px_and_py_and_pz_and_pr_and_px2.c -o example_series_px_and_py_and_pz_and_pr_and_px2

./example_series_px_and_py
./example_series_px_and_py_and_pz
./example_series_px_and_py_and_pz_and_pr
./example_series_px_and_py_and_pz_and_pr_and_px2

Example Drones Monolithic

// Delete all xml files before running these (or else you may try compile the series XML)
make c_enf PROJECT=drones FILE=px_and_py  PARSEARGS=-product
make c_enf PROJECT=drones FILE=px_and_py_and_pz  PARSEARGS=-product
make c_enf PROJECT=drones FILE=px_and_py_and_pz_and_pr  PARSEARGS=-product
make c_enf PROJECT=drones FILE=px_and_py_and_pz_and_pr_and_px2  PARSEARGS=-product

gcc example/drones/mono_px_and_py_main.c example/drones/F_px_and_py.c -o example_mono_px_and_py
gcc example/drones/mono_px_and_py_and_pz_main.c example/drones/F_px_and_py_and_pz.c -o example_mono_px_and_py_and_pz
gcc example/drones/mono_px_and_py_and_pz_and_pr_main.c example/drones/F_px_and_py_and_pz_and_pr.c -o example_mono_px_and_py_and_pz_and_pr
gcc example/drones/mono_px_and_py_and_pz_and_pr_and_px2_main.c example/drones/F_px_and_py_and_pz_and_pr_and_px2.c -o example_mono_px_and_py_and_pz_and_pr_and_px2

./example_mono_px_and_py
./example_mono_px_and_py_and_pz
./example_mono_px_and_py_and_pz_and_pr
./example_mono_px_and_py_and_pz_and_pr_and_px2
./example_series_px_and_py
./example_series_px_and_py_and_pz
./example_series_px_and_py_and_pz_and_pr
./example_series_px_and_py_and_pz_and_pr_and_px2

