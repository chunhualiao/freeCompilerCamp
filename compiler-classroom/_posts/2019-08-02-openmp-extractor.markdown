---
layout: post
title:  "OpenMP Extractor - a plugin to extract OpenMP directives."
author: "@gleisonsdm"
date:   2019-08-02
categories: beginner
tags: [AutoParBench,openmp,framework,parallelization]
---

This is a clang plugin that can parse programs with OpenMP directives and generate Json files containing a description about loops.

---

Concepts in this exercise:
# A. Overview

# B. Building OpenMP Extractor

# C. Usage Examples

---

## A. Overview

This tool is an implementation of an external clang plugin. It can automatically parse the source code and extract OpenMP directives into Json files. For input the plugin use programs with existing OpenMP directives, the tool will parser them and extract relevant information about the code, providing facilities to check the correcntess of directives.


## B. Building OpenMP Extractor

The source code is provided in the following repository:
https://github.com/chunhualiao/AutoParBench

To download it:
```.term1
git clone https://github.com/chunhualiao/AutoParBench.git /usr/src/OMP_Extractor
```

Create ```OMP_Extractor``` build folder.
```.term1
mkdir /usr/src/OMP_Extractor/lib
```

Enter ```OMP_Extractor``` build folder.
```.term1
cd /usr/src/OMP_Extractor/lib
```

Create scripts to building.
```.term1
CXX=g++ cmake -DLLVM_DIR=${LLVM_INSTALL}/lib/cmake/llvm /usr/src/OMP_Extractor/clangPlugin/
```

Then the binaries of this library will be installed to ```/usr/src/OMP_Extractor/lib```.
```.term1
make -j4
```

## C. Usage Examples

We provide an example file to testing. Feel free to try other programs.
```.term1
echo "#include <stdio.h>" &> test.c 
echo "#include <stdlib.h>" &>> test.c
echo "int main(int argc, char* argv[]) " &>> test.c
echo "{ " &>> test.c  
echo "  int i; " &>> test.c 
echo "  int len = 1000; " &>> test.c 
echo " " &>> test.c 
echo "  int a[1000]; " &>> test.c 
echo " " &>> test.c 
echo "  #pragma omp parallel for " &>> test.c 
echo "  for (i=0; i<len; i++) " &>> test.c 
echo "    a[i]= i;  " &>> test.c 
echo " " &>> test.c 
echo "  for (i=0;i< len -1 ;i++) " &>> test.c 
echo "    a[i]=a[i+1]+1; " &>> test.c 
echo " " &>> test.c 
echo '  printf ("a[500]=%d\n", a[500] ); ' &>> test.c 
echo '  return 0; ' &>> test.c 
echo "} " &>> test.c 
```

Run OpenMP Extractor to run the plugin, you should load the library to run the analysis on clang.
```.term1
clang -Xclang -load -Xclang /usr/src/OMP_Extractor/lib/ompextractor/libCLANGOMPExtractor.so -Xclang -add-plugin -Xclang -extract-omp -fopenmp -g -O0 -c -fsyntax-only test.c
```

Checkout if the Json file was created:
```.term1
ls 
```

Checkout the output:
```.term1
cat test.c.json
```


