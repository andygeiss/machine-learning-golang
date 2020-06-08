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

## Benchmark

Try a TensorFlow benchmark of a Matrix Subtraction:

    go test -bench mat_sub_tf_test.go
    
    goos: linux
    goarch: amd64
    pkg: github.com/andygeiss/machine-learning-golang/project
    Benchmark_TensorFlow_NewScope-4                   140017             10038 ns/op
    Benchmark_TensorFlow_Scope_Finalize-4           455846677                2.42 ns/op
    Benchmark_TensorFlow_NewSession-4                    655           1741003 ns/op
    Benchmark_TensorFlow_Session_Run-4                 62317             18037 ns/op
    Benchmark_TensorFlow_Mat_Sub_100x100-4            237834              4982 ns/op
    Benchmark_TensorFlow_Mat_Sub_1000x1000-4          201316              5131 ns/op
    Benchmark_GoNum_Mat_Sub_100x100-4                  42163             28339 ns/op
    Benchmark_GoNum_Mat_Sub_1000x1000-4                  290           3910920 ns/op
    PASS
    ok      github.com/andygeiss/machine-learning-golang/project    13.498s

