---
layout: post
title:  "Extracting OpenMP Information from Programs using a Clang Plugin"
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

# D. Extending the plugin

---

## A. Overview

This tool is an implementation of an external clang plugin. It can automatically parse the source code and extract OpenMP directives into Json files. For input the plugin use programs with existing OpenMP directives, the tool will parser them and extract relevant information about the code, providing facilities to check the correcntess of directives.


## B. Building OpenMP Extractor

The source code is provided in the following repository:
https://github.com/chunhualiao/AutoParBench

To download it:
```.term1
git clone https://github.com/gleisonsdm/OpenMP-Extractor.git /usr/src/OMP_Extractor
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
cat << EOF > test.c
#include <stdlib.h>
#include <stdio.h>
int main (void)
{
  int sum=0;
  #pragma omp parallel for reduction(+:sum)
  for (int i = 0; i < 100; i++)
  {
    sum += 1;
  }
  printf ("sum = %d\n",sum);
  return 0;
}
EOF
```

Run OpenMP Extractor to run the plugin, you should load the library to run the analysis on clang.
To send flags to clang, is necessary to use "-Xclang" before each argument.
 - -load: Necessary to load libraries on clang.
 - /usr/src/OMP_Extractor/lib/ompextractor/libCLANGOMPExtractor.so: This is the library containing the plugin.
 - -add-plugin: Flag to add an out of tree plugin to clang.
 - -extract-omp: This flag ask clang to run the plugin.
 - -fopenmp: Provide objects in clang's frontend about OpenMP directives and clauses.
 - -g: Flag to provide debug information about the source file.
 - -O0: Disable optimizations from clang.
 - -c: Flag to create an object file, then clang does not need to find the source code for included files.
 - -fsyntax-only : Prevents the compiler to write an object file. We use this avoid the the creation of an intermediate file, as it is not necessary.
 
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

In the end, the Json is a format to store OpenMP information extracted by Clang/LLVM.

## D. Extending the plugin

#### Changing the plugin

First, let's open the plugin source file. The plugin was build in one for simplicity.
```.term1
vim /usr/src/OMP_Extractor/clangPlugin/ompextractor/ompextractor.cpp
```

Then, modify the clauses to recognize the new directive. This tutorial will show us how to add suport to the ```num_threads``` OpenMP clause. To modify the file, please, go to the line 357 (or before the comment "/*Final or If clauses are marked as multiversioned.*/" in the function "ClassifyClause") and insert the following code:
```cpp
      /*NumThreads clause*/
      if (OMPNumThreadsClause *OMPCc = dyn_cast<OMPNumThreadsClause>(clause)) {
        clauseType["num_threads"] = getStrForStmt(OMPCc->getNumThreads());
      }

```


The next step is to modify the Json builder to write the new feature (i.e. the clause). Please, go to the line 480 (or before line with ```currFile.labels += "\"body\":[" + snippet + "]\n";``` in the function "CreateLoopDirectiveNode") and insert the code as described below.
```cpp
currFile.labels += "\"num threads\":\"" + ((clauseType.count("num_threads") > 0) ? (clauseType["num_threads"]) : "") + "\",\n";
```

The last step before compile the modified code is to generate similar attribute for loops without openmp directives associated on it. To do that, please, go to the line 164 (or before line with ```currFile.labels += "\"body\":[" + snippet + "]\n";``` in the function "CreateLoopNode") and insert the following code:
```cpp
currFile.labels += "\"num threads\":\"\",\n";
```  

Now, it is done. Just close the file typing ":wq", then type enter.

#### Using an example

First, let's create the example file. We suggest you to use the following source file.
```.term1
cat << EOF > test.c
#include <stdlib.h>
#include <stdio.h>

int main() {
  int i, sum = 0;
  #pragma omp parallel for reduction(+:sum) num_threads(5)
  for (int i = 0; i < 100; i++) {
    sum += 1;
  }

  for (int i = 0; i < 100; i++) {
    sum += 1;
  }

  printf ("sum = %d\n",sum);
  return 0;
}
EOF
```

Enter ```OMP_Extractor``` build folder.
```.term1
cd /usr/src/OMP_Extractor/lib
```

Then the binaries of this library needs to be updated into ```/usr/src/OMP_Extractor/lib```.
```.term1
make -j4
```
To run the example, is just type the following command line. 
```.term1
clang -Xclang -load -Xclang /usr/src/OMP_Extractor/lib/ompextractor/libCLANGOMPExtractor.so -Xclang -add-plugin -Xclang -extract-omp -fopenmp -g -O0 -c -fsyntax-only test.c
```

Checkout if the Json file was created:
```.term1
ls
```

The last step is check the output:
```.term1
cat test.c.json
```

Well done, you got it. 
