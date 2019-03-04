package classfile
type ConstantInfo interface{
	readInfo(reader *ClassReader)
} 

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo{
	tag := reader.readerUint8()
	constanInfo := newConstantInfo(tag, cp)
	constanInfo.readInfo(reader)
	return constanInfo
}

func newConstantInfo(tag uint8, cp ConstantPool){
	switch tag {
		case CONSTANT_Integer: return &ConstantIntegerInfo{}
		case CONSTANT_Float: return &ConstantFloatInfo{}
		case CONSTANT_Long: return &ConstantLongInfo{}
		case CONSTANT_Double: return &ConstantDoubleInfo{}
		case CONSTANT_Utf8: return &ConstantUtf8Info{}
		case CONSTANT_String: return &ConstantStringInfo{cp: cp}
		case CONSTANT_Class: return &ConstantClassInfo{cp: cp}
		case CONSTANT_Fieldref:
			return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
		case CONSTANT_Methodref:
			return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
		case CONSTANT_InterfaceMethodref:
			return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
		case CONSTANT_NameAndType: return &ConstantNameAndTypeInfo{}
		case CONSTANT_MethodType: return &ConstantMethodTypeInfo{}
		case CONSTANT_MethodHandle: return &ConstantMethodHandleInfo{}
		case CONSTANT_InvokeDynamic: return &ConstantInvokeDynamicInfo{}
		default: panic("java.lang.ClassFormatError: constant pool tag!")
}
}


type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader)  ConstantPool{
	cpCount := int(reader.readerUint16())
	cp := make([]ConstantInfo, cpCount)
	for i:=1;i<cpCount;i++{ //i start from 1
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo{
	if cpInfo:=self[index];cpInfo!=nil{
		return cpInfo
	}
	panic("Invalid constant pool index")
}

func (self ConstantPool)getNameAndType(index uint16) (string, string){
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := slef.getUtf8(ntInfo.descriptorIndex)
	return name,_type
}

func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}