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

## 48. Reference vs Value Types
> Arrays
>
>> Primitive data structure
>
>> Can't be resized
>
>> Rarely used directly
> -------------------------------

> Slices
>
>> Can grow and shrink
>
>> Used 99% of the time for lists of elements
> -------------------------------------

- Anytime we make a slice, Go makes a slice(containing information, such as the the pointer to the slice's content - pointer to head, its capacity and its length) and an array with the slice content;
- Although Go makes a copy of the slice information when a slice is passed to a function, the pointer still points to the original slice.

> VALUE TYPES:
>
> Use pointers to change these things in a function
>
>> int
>
>> float
>
>> string
>
>> bool
>
>> structs
> --------------------------------------------

> REFERENCE TYPES:
>
> Don't worry about pointers with these
>
>> slices
>
>> maps
>
>> channels
>
>> pointers
>
>> functions
> -------------------------------------