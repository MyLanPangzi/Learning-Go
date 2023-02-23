package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
//func checkImplementation(i any, targets ...any) {
//	t := reflect.TypeOf(i)
//	if t.Kind() == reflect.Pointer && t.Elem().Kind() == reflect.Interface {
//		it := t.Elem()
//		for _, target := range targets {
//			tt := reflect.TypeOf(target)
//			fmt.Printf("%v implements %v %v\n", tt, it, tt.Implements(it))
//		}
//	}
//}
//func main() {
//	checkImplementation((*CurrencyItem)(nil), Product{}, &Product{}, &Purchase{})
//	var p CurrencyItem
//	var ptr *Product
//	p = &Product{}
//	//ptr = &Product{}
//	pt := reflect.TypeOf(p)
//	fmt.Println(pt, pt.Elem(), pt.Elem().Kind())
//	p = nil
//	pt = reflect.TypeOf(p)
//	fmt.Println(pt, p == nil, ptr == nil)
//	p = (*Product)(nil)
//	pt = reflect.TypeOf(p)
//	fmt.Println(pt, p == nil, ptr == nil)
//
//	pt = reflect.TypeOf((*CurrencyItem)(nil))
//	fmt.Println(pt.Elem(), pt.Elem().Kind())
//	//pt = reflect.TypeOf((CurrencyItem)(nil))
//	//fmt.Println(pt, pt.Kind())
//}
//
////func checkImplementation(check interface{}, targets ...interface{}) {
////    checkType := reflect.TypeOf(check)
////    if (checkType.Kind() == reflect.Ptr &&
////            checkType.Elem().Kind() == reflect.Interface) {
////        checkType := checkType.Elem()
////        for _, target := range targets {
////            targetType := reflect.TypeOf(target)
////            Printfln("Type %v implements %v: %v",
////                targetType, checkType, targetType.Implements(checkType))
////        }
////    }
////}
////func main() {
////    currencyItemType := (*CurrencyItem)(nil)
////    checkImplementation(currencyItemType, Product{}, &Product{}, &Purchase{})
////}
