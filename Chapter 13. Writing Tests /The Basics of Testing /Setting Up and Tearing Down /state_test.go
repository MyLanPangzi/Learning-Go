package Setting_Up_and_Tearing_Down_

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

func TestMain(m *testing.M) {
	fmt.Println("set up stuff for tests here")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("clean up stuff after tests here")
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
	fmt.Println("TestFirst uses stuff set up in TestMain", testTime)
}

func TestSecond(t *testing.T) {
	fmt.Println("TestSecond uses stuff set up in TestMain", testTime)
}

func createFile(t *testing.T) (string, error) {
	f, err := os.Create("tempFIle")
	if err != nil {
		return "", err
	}
	t.Cleanup(func() {
		os.Remove(f.Name())
	})
	return f.Name(), nil
}

func TestFileProcessing(t *testing.T) {
	fileName, err := createFile(t)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fileName)
}
