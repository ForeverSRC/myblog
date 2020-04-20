package functions

import (
	"strings"
)

func EncodingLabels(label string) string{
	//传入参数："go gogin";最终结果：$go#$gogin#
	label="$"+strings.Join(
		strings.Split(label," "),"#$")+"#"
	return label
}

func DecodingLabels(label string) string {
	////传入参数：$go#$gogin#;最终结果："go gogin"
	label=strings.Replace(strings.Replace(
		label,"#"," ",-1),"$","",-1)
	return label
}

func DecodingLabelsIterable(label string) []string {
	////传入参数：$go#$gogin#;最终结果："go gogin"
	label=strings.Replace(strings.Replace(
		label,"#"," ",-1),"$","",-1)
	return strings.Split(label," ")
}

