ARG VERSION=docker:stable-dind
FROM ${VERSION}

RUN apk add --no-cache git tmux py2-pip apache2-utils vim build-base gettext-dev curl bash-completion bash util-linux jq openssh openssl tree

ENV GOPATH /root/go
ENV PATH $PATH:$GOPATH


# Compile and install httping
# (used in orchestration workshop, and very useful anyway)
RUN mkdir -p /opt && cd /opt && \
    curl https://vanheusden.com/httping/httping-2.5.tgz | \
    tar -zxf- && cd httping-2.5 && \
    ./configure && make install LDFLAGS=-lintl && \
    rm -rf httping-2.5

# customize packages
ENV LLVM_SRC /opt/llvm_src
ENV LLVM_PATH /opt/llvm_install
RUN apk add --no-cache cmake gcc g++
RUN mkdir -p /opt/llvm_src && \
    mkdir -p /opt/llvm_install && \
    mkdir -p /opt/llvm_src/build && \
    cd /opt/llvm_src && \
    wget http://releases.llvm.org/8.0.0/llvm-8.0.0.src.tar.xz && \
    tar -xf llvm-8.0.0.src.tar.xz && \
    cd llvm-8.0.0.src/tools && \
    wget http://releases.llvm.org/8.0.0/cfe-8.0.0.src.tar.xz && \
    tar -xf cfe-8.0.0.src.tar.xz && \
    cd /opt/llvm_src/build && \
    cmake -DCMAKE_BUILD_TYPE=RELEASE -DCMAKE_INSTALL_PREFIX=$LLVM_PATH $LLVM_SRC/llvm-8.0.0.src && \
    make -j4 && \
    make install

ENV PATH $LLVM_PATH:$PATH
ENV LD_LIBRARY_PATH $LLVM_PATH/libexec:$LD_LIBRARY_PATH
ENV LD_LIBRARY_PATH $LLVM_PATH/lib:$LD_LIBRARY_PATH
ENV LIBRARY_PATH $LLVM_PATH/libexec:$LIBRARY_PATH
ENV LIBRARY_PATH $LLVM_PATH/lib:$LIBRARY_PATH
ENV MANPATH $LLVM_PATH/share/man:$MANPATH
ENV C_INCLUDE_PATH $LLVM_PATH/include:$C_INCLUDE_PATH
ENV CPLUS_INCLUDE_PATH $LLVM_PATH/include:CPLUS_INCLUDE_PATH

ENV DOCKERAPP_VERSION=v0.6.0
ENV COMPOSE_VERSION=1.23.2
# Install Compose and Machine
RUN pip install docker-compose==${COMPOSE_VERSION}


RUN curl -fsSL --output /tmp/docker-app-linux.tar.gz https://github.com/docker/app/releases/download/${DOCKERAPP_VERSION}/docker-app-linux.tar.gz \
    && tar xf "/tmp/docker-app-linux.tar.gz" -C /usr/bin/ && mv /usr/bin/docker-app-linux /usr/bin/docker-app && rm /tmp/docker-app-linux.tar.gz

# Add bash completion and set bash as default shell
RUN mkdir /etc/bash_completion.d \
    && curl https://raw.githubusercontent.com/docker/cli/master/contrib/completion/bash/docker -o /etc/bash_completion.d/docker \
    && sed -i "s/ash/bash/" /etc/passwd
 
# Replace modprobe with a no-op to get rid of spurious warnings
# (note: we can't just symlink to /bin/true because it might be busybox)
RUN rm /sbin/modprobe && echo '#!/bin/true' >/sbin/modprobe && chmod +x /sbin/modprobe

# Install a nice vimrc file and prompt (by soulshake)
COPY ["docker-prompt", "sudo", "/usr/local/bin/"]
COPY [".vimrc", ".profile", ".inputrc", ".gitconfig", "./root/"]
COPY ["motd", "/etc/motd"]
COPY ["daemon.json", "/etc/docker/"]


# Move to our home
WORKDIR /root

# Setup certs and ssh keys
RUN mkdir -p /var/run/pwd/certs && mkdir -p /var/run/pwd/uploads \
    && ssh-keygen -N "" -t rsa -f  /etc/ssh/ssh_host_rsa_key >/dev/null \
    && mkdir ~/.ssh && ssh-keygen -N "" -t rsa -f ~/.ssh/id_rsa \
    && cat ~/.ssh/id_rsa.pub > ~/.ssh/authorized_keys

# Remove IPv6 alias for localhost and start docker in the background ...
CMD cat /etc/hosts >/etc/hosts.bak && \
    sed 's/^::1.*//' /etc/hosts.bak > /etc/hosts && \
    sed -i "s/\PWD_IP_ADDRESS/$PWD_IP_ADDRESS/" /etc/docker/daemon.json && \
    sed -i "s/\DOCKER_TLSENABLE/$DOCKER_TLSENABLE/" /etc/docker/daemon.json && \
    sed -i "s/\DOCKER_TLSCACERT/$DOCKER_TLSCACERT/" /etc/docker/daemon.json && \
    sed -i "s/\DOCKER_TLSCERT/$DOCKER_TLSCERT/" /etc/docker/daemon.json && \
    sed -i "s/\DOCKER_TLSKEY/$DOCKER_TLSKEY/" /etc/docker/daemon.json && \
    mount -t securityfs none /sys/kernel/security && \
    echo "root:root" | chpasswd &> /dev/null && \
    /usr/sbin/sshd -o PermitRootLogin=yes -o PrintMotd=no 2>/dev/null && \
    dockerd &>/docker.log & \
    while true ; do script -q -c "/bin/bash -l" /dev/null ; done
# ... and then put a shell in the foreground, restarting it if it exits
