FROM phusion/baseimage:0.9.22

# Use baseimage-docker's init system.
CMD ["/sbin/my_init"]

# Some system setup
RUN adduser --no-create-home --disabled-password --disabled-login --gecos '' gouser
EXPOSE 8080

# Add a server's binary
COPY go-test-task /usr/bin/go-test-task

# Make the service be launched by the init system
RUN mkdir /etc/service/go-test-task
COPY go-test-task-run.sh /etc/service/go-test-task/run

# Clean up APT when done.
RUN apt-get clean && rm -vrf /var/lib/apt/lists/* /tmp/* /var/tmp/*
