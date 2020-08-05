# generate docker run environment file
dockerRunEnvList=./env_docker_run.list
bash -x env.sh 2>${dockerRunEnvList}
sed -i 's/+ //' ${dockerRunEnvList}
sed -i '/^export /d' ${dockerRunEnvList}
sed -i "s/'//g" ${dockerRunEnvList}
