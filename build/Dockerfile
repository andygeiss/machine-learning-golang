# Golang 1.14
FROM golang:latest AS golang
# TensorFlow 1.15
FROM tensorflow/tensorflow:latest-gpu AS tensorflow
COPY --from=golang /usr/local/go/ /usr/local/go/
RUN ln -sf /usr/local/go/bin/* /usr/local/bin/
RUN pip install tensorflow-gpu==1.15
# NVIDIA CUDA 10.0
RUN curl https://developer.download.nvidia.com/compute/cuda/repos/ubuntu1804/x86_64/cuda-ubuntu1804.pin -o cuda-ubuntu1804.pin && \
    mv cuda-ubuntu1804.pin /etc/apt/preferences.d/cuda-repository-pin-600 && \
    curl http://developer.download.nvidia.com/compute/cuda/10.2/Prod/local_installers/cuda-repo-ubuntu1804-10-2-local-10.2.89-440.33.01_1.0-1_amd64.deb -o cuda-repo-ubuntu1804-10-2-local-10.2.89-440.33.01_1.0-1_amd64.deb && \
    dpkg -i cuda-repo-ubuntu1804-10-2-local-10.2.89-440.33.01_1.0-1_amd64.deb && \
    apt-key add /var/cuda-repo-10-2-local-10.2.89-440.33.01/7fa2af80.pub
RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install --no-install-recommends -y cuda-10-0 libcudnn7
# C/GO bindings for TensorFlow
RUN curl https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-gpu-linux-x86_64-1.15.0.tar.gz -o libtensorflow-gpu-linux-x86_64-1.15.0.tar.gz && \
    tar -C /usr/local -xzf libtensorflow-gpu-linux-x86_64-1.15.0.tar.gz
# Add python libraries
RUN pip install numpy pandas
# Cleanup
RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    mkdir /.cache && \
    mkdir /.local && \
    chown 1000:1000 -R /.cache && \
    chown 1000:1000 -R /.local