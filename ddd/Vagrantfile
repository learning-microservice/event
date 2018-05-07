# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.box = "ubuntu/trusty64"

  config.vm.provider "virtualbox" do |v|
    v.name = "learning-microservice-event-ddd"
    v.memory = 2048
    v.cpus = 2
  end

  config.vm.box_check_update = true

#  config.vm.synced_folder ".", "/vagrant", :nfs => true

  config.vm.network "private_network", ip: "192.168.33.101"

  # port forwardings
  config.vm.network "forwarded_port", id: "portainer ",      host: 9000,  guest: 9000,  auto_correct: true
  config.vm.network "forwarded_port", id: "mysql ",          host: 3306,  guest: 3306,  auto_correct: true

  # ssh
  config.ssh.forward_agent = true
  config.ssh.shell = "bash -c 'BASH_ENV=/etc/profile exec bash'"

  # provision
  config.vm.provision "shell", inline: <<-SHELL
    touch /var/lib/cloud/instance/locale-check.skip
    apt-key adv --keyserver hkp://pgp.mit.edu --recv-keys 58118E89F3A912897C070ADBF76221572C52609D
    sh -c 'echo "deb https://apt.dockerproject.org/repo ubuntu-trusty main" > /etc/apt/sources.list.d/docker.list'
    apt-cache policy docker-engine
    apt-get update
    apt-get upgrade -y
    apt-get install -y docker-engine python-pip jq
    pip install docker-compose
    usermod -aG docker vagrant
    sudo -u vagrant pip install docker-compose
    docker-compose -f /vagrant/docker-compose.yml up -d --build
    echo "cd /vagrant" >> /home/vagrant/.profile
  SHELL

  if File.file?("custom-provision.sh")
    config.vm.provision "shell", path: "custom-provision.sh"
  end
end
