package main

import (
	"fmt"
	"strings"
)

// part A was simple, part B is causing issues as
// it looks like another LCM issue, but the items are not that regular

type pulse_module struct {
	typ rune // 'b','%' or '&'
	outputs []string // list of modules to send to
	ff_state bool // flipflop state, false if off
	conj_input map[string]bool // conjunction inputs
}

type pulse struct {
	from string
	high bool // true if high
	to string
}

func parse_modules(lines []string) map[string]pulse_module {
	result:=make(map[string]pulse_module)
	for _,line:=range lines {
		kv:=strings.Split(line," -> ")
		pm:=pulse_module{}
		var name string
		if kv[0][0] == '%' || kv[0][0]=='&' {
			pm.typ=rune(kv[0][0])
			name = kv[0][1:]
		} else {
			pm.typ=rune(' ')
			name=kv[0]
		}
		pm.outputs=strings.Split(kv[1],", ")
		pm.ff_state=false
		result[name]=pm
	}
	// fill conjunction inputs
	for n,m:=range result {
		if m.typ=='&' {
			m.conj_input=make(map[string]bool)
			// find all items and add
			for n2,m2:=range result {
				for _,op:=range m2.outputs {
					if op == n {
						m.conj_input[n2]=false
					}
				}
			}
			// copy back
			result[n]=m
		}
	}
	return result
}

func trigger_pulses(modules map[string]pulse_module) (low,high int, high_to_lx map[string]bool)  {
	low,high=0,0
	high_to_lx=make(map[string]bool)
	todo:=[]pulse{{"button",false,"broadcaster"}}
	for len(todo)>0 {
		var p pulse
		p,todo=todo[0],todo[1:]
		//pulsestr:="low"
		if p.high {
		//	pulsestr="high"
			high++
		}else{
			low++
		}
		//fmt.Printf("%s -%s-> %s\n", p.from,pulsestr,p.to)
		if p.high==true && p.to=="lx"{
			//fmt.Printf("high pulse sent to %s from %s in cycle %d\n",p.to,p.from,count)
			high_to_lx[p.from]=true
		}

		mod,_:=modules[p.to]
		if mod.typ==' '{
			for _,op:=range mod.outputs{
				todo=append(todo, pulse{p.to,p.high,op})
			}
		} else if mod.typ == '%' {
			if p.high==false {
				if mod.ff_state {
					mod.ff_state=false
					for _,op:=range mod.outputs{
						todo=append(todo, pulse{p.to,false,op})
					}
				}else{
					mod.ff_state=true
					for _,op:=range mod.outputs{
						todo=append(todo, pulse{p.to,true,op})
					}
				}
				modules[p.to]=mod // copy back
			}
		}  else if mod.typ=='&'{
			mod.conj_input[p.from]=p.high
			modules[p.to]=mod // copy back
			all_high:=true
			for _,v:=range mod.conj_input{
				if v==false{
					all_high=false
					break
				}
			}
			for _,op:=range mod.outputs{
				todo=append(todo, pulse{p.to,!all_high,op})
			}
		}
	}
	return
}

func day20a(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	modules:=parse_modules(lines)
	fmt.Printf("%v\n",modules)

	low,high:=0,0
	for i:=0;i<1000;i++{
		l,h,_:=trigger_pulses(modules)
		low+=l
		high+=h
	}
	fmt.Printf("result %d\n",low*high)
}

func day20b(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%v\n", lines)
	modules := parse_modules(lines)
	high_pulses:=make(map[string][]int)

	for count := 1; count <100*1000; count++ {
		_,_,highs:=trigger_pulses(modules)
		// rx module is not actually listed:
		// but its few from conjunction lx which is fed from: cl,rp,lb,nj
		// so look at high pulses sent to lx
		// PS. got multiple items as I believed that it was not a simple sequence
		// looking at https://mliezun.github.io/2023/12/25/favourite-advent-of-code-2023.html proved me wrong
		for id:=range highs{
			pulses,found:=high_pulses[id]
			if found==false {
				pulses=[]int{count}
			} else {
				pulses=append(pulses,count)
			}
			high_pulses[id]=pulses
		}

	}
	fmt.Printf("high %v\n",high_pulses)

	// get the product of all the values:
	// tried LCM is hangs
	result:=int64(1) // lcm(1,x)=x
	for _,sec:= range high_pulses{
		val:=int64(sec[0])
		result*=val

	}
	fmt.Printf("final result %d\n",result)
}