go方法分为value receivers和pointer receivers  
1、两者大多时候都能自动转换，比如指针能调用value方法，数值能调用pointer方法  
2、但是如果要修改原数据，只能使用pointer方法，因为value方法是拷贝。  
3、对于大结构，pointer方法能减少拷贝消耗，但是使用不当可能会导致引用外泄增大内存消耗  
4、对于立即数或者其他不可更改的数值（unaddressable value），只能调用value方法。