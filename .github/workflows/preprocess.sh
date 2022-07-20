if [ -z "$repo" ]; then
  repo=$PWD
fi

cat ${repo}/openapi.json > ${repo}/openapi.json.tmp
preprocessCount=$(cat ${repo}/bindgen-config.json | json generate.preprocess.length)
for (( i=0; i<$preprocessCount; i++))
do
  current=$(cat ${repo}/bindgen-config.json | json generate.preprocess.$i)
  key=$(echo "${current}" | json key)
  type=$(echo "${current}" | json type)
  value=$(echo "${current}" | json value)
  valueIsString=$(echo "${current}" | json -e "console.log(typeof this.value === 'string')" | head -n 1)
  if [ $valueIsString = true ]; then
    value='"'${value}'"'
  fi
  if [ $type = remove ]; then
    json -I -f ${repo}/openapi.json.tmp -e 'try { this.'"${key}"'=undefined; } catch(err) { console.log("Could not remove key: '${key}'"); }'
  elif [ $type = update ]; then
    json -I -f ${repo}/openapi.json.tmp -e 'try { this.'"${key}"'='"${value}"'; } catch(err) { console.log("Could not update key: '${key}'"); }'
  else
    echo Unknown preprocessor replacement type "${type}" for key "${key}"
    exit 1
  fi
done
changedVersion=$(cat ${repo}/openapi.json | json info.version)
echo '::set-output name=change::'"${changedVersion}"   
echo '::set-output name=random-ext::'"${RANDOM}"
