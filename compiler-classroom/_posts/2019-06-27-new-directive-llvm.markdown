---
layout: post
title:  "How to add a new OpenMP directive in clang/llvm compiler"
author: "@alokmishra.besu"
date:   2019-06-27
categories: beginner
tags: [llvm,clang,openmp,directive]
---
### Features

In this tutorial we will cover how to add a new OpenMP directive in clang/LLVM compiler. The goal of this tutorial is to add a new OpenMP directive hello (`#pragma omp hello`) which will give a failed message if the next statement is not a for loop.

---

## Step 1 - Locate and go to clang directory
First, let's enter the `LLVM` source folder to look around. There are bunch of files there. For now only two of them will be modified for our goal. Both files are in the clang sub-project of the LLVM source code. In this tutorials's environment, the clang project is located at `$LLVM_SRC/llvm-8.0.0.src/tools/cfe-8.0.0.src`. In your machine you should locate the clang project and switch to that directory.
```.term1
cd $LLVM_SRC/llvm-8.0.0.src/tools/cfe-8.0.0.src
```

Now let us update the compiler, such that it just identifies the new directive. For this we need to update two files:
1. *OpenMPKinds.def* -- which defines the list of supported OpenMP directives and clauses.
2. *ParseOpenMP.cpp* -- which implements parsing of all OpenMP directives and clauses.

## Step 2 - Define the new directive
To define the new directive we will modify the file `OpenMPKinds.def`, located in `include/clang/Basic`. So open the file using your favorite editor. In this tutorial we will be using the vim editor.
```.term1
vim include/clang/Basic/OpenMPKinds.def
```

Now in this file go towards the end (to line 891 or before `#undef OPENMP_DIRECTIVE_EXT` is called) and add the following line
```
OPENMP_DIRECTIVE_EXT(hello, "hello")
```

In our current example hello does not have any clause associated with it, so we do not need to define `OPENMP_HELLO_CLAUSE`.

This way we are able to define the new directive `#pragma omp hello`.

## Step 3 - Implements parsing
To implement the parsing of this new directive we will modify the file `ParseOpenMP.cpp`, located in `lib/Parse`. So open the file using your favorite editor.
```.term1
vim lib/Parse/ParseOpenMP.cpp
```

Now in this file go to the function `ParseOpenMPDeclarativeOrExecutableDirective` and add the new case for the hello directive. Identify where the case for `OMPD_parallel` is defined and add your new case right before it. Here we will be reusing the `OMPD_parallel` code, so do not break your case.
```
  // Add new code here
  case OMPD_hello:
    std::cout <<"HELLO is caught\n";
  // New code ends here. Use the body of OMPD_parallel for our case
  case OMPD_parallel:

```
To support the std::cout include `iostream` at the start of the file.

That's it for now. Now let us build and test our code.

## Step 4 - Building LLVM and testing code
To build `LLVM` go to the `LLVM_BUILD` directory and run make. We are redirecting the output of make to /dev/null to have a clean output. It will still show errors.

```.term1
cd $LLVM_BUILD && make -j8 install > /dev/null
```

Once the code builds successfully and is installed, its time to test a small program.

```.term1
cd $EXAMPLE_DIR && cat test_hello.c
```

You should get the output of the file `test_hello.c` which uses the `hello` directive. Build this file using our clang compiler.

```.term1
clang -fopenmp test_hello.c
```

You should get an output `HELLO is caught`. Congratulations you were successfully able to add a new directive to OpenMP in clang compiler.

## Step 5 - Semantic analysis
In the above implementation we used the body of the `OMPD_parallel` case. Now we will implement our own parsing of the directive.

To parse a 


