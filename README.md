# Machine Learning with Golang

This repository provides a basic setup to do Machine Learning with Golang and Python, **TensorFlow 1.15** and **CUDA 10.0**.

It is such a hassle to install a working version combination of Python, Golang, TensorFlow and CUDA.
Thus, I created two scripts to build and run a machine learning environment by using a local Docker installation.

## Why Golang?

Go provides a strong foundation for data manipulation and parsing because of the static typing and explicit error handling. 
So we are able to maintain integrity of our data by handling missing or incorrect values in context of our use case.

## Build

First of all we will build a **Docker image** by using the following command:

    ./scripts/build-image.sh
    
We must use **TensorFlow 1.15** because the corresponding C-Library used for Golang-Bindings is currently not available for TensorFlow 2. 

## Run

Then we are able to use this image to build and run our Go sources by entering a **Bash Shell** as follows:

    ./scripts/run-container.sh

This command will map our current **GOPATH** into the container and change the working directory to **project**.

    ________                               _______________                
    ___  __/__________________________________  ____/__  /________      __
    __  /  _  _ \_  __ \_  ___/  __ \_  ___/_  /_   __  /_  __ \_ | /| / /
    _  /   /  __/  / / /(__  )/ /_/ /  /   _  __/   _  / / /_/ /_ |/ |/ / 
    /_/    \___//_/ /_//____/ \____//_/    /_/      /_/  \____/____/|__/
    
    
    You are running this container as user with ID 1000 and group 1000,
    which should map to the ID and group for your user on the Docker host. Great!
    
    tf-docker /home/andygeiss/go/src/github.com/andygeiss/machine-learning-golang/project > 

Now run your first TensorFlow program with Go:

    go run hello_tf.go
    
    Hello from TensorFlow version 1.15.0
