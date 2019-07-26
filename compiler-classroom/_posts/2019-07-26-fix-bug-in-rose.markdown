---
layout: post
title:  "Fix a bug in ROSE OpenMP lowering"
author: "@ouankou"
date:   2019-07-26
categories: beginner
tags: [rose,openmp,lowering,parallelization]
---

This tutorial is to show how to fix a bug in ROSE compiler.


---

Concepts in this exercise:
# A. Overview

# B. Building autoPar

# C. Usage Examples

---

## A. Overview

The bug is mentioned in the issue 35 in official ROSE repository on GitHub (https://github.com/rose-compiler/rose/issues/35).

When ROSE transforms the ```parallel for``` directive, it outlines the pragma body and make multiple calls to this outlined function for different portion of the ```for``` loop. The beginning, ending index and step of ```for``` loop has to be calculated carefully to enforce correct result.



## B. Existing Bug

1. The declaration of function ```XOMP_loop_default``` and the call to this function have mismatched parameters.

2. The step of ```for``` loop is calculated incorrectly.

The variables in the ```private``` clause are not initialized, but it's not a bug. The initialization should be done by user but not by compiler.

THe source code is provided in the sandbox. The building configuration has been completed as well. User could follow the steps below to build and try ```autoPar``` directly.

#### Show the input

```.term1
cd $EXAMPLE_DIR && cat rose_bug_parallel_for.c 
```

#### Generate the incorrect output

```.term1
rose-compiler -rose:openmp:lowering -lomp -lxomp -lgomp bug_parallel_for_in_rose.c
```

#### Show the incorrect output

```.term1
cat rose_bug_parallel_for_in_rose.c 
```


## C. Fix the Bug


##### Fix the mismatched parameters

Modify the function declaration in the header file:
```.term1
vim $ROSE_SRC/src/midend/programTransformation/ompLowering/libxomp.h
```
On the line 72, the data type for ```lower```, ```upper```` and ```stride``` should be all changed from ```int``` to ```unsigned int```. Then use ```:wq``` to save and quit.

Modify the function in the ```.c``` file:
```.term1
vim $ROSE_SRC/src/midend/programTransformation/ompLowering/xomp.c
```
On the line 574, the data type for ```lower```, ```upper```` and ```stride``` should be all changed from ```int``` to ```unsigned int```. Then use ```:wq``` to save and quit.


#### Fix the miscalculated loop step

```.term1
vim $ROSE_SRC/src/frontend/SageIII/sageInterface/sageInterface.C
```

On the line 11602 and 11607, change the variable ```incr``` to ```arithOp```. Use ```:wq``` to save and quit.

#### Rebuild and test

First we need to rebuild ROSE to make modification effective.
```.term1
cd $ROSE_BUILD && make core -j4 && make install-core
```

#### Generate the output
```.term1
cd $EXAMPLE_DIR && rose-compiler -rose:openmp:lowering -lomp -lxomp -lgomp bug_parallel_for_in_rose.c
```

#### Show the correct output
```.term1
cat rose_bug_parallel_for_in_rose.c 
```


