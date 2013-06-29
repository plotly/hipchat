package xmpp

import (
	"encoding/xml"
	"log"
)

type Element struct {
	StartElement *xml.StartElement
	EndElement   *xml.EndElement
	CharData     *xml.CharData
	Comment      *xml.Comment
	ProcInst     *xml.ProcInst
	Directive    *xml.Directive
	Children     []*Element
	Parent       *Element
}

func NewElement(parent *Element) *Element {
	elem := new(Element)
	elem.Children = make([]*Element, 1)
	if parent != nil {
		parent.Children = append(parent.Children, elem)
		elem.Parent = parent
	}
	return elem
}

func (elem *Element) Print() {
	log.Println("(")
	log.Println(*elem.StartElement)
	if elem.CharData != nil {
		log.Println(string(*elem.CharData))
	}
	for _, child := range elem.Children {
		if child == nil {
			continue
		}
		child.Print()
	}
	log.Println(")")
}
