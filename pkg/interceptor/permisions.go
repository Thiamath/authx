/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package interceptor

type Permission struct {
	Must    [] string `json:"must,omitempty"`
	Should  [] string `json:"should,omitempty"`
	MustNot [] string `json:"must_not,omitempty"`
}

func (p *Permission) Check(primitives [] string) bool{
	for _,must := range p.Must {
		check:=false
		for _,pri := range primitives {
			if !check && pri == must {
				check = true
			}
			if !check {
				return false
			}
		}
	}
	for _,should := range p.Should {
		counter:=0
		for _,pri := range primitives {
			if pri == should {
				counter += 1
			}
			if counter == 0 {
				return false
			}
		}
	}
	for _,mustNo := range p.MustNot {
		for _,pri := range primitives {
			if pri == mustNo {
				return false
			}
		}
	}
	return true
}
