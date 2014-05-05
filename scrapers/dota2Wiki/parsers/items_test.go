package parsers

import (
	"testing"
)

var tango_data string
var clarity_data string
var arcane_data string

func init() {
	tango_data = getDota2Wiki("Tango")
	clarity_data = getDota2Wiki("Clarity")
	arcane_data = getDota2Wiki("Arcane_Boots")
}

func Test_GetItemName_Err(t *testing.T) {
	_, err := ParseItemName("")
	if err.Error() != "no item name found" {
		t.Errorf("%+v", err)
	}
}

func Test_GetItemName_Tango(t *testing.T) {
	ans, err := ParseItemName(tango_data)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if ans != "Tango" {
		t.Errorf("%+v", ans)
	}
}

func Test_GetItemName_Clarity(t *testing.T) {
	ans, err := ParseItemName(clarity_data)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if ans != "Clarity" {
		t.Errorf("%+v", ans)
	}
}

func Test_GetItemCost_Err(t *testing.T) {
	_, err := ParseItemCost("")
	if err.Error() != "no item cost found" {
		t.Errorf("%+v", err)
	}
}

func Test_GetItemCost_Tango(t *testing.T) {
	ans, err := ParseItemCost(tango_data)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if ans != 125 {
		t.Errorf("%+v", ans)
	}
}

func Test_GetItemCost_Clarity(t *testing.T) {
	ans, err := ParseItemCost(clarity_data)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if ans != 50 {
		t.Errorf("%+v", ans)
	}
}

func Test_GetItemDisassemble_Err(t *testing.T) {
	_, err := ParseItemDisassemble("")
	if err.Error() != "no item disassemble found" {
		t.Errorf("%+v", err)
	}
}

func Test_GetItemDisassemble_Tango(t *testing.T) {
	ans, err := ParseItemDisassemble(tango_data)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if ans {
		t.Errorf("%+v", ans)
	}
}

func Test_GetItemDisassemble_Arcane(t *testing.T) {
	ans, err := ParseItemDisassemble(arcane_data)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if !ans {
		t.Errorf("%+v", ans)
	}
}
