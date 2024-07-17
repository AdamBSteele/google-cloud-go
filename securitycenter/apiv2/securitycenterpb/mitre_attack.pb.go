// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/cloud/securitycenter/v2/mitre_attack.proto

package securitycenterpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MITRE ATT&CK tactics that can be referenced by SCC findings.
// See: https://attack.mitre.org/tactics/enterprise/
type MitreAttack_Tactic int32

const (
	// Unspecified value.
	MitreAttack_TACTIC_UNSPECIFIED MitreAttack_Tactic = 0
	// TA0043
	MitreAttack_RECONNAISSANCE MitreAttack_Tactic = 1
	// TA0042
	MitreAttack_RESOURCE_DEVELOPMENT MitreAttack_Tactic = 2
	// TA0001
	MitreAttack_INITIAL_ACCESS MitreAttack_Tactic = 5
	// TA0002
	MitreAttack_EXECUTION MitreAttack_Tactic = 3
	// TA0003
	MitreAttack_PERSISTENCE MitreAttack_Tactic = 6
	// TA0004
	MitreAttack_PRIVILEGE_ESCALATION MitreAttack_Tactic = 8
	// TA0005
	MitreAttack_DEFENSE_EVASION MitreAttack_Tactic = 7
	// TA0006
	MitreAttack_CREDENTIAL_ACCESS MitreAttack_Tactic = 9
	// TA0007
	MitreAttack_DISCOVERY MitreAttack_Tactic = 10
	// TA0008
	MitreAttack_LATERAL_MOVEMENT MitreAttack_Tactic = 11
	// TA0009
	MitreAttack_COLLECTION MitreAttack_Tactic = 12
	// TA0011
	MitreAttack_COMMAND_AND_CONTROL MitreAttack_Tactic = 4
	// TA0010
	MitreAttack_EXFILTRATION MitreAttack_Tactic = 13
	// TA0040
	MitreAttack_IMPACT MitreAttack_Tactic = 14
)

// Enum value maps for MitreAttack_Tactic.
var (
	MitreAttack_Tactic_name = map[int32]string{
		0:  "TACTIC_UNSPECIFIED",
		1:  "RECONNAISSANCE",
		2:  "RESOURCE_DEVELOPMENT",
		5:  "INITIAL_ACCESS",
		3:  "EXECUTION",
		6:  "PERSISTENCE",
		8:  "PRIVILEGE_ESCALATION",
		7:  "DEFENSE_EVASION",
		9:  "CREDENTIAL_ACCESS",
		10: "DISCOVERY",
		11: "LATERAL_MOVEMENT",
		12: "COLLECTION",
		4:  "COMMAND_AND_CONTROL",
		13: "EXFILTRATION",
		14: "IMPACT",
	}
	MitreAttack_Tactic_value = map[string]int32{
		"TACTIC_UNSPECIFIED":   0,
		"RECONNAISSANCE":       1,
		"RESOURCE_DEVELOPMENT": 2,
		"INITIAL_ACCESS":       5,
		"EXECUTION":            3,
		"PERSISTENCE":          6,
		"PRIVILEGE_ESCALATION": 8,
		"DEFENSE_EVASION":      7,
		"CREDENTIAL_ACCESS":    9,
		"DISCOVERY":            10,
		"LATERAL_MOVEMENT":     11,
		"COLLECTION":           12,
		"COMMAND_AND_CONTROL":  4,
		"EXFILTRATION":         13,
		"IMPACT":               14,
	}
)

func (x MitreAttack_Tactic) Enum() *MitreAttack_Tactic {
	p := new(MitreAttack_Tactic)
	*p = x
	return p
}

func (x MitreAttack_Tactic) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MitreAttack_Tactic) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_securitycenter_v2_mitre_attack_proto_enumTypes[0].Descriptor()
}

func (MitreAttack_Tactic) Type() protoreflect.EnumType {
	return &file_google_cloud_securitycenter_v2_mitre_attack_proto_enumTypes[0]
}

func (x MitreAttack_Tactic) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MitreAttack_Tactic.Descriptor instead.
func (MitreAttack_Tactic) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_mitre_attack_proto_rawDescGZIP(), []int{0, 0}
}

// MITRE ATT&CK techniques that can be referenced by SCC findings.
// See: https://attack.mitre.org/techniques/enterprise/
// Next ID: 63
type MitreAttack_Technique int32

const (
	// Unspecified value.
	MitreAttack_TECHNIQUE_UNSPECIFIED MitreAttack_Technique = 0
	// T1036
	MitreAttack_MASQUERADING MitreAttack_Technique = 49
	// T1036.005
	MitreAttack_MATCH_LEGITIMATE_NAME_OR_LOCATION MitreAttack_Technique = 50
	// T1037
	MitreAttack_BOOT_OR_LOGON_INITIALIZATION_SCRIPTS MitreAttack_Technique = 37
	// T1037.005
	MitreAttack_STARTUP_ITEMS MitreAttack_Technique = 38
	// T1046
	MitreAttack_NETWORK_SERVICE_DISCOVERY MitreAttack_Technique = 32
	// T1057
	MitreAttack_PROCESS_DISCOVERY MitreAttack_Technique = 56
	// T1059
	MitreAttack_COMMAND_AND_SCRIPTING_INTERPRETER MitreAttack_Technique = 6
	// T1059.004
	MitreAttack_UNIX_SHELL MitreAttack_Technique = 7
	// T1059.006
	MitreAttack_PYTHON MitreAttack_Technique = 59
	// T1069
	MitreAttack_PERMISSION_GROUPS_DISCOVERY MitreAttack_Technique = 18
	// T1069.003
	MitreAttack_CLOUD_GROUPS MitreAttack_Technique = 19
	// T1071
	MitreAttack_APPLICATION_LAYER_PROTOCOL MitreAttack_Technique = 45
	// T1071.004
	MitreAttack_DNS MitreAttack_Technique = 46
	// T1072
	MitreAttack_SOFTWARE_DEPLOYMENT_TOOLS MitreAttack_Technique = 47
	// T1078
	MitreAttack_VALID_ACCOUNTS MitreAttack_Technique = 14
	// T1078.001
	MitreAttack_DEFAULT_ACCOUNTS MitreAttack_Technique = 35
	// T1078.003
	MitreAttack_LOCAL_ACCOUNTS MitreAttack_Technique = 15
	// T1078.004
	MitreAttack_CLOUD_ACCOUNTS MitreAttack_Technique = 16
	// T1090
	MitreAttack_PROXY MitreAttack_Technique = 9
	// T1090.002
	MitreAttack_EXTERNAL_PROXY MitreAttack_Technique = 10
	// T1090.003
	MitreAttack_MULTI_HOP_PROXY MitreAttack_Technique = 11
	// T1098
	MitreAttack_ACCOUNT_MANIPULATION MitreAttack_Technique = 22
	// T1098.001
	MitreAttack_ADDITIONAL_CLOUD_CREDENTIALS MitreAttack_Technique = 40
	// T1098.004
	MitreAttack_SSH_AUTHORIZED_KEYS MitreAttack_Technique = 23
	// T1098.006
	MitreAttack_ADDITIONAL_CONTAINER_CLUSTER_ROLES MitreAttack_Technique = 58
	// T1105
	MitreAttack_INGRESS_TOOL_TRANSFER MitreAttack_Technique = 3
	// T1106
	MitreAttack_NATIVE_API MitreAttack_Technique = 4
	// T1110
	MitreAttack_BRUTE_FORCE MitreAttack_Technique = 44
	// T1129
	MitreAttack_SHARED_MODULES MitreAttack_Technique = 5
	// T1134
	MitreAttack_ACCESS_TOKEN_MANIPULATION MitreAttack_Technique = 33
	// T1134.001
	MitreAttack_TOKEN_IMPERSONATION_OR_THEFT MitreAttack_Technique = 39
	// T1190
	MitreAttack_EXPLOIT_PUBLIC_FACING_APPLICATION MitreAttack_Technique = 27
	// T1484
	MitreAttack_DOMAIN_POLICY_MODIFICATION MitreAttack_Technique = 30
	// T1485
	MitreAttack_DATA_DESTRUCTION MitreAttack_Technique = 29
	// T1489
	MitreAttack_SERVICE_STOP MitreAttack_Technique = 52
	// T1490
	MitreAttack_INHIBIT_SYSTEM_RECOVERY MitreAttack_Technique = 36
	// T1496
	MitreAttack_RESOURCE_HIJACKING MitreAttack_Technique = 8
	// T1498
	MitreAttack_NETWORK_DENIAL_OF_SERVICE MitreAttack_Technique = 17
	// T1526
	MitreAttack_CLOUD_SERVICE_DISCOVERY MitreAttack_Technique = 48
	// T1528
	MitreAttack_STEAL_APPLICATION_ACCESS_TOKEN MitreAttack_Technique = 42
	// T1531
	MitreAttack_ACCOUNT_ACCESS_REMOVAL MitreAttack_Technique = 51
	// T1539
	MitreAttack_STEAL_WEB_SESSION_COOKIE MitreAttack_Technique = 25
	// T1543
	MitreAttack_CREATE_OR_MODIFY_SYSTEM_PROCESS MitreAttack_Technique = 24
	// T1548
	MitreAttack_ABUSE_ELEVATION_CONTROL_MECHANISM MitreAttack_Technique = 34
	// T1552
	MitreAttack_UNSECURED_CREDENTIALS MitreAttack_Technique = 13
	// T1556
	MitreAttack_MODIFY_AUTHENTICATION_PROCESS MitreAttack_Technique = 28
	// T1562
	MitreAttack_IMPAIR_DEFENSES MitreAttack_Technique = 31
	// T1562.001
	MitreAttack_DISABLE_OR_MODIFY_TOOLS MitreAttack_Technique = 55
	// T1567
	MitreAttack_EXFILTRATION_OVER_WEB_SERVICE MitreAttack_Technique = 20
	// T1567.002
	MitreAttack_EXFILTRATION_TO_CLOUD_STORAGE MitreAttack_Technique = 21
	// T1568
	MitreAttack_DYNAMIC_RESOLUTION MitreAttack_Technique = 12
	// T1570
	MitreAttack_LATERAL_TOOL_TRANSFER MitreAttack_Technique = 41
	// T1578
	MitreAttack_MODIFY_CLOUD_COMPUTE_INFRASTRUCTURE MitreAttack_Technique = 26
	// T1578.001
	MitreAttack_CREATE_SNAPSHOT MitreAttack_Technique = 54
	// T1580
	MitreAttack_CLOUD_INFRASTRUCTURE_DISCOVERY MitreAttack_Technique = 53
	// T1588
	MitreAttack_OBTAIN_CAPABILITIES MitreAttack_Technique = 43
	// T1595
	MitreAttack_ACTIVE_SCANNING MitreAttack_Technique = 1
	// T1595.001
	MitreAttack_SCANNING_IP_BLOCKS MitreAttack_Technique = 2
	// T1613
	MitreAttack_CONTAINER_ADMINISTRATION_COMMAND MitreAttack_Technique = 60
	// T1611
	MitreAttack_ESCAPE_TO_HOST MitreAttack_Technique = 61
	// T1613
	MitreAttack_CONTAINER_AND_RESOURCE_DISCOVERY MitreAttack_Technique = 57
	// T1649
	MitreAttack_STEAL_OR_FORGE_AUTHENTICATION_CERTIFICATES MitreAttack_Technique = 62
)

// Enum value maps for MitreAttack_Technique.
var (
	MitreAttack_Technique_name = map[int32]string{
		0:  "TECHNIQUE_UNSPECIFIED",
		49: "MASQUERADING",
		50: "MATCH_LEGITIMATE_NAME_OR_LOCATION",
		37: "BOOT_OR_LOGON_INITIALIZATION_SCRIPTS",
		38: "STARTUP_ITEMS",
		32: "NETWORK_SERVICE_DISCOVERY",
		56: "PROCESS_DISCOVERY",
		6:  "COMMAND_AND_SCRIPTING_INTERPRETER",
		7:  "UNIX_SHELL",
		59: "PYTHON",
		18: "PERMISSION_GROUPS_DISCOVERY",
		19: "CLOUD_GROUPS",
		45: "APPLICATION_LAYER_PROTOCOL",
		46: "DNS",
		47: "SOFTWARE_DEPLOYMENT_TOOLS",
		14: "VALID_ACCOUNTS",
		35: "DEFAULT_ACCOUNTS",
		15: "LOCAL_ACCOUNTS",
		16: "CLOUD_ACCOUNTS",
		9:  "PROXY",
		10: "EXTERNAL_PROXY",
		11: "MULTI_HOP_PROXY",
		22: "ACCOUNT_MANIPULATION",
		40: "ADDITIONAL_CLOUD_CREDENTIALS",
		23: "SSH_AUTHORIZED_KEYS",
		58: "ADDITIONAL_CONTAINER_CLUSTER_ROLES",
		3:  "INGRESS_TOOL_TRANSFER",
		4:  "NATIVE_API",
		44: "BRUTE_FORCE",
		5:  "SHARED_MODULES",
		33: "ACCESS_TOKEN_MANIPULATION",
		39: "TOKEN_IMPERSONATION_OR_THEFT",
		27: "EXPLOIT_PUBLIC_FACING_APPLICATION",
		30: "DOMAIN_POLICY_MODIFICATION",
		29: "DATA_DESTRUCTION",
		52: "SERVICE_STOP",
		36: "INHIBIT_SYSTEM_RECOVERY",
		8:  "RESOURCE_HIJACKING",
		17: "NETWORK_DENIAL_OF_SERVICE",
		48: "CLOUD_SERVICE_DISCOVERY",
		42: "STEAL_APPLICATION_ACCESS_TOKEN",
		51: "ACCOUNT_ACCESS_REMOVAL",
		25: "STEAL_WEB_SESSION_COOKIE",
		24: "CREATE_OR_MODIFY_SYSTEM_PROCESS",
		34: "ABUSE_ELEVATION_CONTROL_MECHANISM",
		13: "UNSECURED_CREDENTIALS",
		28: "MODIFY_AUTHENTICATION_PROCESS",
		31: "IMPAIR_DEFENSES",
		55: "DISABLE_OR_MODIFY_TOOLS",
		20: "EXFILTRATION_OVER_WEB_SERVICE",
		21: "EXFILTRATION_TO_CLOUD_STORAGE",
		12: "DYNAMIC_RESOLUTION",
		41: "LATERAL_TOOL_TRANSFER",
		26: "MODIFY_CLOUD_COMPUTE_INFRASTRUCTURE",
		54: "CREATE_SNAPSHOT",
		53: "CLOUD_INFRASTRUCTURE_DISCOVERY",
		43: "OBTAIN_CAPABILITIES",
		1:  "ACTIVE_SCANNING",
		2:  "SCANNING_IP_BLOCKS",
		60: "CONTAINER_ADMINISTRATION_COMMAND",
		61: "ESCAPE_TO_HOST",
		57: "CONTAINER_AND_RESOURCE_DISCOVERY",
		62: "STEAL_OR_FORGE_AUTHENTICATION_CERTIFICATES",
	}
	MitreAttack_Technique_value = map[string]int32{
		"TECHNIQUE_UNSPECIFIED":                      0,
		"MASQUERADING":                               49,
		"MATCH_LEGITIMATE_NAME_OR_LOCATION":          50,
		"BOOT_OR_LOGON_INITIALIZATION_SCRIPTS":       37,
		"STARTUP_ITEMS":                              38,
		"NETWORK_SERVICE_DISCOVERY":                  32,
		"PROCESS_DISCOVERY":                          56,
		"COMMAND_AND_SCRIPTING_INTERPRETER":          6,
		"UNIX_SHELL":                                 7,
		"PYTHON":                                     59,
		"PERMISSION_GROUPS_DISCOVERY":                18,
		"CLOUD_GROUPS":                               19,
		"APPLICATION_LAYER_PROTOCOL":                 45,
		"DNS":                                        46,
		"SOFTWARE_DEPLOYMENT_TOOLS":                  47,
		"VALID_ACCOUNTS":                             14,
		"DEFAULT_ACCOUNTS":                           35,
		"LOCAL_ACCOUNTS":                             15,
		"CLOUD_ACCOUNTS":                             16,
		"PROXY":                                      9,
		"EXTERNAL_PROXY":                             10,
		"MULTI_HOP_PROXY":                            11,
		"ACCOUNT_MANIPULATION":                       22,
		"ADDITIONAL_CLOUD_CREDENTIALS":               40,
		"SSH_AUTHORIZED_KEYS":                        23,
		"ADDITIONAL_CONTAINER_CLUSTER_ROLES":         58,
		"INGRESS_TOOL_TRANSFER":                      3,
		"NATIVE_API":                                 4,
		"BRUTE_FORCE":                                44,
		"SHARED_MODULES":                             5,
		"ACCESS_TOKEN_MANIPULATION":                  33,
		"TOKEN_IMPERSONATION_OR_THEFT":               39,
		"EXPLOIT_PUBLIC_FACING_APPLICATION":          27,
		"DOMAIN_POLICY_MODIFICATION":                 30,
		"DATA_DESTRUCTION":                           29,
		"SERVICE_STOP":                               52,
		"INHIBIT_SYSTEM_RECOVERY":                    36,
		"RESOURCE_HIJACKING":                         8,
		"NETWORK_DENIAL_OF_SERVICE":                  17,
		"CLOUD_SERVICE_DISCOVERY":                    48,
		"STEAL_APPLICATION_ACCESS_TOKEN":             42,
		"ACCOUNT_ACCESS_REMOVAL":                     51,
		"STEAL_WEB_SESSION_COOKIE":                   25,
		"CREATE_OR_MODIFY_SYSTEM_PROCESS":            24,
		"ABUSE_ELEVATION_CONTROL_MECHANISM":          34,
		"UNSECURED_CREDENTIALS":                      13,
		"MODIFY_AUTHENTICATION_PROCESS":              28,
		"IMPAIR_DEFENSES":                            31,
		"DISABLE_OR_MODIFY_TOOLS":                    55,
		"EXFILTRATION_OVER_WEB_SERVICE":              20,
		"EXFILTRATION_TO_CLOUD_STORAGE":              21,
		"DYNAMIC_RESOLUTION":                         12,
		"LATERAL_TOOL_TRANSFER":                      41,
		"MODIFY_CLOUD_COMPUTE_INFRASTRUCTURE":        26,
		"CREATE_SNAPSHOT":                            54,
		"CLOUD_INFRASTRUCTURE_DISCOVERY":             53,
		"OBTAIN_CAPABILITIES":                        43,
		"ACTIVE_SCANNING":                            1,
		"SCANNING_IP_BLOCKS":                         2,
		"CONTAINER_ADMINISTRATION_COMMAND":           60,
		"ESCAPE_TO_HOST":                             61,
		"CONTAINER_AND_RESOURCE_DISCOVERY":           57,
		"STEAL_OR_FORGE_AUTHENTICATION_CERTIFICATES": 62,
	}
)

func (x MitreAttack_Technique) Enum() *MitreAttack_Technique {
	p := new(MitreAttack_Technique)
	*p = x
	return p
}

func (x MitreAttack_Technique) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MitreAttack_Technique) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_securitycenter_v2_mitre_attack_proto_enumTypes[1].Descriptor()
}

func (MitreAttack_Technique) Type() protoreflect.EnumType {
	return &file_google_cloud_securitycenter_v2_mitre_attack_proto_enumTypes[1]
}

func (x MitreAttack_Technique) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MitreAttack_Technique.Descriptor instead.
func (MitreAttack_Technique) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_mitre_attack_proto_rawDescGZIP(), []int{0, 1}
}

// MITRE ATT&CK tactics and techniques related to this finding.
// See: https://attack.mitre.org
type MitreAttack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The MITRE ATT&CK tactic most closely represented by this finding, if any.
	PrimaryTactic MitreAttack_Tactic `protobuf:"varint,1,opt,name=primary_tactic,json=primaryTactic,proto3,enum=google.cloud.securitycenter.v2.MitreAttack_Tactic" json:"primary_tactic,omitempty"`
	// The MITRE ATT&CK technique most closely represented by this finding, if
	// any. primary_techniques is a repeated field because there are multiple
	// levels of MITRE ATT&CK techniques.  If the technique most closely
	// represented by this finding is a sub-technique (e.g. `SCANNING_IP_BLOCKS`),
	// both the sub-technique and its parent technique(s) will be listed (e.g.
	// `SCANNING_IP_BLOCKS`, `ACTIVE_SCANNING`).
	PrimaryTechniques []MitreAttack_Technique `protobuf:"varint,2,rep,packed,name=primary_techniques,json=primaryTechniques,proto3,enum=google.cloud.securitycenter.v2.MitreAttack_Technique" json:"primary_techniques,omitempty"`
	// Additional MITRE ATT&CK tactics related to this finding, if any.
	AdditionalTactics []MitreAttack_Tactic `protobuf:"varint,3,rep,packed,name=additional_tactics,json=additionalTactics,proto3,enum=google.cloud.securitycenter.v2.MitreAttack_Tactic" json:"additional_tactics,omitempty"`
	// Additional MITRE ATT&CK techniques related to this finding, if any, along
	// with any of their respective parent techniques.
	AdditionalTechniques []MitreAttack_Technique `protobuf:"varint,4,rep,packed,name=additional_techniques,json=additionalTechniques,proto3,enum=google.cloud.securitycenter.v2.MitreAttack_Technique" json:"additional_techniques,omitempty"`
	// The MITRE ATT&CK version referenced by the above fields. E.g. "8".
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *MitreAttack) Reset() {
	*x = MitreAttack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v2_mitre_attack_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MitreAttack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MitreAttack) ProtoMessage() {}

func (x *MitreAttack) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v2_mitre_attack_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MitreAttack.ProtoReflect.Descriptor instead.
func (*MitreAttack) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_mitre_attack_proto_rawDescGZIP(), []int{0}
}

func (x *MitreAttack) GetPrimaryTactic() MitreAttack_Tactic {
	if x != nil {
		return x.PrimaryTactic
	}
	return MitreAttack_TACTIC_UNSPECIFIED
}

func (x *MitreAttack) GetPrimaryTechniques() []MitreAttack_Technique {
	if x != nil {
		return x.PrimaryTechniques
	}
	return nil
}

func (x *MitreAttack) GetAdditionalTactics() []MitreAttack_Tactic {
	if x != nil {
		return x.AdditionalTactics
	}
	return nil
}

func (x *MitreAttack) GetAdditionalTechniques() []MitreAttack_Technique {
	if x != nil {
		return x.AdditionalTechniques
	}
	return nil
}

func (x *MitreAttack) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_google_cloud_securitycenter_v2_mitre_attack_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v2_mitre_attack_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x32,
	0x2f, 0x6d, 0x69, 0x74, 0x72, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x32, 0x22, 0xc0, 0x13, 0x0a, 0x0b, 0x4d, 0x69, 0x74, 0x72, 0x65, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x12, 0x59, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x74,
	0x61, 0x63, 0x74, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x69, 0x74,
	0x72, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x54, 0x61, 0x63, 0x74, 0x69, 0x63, 0x52,
	0x0d, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x54, 0x61, 0x63, 0x74, 0x69, 0x63, 0x12, 0x64,
	0x0a, 0x12, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x69, 0x74, 0x72,
	0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x52, 0x11, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x73, 0x12, 0x61, 0x0a, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x74, 0x61, 0x63, 0x74, 0x69, 0x63, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x32, 0x2e, 0x4d, 0x69, 0x74, 0x72, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x54, 0x61,
	0x63, 0x74, 0x69, 0x63, 0x52, 0x11, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x54, 0x61, 0x63, 0x74, 0x69, 0x63, 0x73, 0x12, 0x6a, 0x0a, 0x15, 0x61, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x69, 0x74, 0x72, 0x65, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x2e, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x52, 0x14, 0x61,
	0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xb4, 0x02,
	0x0a, 0x06, 0x54, 0x61, 0x63, 0x74, 0x69, 0x63, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x41, 0x43, 0x54,
	0x49, 0x43, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x43, 0x4f, 0x4e, 0x4e, 0x41, 0x49, 0x53, 0x53, 0x41, 0x4e,
	0x43, 0x45, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x12,
	0x0a, 0x0e, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x45, 0x52, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x43, 0x45,
	0x10, 0x06, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x49, 0x56, 0x49, 0x4c, 0x45, 0x47, 0x45, 0x5f,
	0x45, 0x53, 0x43, 0x41, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f,
	0x44, 0x45, 0x46, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x45, 0x56, 0x41, 0x53, 0x49, 0x4f, 0x4e, 0x10,
	0x07, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x52, 0x45, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c, 0x5f,
	0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x49, 0x53, 0x43,
	0x4f, 0x56, 0x45, 0x52, 0x59, 0x10, 0x0a, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x41, 0x54, 0x45, 0x52,
	0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x56, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x0b, 0x12, 0x0e, 0x0a,
	0x0a, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0c, 0x12, 0x17, 0x0a,
	0x13, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x43, 0x4f, 0x4e,
	0x54, 0x52, 0x4f, 0x4c, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x58, 0x46, 0x49, 0x4c, 0x54,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0d, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x4d, 0x50, 0x41,
	0x43, 0x54, 0x10, 0x0e, 0x22, 0xcf, 0x0d, 0x0a, 0x09, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x45, 0x43, 0x48, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x4d, 0x41, 0x53, 0x51, 0x55, 0x45, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x31, 0x12,
	0x25, 0x0a, 0x21, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x4c, 0x45, 0x47, 0x49, 0x54, 0x49, 0x4d,
	0x41, 0x54, 0x45, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x52, 0x5f, 0x4c, 0x4f, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x32, 0x12, 0x28, 0x0a, 0x24, 0x42, 0x4f, 0x4f, 0x54, 0x5f, 0x4f,
	0x52, 0x5f, 0x4c, 0x4f, 0x47, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49,
	0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x53, 0x10, 0x25,
	0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x52, 0x54, 0x55, 0x50, 0x5f, 0x49, 0x54, 0x45, 0x4d,
	0x53, 0x10, 0x26, 0x12, 0x1d, 0x0a, 0x19, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59,
	0x10, 0x20, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x44, 0x49,
	0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x10, 0x38, 0x12, 0x25, 0x0a, 0x21, 0x43, 0x4f, 0x4d,
	0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x50, 0x52, 0x45, 0x54, 0x45, 0x52, 0x10, 0x06,
	0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x49, 0x58, 0x5f, 0x53, 0x48, 0x45, 0x4c, 0x4c, 0x10, 0x07,
	0x12, 0x0a, 0x0a, 0x06, 0x50, 0x59, 0x54, 0x48, 0x4f, 0x4e, 0x10, 0x3b, 0x12, 0x1f, 0x0a, 0x1b,
	0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50,
	0x53, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x10, 0x12, 0x12, 0x10, 0x0a,
	0x0c, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x53, 0x10, 0x13, 0x12,
	0x1e, 0x0a, 0x1a, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c,
	0x41, 0x59, 0x45, 0x52, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x10, 0x2d, 0x12,
	0x07, 0x0a, 0x03, 0x44, 0x4e, 0x53, 0x10, 0x2e, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x4f, 0x46, 0x54,
	0x57, 0x41, 0x52, 0x45, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x4f, 0x4f, 0x4c, 0x53, 0x10, 0x2f, 0x12, 0x12, 0x0a, 0x0e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x53, 0x10, 0x0e, 0x12, 0x14, 0x0a, 0x10, 0x44,
	0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x53, 0x10,
	0x23, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x53, 0x10, 0x0f, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x41,
	0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x53, 0x10, 0x10, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x52, 0x4f,
	0x58, 0x59, 0x10, 0x09, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c,
	0x5f, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x55, 0x4c, 0x54,
	0x49, 0x5f, 0x48, 0x4f, 0x50, 0x5f, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x10, 0x0b, 0x12, 0x18, 0x0a,
	0x14, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4d, 0x41, 0x4e, 0x49, 0x50, 0x55, 0x4c,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x16, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x44, 0x44, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x43, 0x52, 0x45, 0x44,
	0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c, 0x53, 0x10, 0x28, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x53, 0x48,
	0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x5f, 0x4b, 0x45, 0x59, 0x53,
	0x10, 0x17, 0x12, 0x26, 0x0a, 0x22, 0x41, 0x44, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c,
	0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45, 0x52, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54,
	0x45, 0x52, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x53, 0x10, 0x3a, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e,
	0x47, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x4f, 0x4f, 0x4c, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53,
	0x46, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f,
	0x41, 0x50, 0x49, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x52, 0x55, 0x54, 0x45, 0x5f, 0x46,
	0x4f, 0x52, 0x43, 0x45, 0x10, 0x2c, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x48, 0x41, 0x52, 0x45, 0x44,
	0x5f, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x53, 0x10, 0x05, 0x12, 0x1d, 0x0a, 0x19, 0x41, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4d, 0x41, 0x4e, 0x49, 0x50,
	0x55, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x21, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x4f, 0x4b,
	0x45, 0x4e, 0x5f, 0x49, 0x4d, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4f, 0x52, 0x5f, 0x54, 0x48, 0x45, 0x46, 0x54, 0x10, 0x27, 0x12, 0x25, 0x0a, 0x21, 0x45,
	0x58, 0x50, 0x4c, 0x4f, 0x49, 0x54, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x5f, 0x46, 0x41,
	0x43, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x1b, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x50, 0x4f, 0x4c,
	0x49, 0x43, 0x59, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x1e, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x52,
	0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x1d, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x34, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e,
	0x48, 0x49, 0x42, 0x49, 0x54, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x52, 0x45, 0x43,
	0x4f, 0x56, 0x45, 0x52, 0x59, 0x10, 0x24, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x48, 0x49, 0x4a, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x08, 0x12,
	0x1d, 0x0a, 0x19, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x41,
	0x4c, 0x5f, 0x4f, 0x46, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x11, 0x12, 0x1b,
	0x0a, 0x17, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x10, 0x30, 0x12, 0x22, 0x0a, 0x1e, 0x53,
	0x54, 0x45, 0x41, 0x4c, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x2a, 0x12,
	0x1a, 0x0a, 0x16, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x41, 0x4c, 0x10, 0x33, 0x12, 0x1c, 0x0a, 0x18, 0x53,
	0x54, 0x45, 0x41, 0x4c, 0x5f, 0x57, 0x45, 0x42, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x43, 0x4f, 0x4f, 0x4b, 0x49, 0x45, 0x10, 0x19, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x5f, 0x4f, 0x52, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x53, 0x59,
	0x53, 0x54, 0x45, 0x4d, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x10, 0x18, 0x12, 0x25,
	0x0a, 0x21, 0x41, 0x42, 0x55, 0x53, 0x45, 0x5f, 0x45, 0x4c, 0x45, 0x56, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x4d, 0x45, 0x43, 0x48, 0x41, 0x4e,
	0x49, 0x53, 0x4d, 0x10, 0x22, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x4e, 0x53, 0x45, 0x43, 0x55, 0x52,
	0x45, 0x44, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c, 0x53, 0x10, 0x0d,
	0x12, 0x21, 0x0a, 0x1d, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x45,
	0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x1c, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4d, 0x50, 0x41, 0x49, 0x52, 0x5f, 0x44, 0x45,
	0x46, 0x45, 0x4e, 0x53, 0x45, 0x53, 0x10, 0x1f, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x49, 0x53, 0x41,
	0x42, 0x4c, 0x45, 0x5f, 0x4f, 0x52, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x54, 0x4f,
	0x4f, 0x4c, 0x53, 0x10, 0x37, 0x12, 0x21, 0x0a, 0x1d, 0x45, 0x58, 0x46, 0x49, 0x4c, 0x54, 0x52,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x57, 0x45, 0x42, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x14, 0x12, 0x21, 0x0a, 0x1d, 0x45, 0x58, 0x46, 0x49,
	0x4c, 0x54, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x4c, 0x4f, 0x55,
	0x44, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x10, 0x15, 0x12, 0x16, 0x0a, 0x12, 0x44,
	0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x55, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x0c, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x41, 0x54, 0x45, 0x52, 0x41, 0x4c, 0x5f, 0x54,
	0x4f, 0x4f, 0x4c, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x10, 0x29, 0x12, 0x27,
	0x0a, 0x23, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x55, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x46, 0x52, 0x41, 0x53, 0x54, 0x52, 0x55,
	0x43, 0x54, 0x55, 0x52, 0x45, 0x10, 0x1a, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x5f, 0x53, 0x4e, 0x41, 0x50, 0x53, 0x48, 0x4f, 0x54, 0x10, 0x36, 0x12, 0x22, 0x0a, 0x1e,
	0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x49, 0x4e, 0x46, 0x52, 0x41, 0x53, 0x54, 0x52, 0x55, 0x43,
	0x54, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x10, 0x35,
	0x12, 0x17, 0x0a, 0x13, 0x4f, 0x42, 0x54, 0x41, 0x49, 0x4e, 0x5f, 0x43, 0x41, 0x50, 0x41, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x49, 0x45, 0x53, 0x10, 0x2b, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x45, 0x5f, 0x53, 0x43, 0x41, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x16,
	0x0a, 0x12, 0x53, 0x43, 0x41, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x49, 0x50, 0x5f, 0x42, 0x4c,
	0x4f, 0x43, 0x4b, 0x53, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49,
	0x4e, 0x45, 0x52, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x10, 0x3c, 0x12, 0x12, 0x0a, 0x0e,
	0x45, 0x53, 0x43, 0x41, 0x50, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x10, 0x3d,
	0x12, 0x24, 0x0a, 0x20, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45, 0x52, 0x5f, 0x41, 0x4e,
	0x44, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f,
	0x56, 0x45, 0x52, 0x59, 0x10, 0x39, 0x12, 0x2e, 0x0a, 0x2a, 0x53, 0x54, 0x45, 0x41, 0x4c, 0x5f,
	0x4f, 0x52, 0x5f, 0x46, 0x4f, 0x52, 0x47, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54,
	0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43,
	0x41, 0x54, 0x45, 0x53, 0x10, 0x3e, 0x42, 0xea, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x42, 0x10, 0x4d,
	0x69, 0x74, 0x72, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x32, 0xca, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x32, 0xea,
	0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a,
	0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v2_mitre_attack_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v2_mitre_attack_proto_rawDescData = file_google_cloud_securitycenter_v2_mitre_attack_proto_rawDesc
)

func file_google_cloud_securitycenter_v2_mitre_attack_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v2_mitre_attack_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v2_mitre_attack_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v2_mitre_attack_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v2_mitre_attack_proto_rawDescData
}

var file_google_cloud_securitycenter_v2_mitre_attack_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_securitycenter_v2_mitre_attack_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_securitycenter_v2_mitre_attack_proto_goTypes = []any{
	(MitreAttack_Tactic)(0),    // 0: google.cloud.securitycenter.v2.MitreAttack.Tactic
	(MitreAttack_Technique)(0), // 1: google.cloud.securitycenter.v2.MitreAttack.Technique
	(*MitreAttack)(nil),        // 2: google.cloud.securitycenter.v2.MitreAttack
}
var file_google_cloud_securitycenter_v2_mitre_attack_proto_depIdxs = []int32{
	0, // 0: google.cloud.securitycenter.v2.MitreAttack.primary_tactic:type_name -> google.cloud.securitycenter.v2.MitreAttack.Tactic
	1, // 1: google.cloud.securitycenter.v2.MitreAttack.primary_techniques:type_name -> google.cloud.securitycenter.v2.MitreAttack.Technique
	0, // 2: google.cloud.securitycenter.v2.MitreAttack.additional_tactics:type_name -> google.cloud.securitycenter.v2.MitreAttack.Tactic
	1, // 3: google.cloud.securitycenter.v2.MitreAttack.additional_techniques:type_name -> google.cloud.securitycenter.v2.MitreAttack.Technique
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v2_mitre_attack_proto_init() }
func file_google_cloud_securitycenter_v2_mitre_attack_proto_init() {
	if File_google_cloud_securitycenter_v2_mitre_attack_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v2_mitre_attack_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MitreAttack); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_securitycenter_v2_mitre_attack_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v2_mitre_attack_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v2_mitre_attack_proto_depIdxs,
		EnumInfos:         file_google_cloud_securitycenter_v2_mitre_attack_proto_enumTypes,
		MessageInfos:      file_google_cloud_securitycenter_v2_mitre_attack_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v2_mitre_attack_proto = out.File
	file_google_cloud_securitycenter_v2_mitre_attack_proto_rawDesc = nil
	file_google_cloud_securitycenter_v2_mitre_attack_proto_goTypes = nil
	file_google_cloud_securitycenter_v2_mitre_attack_proto_depIdxs = nil
}