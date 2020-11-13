package main

import (
	"fmt"
	"math/big"
)

const searchDepthBits = 8

func bin2str(s []byte) string {
	r := ""
	for i := 0; i<len(s); i++ {
		r=r+string(48+s[len(s)-1-i])
	}
	return r
}

func binArr2BInt(arr []byte) *big.Int {
	r := big.NewInt(0)
	for i:=len(arr)-1;i>=0 ;i-- {
		r.Mul(r,big.NewInt(2))
		r.Add(r,big.NewInt((int64)(arr[i])))
	}
	return r
}

func BInt2binArr(number *big.Int,l int) []byte {
	r := []byte{}
	current := new(big.Int).Set(number)
	
	for i:=0;i<l;i++{
		_,m :=current.DivMod(current,big.NewInt(2),big.NewInt(2))
		r = append(r,(uint8)(m.Int64()))
	}

	return r
}

func bitMask(length int) *big.Int {
	r := big.NewInt(0)
	for i:=0; i<length ;i++ {
		r=r.Mul(r,big.NewInt(2))
		r=r.Add(r,big.NewInt(1))
	}
	return r;
}

func equation_1 (first []byte, second []byte) bool {
	a := binArr2BInt(first)
	b := binArr2BInt(second)


	x := new(big.Int)
	x.Set(a)
	x = x.Add(x,b)
	
	y := new(big.Int)
	y.Set(b)
	y = y.Sub(y,a)
		
	xy := new(big.Int).Set(x)
	xy = xy.Mul(xy,y)
	
	mask := bitMask(len(first))

	xy_f := new(big.Int).Set(xy).And(mask,xy)
	t_f := new(big.Int).Set(p).And(mask,p)
	
	areEqual := xy_f.Cmp(t_f) == 0
	
	
	fmt.Println(fmt.Sprintf("e_1:[%v,%v](%v,%v)->(%v,%v)->%v && %b->  %v[%b] ?= %v[%b]  => %t",
		bin2str(first),bin2str(second),
		a,b,
		x,y,
		xy,mask,
		xy_f,xy_f,
		t_f,t_f,
		areEqual));
	
	
	return areEqual
}

func equation_2 (first []byte, second []byte) bool {
	a := binArr2BInt(first)
	b := binArr2BInt(second)

	aa := new(big.Int).Set(a)
	aa = aa.Mul(aa,aa)
	
	bb := new(big.Int).Set(b)
	bb = bb.Mul(bb,bb)
	
	mask := bitMask(len(first))
	
	aa_c := new(big.Int).Set(aa).Add(aa,p)
	
	aa_c_f := new(big.Int).Set(aa_c).And(mask,aa_c)
	
	bb_f := new(big.Int).Set(bb).And(mask,bb)

	areEqual := aa_c_f.Cmp(bb_f) == 0
	
	fmt.Println(fmt.Sprintf("e_2:[%v,%v](%v,%v)^2->%v,%v->+%v[%b] = %v[%b] && [%b]-> %v[%b] ?= %v[%b]  => %t",
		bin2str(first),bin2str(second),
		a,b,
		aa,bb,
		p,p,
		aa_c,aa_c,
		mask,
		aa_c_f,aa_c_f,
		bb_f,bb_f,
		areEqual));
		
	return areEqual
}

func equation_3 (first []byte, second []byte) bool {
	a := binArr2BInt(first)
	b := binArr2BInt(second)

	aa := new(big.Int).Set(a)
	aa = aa.Mul(aa,a)

	sum :=aa.Add(aa,p)

	bb := new(big.Int).Set(b)
	bb = bb.Mul(bb,b)

	//fmt.Println(fmt.Sprintf("%v, %v",bin2str(first),bin2str(second)));
	return sum.Cmp(bb) == 0
}

func showSolution(n1 *big.Int, n2 *big.Int) {
	fmt.Println(fmt.Sprintf("\nSOLUTION:%v,%v",n1,n2));
}

func isCorrectEquationSystem(first []byte, second []byte) (bool, bool) {
	
	if (len(first) > searchDepthBits){
		return false, false;
	}

	eq_1 := equation_1(first, second)
	eq_2 := equation_2(first, second)
	
	if eq_1 && eq_2 {

		equal := equation_3(first,second)
		if equal {
			return true, true;
		
		}
		
		return false, true;
	}
	return false, false;
}


func dfs(filter func (first []byte, second []byte) (bool,bool),showSolution func (n1 *big.Int, n2 *big.Int)) {
	
	counter := []byte{0}
	
	notFinished: for {
		first:=[]byte{};
		second :=[]byte{};

		for _, value :=range counter {
			first = append(first,value & 1)
			second = append(second,(value/2) & 1)
		}

		founded, isCorrect := filter(first, second)
		if founded {
		
			solution1:= new(big.Int).Set(binArr2BInt(first))
			solution2:= new(big.Int).Set(binArr2BInt(second))
			
			showSolution (solution1, solution2)
		}

		if isCorrect {  			// correct, move to next register
			counter = append(counter,0)
		} else {					// incorrect, increment last one or return to previous register
			nextOne: for {
				item := counter[len(counter)-1]
				if (item > 2) {
					counter=counter[:len(counter)-1] // remove last one
				} else {
					item++
					counter[len(counter)-1] = item
					break nextOne
				}

				if (len(counter)==0) {
					break notFinished
				}
				
				
			}
		}
		if (len(counter) == 0) {break}
	}
}


type pair struct {
	first *big.Int
	second *big.Int
}

func equation_4 (first []byte, second []byte) bool {
	a := binArr2BInt(first)
	d := binArr2BInt(second)

	aa := new(big.Int).Set(a)
	aa = aa.Mul(aa,a)

	sum :=aa.Add(aa,p)
	
	d_a := d.Add(d,a)
	d_ad_a := d_a.Mul(d_a,d_a)
	
	return sum.Cmp(d_ad_a) == 0
}

func equation_5 (first []byte, second []byte) bool {
	a := binArr2BInt(first)
	d := binArr2BInt(second)

	aa := new(big.Int).Set(a)
	aa = aa.Mul(aa,a)

	sum :=aa.Add(aa,p)
	
	d_a := d.Add(d,a)
	d_ad_a := d_a.Mul(d_a,d_a)	
	
	mask := bitMask(len(first))
	
	sum_f := new(big.Int).Set(sum).And(mask,sum)
	d_ad_a_f := new(big.Int).Set(d_ad_a).And(mask,d_ad_a)
	
	return sum_f.Cmp(d_ad_a_f) == 0
}

func equation_6 (first []byte, second []byte) bool {
	a := binArr2BInt(first)
	d := binArr2BInt(second)

	a_a := new(big.Int).Set(a)
	a_a = a_a.Add(a_a,a_a)
	a_a_d := a_a.Add(a_a,d)
	
	a_a_dd := a_a_d.Mul(a_a_d,d)

	mask := bitMask(len(first))
	
	a_a_dd_f := new(big.Int).Set(a_a_dd).And(mask,a_a_dd)
	p_f := new(big.Int).Set(p).And(mask,p)
	
	return a_a_dd_f.Cmp(p_f) == 0
}



func bfs(filter func (first []byte, second []byte) (bool,bool),showSolution func (n1 *big.Int, n2 *big.Int)) {
	
	layer := []pair{
		pair{
			first: big.NewInt(0),
			second: big.NewInt(0),},
	}
	
	for n:=0;n<searchDepthBits;n++ {

		var newLayer = []pair{};

		fmt.Println(fmt.Sprintf("LAYER %v", n))		
		
		for _,p := range layer {

			for i:=0;i<4;i++ {
				bit1 := i & 1
				bit2 := (i/2) & 1
			
				first := new(big.Int).Exp(big.NewInt(2),big.NewInt(int64(n)),nil)
				first = first.Mul(first,big.NewInt(int64(bit1)))
				first.Add(first, p.first)

				second := new(big.Int).Exp(big.NewInt(2),big.NewInt(int64(n)),nil)
				second = second.Mul(second,big.NewInt(int64(bit2)))
				second.Add(second, p.second)

				f := BInt2binArr(first,n+1)
				s := BInt2binArr(second,n+1)

				if equation_5(f, s) && equation_6(f, s){
					
					newLayer = append(newLayer, pair {
						first: first,
						second: second})
						
					//fmt.Println(fmt.Sprintf("%v, %v",bin2str(f),bin2str(s)));
						
				}
				

				found := equation_4(f, s)
				
				if found {
					fmt.Println(fmt.Sprintf("SOLUTION ---------------------------   %v[%b],  %v[%b]",first,first,second,second));
				}

			} 
		}

		fmt.Println(fmt.Sprintf("            COUNT: %v",len(newLayer)))
		
		layer = newLayer;
	}
}

var p *big.Int

func main() {
	var x, _ = new(big.Int).SetString("197", 10) //
	var y, _ = new(big.Int).SetString("173", 10) //
	p = big.NewInt(0).Mul(x, y)

	a := new(big.Int).Set(x).Sub(x,y)
	a = a.Div(a,big.NewInt(2))
	
	b := new(big.Int).Set(x).Add(x,y)
	b = b.Div(b,big.NewInt(2))
	
	d := new(big.Int).Set(b).Sub(b,a)
	
	fmt.Println(fmt.Sprintf("x  :%32b=[%10X]=%v", x,x,x));
	fmt.Println(fmt.Sprintf("y  :%32b=[%10X]=%v", y,y,y));
	fmt.Println(fmt.Sprintf("x*y:%32b=[%10X]=%v", p,p,p));


	fmt.Println(fmt.Sprintf("a  :%32b=[%10X]=%v", a,a,a));
	fmt.Println(fmt.Sprintf("b  :%32b=[%10X]=%v", b,b,b));
	fmt.Println(fmt.Sprintf("d  :%32b=[%10X]=%v", d,d,d));
	
	
	
	fmt.Println("-----------------------------------");
	
	//dfs(isCorrectEquationSystem,showSolution);
	
	bfs(isCorrectEquationSystem,showSolution)
	
			
}