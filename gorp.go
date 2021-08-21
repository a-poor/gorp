package main

import (
	"fmt"
)

// const definitionFile = "./def.json"
// const outputDir = "./output/"

type ProtoRule int

const (
	ProtoOptional ProtoRule = iota
	ProtoRepeated
	ProtoRequired
)

func (r ProtoRule) String() string {
	switch r {
	case ProtoOptional:
		return "optional"
	case ProtoRequired:
		return "required"
	case ProtoRepeated:
		return "repeated"
	default:
		return ""
	}
}

type MessageField struct {
	Name    string    // The name of the field
	Type    string    // The name of the field type
	Rule    ProtoRule // Is the field optional, repeated, or required?
	Comment string    // Optional comment
}

func (f MessageField) String(id int) string {
	var c string
	if len(f.Comment) > 0 {
		c = fmt.Sprintf(" // %s", f.Comment)
	}
	return fmt.Sprintf("%s %s %s = %d;%s", f.Rule.String(), f.Type, f.Name, id, c)
}

type Message struct {
	Name   string
	Fields []MessageField
}

func (msg Message) String() string {
	s := fmt.Sprintf("message %s {\n", msg.Name)
	for i, f := range msg.Fields {
		s += fmt.Sprintf("  %s\n", f.String(i))
	}
	s += "}"
	return s
}

func main() {
	// log.Println("Starting...")
	// log.Printf("Definition file: %s\n", definitionFile)
	// log.Printf("Output directory: %s\n", outputDir)
	// fmt.Println()

	// fmt.Println()
	// log.Println("Done.")

	fields := make([]MessageField, 0)
	fields = append(fields, MessageField{"id", "string", ProtoRequired, "The user's id in the database"})
	fields = append(fields, MessageField{Name: "name", Type: "string"})
	fields = append(fields, MessageField{"fave_numbers", "int32", ProtoRepeated, "The user's favorite numbers"})

	m := Message{
		Name:   "user",
		Fields: fields,
	}

	fmt.Println(m.String())
}
