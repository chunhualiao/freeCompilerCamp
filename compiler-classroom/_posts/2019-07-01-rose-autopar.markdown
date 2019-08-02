---
layout: post
title:  "autoPar - Auto Parallization Tool in ROSE"
author: "@ouankou"
date:   2019-07-01
categories: beginner
tags: [rose,openmp,autopar,parallelization]
---

This is a tool which can automatically insert OpenMP pragmas into input serial C/C++ codes.


---

# A. Overview

# B. Building autoPar

# C. Usage Examples

---

## A. Overview

The goal of this tutorial is to learn how to use an automatic parallelizaiton tool, autoPar, built using ROSE. autoPar can automatically insert OpenMP directives into input serial C/C++ codes. For input programs with existing OpenMP directives, the tool can also double check the correctness when the right option is turned on.

The source files are currently located in ```$ROSE_SRC/projects/autoParallelization```.
A standalone executable program (named ```autoPar``` ) is generated and installed to the installation tree of ROSE (under ```$ROSE_PATH/bin```).
Test input files are located at ```$ROSE_SRC/projects/autoParallelization/tests```.
You can test the tool in ```$ROSE_BUILD/projects/autoParallelization``` by typing ```make check```.
There is a section in ROSE manual: 12.7 Automatic Parallelization *pdf(http://rosecompiler.org/docs/snapshots/Edited%20ROSE-UserManual%209_10_231.pdf)*. 
It is used to explore semantics-aware automatic parallelization, as described in our papers:

A workshop paper: Chunhua Liao, Daniel J. Quinlan, Jeremiah J. Willcock and Thomas Panas, Extending Automatic Parallelization to Optimize High-Level Abstractions for Multicore, In Proceedings of the 5th international Workshop on OpenMP: Evolving OpenMP in An Age of Extreme Parallelism (Dresden, Germany, June 3–05, 2009). *pdf(https://e-reports-ext.llnl.gov/pdf/368494.pdf)*

A journal version of the paper: Chunhua Liao, Daniel J. Quinlan, Jeremiah J. Willcock and Thomas Panas, Semantic-Aware Automatic Parallelization of Modern Applications Using High-Level Abstractions, Journal of Parallel Programming, Accepted in Jan. 2010 *pdf(https://e-reports-ext.llnl.gov/pdf/384220.pdf)*

Similar to ROSE, autoPar is released under the BSD license.


## B. Building autoPar

THe source code is provided in the sandbox. The building configuration has been completed as well. User could follow the steps below to build and try ```autoPar``` directly.

Enter ```autoPar``` build folder.
```.term1
cd $ROSE_BUILD/projects/autoParallelization
```
Start building.
```.term1
make -j4
```
Install the binaries
```.term1
make install
```

Then the binaries will be installed to ```$ROSE_PATH/bin```.
The tool can be tested by the following code.
```.term1
make check
```

Command line options:
```.term1
autoPar --help
```

Additional useful ROSE flags:
```
-rose:skipfinalCompileStep // skip invoking the backend compiler to compile the transformed code, this is useful to workaround some bugs
--edg:no_warnings // suppress warnings from the EDG C++ frontend
```

## C. Usage Examples

Testing input files can be found at https://github.com/rose-compiler/rose/tree/master/projects/autoParallelization/tests

The corresponding generated testing output files can be found at: https://github.com/chunhualiao/autoPar-demo

We provide two samples below:

##### Without using annotations

Checkout the input:
```.term1
cat $ROSE_SRC/projects/autoParallelization/tests/inner_only.c
```

Conduct auto parallelization:
```.term1
autoPar -rose:C99 --edg:no_warnings -w -rose:verbose 0 --edg:restrict -rose:autopar:unique_indirect_index -rose:autopar:enable_patch -I$ROSE_SRC/src/frontend/SageIII -c $ROSE_SRC/projects/autoParallelization/tests/inner_only.c
```

Checkout the output:
```.term1
cat ./rose_inner_only.c
```
