#!/bin/bash

declare -A comments;

while read line; do
	comments["$(echo "$line" | cut -d' ' -f1)"]="$line";
done < "comments.gen";

function getComment {
	(
		IFS=" ";
		comment="${comments[$1]}";
		if [ -z "$comment" ]; then
			comment="$1";
		fi;
		echo -n "//";
		local lineLength=2;
		for word in $comment; do
			let "lineLength += ${#word} + 1";
			if [ $lineLength -gt 80 ]; then
				echo;
				echo -n "//";
				let "lineLength = ${#word} + 3";
			fi;
			echo -n " $word";
		done;
		echo ".";
	)
}
