package eft

import (
	"io/ioutil"
	"testing"
	"fmt"
	"os"
)

func TestFullRoundtrip(tt *testing.T) {
	eft_dir := TmpRandomName()
	hi0_txt := TmpRandomName()
	hi1_txt := TmpRandomName()

	defer func() {
		if len(eft_dir) > 8 {
			//os.RemoveAll(eft_dir)
			fmt.Println("XX - eft_dir =", eft_dir)
			os.Remove(hi0_txt)
			os.Remove(hi1_txt)
		}
	}()

	err := ioutil.WriteFile(hi0_txt, []byte("hai there"), 0600)
	if err != nil {
		panic(err)
	}

	key := [32]byte{}
	eft := EFT{Key: key, Dir: eft_dir} 

	info0, err := FastItemInfo(hi0_txt)
	if err != nil {
		panic(err)
	}

	err = eft.Put(info0, hi0_txt)
	if err != nil {
		panic(err)
	}

	dead, err := eft.Collect()
	if err != nil {
		panic(err)
	}
	os.Remove(dead)

	fmt.Println("XX - About to Get")

	info1, err := eft.Get(info0.Path, hi1_txt)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("XX - Done with Get")

	if info0 != info1 {
		fmt.Println("Item info mismatch")
		tt.Fail()
	}

	data, err := ioutil.ReadFile(hi1_txt)
	if err != nil {
		panic(err)
	}

	if string(data) != "hai there" {
		fmt.Println("Item data mismatch")
		tt.Fail()
	}

	/*
	fmt.Println("dir:", eft.Dir)
	fmt.Println(eft.ListDir("/"))
	fmt.Println(eft.ListDir("/tmp"))
	*/
}
