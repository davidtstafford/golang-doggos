# Get postgres docker image
docker pull postgres

#Directory to mount postgres storage
WORKING_DIR=$HOME/docker/volumes/postgres
if [ -d "$WORKING_DIR" ]; then rm -Rf $WORKING_DIR; fi
mkdir -p $WORKING_DIR

#Launch docker
docker run --rm   --name pg-docker -e POSTGRES_PASSWORD=postgres_docker -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data  postgres
