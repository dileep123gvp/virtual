// Code generated by "stringer -type Opcode enum.go"; DO NOT EDIT.

package virtual

import "fmt"

const _Opcode_name = "NopAPAddF32AddF64AddC64AddC128AddI32AddI64AddPtrAddPtrsAddSPAnd16And32And64And8ArgumentArgument16Argument32Argument64Argument8ArgumentsArgumentsFPBPBitfieldI8BitfieldI16BitfieldI32BitfieldI64BitfieldU8BitfieldU16BitfieldU32BitfieldU64BoolC128BoolF32BoolF64BoolI16BoolI32BoolI64BoolI8CallCallFPConvC64C128ConvF32C128ConvF32C64ConvF32F64ConvF32I32ConvF32I64ConvF32U32ConvF64C128ConvF64F32ConvF64I32ConvF64I64ConvF64I8ConvF64U16ConvF64U32ConvF64U64ConvI16I32ConvI16I64ConvI16U32ConvI32C128ConvI32C64ConvI32F32ConvI32F64ConvI32I16ConvI32I64ConvI32I8ConvI64ConvI64F64ConvI64I16ConvI64I32ConvI64I8ConvI64U16ConvI8I16ConvI8I32ConvI8I64ConvI8F64ConvI8U32ConvU16I32ConvU16I64ConvU16U32ConvU16U64ConvU32F32ConvU32F64ConvU32I16ConvU32I64ConvU32U8ConvU8I16ConvU8I32ConvU8U32ConvU8U64CopyCpl32Cpl64Cpl8DSDSC128DSI16DSI32DSI64DSI8DSNDivC128DivC64DivF32DivF64DivI32DivI64DivU32DivU64Dup32Dup64Dup8EqF32EqF64EqI32EqI64EqI8ExtFFIReturnFPField16Field64Field8FuncGeqF32GeqF64GeqI32GeqI64GeqI8GeqU32GeqU64GtF32GtF64GtI32GtI64GtU32GtU64IndexIndexI16IndexU16IndexI32IndexI64IndexI8IndexU32IndexU64IndexU8JmpJmpPJnzJzLabelLeqF32LeqF64LeqI32LeqI64LeqI8LeqU32LeqU64LoadLoad16Load32Load64Load8LshI16LshI32LshI64LshI8LtF32LtF64LtI32LtI64LtU32LtU64MulC128MulC64MulF32MulF64MulI32MulI64NegF32NegF64NegI16NegI32NegI64NegI8NegIndexI32NegIndexI64NegIndexU16NegIndexU32NegIndexU64NeqC128NeqC64NeqF32NeqF64NeqI32NeqI64NeqI8NotOr32Or64PanicPostIncF64PostIncI16PostIncI32PostIncI64PostIncI8PostIncPtrPostIncU32BitsPostIncU64BitsPreIncI16PreIncI32PreIncI64PreIncI8PreIncPtrPreIncU32BitsPreIncU64BitsPtrDiffPush16Push32Push64Push8PushC128RemI32RemI64RemU32RemU64ReturnRshI16RshI32RshI64RshI8RshU16RshU32RshU64RshU8StoreStore16Store32Store64StoreC128Store8StoreBits16StoreBits32StoreBits64StoreBits8StrNCopySubF32SubF64SubI32SubI64SubPtrsTextVariableVariable16Variable32Variable64Variable8Xor32Xor64Zero8Zero16Zero32Zero64__signbit__signbitfabortabsaccessacosallocaasinatanbswap64builtincallocceilcimagfclose_clrsbclrsblclrsbllclzclzlclzllcopysigncoscoshcrealfctzctzlctzlldlclosedlerrordlopendlsymerrno_locationexitexpfabsfchmodfchownfclosefcntlfflushffsffslffsllfgetcfgetsfloorfopen64fprintfframeAddressfreadfreefstat64fsyncftruncate64fwritegetcwdgetenvgeteuidgetpidgettimeofdayisinfisinffisinflisprintlocaltimeloglog10longjmplseek64lstat64mallocmemcmpmemcpymemmovemempcpymemsetmkdirmmap64munmapopen64parityparitylparityllpopcountpopcountlpopcountllpowprintfpthread_createpthread_equalpthread_joinpthread_mutex_destroypthread_mutex_initpthread_mutex_lockpthread_mutex_trylockpthread_mutex_unlockpthread_mutexattr_destroypthread_mutexattr_initpthread_mutexattr_settypepthread_selfqsortreadreadlinkreallocreturnAddressrmdirroundsetjmpsinsinhsleepsprintfsqrtstat64register_stdfilesstrcatstrchrstrcmpstrcpystrlenstrncmpstrncpystrrchrsysconftantanhtimetolowerunlinkutimesvfprintfvprintfwrite"

var _Opcode_index = [...]uint16{0, 3, 5, 11, 17, 23, 30, 36, 42, 48, 55, 60, 65, 70, 75, 79, 87, 97, 107, 117, 126, 135, 146, 148, 158, 169, 180, 191, 201, 212, 223, 234, 242, 249, 256, 263, 270, 277, 283, 287, 293, 304, 315, 325, 335, 345, 355, 365, 376, 386, 396, 406, 415, 425, 435, 445, 455, 465, 475, 486, 496, 506, 516, 526, 536, 545, 552, 562, 572, 582, 591, 601, 610, 619, 628, 637, 646, 656, 666, 676, 686, 696, 706, 716, 726, 735, 744, 753, 762, 771, 775, 780, 785, 789, 791, 797, 802, 807, 812, 816, 819, 826, 832, 838, 844, 850, 856, 862, 868, 873, 878, 882, 887, 892, 897, 902, 906, 909, 918, 920, 927, 934, 940, 944, 950, 956, 962, 968, 973, 979, 985, 990, 995, 1000, 1005, 1010, 1015, 1020, 1028, 1036, 1044, 1052, 1059, 1067, 1075, 1082, 1085, 1089, 1092, 1094, 1099, 1105, 1111, 1117, 1123, 1128, 1134, 1140, 1144, 1150, 1156, 1162, 1167, 1173, 1179, 1185, 1190, 1195, 1200, 1205, 1210, 1215, 1220, 1227, 1233, 1239, 1245, 1251, 1257, 1263, 1269, 1275, 1281, 1287, 1292, 1303, 1314, 1325, 1336, 1347, 1354, 1360, 1366, 1372, 1378, 1384, 1389, 1392, 1396, 1400, 1405, 1415, 1425, 1435, 1445, 1454, 1464, 1478, 1492, 1501, 1510, 1519, 1527, 1536, 1549, 1562, 1569, 1575, 1581, 1587, 1592, 1600, 1606, 1612, 1618, 1624, 1630, 1636, 1642, 1648, 1653, 1659, 1665, 1671, 1676, 1681, 1688, 1695, 1702, 1711, 1717, 1728, 1739, 1750, 1760, 1768, 1774, 1780, 1786, 1792, 1799, 1803, 1811, 1821, 1831, 1841, 1850, 1855, 1860, 1865, 1871, 1877, 1883, 1892, 1902, 1907, 1910, 1916, 1920, 1926, 1930, 1934, 1941, 1948, 1954, 1958, 1964, 1970, 1975, 1981, 1988, 1991, 1995, 2000, 2008, 2011, 2015, 2021, 2024, 2028, 2033, 2040, 2047, 2053, 2058, 2072, 2076, 2079, 2083, 2089, 2095, 2101, 2106, 2112, 2115, 2119, 2124, 2129, 2134, 2139, 2146, 2153, 2165, 2170, 2174, 2181, 2186, 2197, 2203, 2209, 2215, 2222, 2228, 2240, 2245, 2251, 2257, 2264, 2273, 2276, 2281, 2288, 2295, 2302, 2308, 2314, 2320, 2327, 2334, 2340, 2345, 2351, 2357, 2363, 2369, 2376, 2384, 2392, 2401, 2411, 2414, 2420, 2434, 2447, 2459, 2480, 2498, 2516, 2537, 2557, 2582, 2604, 2629, 2641, 2646, 2650, 2658, 2665, 2678, 2683, 2688, 2694, 2697, 2701, 2706, 2713, 2717, 2723, 2740, 2746, 2752, 2758, 2764, 2770, 2777, 2784, 2791, 2798, 2801, 2805, 2809, 2816, 2822, 2828, 2836, 2843, 2848}

func (i Opcode) String() string {
	if i < 0 || i >= Opcode(len(_Opcode_index)-1) {
		return fmt.Sprintf("Opcode(%d)", i)
	}
	return _Opcode_name[_Opcode_index[i]:_Opcode_index[i+1]]
}
