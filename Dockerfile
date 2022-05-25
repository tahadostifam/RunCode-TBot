FROM ubuntu

RUN apt update

RUN apt install openssh-server net-tools nano -y

RUN apt install python3 ruby -y

CMD ["/usr/sbin/sshd","-D"]

EXPOSE 25