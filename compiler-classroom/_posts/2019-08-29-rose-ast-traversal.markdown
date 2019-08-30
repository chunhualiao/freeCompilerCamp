---
layout: post
title:  "ROSE AST - Traversal"
author: "@chunhualiao"
date:   2019-08-29
categories: beginner
tags: [rose,ast,traversal]
---

# Tips:

Code snippets are shown in one of three ways throughout this environment:

1. Code that looks like `this` is sample code snippets that is usually part of an explanation.
2. Code that appears in box like the one below can be clicked on and it will automatically be typed in to the appropriate terminal window:
```.term1
vim readme.txt
```

3. Code appearing in windows like the one below is code that you should type in yourself. Usually there will be a unique ID or other bit your need to enter which we cannot supply. Items appearing in <> are the pieces you should substitute based on the instructions.
```
Add your name here - <name>
```

## Features
This is a tutorial to build your own tool traversing ROSE AST to find things of your interests.

---

# A. Overview
An essential task in compiler development is to walk the AST 
to find nodes of interests, in order to gather information and/or modify
AST in the context of building program analysis and transformation..  
ROSE includes different
sorts of traversal APIs to help this task.

The goal of this tutorial is to learn how to use a visitor pattern traversal API to walk the AST and find for loops in an input program. 

# B. Get the source files and makefile

Get the example ROSE-based analyzer traversing AST to find loops. Rename it to be demo.C:
```.term1
wget https://raw.githubusercontent.com/rose-compiler/rose/develop/tutorial/visitorTraversal.C
mv visitorTraversal.C demo.C

```

Get a sample makefile
```.term1
wget https://raw.githubusercontent.com/rose-compiler/rose/develop/tutorial/SampleMakefile
```

Get an example input code for the analyzer:
```.term1
wget https://raw.githubusercontent.com/rose-compiler/rose/develop/tutorial/inputCode_ExampleTraversals.C
```

# C. Build the analyzer using the makefile

Prepare the environment variable used to specify where ROSE is installed.
```.term1
export ROSE_HOME=/opt/install/rose_install
```

Build the analyzer
```.term1
make -f SampleMakefile
```
There should be an executable file named demo under the current directory:
```.term1
ls demo
```

Finally, run the demo analyzer to process the example input code:

```.term1
./demo -c inputCode_ExampleTraversals.C
```
The analyzer should find two for loops and report the end of the traveral.

```
find a loop ..
find a loop ..


```

# References

For more information about AST traversal, please check
* http://rosecompiler.org/ROSE_Tutorial/ROSE-Tutorial.pdf Chapter 7
