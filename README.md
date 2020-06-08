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
    
    2020-06-08 14:31:19.816722: I tensorflow/core/platform/cpu_feature_guard.cc:142] Your CPU supports instructions that this TensorFlow binary was not compiled to use: AVX2 FMA
    2020-06-08 14:31:19.822836: I tensorflow/core/platform/profile_utils/cpu_utils.cc:94] CPU Frequency: 2499950000 Hz
    2020-06-08 14:31:19.823141: I tensorflow/compiler/xla/service/service.cc:168] XLA service 0x167c270 initialized for platform Host (this does not guarantee that XLA will be used). Devices:
    2020-06-08 14:31:19.823153: I tensorflow/compiler/xla/service/service.cc:176]   StreamExecutor device (0): Host, Default Version
    2020-06-08 14:31:19.824114: I tensorflow/stream_executor/platform/default/dso_loader.cc:44] Successfully opened dynamic library libcuda.so.1
    2020-06-08 14:31:19.848927: I tensorflow/core/common_runtime/gpu/gpu_device.cc:1618] Found device 0 with properties: 
    name: GeForce GTX 1050 major: 6 minor: 1 memoryClockRate(GHz): 1.493
    pciBusID: 0000:01:00.0
    2020-06-08 14:31:19.849279: I tensorflow/stream_executor/platform/default/dso_loader.cc:44] Successfully opened dynamic library libcudart.so.10.0
    2020-06-08 14:31:19.850529: I tensorflow/stream_executor/platform/default/dso_loader.cc:44] Successfully opened dynamic library libcublas.so.10.0
    2020-06-08 14:31:19.851510: I tensorflow/stream_executor/platform/default/dso_loader.cc:44] Successfully opened dynamic library libcufft.so.10.0
    2020-06-08 14:31:19.851781: I tensorflow/stream_executor/platform/default/dso_loader.cc:44] Successfully opened dynamic library libcurand.so.10.0
    2020-06-08 14:31:19.853095: I tensorflow/stream_executor/platform/default/dso_loader.cc:44] Successfully opened dynamic library libcusolver.so.10.0
    2020-06-08 14:31:19.854137: I tensorflow/stream_executor/platform/default/dso_loader.cc:44] Successfully opened dynamic library libcusparse.so.10.0
    2020-06-08 14:31:19.857563: I tensorflow/stream_executor/platform/default/dso_loader.cc:44] Successfully opened dynamic library libcudnn.so.7
    2020-06-08 14:31:19.858274: I tensorflow/core/common_runtime/gpu/gpu_device.cc:1746] Adding visible gpu devices: 0
    2020-06-08 14:31:19.858311: I tensorflow/stream_executor/platform/default/dso_loader.cc:44] Successfully opened dynamic library libcudart.so.10.0
    2020-06-08 14:31:19.964502: I tensorflow/core/common_runtime/gpu/gpu_device.cc:1159] Device interconnect StreamExecutor with strength 1 edge matrix:
    2020-06-08 14:31:19.964545: I tensorflow/core/common_runtime/gpu/gpu_device.cc:1165]      0 
    2020-06-08 14:31:19.964566: I tensorflow/core/common_runtime/gpu/gpu_device.cc:1178] 0:   N 
    2020-06-08 14:31:19.965469: I tensorflow/core/common_runtime/gpu/gpu_device.cc:1304] Created TensorFlow device (/job:localhost/replica:0/task:0/device:GPU:0 with 1250 MB memory) -> physical GPU (device: 0, name: GeForce GTX 1050, pci bus id: 0000:01:00.0, compute capability: 6.1)
    Hello from TensorFlow version 1.15.0
