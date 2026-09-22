# Day 8 Build-in-Public Draft — First Linux EC2 Lab

## LinkedIn

Day 8 of my AWS cloud engineering journey was my first short-lived Linux EC2 lab.

I launched an Amazon Linux 2023 `t3.micro` in `eu-west-1` with an EC2 instance role carrying `AmazonSSMManagedInstanceCore`. I connected through Session Manager without creating an SSH key or opening an inbound security-group rule. The group allowed outbound traffic so the SSM Agent could reach AWS.

Inside the instance, I confirmed the SSM Agent was active, inspected the root filesystem (21% used), and counted two listening TCP sockets. I also learned that `ID_LIKE=fedora` in `/etc/os-release` describes a relationship to Fedora; the operating system was Amazon Linux 2023.

After the inspection, I terminated the instance and checked that no Day 8 EBS volume remained. The IAM role and security group are ready to reuse for the next lab.

My takeaway: instance access is a combination of an IAM role, network egress, and an agent running on the host. Session Manager let me inspect the machine without exposing SSH to the internet, while termination and volume checks closed the cost loop.

Next, I’m putting a small Go health service under `systemd` and checking it with `curl`, `ss`, and `journalctl`.

#AWS #CloudEngineering #EC2 #Linux #SystemsManager

## X

AWS Day 8: launched Amazon Linux 2023 on a t3.micro, connected through Session Manager with no SSH key or inbound rule, checked the host, then terminated it and confirmed no EBS volume remained. Next: a Go health service under systemd. #AWS #EC2 #Linux
