# LECTURES

## 43. Pass by Value
> RAM
>
> Address            Value
>
> 0000              
>
> 0001               jim person (jim := person{...})
>
> 0002
>
> 0003               p person (jim.updateFirstName(...))
>
> 0004
>
> -----------------------------------------

## 45. Pointer Operations
- &variable: give me the memory address of the value this variable is pointing at
- *pointer: give me the value this memory address is pointing at

- Whenever we use *, we are referencing an address, but when we use *variableType, it means we are working with a pointer to that variable type (type description). And *pointerOfVariable means we want to manipulate the value the pointer is referencing (operator).

- (*pointerOfVariable): takes the actual value of the variable.

- TURN address INTO value WITH *address;
- TURN value INTO address WITH &value.