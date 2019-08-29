# freeCompilerCamp
Goal: an open, extensive online platform to automatically train and certify compiler researchers and developers.
* compiler-classroom: the website's text content
** index.html  : the home page of the webpage
** _posts  : a directory contains the markdown files for individual tutorials
* play-with-compiler: the online sandbox based on play-with-docker
** dockerfiles/dind/ : a directory contains docker files for our online sandboxed terminal
*** Dockerfile.base - This is the base docker file which loads the ubuntu environment.
*** Dockerfile.middle - In this docker file we add the user and user group required for the docker environment. This file is dependent upon fcc_docker:0.1, which is an image built from Dockerfile.base.
*** Dockerfile.dind - This is the main docker file. This is the file where we up setup our environment, add required tools, like LLVM and ROSE, etc. This file is dependent upon fcc_dind:0.1, which is an image buikt from Dockerfile.middle.
