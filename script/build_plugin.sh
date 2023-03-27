#!/bin/bash
plugin_file=./script/plugin_list
if [ ! -f "$plugin_file" ]; then
  echo "plugin_list is not exist"
  exit 0
fi
num=0
for line in `cat $plugin_file`
do
  account=$line
  accounts[$num]=$account
  ((num++))
done
if [ $num -eq 0 ]; then
    echo "plugin_list is null"
    exit 0
fi
cmd="./answer build "
for repo in ${accounts[@]}
do
echo ${repo}
cmd=$cmd" --with "${repo}
done
$cmd
if [ ! -f "./new_answer" ]; then
  echo "new_answer is not exist build failed"
  exit 0
fi
rm answer
mv new_answer answer
./answer plugin