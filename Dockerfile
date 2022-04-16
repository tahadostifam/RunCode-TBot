FROM ubuntu

RUN apt update

RUN apt install openssh-server net-tools nano -y

RUN apt install python3 ruby -y

RUN rm -f /etc/ssh/sshd_config

COPY sshd_config /etc/ssh/sshd_config

RUN service ssh start

EXPOSE 25