# 電卓

## BNF

```
   expr := factor | expr ADD factor | expr SUB factor;
   factor := primary | factor MUL primary | factor DIV primary;
   primary := INT_NUM | LP expr RP;
   
```