#!/bin/bash

source "names.sh";
source "comments.sh";

(
	declare -A regexes;
	echo "package ics";
	echo;
	echo "// File automatically generated with ./genParams.sh";
	echo;
	echo "import (";
	echo "	\"errors\"";
	echo "	\"fmt\"";
	echo "	\"regexp\"";
	echo "	\"strings\"";
	echo "	\"unicode/utf8\"";
	echo;
	echo "	\"vimagination.zapto.org/parser\"";
	echo ")";
	echo;
	{
		while read line; do
			keyword="$(echo "$line" | cut -d'=' -f1)";
			type="$(getName "$keyword")";
			values="$(echo "$line" | cut -d'=' -f2)";

			getComment "$type";
			echo -n "type Param$type ";

			declare multiple=false;
			declare freeChoice=false;
			declare doubleQuote=false;
			declare regex="";
			declare vType="";
			declare string=false;
			declare -a choices=();
			fc="${values:0:1}";
			if [ "$fc" = "*" ]; then
				echo -n "[]";
				multiple=true
				values="${values:1}";
				fc="${values:0:1}";
			fi;
			if [ "$fc" = "?" ]; then
				freeChoice=true
				values="${values:1}";
				fc="${values:0:1}";
			fi;
			if [ "$fc" = '"' ]; then
				doubleQuote=true;
				values="${values:1}";
				fc="${values:0:1}";
				string=true;
			elif [ "$fc" = "'" ]; then
				values="${values:1}";
				string=true;
				fc="${values:0:1}";
			elif [ "$fc" = "~" ]; then
				regex="${values:1}";
				string=true;
				values="${values:1}";
				fc="${values:0:1}";
			fi;
			if [ "$fc" = "!" ]; then
				values="${values:1}";
				echo "$values";
				vType="$values";
				if [ "$vType" = "Boolean" ]; then
					echo;
					echo "// New$type returns a *Param$type for ease of use with optional values";
					echo "func New$type(v Param$type) *Param$type {";
					echo "	return &v";
					echo "}";
				fi;
			elif $string; then
				echo "string";
				echo;
				echo "// New$type returns a *Param$type for ease of use with optional values";
				echo "func New$type(v Param$type) *Param$type {";
				echo "	return &v";
				echo "}";
				if [ ! -z "$regex" ]; then
					echo;
					echo "var regex$type *regexp.Regexp";
					regexes[$type]="$values";
				fi;
			else
				if $freeChoice; then
					choices=( $(echo "Unknown|$values" | tr "|" " ") );
				else
					choices=( $(echo "$values" | tr "|" " ") );
				fi;
				case ${#choices[@]} in
				1)
					echo "struct{}";;
				*)
					echo "uint8";
					echo;
					echo "// $type constant values";
					echo "const (";
					declare first=true;
					for choice in ${choices[@]};do
						echo -n "	$type$(getName "$choice")";
						if $first; then
							echo -n " Param$type = iota";
							first=false;
						fi;
						echo;
					done;
					echo ")";
					echo;
					echo "// New returns a pointer to the type (used with constants for ease of use with";
					echo "// optional values)";
					echo "func (t Param$type) New() *Param$type {";
					echo "	return &t";
					echo "}";
				esac;
				choices=( $(echo "$values" | tr "|" " ") );
			fi;
			echo;

			# decoder

			echo "func (t *Param$type) decode(vs []parser.Token) error {";
			declare indent="";
			declare vName="vs[0]";
			if $multiple; then
				echo "	for _, v := range vs {";
				indent="	";
				vName="v";
			else
				echo "	if len(vs) != 1 {";
				echo "		return fmt.Errorf(errDecodingType, c$type, ErrInvalidParam)";
				echo "	}";
			fi;
			if $doubleQuote; then
				echo "$indent	if ${vName}.Type != tokenParamQuotedValue {";
				echo "$indent		return fmt.Errorf(errDecodingType, c$type, ErrInvalidParam)";
				echo "$indent	}";
			fi;
			if [ ! -z "$vType" ]; then
				echo "$indent	var q $vType";
				echo "$indent	if err := q.decode(nil, ${vName}.Data); err != nil {";
				echo "$indent		return fmt.Errorf(errDecodingType, c$type, err)";
				echo "$indent	}";
				if $multiple; then
					echo "		*t = append(*t, q)";
				else
					echo "	*t = Param$type(q)";
				fi;
			elif [ ${#choices[@]} -eq 1 ]; then
				echo "	if strings.ToUpper(${vName}.Data) != \"${choices[0]}\" {";
				echo "		return fmt.Errorf(errDecodingType, c$type, ErrInvalidParam)";
				echo "	}";
			elif [ ${#choices[@]} -gt 1 ]; then
				echo "$indent	switch strings.ToUpper(${vName}.Data) {";
				for choice in ${choices[@]}; do
					echo "$indent	case \"$choice\":";
					if $multiple; then
						echo "		*t = append(*t, $type$(getName "$choice")";
					else
						echo "		*t = $type$(getName "$choice")";
					fi;
				done;
				echo "$indent	default:";
				if $freeChoice; then
					if $multiple; then
						echo "		*t = append(*t, {$type}Unknown)";
					else
						echo "		*t = ${type}Unknown";
					fi;
				else
					echo "$indent		return fmt.Errorf(errDecodingType, c$type, ErrInvalidParam)";
				fi;
				echo "$indent	}";
			else
				if [ -z "$regex" ]; then
					if $multiple; then
						echo "		*t = append(*t, decode6868(${vName}.Data))";
					else
						echo "	*t = Param$type(decode6868(${vName}.Data))";
					fi;
				else
					echo "$indent	if !regex${type}.MatchString(${vName}.Data) {";
					echo "$indent		return fmt.Errorf(errDecodingType, c$type, ErrInvalidParam)";
					echo "$indent	}";
					echo "$indent	*t = Param$type(${vName}.Data)";
				fi;
			fi;
			if $multiple; then
				echo "	}";
			fi;
			echo "	return nil";
			echo "}";
			echo;

			#encoder

			echo "func (t Param$type) encode(w writer) {";
			if [ ${#choices} -eq 0 ] || $multiple; then
				if [ "$vType" = "CALADDRESS" -o "$vType" = "URI" ]; then
					echo "	if len(t.String()) == 0 {";
					echo "		return";
					echo "	}";
				elif [ "$vType" = "Boolean" ]; then
					echo "	if !t {";
					echo "		return";
					echo "	}";
				else
					echo "	if len(t) == 0 {";
					echo "		return";
					echo "	}";
				fi;
			fi;
			echo "	w.WriteString(\";${keyword}=\")";
			if $multiple; then
				echo "	for n, v := range t {";
				echo "		if n > 0 {";
				echo "			w.WriteString(\",\")";
				echo "		}";
			else
				vName="t";
			fi;
			if [ ! -z "$vType" ]; then
				echo "$indent	q := $vType($vName)";
				echo "$indent	q.encode(w)";
			elif [ ${#choices[@]} -eq 1 ]; then
				echo "$indent	w.WriteString(\"${choices[0]}\")";
				freeChoice=true;
			elif [ ${#choices[@]} -gt 1 ]; then
				echo "$indent	switch $vName {";
				for choice in ${choices[@]}; do
					echo "$indent	case $type$(getName "$choice"):";
					echo "$indent		w.WriteString(\"$choice\")";
				done;
				if $freeChoice; then
					echo "$indent	default:";
					echo "$indent		w.WriteString(\"UNKNOWN\")";
				fi;
				echo "$indent	}";
			else
				if $doubleQuote; then
					echo "$indent	w.WriteString(\"\\\"\")";
					echo "$indent	w.Write(encode6868(string($vName)))";
					echo "$indent	w.WriteString(\"\\\"\")";
				else
					echo "$indent	if strings.ContainsAny(string($vName), nonsafeChars[32:]) {";
					echo "$indent		w.WriteString(\"\\\"\")";
					echo "$indent		w.Write(encode6868(string($vName)))";
					echo "$indent		w.WriteString(\"\\\"\")";
					echo "$indent	} else {";
					echo "$indent		w.Write(encode6868(string($vName)))";
					echo "$indent	}";
				fi;
			fi;
			if $multiple; then
				echo "	}";
			fi;
			echo "}";
			echo;

			#validator

			echo "func (t Param$type) valid() error {";
			if [ "$vType" = "Boolean" ]; then
				echo "	return nil";
			elif [ ${#choices[@]} -eq 0 ] || ! $freeChoice; then
				if $multiple; then
					echo "	for _, v := range t {";
				fi;
				if [ ! -z "$vType" ]; then
					if $multiple; then
						echo "		if err := v.valid(); err != nil {"
						echo "			return fmt.Errorf(errValidatingType, c$type, err)";
						echo "		}";
					else
						echo "	q := $vType(t)";
						echo "	if err := q.valid(); err != nil {";
						echo "		return fmt.Errorf(errValidatingType, c$type, err)";
						echo "	}";
						echo "	return nil";
					fi;
				elif [ ${#choices[@]} -gt 0 ]; then
					echo "$indent	switch $vName {";
					echo -n "$indent	case ";
					first=false;
					for choice in ${choices[@]}; do
						if $first; then
							echo -n ", ";
						fi;
						first=true;
						echo -n "$type$(getName "$choice")";
					done;
					echo ":";
					echo "$indent	default:";
					echo "$indent		return fmt.Errorf(errValidatingType, c$type, ErrInvalidValue)";
					echo "$indent	}";
				elif [ ! -z "$regex" ]; then
					echo "$indent	if !regex${type}.Match([]byte($vName)) {";
					echo "$indent		return fmt.Errorf(errValidatingType, c$type, ErrInvalidValue)";
					echo "$indent	}";
				else
					echo "$indent	if strings.ContainsAny(string($vName), nonsafeChars[:31]) {";
					echo "$indent		return fmt.Errorf(errValidatingType, c$type, ErrInvalidText)";
					echo "$indent	}";
				fi;
				if $multiple; then
					echo "	}";
				fi;
				if [ -z "$vType" ] || $multiple; then
					echo "	return nil";
				fi;
			else
				echo "	return nil";
			fi;
			echo "}";
			echo;
		done;
	} < params.gen

	cat <<HEREDOC
func decode6868(s string) string {
	t := parser.NewStringTokeniser(s)
	d := make([]byte, 0, len(s))
	var ru [4]byte
Loop:
	for {
		c := t.ExceptRun("^")
		d = append(d, t.Get()...)
		switch c {
		case -1:
			break Loop
		case '^':
			t.Accept("^")
			switch t.Peek() {
			case -1:
				d = append(d, '^')
				break Loop
			case 'n':
				d = append(d, '\n')
			case '\'':
				d = append(d, '"')
			case '^':
				d = append(d, '^')
			default:
				d = append(d, '^')
				l := utf8.EncodeRune(ru[:], c)
				d = append(d, ru[:l]...)
			}
			t.Except("")
		}
	}
	return string(d)
}

func encode6868(s string) []byte {
	t := parser.NewStringTokeniser(s)
	d := make([]byte, 0, len(s))
Loop:
	for {
		c := t.ExceptRun("\n^\"")
		d = append(d, t.Get()...)
		switch c {
		case -1:
			break Loop
		case '\n':
			d = append(d, '^', 'n')
		case '^':
			d = append(d, '^', '^')
		case '"':
			d = append(d, '^', '\'')
		}
	}
	return d
}

HEREDOC

	echo "func init() {";
	for key in ${!regexes[@]}; do
		echo "	regex$key = regexp.MustCompile(\"${regexes[$key]}\")";
	done;
	echo "}";
	echo;
	echo "// Errors";
	echo "var (";
	echo "	ErrInvalidParam = errors.New(\"invalid param value\")";
	echo "	ErrInvalidValue = errors.New(\"invalid value\")";
	echo ")";
	echo;
	echo "const ("
	echo "	errDecodingType            = \"error decoding %s: %w\"";
	echo "	errValidatingType          = \"error decoding %s: %w\"";
	{
		while read line; do
			keyword="$(echo "$line" | cut -d'=' -f1)";
			type="$(getName "$keyword")";
			echo -n "	c$type";
			for i in $(seq $(( 26 - ${#type} ))); do
				echo -n " ";
			done;
			echo "= \"$type\"";
		done;
	} < params.gen
	echo ")";
) > params.go
