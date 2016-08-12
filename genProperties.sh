#!/bin/bash

source "names.sh";

declare currProperty="";
declare valueType=false;
declare -a values;
declare -a params;
function printProperty {
	if [ -z "$currProperty" ]; then
		return;
	fi;

	local tName="$(getName "$currProperty")";

	# typedef

	local mode=0;
	if [ ${#params[@]} -eq 0 ] && ! $valueType; then
		echo "type $tName uint8";
		echo;
		echo "const ("
		local first=true;
		for value in ${values[@]}; do
			echo -n "	$tName$(getName "$value")";
			if $first; then
				echo -n " $tName = iota";
				first=false;
			fi;
			echo;
		done;
		echo ")";
		mode=1;
	elif [ ${#params[@]} -eq 0 ] && $valueType && [ ${#values[@]} -eq 1 ]; then
		echo "type $tName ${values[0]}";
		mode=2;
	else
		echo "type $tName struct {";
		local longest=0;
		for param in ${params[@]}; do
			local n="$(getName "$param")";
			if [ ${#n} -gt $longest ]; then
				longest=${#n};
			fi;
		done;
		for value in ${values[@]}; do
			local n="$(getName "$value")";
			if [ ${#n} -gt $longest ]; then
				longest=${#n};
			fi;
		done;
		for param in ${params[@]}; do
			local n=$(getName "$param")
			echo -n "	$n ";
			for i in $(seq $(( $longest - ${#n} ))); do
				echo -n " ";
			done;
			if [ "$param" = "DELEGATED-FROM" -o "$param" = "DELEGATED-TO" -o "$param" = "MEMBER" ]; then
				echo "$n";
			else
				echo "*$n";
			fi;
		done;
		if [ ${#values[@]} -eq 1 ]; then
			echo "	${values[0]}";
		else
			for value in ${values[@]}; do
				local n=$(getName "$value")
				echo -n "	$n ";
				for i in $(seq $(( $longest - ${#n} ))); do
					echo -n " ";
				done;
				if [ "$value" = "Binary" -o "$value" = "MText" ]; then
					echo "$n";
				else
					echo "*$n";
				fi;
			done;
		fi;
		echo "}";
	fi;
	echo;

	# decoder

	echo "func (p *$tName) decode(params []parser.Token, value string) error {";
	case $mode in
	0)
		if [ ${#values[@]} -gt 1 ]; then
			echo "	vType := -1";
		fi;
		echo "	oParams := make(map[string]string)";
		echo "	var ts []string";
		echo "	for len(params) > 0 {";
		echo "		pName := strings.ToUpper(params[0].Data)";
		echo "		i := 1";
		echo "		for i < len(params) && params[i].Type != tokenParamName {";
		echo "			i++";
		echo "		}";
		echo "		pValues := params[1:i]";
		echo "		params = params[i:]";
		echo "		switch pName {";
		for param in ${params[@]}; do
			local tParam="$(getName "$param")";
			echo "		case \"$param\":";
			echo "			if p.$tParam != nil {";
			echo "				return ErrDuplicateParam";
			echo "			}";
			if [ "$param" != "DELEGATED-FROM" -a "$param" != "DELEGATED-TO" -a "$param" != "MEMBER" ]; then
				echo "			p.$tParam = new($tParam)";
			fi;
			echo "			if err := p.${tParam}.decode(pValues); err != nil {";
			echo "				return err";
			echo "			}";
		done;
		if [ ${#values[@]} -gt 1 ]; then
			echo "		case \"VALUE\":";
			echo "			if len(pValues) != 1 {";
			echo "				return ErrInvalidValue";
			echo "			}";
			echo "			if vType != -1 {";
			echo "				return ErrDuplicateParam";
			echo "			}";
			echo "			switch strings.ToUpper(pValues[0].Data) {";
			local i=0;
			for value in ${values[@]}; do
				echo "			case \"$value\":";
				echo "				vType = $i";
				let "i++";
			done;
			echo "			default:";
			echo "				return ErrInvalidValue";
			echo "			}";
		fi;
		echo "		default:";
		echo "			for _, v := range pValues {";
		echo "				ts = append(ts, v.Data)";
		echo "			}";
		echo "			oParams[pName] = strings.Join(ts, \",\")";
		echo "			ts = ts[:0]";
		echo "		}";
		echo "	}";
		if [ ${#values[@]} -gt 1 ]; then
			echo "	if vType == -1 {";
			echo "		vType = 0";
			echo "	}";
			echo "	switch vType {";
			local i=0;
			for value in ${values[@]}; do
				local tValue="$(getName "$value")";
				echo "	case $i:";
				if [ "$value" != "Binary" -a "$value" != "MText" ]; then
					echo "		p.$tValue = new($tValue)";
				fi;
				echo "		if err := p.${tValue}.decode(oParams, value); err != nil {";
				echo "			return err";
				echo "		}";
				let "i++";
			done;
			echo "	}";
		else
			echo "	if err := p.$(getName "${values[0]}").decode(oParams, value); err != nil {";
			echo "		return err";
			echo "	}";
		fi;;
	1)
		echo "	switch strings.ToUpper(value) {";
		for value in ${values[@]}; do
			echo "	case \"$value\":";
			echo "		*p = $tName$(getName "$value")";
		done;
		echo "	default:";
		echo "		return ErrInvalidValue";
		echo "	}";;
	2)
		echo "	var t ${values[0]}";
		echo "	if err := t.decode(value); err != nil {";
		echo "		return err";
		echo "	}";
		echo "	*p = $tName(t)";
	esac;
	echo "	return nil";
	echo "}";
	echo;

	# encoder

	echo "func (p *$tName) encode(w writer) {";
	case $mode in
	0)
		echo "	w.WriteString(\"$currProperty\")";
		for param in ${params[@]}; do
			tParam="$(getName "$param")";
			echo "	if p.$tParam != nil {";
			echo "		p.${tParam}.encode(w)";
			echo "	}";
		done;
		if [ ${#values[@]} -gt 1 ]; then
			for value in ${values[@]}; do
				tValue="$(getName "$value")";
				echo "	if p.$tValue != nil {";
				echo "		p.${tValue}.aencode(w)";
				echo "	}";
			done;
		else
			echo "	p.$(getName "${values[0]}").aencode(w)";
		fi;;
	1)
		echo "	w.WriteString(\"$currProperty:\")";
		echo "	switch *p {";
		for value in ${values[@]}; do
			echo "	case $(getName "$value"):";
			echo "		w.WriteString(\"$value\")";
		done;
		echo "	}";;
	2)
		echo "	w.WriteString(\"$currProperty\")";
		echo "	t := ${values[0]}(*p)";
		echo "	t.aencode(w)";
	esac;
	echo "	w.WriteString(\"\\r\\n\")";
	echo "}";
	echo;


	# validator


	echo "func (p *$tName) valid() error {";
	case $mode in
	0)
		for param in ${params[@]}; do
			tParam="$(getName "$param")";
			echo "	if p.$tParam != nil {";
			echo "		if err := p.${tParam}.valid(); err != nil {";
			echo "			return err";
			echo "		}";
			echo "	}";
		done;
		if [ ${#values[@]} -gt 1 ]; then
			echo "	c := 0";
			for value in ${values[@]}; do
				tValue="$(getName "$value")";
				echo "	if p.$tValue != nil {";
				echo "		if err := p.${tValue}.valid(); err != nil {";
				echo "			return err";
				echo "		}";
				echo "		c++";
				echo "	}";
			done;
			echo "	if c != 1 {";
			echo "		return ErrInvalidValue";
			echo "	}";
		else
			echo "	if err := p.$(getName "${values[0]}").valid(); err != nil {";
			echo "		return err";
			echo "	}";
		fi;
		echo "	return nil";;
	1)
		echo "	switch *p {";
		echo -n "	case ";
		local first=false;
		for value in ${values[@]}; do
			if $first; then
				echo -n ", ";
			fi;
			first=true;
			echo -n "$(getName "$value")";
		done;
		echo ":";
		echo "	default:";
		echo "		return ErrInvalidValue";
		echo "	}";
		echo "	return nil";;
	2)
		echo "	t := ${values[0]}(*p)";
		echo "	return t.valid()";
	esac;
	echo "}";
	echo;

	# reset

	currProperty="";
	valueType=false;
	values=();
	params=();
}

(
	echo "package ics";
	echo;
	echo "// File automatically generated with ./genParams.sh";
	echo;
	echo "import (";
	echo "	\"strings\"";
	echo ")";
	echo;
	{
		IFS="
";
		while read line;do
			if [ "${line:0:1}" != "	" ]; then
				printProperty;
				currProperty="$(echo "$line" | cut -d':' -f1)";
				vs="$(echo "$line" | cut -d':' -f2)";
				if [ "${vs:0:1}" = "!" ]; then
					valueType=true;
					vs="${vs:1}";
				fi;
				values=( $(echo -n "$vs" | tr "|" "\n") );
			else
				params[${#params[@]}]="${line:1}";
			fi;
		done;
	} < "properties.gen";
	printProperty;
	echo "// Errors";
	echo "var (";
	echo "	ErrInvalidValue = errors.New(\"invalid value\")";
	echo ")";
) > "properties.go";
