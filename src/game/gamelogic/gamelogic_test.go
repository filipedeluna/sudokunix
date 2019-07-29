package gamelogic

import (
	"testing"
)

func TestSeedRotation(t *testing.T) {
	if rotateSeed("123456789", 1, 3) != "741852963" {
		t.Error()
	}

	if rotateSeed("1234", 1, 2) != "3142" {
		t.Error()
	}

	if rotateSeed("1234", 2, 2) != "4321" {
		t.Error()
	}
}

func TestGetFileByDifficulty(t *testing.T) {
	file, err := getFileByDifficulty(1)
	fileName := "../../seeds/easy.txt"

	if err != nil {
		t.Error(err)
	}

	if file.Name() !=  fileName {
		t.Error("Expected" + fileName + ", got '" + file.Name() + "'")
	}

	file, err = getFileByDifficulty(2)
	fileName = "../../seeds/med.txt"

	if err != nil {
		t.Error(err)
	}

	if file.Name() != fileName {
		t.Error("Expected" + fileName + ", got '" + file.Name() + "'")
	}

	file, err = getFileByDifficulty(3)
	fileName = "../../seeds/hard.txt"

	if err != nil  {
		t.Error(err)
	}

	if file.Name() != fileName {
		t.Error("Expected" + fileName + ", got '" + file.Name() + "'")
	}

}