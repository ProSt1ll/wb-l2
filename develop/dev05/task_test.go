package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	desired string
	input   []string
	output  [][]string
	fl      flags
}{
	{
		desired: "самурай",
		input: []string{"Наемными убийцами были ниндзя, а самураи — профессиональными воинами. Их нанимали для охраны дворца, на службу в личную гвардию, для решения земельных конфликтов, усмирения восстаний и защиты внешних границ.",
			"Воинов (в том числе женщин-воительниц) обычно называли буси или цувамоно — «человек с оружием». Самураями их стали называть в период Хэйан (794–1185).",
			"Само слово «самурай» означало «человека на службе» — чиновника ниже шестого (реже пятого) придворного ранга. Поэтому служащие самураи были как среди воинов, так и среди низшей придворной знати.",
			"Именно самурай в образе жалкого чиновника, а не храброго воина, желающий только одного — до отвала наесться каши, описан в повести Акутагавы Рюноскэ «Бататовая каша».",
		},
		output: [][]string{
			{"Само слово «самурай» означало «человека на службе» — чиновника ниже шестого (реже пятого) придворного ранга. Поэтому служащие самураи были как среди воинов, так и среди низшей придворной знати."},
			{"Именно самурай в образе жалкого чиновника, а не храброго воина, желающий только одного — до отвала наесться каши, описан в повести Акутагавы Рюноскэ «Бататовая каша»."},
		},
		fl: flags{},
	},
	{
		desired: "Наемными убийцами были ниндзя, а самураи — профессиональными воинами. Их нанимали для охраны дворца, на службу в личную гвардию, для решения земельных конфликтов, усмирения восстаний и защиты внешних границ.",
		input: []string{"Наемными убийцами были ниндзя, а самураи — профессиональными воинами. Их нанимали для охраны дворца, на службу в личную гвардию, для решения земельных конфликтов, усмирения восстаний и защиты внешних границ.",
			"Воинов (в том числе женщин-воительниц) обычно называли буси или цувамоно — «человек с оружием». Самураями их стали называть в период Хэйан (794–1185).",
			"Само слово «самурай» означало «человека на службе» — чиновника ниже шестого (реже пятого) придворного ранга. Поэтому служащие самураи были как среди воинов, так и среди низшей придворной знати.",
			"Именно самурай в образе жалкого чиновника, а не храброго воина, желающий только одного — до отвала наесться каши, описан в повести Акутагавы Рюноскэ «Бататовая каша».",
		},
		output: [][]string{
			{"Наемными убийцами были ниндзя, а самураи — профессиональными воинами. Их нанимали для охраны дворца, на службу в личную гвардию, для решения земельных конфликтов, усмирения восстаний и защиты внешних границ."},
		},
		fl: flags{Fixed: true},
	},
	{
		desired: "службе",
		input: []string{"Наемными убийцами были ниндзя, а самураи — профессиональными воинами. Их нанимали для охраны дворца, на службу в личную гвардию, для решения земельных конфликтов, усмирения восстаний и защиты внешних границ.",
			"Воинов (в том числе женщин-воительниц) обычно называли буси или цувамоно — «человек с оружием». Самураями их стали называть в период Хэйан (794–1185).",
			"Само слово «самурай» означало «человека на службе» — чиновника ниже шестого (реже пятого) придворного ранга. Поэтому служащие самураи были как среди воинов, так и среди низшей придворной знати.",
			"Именно самурай в образе жалкого чиновника, а не храброго воина, желающий только одного — до отвала наесться каши, описан в повести Акутагавы Рюноскэ «Бататовая каша».",
		},
		output: [][]string{
			{"Само слово «самурай» означало «человека на службе» — чиновника ниже шестого (реже пятого) придворного ранга. Поэтому служащие самураи были как среди воинов, так и среди низшей придворной знати.",
				"Именно самурай в образе жалкого чиновника, а не храброго воина, желающий только одного — до отвала наесться каши, описан в повести Акутагавы Рюноскэ «Бататовая каша»."},
		},
		fl: flags{After: 1},
	},
	{
		desired: "службе",
		input: []string{"Наемными убийцами были ниндзя, а самураи — профессиональными воинами. Их нанимали для охраны дворца, на службу в личную гвардию, для решения земельных конфликтов, усмирения восстаний и защиты внешних границ.",
			"Воинов (в том числе женщин-воительниц) обычно называли буси или цувамоно — «человек с оружием». Самураями их стали называть в период Хэйан (794–1185).",
			"Само слово «самурай» означало «человека на службе» — чиновника ниже шестого (реже пятого) придворного ранга. Поэтому служащие самураи были как среди воинов, так и среди низшей придворной знати.",
			"Именно самурай в образе жалкого чиновника, а не храброго воина, желающий только одного — до отвала наесться каши, описан в повести Акутагавы Рюноскэ «Бататовая каша».",
		},
		output: [][]string{
			{"Воинов (в том числе женщин-воительниц) обычно называли буси или цувамоно — «человек с оружием». Самураями их стали называть в период Хэйан (794–1185).",
				"Само слово «самурай» означало «человека на службе» — чиновника ниже шестого (реже пятого) придворного ранга. Поэтому служащие самураи были как среди воинов, так и среди низшей придворной знати."},
		},
		fl: flags{Before: 1},
	},
	{
		desired: "самурай",
		input: []string{"Наемными убийцами были ниндзя, а самураи — профессиональными воинами. Их нанимали для охраны дворца, на службу в личную гвардию, для решения земельных конфликтов, усмирения восстаний и защиты внешних границ.",
			"Воинов (в том числе женщин-воительниц) обычно называли буси или цувамоно — «человек с оружием». Самураями их стали называть в период Хэйан (794–1185).",
			"Само слово «самурай» означало «человека на службе» — чиновника ниже шестого (реже пятого) придворного ранга. Поэтому служащие самураи были как среди воинов, так и среди низшей придворной знати.",
			"Именно самурай в образе жалкого чиновника, а не храброго воина, желающий только одного — до отвала наесться каши, описан в повести Акутагавы Рюноскэ «Бататовая каша».",
		},
		output: [][]string{
			{"Наемными убийцами были ниндзя, а самураи — профессиональными воинами. Их нанимали для охраны дворца, на службу в личную гвардию, для решения земельных конфликтов, усмирения восстаний и защиты внешних границ."},
			{"Воинов (в том числе женщин-воительниц) обычно называли буси или цувамоно — «человек с оружием». Самураями их стали называть в период Хэйан (794–1185)."},
		},
		fl: flags{
			Invert: true,
		},
	},
}

func TestSort(t *testing.T) {
	for _, tt := range tests {
		assert.Equal(t, Grep(tt.desired, tt.input, tt.fl).strs, tt.output)
	}
}
