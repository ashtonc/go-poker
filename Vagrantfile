# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "ubuntu/xenial64"
  config.vm.box_version = '>= 20160921.0.0'

  config.vm.network "forwarded_port", guest: 80, host: 8000

  config.vm.synced_folder "./", "/vagrant/src/poker"

  config.ssh.shell = "bash -c 'BASH_ENV=/etc/profile exec bash'"

  cpus = "1"
  memory = "512" # MB
  config.vm.provider :virtualbox do |vb|
    vb.customize ["modifyvm", :id, "--cpus", cpus, "--memory", memory]
    vb.customize ["modifyvm", :id, "--uartmode1", "disconnected"] # speed up boot https://bugs.launchpad.net/cloud-images/+bug/1627844
    #vb.gui = true
  end
  
  config.vm.provider "vmware_fusion" do |v, override|
    v.vmx["memsize"] = memory
    v.vmx["numvcpus"] = cpus
  end

  config.vm.provision "chef_solo" do |chef|
    chef.log_level = "debug"
    chef.cookbooks_path = "chef/cookbooks"
    chef.channel = "stable"
    chef.version = "12.10.24"
    chef.add_recipe "baseconfig"
  end
end
