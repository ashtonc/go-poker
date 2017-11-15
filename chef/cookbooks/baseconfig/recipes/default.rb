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
  environment 'GOPATH' => '/vagrant'
  command 'go get -u github.com/lib/pq'
end
execute 'get-mux' do
  environment 'GOPATH' => '/vagrant'
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
  cwd '/vagrant/src/poker/database'
  command 'sudo -u postgres psql pokerdb -f schema.sql'
end

# nginx setup
package "nginx"
execute 'nginx_pid' do
  command 'mkdir -p /run/nginx'
end
cookbook_file "nginx-config" do
  path "/etc/nginx/sites-available/default"
end
execute 'nginx_reload' do
  command 'nginx -s reload'
end

# Install tmux and start the server in the background.
package "tmux"
execute 'create-server-session' do
  cwd '/vagrant/src/poker'
  environment 'GOPATH' => '/vagrant'
  command 'tmux new-session -d -s server'
end
execute 'start-server' do
  command "tmux send-keys -t server 'go run poker.go' C-m"
end
