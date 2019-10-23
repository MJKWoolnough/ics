#!/bin/bash

source "names.sh";
source "comments.sh";

declare -a currSection;
declare -a requirements;
declare sectionName;
declare longest=0;

function addToSection {
	declare name="$(getName "${1%#*}")";
	currSection[${#currSection[@]}]="$name ${1##*#} $2 $3 $4";
	local l=${#name};
	if [ $l -gt $longest ]; then
		longest=$l;
	fi;
}

function printSection {
	declare sName="$(getName $sectionName)";

	# type declaration

	getComment "$sName";
	declare checkRequired=false;
	if [ ${#currSection[@]} -eq 0 ]; then
		echo "type $sName struct{}";
	else
		echo "type $sName struct {";
		IFS="$OFS";
		for tline in "${currSection[@]}"; do
			aline=( $tline ); # 0:name 1:KEYWORD 2:required 3:multiple 4:section 5:requiredAlso 6:requiredInstead
			name="${aline[0]}";
			required=${aline[2]};
			multiple=${aline[3]};
			section=${aline[4]};
			if $required; then
				checkRequired=true;
			fi;
			echo -n "	$name ";
			for i in $(seq $(( $longest - ${#name} ))); do
				echo -n " ";
			done;
			if $multiple; then
				echo -n "[]";
			elif ! $required; then
				echo -n "*";
			fi;
			if ! $section; then
				echo -n "Prop";
			fi;
			echo "$name";
		done;
		echo "}";
	fi;
	echo;

	# decoder

	echo "func (s *$sName) decode(t tokeniser) error {";

	# required bools
	if [ ${#currSection[@]} -gt 0 ]; then
		echo -n "	var";
		declare first=false;
		for tline in "${currSection[@]}"; do
			aline=( $tline ); # 0:name 1:KEYWORD 2:required 3:multiple 4:section 5:requiredAlso 6:requiredInstead
			name="${aline[0]}";
			required=${aline[2]};
			multiple=${aline[3]};
			if $required && ! $multiple; then
				if $first; then
					echo -n ",";
				fi;
				first=true;
				echo -n " required$name";
			fi;
		done;
		echo " bool";
	fi;

	# type switch

	echo "Loop:";
	echo "	for {"
	echo "		p, err := t.GetPhrase()";
	echo "		if err != nil {";
	echo "			return fmt.Errorf(errDecodingType, c$sName, err)";
	echo "		} else if p.Type == parser.PhraseDone {";
	echo "			return fmt.Errorf(errDecodingType, c$sName, io.ErrUnexpectedEOF)";
	echo "		}";
	if [ ${#currSection[@]} -gt 0 ]; then
		echo "		params := p.Data[1 : len(p.Data)-1]";
	fi;
	echo "		value := p.Data[len(p.Data)-1].Data";
	echo "		switch strings.ToUpper(p.Data[0].Data) {";
	echo "		case \"BEGIN\":";
	echo "			switch n := strings.ToUpper(value); n {";

	# BEGIN keywords

	for tline in "${currSection[@]}"; do
		aline=( $tline ); # 0:name 1:KEYWORD 2:required 3:multiple 4:section 5:requiredAlso 6:requiredInstead
		name="${aline[0]}";
		keyword="${aline[1]}";
		required=${aline[2]};
		multiple=${aline[3]};
		section=${aline[4]};
		if ! $section; then
			continue;
		fi;
		echo "			case \"$keyword\":";
		if $required && ! $multiple; then
			echo "				if required$name {";
			echo "					return fmt.Errorf(errMultiple, c$sName, c$name)";
			echo "				}";
			echo "				required$name = true";
			echo "				if err := s.${name}.decode(t); err != nil {";
			echo "					return fmt.Errorf(errDecodingProp, c$sName, c$name, err)";
			echo "				}";
		elif $multiple; then
			echo "				var e $name";
			echo "				if err := e.decode(t); err != nil {";
			echo "					return fmt.Errorf(errDecodingProp, c$sName, c$name, err)";
			echo "				}";
			echo "				s.$name = append(s.$name, e)";
		else
			echo "				if s.$name != nil {";
			echo "					return fmt.Errorf(errMultiple, c$sName, c$name)";
			echo "				}";
			echo "				s.$name = new($name)";
			echo "				if err := s.${name}.decode(t); err != nil {";
			echo "					return fmt.Errorf(errDecodingProp, c$sName, c$name, err)";
			echo "				}";
		fi;
	done;
	echo "			default:";
	echo "				if err := decodeDummy(t, n); err != nil {";
	echo "					return fmt.Errorf(errDecodingType, c$sName, err)";
	echo "				}";
	echo "			}";

	# non-BEGIN keywords

	for tline in "${currSection[@]}"; do
		aline=( $tline ); # 0:name 1:KEYWORD 2:required 3:multiple 4:section 5:requiredAlso 6:requiredInstead
		name="${aline[0]}";
		keyword="${aline[1]}";
		required=${aline[2]};
		multiple=${aline[3]};
		section=${aline[4]};
		if $section; then
			continue;
		fi;
		echo "		case \"$keyword\":";
		if $required && ! $multiple; then
			echo "			if required$name {";
			echo "				return fmt.Errorf(errMultiple, c$sName, c$name)";
			echo "			}";
			echo "			required$name = true";
			echo "			if err := s.${name}.decode(params, value); err != nil {";
			echo "				return fmt.Errorf(errDecodingProp, c$sName, c$name, err)";
			echo "			}";
		elif $multiple; then
			echo "			var e Prop$name";
			echo "			if err := e.decode(params, value); err != nil {";
			echo "				return fmt.Errorf(errDecodingProp, c$sName, c$name, err)";
			echo "			}";
			echo "			s.$name = append(s.$name, e)";
		else
			echo "			if s.$name != nil {";
			echo "				return fmt.Errorf(errMultiple, c$sName, c$name)";
			echo "			}";
			echo "			s.$name = new(Prop$name)";
			echo "			if err := s.${name}.decode(params, value); err != nil {";
			echo "				return fmt.Errorf(errDecodingProp, c$sName, c$name, err)";
			echo "			}";
		fi;
	done;
	echo "		case \"END\":"
	if [ "${sectionName:0:6}" = "VALARM" ]; then
		echo "			if value != \"VALARM\" {";
	else
		echo "			if value != \"$sectionName\" {";
	fi;
	echo "				return fmt.Errorf(errDecodingType, c$sName, ErrInvalidEnd)"
	echo "			}";
	echo "			break Loop";
	echo "		}";
	echo "	}";

	# check required bools
	
	if $checkRequired; then
		first=false;
		echo -n "	if";
		for tline in "${currSection[@]}"; do
			aline=( $tline ); # 0:name 1:KEYWORD 2:required 3:multiple 4:section 5:requiredAlso 6:requiredInstead
			name="${aline[0]}";
			required=${aline[2]};
			multiple=${aline[3]};
			if $required; then
				if $first; then
					echo -n " ||";
				fi;
				first=true;
				if $multiple; then
					echo -n " len(s.$name) == 0";
				else
					echo -n " !required$name";
				fi;
			fi;
		done;
		echo " {";
		echo "		return fmt.Errorf(errDecodingType, c$sName, ErrMissingRequired)";
		echo "	}";
	fi;

	# check other requirements
	
	for req in "${requirements[@]}"; do
		first=false
		declare second=false
		echo -n "	if";
		reqs=( $req );
		typ=${reqs[0]};
		reqs[0]="";
		reqs=( ${reqs[@]} );
		if [ "$typ" = "AND" ]; then
			for r in ${reqs[@]}; do
				if $second; then
					echo -n " &&";
				fi;
				if $first; then
					echo -n " t == (s.$(getName "$r") == nil)";
					second=true;
				else
					echo -n " t := s.$(getName "$r") == nil;";
					first=true;
				fi;
			done;
		elif [ "$typ" = "OR" ]; then
			for r in ${reqs[@]}; do
				if $first; then
					echo -n " &&";
				fi;
				first=true;
				echo -n " s.$(getName "$r") == nil";
			done;
		elif [ "$typ" = "ONE" ]; then
			for r in ${reqs[@]}; do
				if $first; then
					echo -n " &&";
				fi;
				first=true;
				echo -n " s.$(getName "$r") != nil";
			done;
		elif [ "$typ" = "ERGO" ]; then
			for r in ${reqs[@]}; do
				if $second; then
					echo -n " || ";
				fi;
				if $first; then
					echo -n "s.$(getName "$r") == nil";
					second=true;
				else
					echo -n " s.$(getName "$r") != nil && (";
					first=true;
				fi;
			done;
			echo -n ")";
		fi;
		echo " {";
		echo "		return fmt.Errorf(errDecodingType, c$sName, ErrRequirementNotMet)";
		echo "	}";
	done;

	# end of decoder

	echo "	return nil";
	echo "}";
	echo;

	# encoder

	echo "func (s *$sName) encode(w writer) {"
	if [ "${sectionName:0:6}" != "VALARM" ]; then
		echo "	w.WriteString(\"BEGIN:$sectionName\r\n\")";
	fi;
	for tline in "${currSection[@]}"; do
		aline=( $tline ); # 0:name 1:KEYWORD 2:required 3:multiple 4:section 5:requiredAlso 6:requiredInstead
		name="${aline[0]}";
		required=${aline[2]};
		multiple=${aline[3]};
		if $multiple; then
			echo "	for n := range s.$name {";
			echo "		s.$name[n].encode(w)";
			echo "	}";
		elif $required; then
			echo "	s.${name}.encode(w)";
		else
			echo "	if s.$name != nil {";
			echo "		s.${name}.encode(w)";
			echo "	}";
		fi;
	done;
	if [ "${sectionName:0:6}" != "VALARM" ]; then
		echo "	w.WriteString(\"END:$sectionName\r\n\")";
	fi;
	echo "}";
	echo;

	# validator

	echo "func (s *$sName) valid() error {";
	for tline in "${currSection[@]}"; do
		aline=( $tline ); # 0:name 1:KEYWORD 2:required 3:multiple 4:section 5:requiredAlso 6:requiredInstead
		name="${aline[0]}";
		required=${aline[2]};
		multiple=${aline[3]};
		if $multiple; then
			if $required; then
				echo "	if len(s.$name) == 0 {";
				echo "		return fmt.Errorf(errMissing, c$sName, c$name)";
				echo "	}";
			fi;
			echo "	for n := range s.$name {";
			echo "		if err := s.$name[n].valid(); err != nil {";
			echo "			return fmt.Errorf(errValidatingProp, c$sName, c$name, err)";
			echo "		}";
			echo "	}";
		else
			if $required; then
				echo "	if err := s.${name}.valid(); err != nil {";
				echo "		return fmt.Errorf(errValidatingProp, c$sName, c$name, err)";
				echo "	}";
			else
				echo "	if s.$name != nil {";
				echo "		if err := s.${name}.valid(); err != nil {";
				echo "			return fmt.Errorf(errValidatingProp, c$sName, c$name, err)";
				echo "		}";
				echo "	}";
			fi;
		fi;
	done;
	echo "	return nil";
	echo "}";
	echo;
	currSection=();
	requirements=();
	IFS=$(echo);
	longest=0;
}

OFS="$IFS";
(
	echo "package ics";
	echo;
	echo "// File automatically generated with ./genSections.sh";
	echo;
	echo "import (";
	echo "	\"fmt\"";
	echo "	\"io\"";
	echo "	\"strings\"";
	echo;
	echo "	\"vimagination.zapto.org/errors\"";
	echo "	\"vimagination.zapto.org/parser\"";
	echo ")";
	echo;

	{
		read sectionName;
		OFS="$IFS";
		IFS=$(echo);
		while read line; do
			if [ "${line:0:1}" != "	" ]; then
				printSection
				sectionName="$line"
				continue;
			fi;
			field="${line:1}";
			required=false;
			multiple=false;
			section=false;
			fc="${field:0:1}"
			if [ "$fc" = "?" ]; then
				field="${field:1}";
				if [ ! -z "$(echo "$field" | grep "!")" ]; then
					requirements[${#requirements}]="ONE $(echo "$field" | tr "!" " ")";
				elif [ ! -z "$(echo "$field" | grep "+")" ]; then
					requirements[${#requirements}]="AND $(echo "$field" | tr "+" " ")";
				elif [ ! -z "$(echo "$field" | grep ">")" ]; then
					requirements[${#requirements}]="ERGO $(echo "$field" | tr ">" " ")";
				elif [ ! -z "$(echo "$field" | grep "|")" ]; then
					requirements[${#requirements}]="OR $(echo "$field" | tr "|" " ")";
				fi;
				continue;
			elif [ "$fc" = "!" ]; then
				required=true;
				field="${field:1}";
			elif [ "$fc" = "+" ]; then
				multiple=true;
				required=true;
				field="${field:1}";
			elif [ "$fc" = "*" ]; then 
				multiple=true;
				field="${field:1}";
			fi;
			if [ "${field:0:6}" = "BEGIN:" ]; then
				section=true;
				field="${field:6}";
			else
				section=false;
			fi;
			addToSection $field $required $multiple $section;
		done;
	}< sections.gen;
	printSection
	cat <<HEREDOC
// decodeDummy reads unknown sections, discarding the data
func decodeDummy(t tokeniser, n string) error {
	for {
		p, err := t.GetPhrase()
		if err != nil {
			return err
		} else if p.Type == parser.PhraseDone {
			return io.ErrUnexpectedEOF
		}
		switch strings.ToUpper(p.Data[0].Data) {
		case "BEGIN":
			if err := decodeDummy(t, p.Data[len(p.Data)-1].Data); err != nil {
				return err
			}
		case "END":
			if strings.ToUpper(p.Data[len(p.Data)-1].Data) == n {
				return nil
			}
			return ErrInvalidEnd
		}
	}
}

// Errors
const (
	ErrInvalidEnd        errors.Error = "invalid end of section"
	ErrMissingRequired   errors.Error = "required property missing"
	ErrRequirementNotMet errors.Error = "requirement not met"
	errMultiple                       = "error decoding %s: multiple %s"
	errMissing                        = "error validating %s: missing %s"
HEREDOC
	while read line; do
		if [ "${line:0:1}" = "	" -o "$line" = "VFREEBUSY" ]; then
			continue;
		fi;
		type="$(getName $line)";
		echo -n "	c$type";
		IFS="$OFS";
		for i in $(seq $(( 33 - ${#type} ))); do
			echo -n " ";
		done;
		IFS="$(echo)";
		echo "= \"$type\"";
	done < sections.gen;
	echo ")";
) > sections.go
