# Overview
[FreeCompilerCamp](http://freecompilercamp.org) (currently alive) is a free and open online training platform aimed to automate the training of compiler developers. Our platform allows anyone who is interested in developing compilers to learn the necessary skills for free. The platform is built on top of Play-With-Docker, a docker playground for users to conduct experiments in a sandbox. We welcome anyone to try out our system, give us feedback, contribute new training courses, or enhance the training platform to make it an effective learning resource for the compiler community.

While this platform can be used to host any compiler tutorials, we specially collect some tutorials for OpenMP compilers. We have created some initial tutorials to train users to learn how to use the ROSE or Clang/LLVM compiler to support OpenMP.

# Directory Layout
Goal: an open, extensive online platform to automatically train and certify compiler researchers and developers.
* compiler-classroom: the website's text content
 index.html  : the home page of the webpage
  * _posts  : a directory contains the markdown files for individual tutorials
* play-with-compiler: the online sandbox based on play-with-docker
  * dockerfiles/dind/ : a directory contains docker files for our online sandboxed terminal
    * Dockerfile.base - This is the base docker file which loads the ubuntu environment.
    * Dockerfile.middle - In this docker file we add the user and user group required for the docker environment. This file is dependent upon fcc_docker:0.1, which is an image built from Dockerfile.base.
    * Dockerfile.dind - This is the main docker file. This is the file where we up setup our environment, add required tools, like LLVM and ROSE, etc. This file is dependent upon fcc_dind:0.1, which is an image buikt from Dockerfile.middle.

# Installation
You can install your own instance of this website. Please follow instructions at 
* https://github.com/chunhualiao/freeCompilerCamp/wiki/Deploy-FreeCC-to-AWS

# Contact Us
This work was performed under the auspices of the U.S. Department of Energy by Lawrence Livermore National Laboratory under Contract DE-AC52-07NA27344, and partially supported by the U.S. Dept. of Energy, Office of Science, ASCR SC-21), under contract DE-AC02-06CH11357. 

This website is still under development (LLNL-WEB-789932). For questions and comments, please file issue tickets to this git repo. Alternatively, you can contact liao6@llnl.gov .

# Publication
List
* Anjia Wang, Alok Mishra, Chunhua Liao, Yonghong Yan, Barbara Chapman, FreeCompilerCamp.org: Training for OpenMP Compiler Development from Cloud, Sixth SC Workshop on Best Practices for HPC Training and Education: BPHTE19, 2019
* Alok Mishra, Anjia Wang, Chunhua Liao, Yonghong Yan, Barbara Chapman, FreeCompilerCamp: Online Training for Extending Compilers, SC'19 Research Poster submission, accepted (also selected as a Best Poster nominee).
