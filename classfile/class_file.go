package classfile
import "fmt"

type ClassFile struct{
	// magic		uint32
	minorVersion	uint16
	majorVersion	uint16
	constantPool	ConstantPool
	accessFlags		uint16
	thisClass		uint16
	superClass		uint16
	interfaces		uint16
	fields			[]*MemberInfo
	methods			[]*MemberInfo
	attributes		[]AttributeInfo
}

func Parse(classData []byte)  (cf *ClassFile, err error){
	defer func(){
		if r:=recover();r!=nil{
			var ok bool
			err, ok = r.(error)
			if !ok{
				err = fmt.Errorf("%v",r)
			}
		}
	}()
	cr := &ClassReader(classData)
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader){
	self.readAndCheckMagic(reader)

	self.readAndCheckVersion(reader)

	self.constantPool = readConstantPool(reader)

	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)

	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader){
	magic := reader.readUint32()
	if magic != 0xCAFEBABE{
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (self*ClassFile)readAndCheckVersion()  {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
		case 45:
			return
		case 46,47,48,49,50,51,52:
			if self.minorVersion == 0{
				return
			}	
	}
	panic("java.lang.UnsupportedClassVersion")
}

//getter
func (self* ClassFile) MinorVersion() uint16 {
	return this.minorVersion
}
//getter
func (self* ClassFile) MajorVersion() uint16 {
	return this.majorVersion
}
//getter
func (self* ClassFile) ConstantPool() ConstantPool {
	return this.constantPool
}
//getter
func (self* ClassFile) AccessFlags() uint16 {
	return this.accessFlags
}
//getter
func (self* ClassFile) Fields() []*MemberInfo {
	return this.fields
}
//getter
func (self* ClassFile) Methods() []*MemberInfo {
	return this.methods
}
//getter
func (self* ClassFile) ClassName() string {
	return this.constantPool.getClassName(self.thisClass)
}
//getter
func (self* ClassFile) SuperClassName() string {
	if self.superClass >0 {
		return this.constantPool.getClassName(self.thisClass)
	}
	return ""
}
//getter
func (self* ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces{
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}