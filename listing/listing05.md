Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Вывод: error

Почему:
В Go интерфейсы состоят из двух частей: конкретного значения и конкретного типа. Даже если конкретное значение интерфейса равно nil, интерфейс может не быть nil, если конкретный тип не является nil.

В функции test возвращается *customError, которая является указателем на структуру customError. Несмотря на то, что значение, возвращаемое функцией test, равно nil, его тип - это *customError.

Данный пример похож на listing03. Интерфейс err не является nil, так как его типовая часть не является nil.
```