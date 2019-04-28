package classfile

import (
	"bufio"
	"fmt"
)

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type ClassFile struct {
	// magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags uint16
	thisClass uint16
	superClass uint16
	interfaces []uint16
	fields []*MemberInfo
	methods []*MemberInfo
	attributes []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {

}

func (cf *ClassFile) read(reader *ClassReader) {

}

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader)  {

}

func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {

}

func (cf *ClassFile) MinorVersion() uint16 {

}

func (cf *ClassFile) MajorVersion() uint16 {

}

func (cf *ClassFile) ConstantPool() ConstantPool {

}

func (cf *ClassFile) AccessFlags() uint16 {

}

func (cf *ClassFile) Fields() []*MemberInfo {

}

func (cf *ClassFile) Methods() []*MemberInfo {

}

func (cf *ClassFile) ClassName() string {

}

func (cf *ClassFile) SuperClassName() string {

}

func (cf *ClassFile) InterfaceNames() []string {

}
