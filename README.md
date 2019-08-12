# Build the Docker image
1. In the root directory run `docker build -t accountservice .` to build the docker image
1. Run `docker run --rm accountservice` to confirm that the service is running on a container using the image that was just built

# Deploying on Docker Swarm
1. Ensure you have docker-machine installed and running by executing `docker-machine -v`
1. Run `docker-machine create swarm-manager` to create the swarm manager node
1. Run `docker $(docker-machine config swarm-manager) swarm init --advertise-addr $(docker-machine ip swarm-manager)` to initialize the swarm and register the `swarm-manager` docker machine as the manager node
1. Run `docker network create --driver overlay service-mesh` to create a network that the services will communicate over
1. Run `docker service create --name=accountservice --replicas=1 --network=service-mesh -p=6767:6767 accountservice`
1. You should be able to run `curl localhost:6767/accounts/10000` and get a JSON response back with the account struct
