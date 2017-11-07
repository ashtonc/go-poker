# Make sure the apt package lists are up to date.
cookbook_file "apt-sources.list" do
  path "/etc/apt/sources.list"
end
execute 'apt_update' do
  command 'apt-get update'
end

# Base configuration recipe in Chef.
package "wget"
package "ntp"
cookbook_file "ntp.conf" do
  path "/etc/ntp.conf"
end
execute 'ntp_restart' do
  command 'service ntp restart'
end

# Go installation and Go dependencies installation.
package "golang"
execute 'get-pq' do
  environment 'GOPATH' => '/go'
  command 'go get -u github.com/lib/pq'
end
execute 'get-mux' do
  environment 'GOPATH' => '/go'
  command 'go get -u github.com/gorilla/mux'
end

# Postgres setup.
package "postgresql"
execute 'postgres-setup' do
  command 'echo "CREATE DATABASE pokerdb;" | sudo -u postgres psql'
end
execute 'postgres-set-password' do
  command 'echo "ALTER USER postgres WITH PASSWORD \'postgres\';" | sudo -u postgres psql'
end
execute 'database-setup' do
  cwd '/vagrant'
  command 'sudo -u postgres psql pokerdb -f schema.sql'
end

# Install tmux and start the server in the background.
package "tmux"
execute 'create-server-session' do
  cwd '/vagrant'
  environment 'GOPATH' => '/go'
  command 'tmux new-session -d -s server'
end
execute 'start-server' do
  command "tmux send-keys -t server 'go run poker.go' C-m"
end

# To access the program/prompt from inside the vm:
# vagrant ssh
# sudo su
# tmux attach -t server