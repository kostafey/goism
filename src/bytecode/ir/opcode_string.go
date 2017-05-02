// Code generated by "stringer -type=Opcode"; DO NOT EDIT.

package ir

import "fmt"

const _Opcode_name = "OpReturnOpCallOpConstRefOpStackRefOpStackSetOpDropOpVarRefOpVarSetOpSetCarOpSetCdrOpArrayRefOpArraySetOpSubstrOpConcatOpStringEqOpStringLtOpToLowerOpToUpperOpIsConsOpIsStringOpIsNumOpIsIntOpNumAddOpNumSubOpNumMulOpNumDivOpNumEqOpNumLtOpNumLteOpNumGtOpNumGteOpNumNegOpNumMaxOpNumMinOpRemOpEqOpEqualOpNotOpMakeListOpMakeConsOpJmpOpJmpNilOpJmpNotNilOpJmpNilElsePopOpJmpNotNilElsePopOpRelJmpOpRelJmpNilOpRelJmpNotNilOpRelJmpNilElsePopOpRelJmpNotNilElsePopOpCatch"

var _Opcode_index = [...]uint16{0, 8, 14, 24, 34, 44, 50, 58, 66, 74, 82, 92, 102, 110, 118, 128, 138, 147, 156, 164, 174, 181, 188, 196, 204, 212, 220, 227, 234, 242, 249, 257, 265, 273, 281, 286, 290, 297, 302, 312, 322, 327, 335, 346, 361, 379, 387, 398, 412, 430, 451, 458}

func (i Opcode) String() string {
	if i < 0 || i >= Opcode(len(_Opcode_index)-1) {
		return fmt.Sprintf("Opcode(%d)", i)
	}
	return _Opcode_name[_Opcode_index[i]:_Opcode_index[i+1]]
}
