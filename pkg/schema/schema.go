// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "fmt"
import "reflect"
import "encoding/json"

// UnmarshalJSON implements json.Unmarshaler.
func (j *Omit) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain Omit
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Omit(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObfuscateType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ObfuscateType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ObfuscateType, v)
	}
	*j = ObfuscateType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SchemaJsonConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain SchemaJsonConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if len(plain.Obfuscate) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "obfuscate", 1)
	}
	*j = SchemaJsonConfig(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObfuscateReplacementType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ObfuscateReplacementType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ObfuscateReplacementType, v)
	}
	*j = ObfuscateReplacementType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OmitType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_OmitType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_OmitType, v)
	}
	*j = OmitType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObfuscateTarget) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ObfuscateTarget {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ObfuscateTarget, v)
	}
	*j = ObfuscateTarget(v)
	return nil
}

type Obfuscate struct {
	// The list of domains and their subdomains which should be obfuscated in the
	// output.
	Domains []string `json:"domains,omitempty"`

	// when replacementType=regex is used, the supplied  regex (Golang regexp
	// https://pkg.go.dev/regexp) will be used to detect the string that should be
	// replaced. The regex is line based, spanning multi-line regex statements is not
	// supported.
	Regex *string `json:"regex,omitempty"`

	// on any replacement type, this will override a given input string with another
	// output string. On duplicate keys it will use the last defined value as
	// replacement. The input values are matched in a case-sensitive fashion and only
	// as a full words, substrings must be matched using a regex. With type regex, the
	// keys will be considered as capture group values, if there is no mapping
	// available it will be replaced with a random string of the same size.
	Replacement ObfuscateReplacement `json:"replacement,omitempty"`

	// This defines how the detected string will be replaced. Type 'Consistent' will
	// guarantee the same input will always create the same output string and 'Random'
	// will just create a random replacement string of the same length as the input.
	// 'Static' relies solely on the replacement object to define an input/output
	// mapping.
	ReplacementType ObfuscateReplacementType `json:"replacementType"`

	// This determines if the obfuscation should be performed on the file name or on
	// the file contents. The file contents are obfuscated by default
	Target ObfuscateTarget `json:"target,omitempty"`

	// type defines the kind of detection you want to use. For example IP will find IP
	// addresses, whereas Keywords will find predefined keywords.
	Type ObfuscateType `json:"type"`
}

// on any replacement type, this will override a given input string with another
// output string. On duplicate keys it will use the last defined value as
// replacement. The input values are matched in a case-sensitive fashion and only
// as a full words, substrings must be matched using a regex. With type regex, the
// keys will be considered as capture group values, if there is no mapping
// available it will be replaced with a random string of the same size.
type ObfuscateReplacement map[string]string

type ObfuscateReplacementType string

const ObfuscateReplacementTypeConsistent ObfuscateReplacementType = "Consistent"
const ObfuscateReplacementTypeRandom ObfuscateReplacementType = "Random"
const ObfuscateReplacementTypeStatic ObfuscateReplacementType = "Static"

type ObfuscateTarget string

const ObfuscateTargetAll ObfuscateTarget = "All"
const ObfuscateTargetFileContents ObfuscateTarget = "FileContents"
const ObfuscateTargetFileName ObfuscateTarget = "FileName"

type ObfuscateType string

const ObfuscateTypeDomain ObfuscateType = "Domain"

// UnmarshalJSON implements json.Unmarshaler.
func (j *Obfuscate) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain Obfuscate
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["replacementType"]; !ok || v == nil {
		plain.ReplacementType = "Random"
	}
	if v, ok := raw["target"]; !ok || v == nil {
		plain.Target = "FileContents"
	}
	*j = Obfuscate(plain)
	return nil
}

const ObfuscateTypeIP ObfuscateType = "IP"
const ObfuscateTypeKeywords ObfuscateType = "Keywords"
const ObfuscateTypeMAC ObfuscateType = "MAC"
const ObfuscateTypeRegex ObfuscateType = "Regex"

type Omit struct {
	// KubernetesResource corresponds to the JSON schema field "kubernetesResource".
	KubernetesResource *OmitKubernetesResource `json:"kubernetesResource,omitempty"`

	// A file glob pattern on file paths relative to the must-gather root. The pattern
	// should be as described in https://pkg.go.dev/path/filepath#Match
	Pattern *string `json:"pattern,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type OmitType `json:"type"`
}

type OmitKubernetesResource struct {
	// This defines the apiVersion of the kubernetes resource. That can be used to
	// further refine specific versions of a resource that should be omitted.
	ApiVersion *string `json:"apiVersion,omitempty"`

	// This defines the kind of kubernetes resource that should be omitted. This can
	// be further specified with the apiVersion and namespaces.
	Kind *string `json:"kind,omitempty"`

	// This defines the namespaces which are supposed to be omitted. When used
	// together with kind and apiVersions, it becomes a filter. Standalone it will be
	// used as a filter for all resources in a given namespace.
	Namespaces []string `json:"namespaces,omitempty"`
}

type OmitType string

const OmitTypeFile OmitType = "File"
const OmitTypeKubernetes OmitType = "Kubernetes"

// This configuration defines the behaviour of the must-gather-clean CLI.
type SchemaJson struct {
	// The config schema defines the behaviour of the must-gather-clean CLI. There are
	// two sections, "omit" which defines the omission behaviour and "obfuscate" which
	// defines the obfuscation behaviour.
	Config SchemaJsonConfig `json:"config"`
}

// The config schema defines the behaviour of the must-gather-clean CLI. There are
// two sections, "omit" which defines the omission behaviour and "obfuscate" which
// defines the obfuscation behaviour.
type SchemaJsonConfig struct {
	// The obfuscation schema determines what is being detected and how it is being
	// replaced. We ship with several built-in replacements for common types such as
	// IP or MAC, Keywords and Regex. The replacements are done in order of the whole
	// list, so you can define chains of replacements that built on top of one another
	// - for example replacing a keyword and later matching its replacement with a
	// regex. The input to the given replacements are always a line of text (string).
	// Since file names can also have private content in them, they are also processed
	// as a line - exactly as they would with file content.
	Obfuscate []Obfuscate `json:"obfuscate,omitempty"`

	// The omission schema defines what kind of files shall not be included in the
	// final must-gather. This can be seen as a filter and can operate on file paths
	// or Kubernetes and OpenShift and other custom resources. Omissions are settled
	// first in the process of obfuscating a must-gather, so its content won't be
	// scanned and replaced.
	Omit []Omit `json:"omit,omitempty"`
}

var enumValues_ObfuscateReplacementType = []interface{}{
	"Consistent",
	"Random",
	"Static",
}
var enumValues_ObfuscateTarget = []interface{}{
	"FileName",
	"FileContents",
	"All",
}
var enumValues_ObfuscateType = []interface{}{
	"Domain",
	"IP",
	"Keywords",
	"MAC",
	"Regex",
}
var enumValues_OmitType = []interface{}{
	"Kubernetes",
	"File",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SchemaJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["config"]; !ok || v == nil {
		return fmt.Errorf("field config: required")
	}
	type Plain SchemaJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SchemaJson(plain)
	return nil
}
