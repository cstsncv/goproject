package monster

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func (mon *Monster) Store(filepath string) error {
	str, err := json.Marshal(&mon)
	if err != nil {
		fmt.Println(err)
		return err
	}
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	write := bufio.NewWriter(file)
	write.WriteString(string(str))
	write.Flush()
	return err
}

func (mon *Monster) Restore(filepath string) error {
	str, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(str, mon)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("读取文件中内容并反序列化后 = ", *mon)
	return err
}
