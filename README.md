Please refer to my blog - https://sql-cubed.com/golang-and-databases-part-3
*getting postgres docker up and running**
```bash
# Get postgres docker image
docker pull postgres

#Directory to mount postgres storage
WORKING_DIR=$HOME/docker/volumes/postgres
if [ -d "$WORKING_DIR" ]; then rm -Rf $WORKING_DIR; fi
mkdir -p $WORKING_DIR

#Launch docker
docker run --rm   --name pg-docker -e POSTGRES_PASSWORD=postgres_docker -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data  postgres
```

Settings envs

```bash
#Postgres  Environment Variables
export pgHost=localhost
export pgPort=5432
export pgUser=postgres
export pgPassword=postgres_docker
export pgDbName=postgres
```

Don't forget to create the demo schema and doggo table in postgres

All of this is covered in the first blog ;)
