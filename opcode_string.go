// Code generated by "stringer -type Opcode enum.go"; DO NOT EDIT

package virtual

import "fmt"

const _Opcode_name = "NopAPAddI32AddPtrAddSPArgument32Argument64ArgumentsBPCallDup32EqI32FuncIndexIndexI32Int32JmpJnzJzLabelLeqI32Load32LtI32MulI32PanicPostIncI32ReturnStore32SubI32TextVariable32abortexitprintf"

var _Opcode_index = [...]uint8{0, 3, 5, 11, 17, 22, 32, 42, 51, 53, 57, 62, 67, 71, 76, 84, 89, 92, 95, 97, 102, 108, 114, 119, 125, 130, 140, 146, 153, 159, 163, 173, 178, 182, 188}

func (i Opcode) String() string {
	if i < 0 || i >= Opcode(len(_Opcode_index)-1) {
		return fmt.Sprintf("Opcode(%d)", i)
	}
	return _Opcode_name[_Opcode_index[i]:_Opcode_index[i+1]]
}
