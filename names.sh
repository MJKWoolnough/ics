#!/bin/bash

declare -A names;

while read line; do
	names["$(echo "$line" | cut -d'=' -f1)"]="$(echo "$line" | cut -d'=' -f2)";
done < "names.gen";

function getName {
	name="${names[$1]}";
	if [ -z "$name" ]; then
		name="$1"
		echo -n "${name:0:1}$(echo "${name:1}" | tr A-Z a-z)";
		return;
	fi;
	echo -n "$name";
}
