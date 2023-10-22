// Code generated by gencel. DO NOT EDIT.

package funcs

import (
	"fmt"
	"log"

	"github.com/flanksource/gomplate/v3/kubernetes"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types/ref"
)

func convertMap(arg ref.Val) (map[string]any, error) {
	switch m := arg.Value().(type) {
	case  map[ref.Val]ref.Val:
		var out = make(map[string]any)
		for key, val := range m {
			out[key.Value().(string)] = val.Value()
		}
		return out, nil
	case  map[string]any:
		return m, nil
	default:
		return nil, 	fmt.Errorf("Not a map %T\n", arg.Value())
	}
}

func transferSlice[K any](arg ref.Val) []K {
	list, ok := arg.Value().([]ref.Val)
	if !ok {
		log.Printf("Not a list %T\n", arg.Value())
		return nil
	}

	var out = make([]K, len(list))
	for i, val := range list {
		out[i] = val.Value().(K)
	}

	return out
}

var CelEnvOption = []cel.EnvOption{
	fromAwsMap,
	arnToMap,
	kubernetes.Lists(),
	kubernetes.URLs(),
	kubernetes.Regex(),
	base64EncodeGen,
	base64DecodeGen,
	base64DecodeBytesGen,
	collSliceGen,
	collHasGen,
	collDictGen,
	collKeysGen,
	collValuesGen,
	collAppendGen,
	collPrependGen,
	collUniqGen,
	collReverseGen,
	collSortGen,
	collJQGen,
	collFlattenGen,
	collPickGen,
	collOmitGen,
	convSliceGen,
	convJoinGen,
	convHasGen,
	convURLGen,
	convToStringsGen,
	convDefaultGen,
	convDictGen,
	cryptoSHA1Gen,
	cryptoSHA224Gen,
	cryptoSHA256Gen,
	cryptoSHA384Gen,
	cryptoSHA512Gen,
	cryptoSHA512_224Gen,
	cryptoSHA512_256Gen,
	cryptoSHA1BytesGen,
	cryptoSHA224BytesGen,
	cryptoSHA256BytesGen,
	cryptoSHA384BytesGen,
	cryptoSHA512BytesGen,
	cryptoSHA512_224BytesGen,
	cryptoSHA512_256BytesGen,
	dataJSONGen,
	dataJSONArrayGen,
	dataYAMLGen,
	dataYAMLArrayGen,
	dataTOMLGen,
	dataCSVGen,
	dataCSVByRowGen,
	dataCSVByColumnGen,
	dataToCSVGen,
	dataToJSONGen,
	dataToJSONPrettyGen,
	dataToYAMLGen,
	dataToTOMLGen,
	urlEncodeGen,
	urlDecodeGen,
	filepathBaseGen,
	filepathCleanGen,
	filepathDirGen,
	filepathExtGen,
	filepathFromSlashGen,
	filepathIsAbsGen,
	filepathJoinGen,
	filepathMatchGen,
	filepathRelGen,
	filepathSplitGen,
	filepathToSlashGen,
	filepathVolumeNameGen,
	k8sIsHealthyGen,
	k8sGetStatusGen,
	k8sGetHealthGen,
	mathIsIntGen,
	mathIsFloatGen,
	mathcontainsFloatGen,
	mathIsNumGen,
	mathAbsGen,
	mathAddGen,
	mathMulGen,
	mathSubGen,
	mathDivGen,
	mathRemGen,
	mathPowGen,
	mathSeqGen,
	mathMaxGen,
	mathMinGen,
	mathCeilGen,
	mathFloorGen,
	mathRoundGen,
	pathBaseGen,
	pathCleanGen,
	pathDirGen,
	pathExtGen,
	pathIsAbsGen,
	pathJoinGen,
	pathMatchGen,
	pathSplitGen,
	randomASCIIGen,
	randomAlphaGen,
	randomAlphaNumGen,
	randomStringGen,
	randomItemGen,
	randomNumberGen,
	randomFloatGen,
	regexpFindGen,
	regexpFindAllGen,
	regexpMatchGen,
	regexpQuoteMetaGen,
	regexpReplaceGen,
	regexpReplaceLiteralGen,
	regexpSplitGen,
	stringsHumanDurationGen,
	stringsHumanSizeGen,
	stringsHumanDurationGen2,
	stringsHumanSizeGen2,
	stringsSemverGen,
	stringsSemverCompareGen,
	stringsAbbrevGen,
	stringsReplaceAllGen,
	stringsContainsGen,
	stringsRepeatGen,
	stringsSortGen,
	stringsSplitNGen,
	stringsTrimPrefixGen,
	stringsTrimSuffixGen,
	stringsTitleGen,
	stringsToUpperGen,
	stringsToLowerGen,
	stringsTrimSpaceGen,
	stringsTruncGen,
	stringsIndentGen,
	stringsSlugGen,
	stringsQuoteGen,
	stringsShellQuoteGen,
	stringsSquoteGen,
	stringsSnakeCaseGen,
	stringsCamelCaseGen,
	stringsKebabCaseGen,
	stringsWordWrapGen,
	stringsRuneCountGen,
	testAssertGen,
	testFailGen,
	testRequiredGen,
	testTernaryGen,
	testKindGen,
	testIsKindGen,
	timeZoneNameGen,
	timeZoneOffsetGen,
	timeParseGen,
	timeParseLocalGen,
	timeParseInLocationGen,
	timeNowGen,
	timeUnixGen,
	timeNanosecondGen,
	timeMicrosecondGen,
	timeMillisecondGen,
	timeSecondGen,
	timeMinuteGen,
	timeHourGen,
	timeParseDurationGen,
	timeSinceGen,
	timeUntilGen,
	uuidV1Gen,
	uuidV4Gen,
	uuidNilGen,
	uuidIsValidGen,
	uuidParseGen,
}
