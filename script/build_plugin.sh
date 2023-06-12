#!/bin/bash
set -e
echo "begin build plugin"
plugin_file=./script/plugin_list
if [ ! -f "$plugin_file" ]; then
  echo "plugin_list is not exist"
  exit 0
fi

echo "plugin_list exist"
cmd="./answer build "
for repo in `cat $plugin_file`
do
  echo ${repo}
  cmd=$cmd" --with "${repo}
done

echo "cmd is "$cmd
$cmd
if [ ! -f "./new_answer" ]; then
  echo "new_answer is not exist build failed"
  exit 1
fi
rm answer
mv new_answer answer
./answer plugin