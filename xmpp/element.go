package xmpp

import (
	"encoding/xml"
	"fmt"
)

type Element struct {
	StartElement xml.StartElement
	EndElement   xml.EndElement
	CharData     string
	Comment      xml.Comment
	ProcInst     xml.ProcInst
	Directive    xml.Directive
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

func (elem *Element) String() string {
	result := fmt.Sprintf("(start:%v", elem.StartElement)
	charData := elem.CharData
	if charData != "" {
		result += fmt.Sprintf(", charData:%v", charData)
	}
	if len(elem.Children) > 0 {
		result += ", children:("
		for _, child := range elem.Children {
			if child == nil {
				continue
			}
			result += child.String()
		}
		result += ")"
	}
	result += ")"
	return result
}
