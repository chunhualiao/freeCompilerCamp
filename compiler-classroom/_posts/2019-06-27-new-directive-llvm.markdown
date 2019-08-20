---
layout: post
title:  "How to add a new OpenMP directive in Clang/LLVM compiler"
author: "@alokmishra.besu"
date:   2019-06-27
categories: beginner
tags: [llvm,clang,openmp,directive]
---
### Features

In this tutorial we will cover how to add a new OpenMP directive in Clang/LLVM compiler. The goal of this tutorial is to add a new OpenMP directive -- metadirective (`#pragma omp metadirective [clause[[,]clause]...]`), defined in OpenMP Specification 5.0, that can specify multiple directive variants of which one may be conditionally selected to replace the metadirective based on the enclosing OpenMP context.

---

## Step 1 - Locate and go to clang directory
First, let's enter the `LLVM` source folder to look around. There are a bunch of files and directories there. For now only interested in the Clang sub-project of the LLVM source code. In this tutorials's environment, the Clang project is located at `$LLVM_SRC/tools/clang`. In your machine you should locate the Clang project and switch to that directory.
```.term1
cd $LLVM_SRC/tools/clang
```

## Step 2 - Define the new directive
The first thing that we should do is let the compiler identify a new directive, which in this tutorial is `metadirective`.

Now let us update the compiler, such that it just identifies the new directive. For this we need to update two files:
1. *OpenMPKinds.def* -- which defines the list of supported OpenMP directives and clauses.
2. *ParseOpenMP.cpp* -- which implements parsing of all OpenMP directives and clauses.

To define the new directive we will modify the file `OpenMPKinds.def`, located in `include/clang/Basic`. So open the file using your favorite editor. In this tutorial we will be using the vim editor.
```.term1
vim include/clang/Basic/OpenMPKinds.def
```

Now in this file go to line 237 (or anywhere before `#undef OPENMP_DIRECTIVE_EXT` is called) and add the following new line after it:
```
OPENMP_DIRECTIVE_EXT(metadirective, "metadirective")
```

In our current state we are not dealing with any clause associated with metadirective, so we do not need to define `OPENMP_METADIRECTIVE_CLAUSE`. We will learn about adding clause later in the tutorial.

This way we are able to define the new directive `#pragma omp metadirective`.

## Step 3 - Implement parsing
Before parsing the lexer will split the source code into multiple tokens. The parser will read these tokens and give a structural representation to them. To implement the parsing of this new directive we need to modify the file `ParseOpenMP.cpp`, located in `lib/Parse`. So open the file using your favorite editor.
```.term1
vim lib/Parse/ParseOpenMP.cpp
```

Now in this file go to the function `ParseOpenMPDeclarativeOrExecutableDirective`, identify the switch case (line 997) and add a new case for `OMPD_metadirective`. Here we will print out <span style="color:blue">**METADIRECTIVE is caught**</span> and then consume the token.
```
  switch (DKind) {
  case OMPD_metadirective: {
    llvm::errs() <<"METADIRECTIVE is caught\n";
    ConsumeToken();
    ConsumeAnnotationToken();
    break;
  }
```

That's it for now. Now let us build and test our code.

## Step 4 - Building LLVM and testing code
To build `LLVM` go to the `LLVM_BUILD` directory and run make. We are redirecting the output of make to /dev/null to have a clean output. It will still show errors.

```.term1
cd $LLVM_BUILD && make -j8 install > /dev/null
```

You might get a couple of warnings about `enumeration value 'ompd_metadirective' not handled in switch`. ignore these warnings for now. we will handle them later. Once the code builds successfully and is installed, its time to test a small program. let us get a new test file

```.term1
wget https://raw.githubusercontent.com/chunhualiao/freecc-examples/master/metadirective/meta.c
```

now you have a new test file `meta.c` which uses the `metadirective` directive. The content of the file should be as follows:
```
int main()
{
#pragma omp metadirective 
    for(int i=0; i<10; i++)
    ;
    return 0;
}
```

Build this file using your Clang compiler.

```.term1
clang -fopenmp meta.c
```

you should get an output `metadirective is caught`. 

<span style="color:green">**Congratulations**</span> you were successfully able to add and identify a new directive to openmp in Clang compiler.
