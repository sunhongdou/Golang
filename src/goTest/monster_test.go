package demo10

import "testing"

func TestStore(t *testing.T){
	//1.创建对象
	monster := &Monster{
		Name : "xiaoming",
		Age : 13,
		Skll: "computer",
	}
	res := monster.Store()
	if !res{
		t.Fatal("monster.Store() failed,expect is: %v,actual is: %v",true,res)
	}
	t.Logf("monster.Store test success!")
}

func TestRestore(t *testing.T){
	var monster Monster
	res := monster.Restore()
	if !res{
		t.Fatal("monster.ReStore() failed,expect is: %v,actual is: %v",true,res)
	}
	if monster.Name != "xiaoming"{
		t.Fatal("monster.Restore() failed,name expect is : %V,actual is : %V","xiaoming",res)
	}
	t.Logf("monster.Store test success!")
}
