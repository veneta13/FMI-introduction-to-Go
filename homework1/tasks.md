## Filter

Функция, която приема предикат p и връща функция, приемаща произволен брой аргументи и филтрираща тези елементи x, за които p(x) е true.

**func Filter(p func(int) bool) func(...int) []int**

#### Пример:

```go
    odds := Filter(func(x int) bool { return x%2 == 1 })
    evens := Filter(func(x int) bool { return x%2 == 0 })

    odds(1, 2, 3, 4, 5)   // [1 3 5]
    evens(1, 2, 3, 4, 5)  // [2 4]
    odds(6, 7, 8, 9, 10)  // [7 9]
    evens(6, 7, 8, 9, 10) // [6 8 10]
```

## Mapper

Функция, която приема функция m, работеща върху цели числа и връща функция, която прилага m върху произволен брой цели числа, подадени като аргументи.

**func Mapper(f func(int) int) func(...int) []int**

#### Пример:

```go
    double := Mapper(func(a int) int { return 2 * a })

    double(1, 2, 3) // [2, 4, 6]
    double(4, 5, 6) // [8, 10, 12]
```

## Reducer

Фунцкия, приемаща начална стойност и редуцираща функция. Трябва да връща функция, която редуцира произволен брой числа и връща текущия резултат от редуцирането.

**func Reducer(initial int, f func(int, int) int) func(...int) int**

#### Пример:

```go
    sum := Reducer(0, func(a, b int) int { return a + b })

    sum(1, 2, 3)       // 6
    sum(5)             // 11
    sum(100, 101, 102) // 314
```

## MapReducer

Функция, която създава map reducer функция за int аргументи с подадени map функция, reduce функция и първоначална стойност initial за reduce функцията.

**func MapReducer(initial int, mapper func(int) int, reducer func(int, int) int) func(...int) int {**

#### Пример:

```go
    powerSum := MapReducer(
        0,
        func(v int) int { return v * v },
        func(a, v int) int { return a + v },
    )

    powerSum(1, 2, 3, 4) // 30
    powerSum(1, 2, 3, 4) // 60
    powerSum()           // 60
```
