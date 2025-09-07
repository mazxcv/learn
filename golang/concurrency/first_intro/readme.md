# Конкурентность в go

https://www.youtube.com/watch?v=4aTt9E-EG-o

Конкурентно на самом деле код может быть запущен
-  асинхронно на том же процессоре и в том же потоке
-  в соседнем потоке
-  на другом ядре процессора 
-  
Это выбирает runtime Go, точнее планировщик, программист этим не управляет


# Что такое gorotine 
функция, которая может запускаться конкурентно вместе с другими функиями

Остается решить три проблемы:
 - sync, проблема синхронизации
 - race, проблема race condition, data race
 - leak, проблема утечек горутин

Модель конкуренции Fork/Join
В ней мы не управляем моментом запуска.
Из-за этого необходимо выполнять синхронизацию горутин

|
|
|
go g1()------------() g1
|                  |
|                  |
|                  | run body of func g1()
|                  |
() join------------()
|
|
\/ end


```golang
wg := &sync.WaitGroup{}

wg.Add(1)

go func() {
    defer wg.Done()

    ...
}

wg.Wait() // (JOIN)

```

# Race condition. Data Race. Состояние гонки 
Когда две или более горутин хотят получить доступ к одной структуре данных
Самое опасное: не всегда видна гонка данных.
`go run --race`
```golang

const count = 50

func main() {

	var money atomic.Int32
	var donation atomic.Int32

	mu := &sync.RWMutex{}

	go func() {
		for {
			mu.RLock()
			m := money.Load()
			d := donation.Load()
			mu.RUnlock()

			if m != d {
				fmt.Println("error", m, d)
				break
			}
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(count)
	for range count {
		go func() {
			defer wg.Done()
			mu.Lock()
			money.Add(1)
			donation.Add(1)
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(money.Load())
	fmt.Println(donation.Load())

}

```



# Каналы 
Канал используется как для синхронизации, так и для доступа к общим ресурсам

Операции чтения и записи в канал - блокирующие 
Запись в канал будет блокирована до тех пор порка из  каналане будет прочитано
Чтение  не раблокируется пока в канал что-то не запишется
Поэтому надо использовать каналы по назначению: они должны работать с горутинами

```golang

	ch := make(chan int)

for 
	go func() {
		ch <- 1
	}()

	val := <-ch

	fmt.Println(val)

``` 

```golang

func main() {

	ch := make(chan int)

	for i := range 5 {
		go func() {
			ch <- i + 1
		}()
	}

	for range 5 {
		val := <-ch

		fmt.Println(val)
	}
}

```