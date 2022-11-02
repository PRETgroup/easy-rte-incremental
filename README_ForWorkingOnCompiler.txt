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

Manual Series ABC
    gcc example/abc/manual_series_ab_and_ac_main.c example/abc/manual_series_F_ab_and_ac.c -o example_abc_manual_series
    ./example_abc_manual_series
