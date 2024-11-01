# Адаптер
Wrapper, Обёртка, Adapter
Адаптер — это структурный паттерн проектирования, который позволяет объектам с несовместимыми интерфейсами работать вместе.

# Проблема
Приложение для торговли на бирже. Оно скачивает биржевые котировки из нескольких источников в XML, а затем рисует графики. В какой-то момент применяется сторонняя библиотека аналитики. Но библиотека поддерживает только формат данных JSON, несовместимый с приложением. Подключить стороннюю библиотеку не выйдет из-за несовместимых форматов данных.

Можно переписать библиотеку, чтобы та поддерживала формат XML. Но, во-первых, это может нарушить работу существующего кода, который уже зависит от библиотеки. А во-вторых,  может не быть доступа к её исходному коду.

# Решение
Можно создать адаптер. Это объект-переводчик, который трансформирует интерфейс или данные одного объекта в такой вид, чтобы он стал понятен другому объекту. При этом адаптер оборачивает один из объектов, так что другой объект даже не знает о наличии первого. Например, можно обернуть объект, работающий в метрах, адаптером, который бы конвертировал данные в футы. 
Адаптеры могут не только переводить данные из одного формата в другой, но и помогать объектам с разными интерфейсами работать сообща. Это работает так: 
1. Адаптер имеет интерфейс, который совместим с одним из объектов. 
2. Поэтому этот объект может свободно вызывать методы адаптера. 
3. Адаптер получает эти вызовы и перенаправляет их второму объекту, но уже в том формате и последовательности, которые понятны второму объекту.
Иногда возможно создать даже двухсторонний адаптер, который работал бы в обе стороны.

Программа может работать со сторонней библиотекой через адаптер.
Таким образом, в приложении биржевых котировок можно создать класс XML_To_JSON_Adapter, который бы оборачивал объект того или иного класса библиотеки аналитики. Код посылал бы адаптеру запросы в формате XML, а адаптер сначала транслировал входящие данные в формат JSON, а затем передавал бы их методам обёрнутого объекта аналитики.

# Структура
# Адаптер объектов
Эта реализация использует агрегацию: объект адаптера «оборачивает», то есть содержит ссылку на служебный объект. Такой подход работает во всех языках программирования. 

1. Клиент — это класс, который содержит существующую бизнес-логику программы. 
2. Клиентский интерфейс описывает протокол, через который клиент может работать с другими классами. 
3. Сервис — это какой-то полезный класс, обычно сторонний. Клиент не может использовать этот класс напрямую, так как сервис имеет непонятный ему интерфейс. 
4. Адаптер — это класс, который может одновременно работать и с клиентом, и с сервисом. Он реализует клиентский интерфейс и содержит ссылку на объект сервиса. Адаптер получает вызовы от клиента через методы клиентского интерфейса, а затем переводит их в вызовы методов обёрнутого объекта в правильном формате. 
5. Работая с адаптером через интерфейс, клиент не привязывается к конкретному классу адаптера. Благодаря этому, вы можете добавлять в программу новые виды адаптеров, независимо от клиентского кода. Это может пригодиться, если интерфейс сервиса вдруг изменится, например, после выхода новой версии сторонней библиотеки.


# Адаптер классов
Эта реализация базируется на наследовании: адаптер наследует оба интерфейса одновременно. Такой подход возможен только в языках, поддерживающих множественное наследование, например, C++.

Структура классов паттерна Адаптер (адаптер классов)
1. Адаптер классов не нуждается во вложенном объекте, так как он может одновременно наследовать и часть существующего класса, и часть сервиса.

# Псевдокод
В этом примере Адаптер преобразует один интерфейс в другой, позволяя совместить квадратные колышки и круглые отверстия.
Адаптер вычисляет наименьший радиус окружности, в которую можно вписать квадратный колышек, и представляет его как круглый колышек с этим радиусом.

```java
// Классы с совместимыми интерфейсами: КруглоеОтверстие и
// КруглыйКолышек.
class RoundHole is
    constructor RoundHole(radius) { ... }

    method getRadius() is
        // Вернуть радиус отверстия.

    method fits(peg: RoundPeg) is
        return this.getRadius() >= peg.getRadius()

class RoundPeg is
    constructor RoundPeg(radius) { ... }

    method getRadius() is
        // Вернуть радиус круглого колышка.


// Устаревший, несовместимый класс: КвадратныйКолышек.
class SquarePeg is
    constructor SquarePeg(width) { ... }

    method getWidth() is
        // Вернуть ширину квадратного колышка.


// Адаптер позволяет использовать квадратные колышки и круглые
// отверстия вместе.
class SquarePegAdapter extends RoundPeg is
    private field peg: SquarePeg

    constructor SquarePegAdapter(peg: SquarePeg) is
        this.peg = peg

    method getRadius() is
        // Вычислить половину диагонали квадратного колышка по
        // теореме Пифагора.
        return peg.getWidth() * Math.sqrt(2) / 2


// Где-то в клиентском коде.
hole = new RoundHole(5)
rpeg = new RoundPeg(5)
hole.fits(rpeg) // TRUE

small_sqpeg = new SquarePeg(5)
large_sqpeg = new SquarePeg(10)
hole.fits(small_sqpeg) // Ошибка компиляции, несовместимые типы

small_sqpeg_adapter = new SquarePegAdapter(small_sqpeg)
large_sqpeg_adapter = new SquarePegAdapter(large_sqpeg)
hole.fits(small_sqpeg_adapter) // TRUE
hole.fits(large_sqpeg_adapter) // FALSE
```
# Применимость
1. Когда надо использовать сторонний класс, но его интерфейс не соответствует остальному коду приложения. Адаптер позволяет создать объект-прокладку, который будет превращать вызовы приложения в формат, понятный стороннему классу. 
2. Когда нужно использовать несколько существующих подклассов, но в них не хватает какой-то общей функциональности, причём расширить суперкласс не получится. Можно попробовать создать ещё один уровень подклассов и добавить в них недостающую функциональность. Но при этом придётся дублировать один и тот же код в обеих ветках подклассов.  Более элегантным решением было бы поместить недостающую функциональность в адаптер и приспособить его для работы с суперклассом. Такой адаптер сможет работать со всеми подклассами иерархии. Это решение будет сильно напоминать паттерн Декоратор.

# Шаги реализации
1. Убедиться, что есть два класса с несовместимыми интерфейсами: 
   1. полезный сервис — служебный класс, который нельзя изменять (он либо сторонний, либо от него зависит другой код);
   2. один или несколько клиентов — существующих классов приложения, несовместимых с сервисом из-за неудобного или несовпадающего интерфейса.
2. Описать клиентский интерфейс, через который классы приложения смогли бы использовать класс сервиса. 
3. Создать класс адаптера, реализовав этот интерфейс. 
4. Поместить в адаптер поле, которое будет хранить ссылку на объект сервиса. Обычно это поле заполняют объектом, переданным в конструктор адаптера. В случае простой адаптации этот объект можно передавать через параметры методов адаптера. 
5. Реализовать все методы клиентского интерфейса в адаптере. Адаптер должен делегировать основную работу сервису. 
6. Приложение должно использовать адаптер только через клиентский интерфейс. Это позволит легко изменять и добавлять адаптеры в будущем.

# Преимущества и недостатки
1. Отделяет и скрывает от клиента подробности преобразования различных интерфейсов. 
2. Усложняет код программы из-за введения дополнительных классов.

# Отношения с другими паттернами
1. Мост проектируют загодя, чтобы развивать большие части приложения отдельно друг от друга. Адаптер применяется постфактум, чтобы заставить несовместимые классы работать вместе. 
2. Адаптер предоставляет совершенно другой интерфейс для доступа к существующему объекту. С другой стороны, при использовании паттерна Декоратор интерфейс либо остается прежним, либо расширяется. Причём Декоратор поддерживает рекурсивную вложенность, чего не скажешь об Адаптере. 
3. С Адаптером вы получаете доступ к существующему объекту через другой интерфейс. Используя Заместитель, интерфейс остается неизменным. Используя Декоратор, вы получаете доступ к объекту через расширенный интерфейс. 
4. Фасад задаёт новый интерфейс, тогда как Адаптер повторно использует старый. Адаптер оборачивает только один класс, а Фасад оборачивает целую подсистему. Кроме того, Адаптер позволяет двум существующим интерфейсам работать сообща, вместо того, чтобы задать полностью новый. 
5. Мост, Стратегия и Состояние (а также слегка и Адаптер) имеют схожие структуры классов — все они построены на принципе «композиции», то есть делегирования работы другим объектам. Тем не менее, они отличаются тем, что решают разные проблемы. Паттерны — это не только рецепт построения кода определённым образом, но и описание проблем, которые привели к данному решению.

# Концептуальный пример
Адаптер — это структурный паттерн, который позволяет подружить несовместимые объекты.Адаптер выступает прослойкой между двумя объектами, превращая вызовы одного в вызовы понятные другому.

Имеется код клиента, ожидающий определенных от объекта определенных качеств (порт Lightning), но другой объект под названием adaptee (ноутбук на Windows), который предоставляет тот же функционал, но через другой интерфейс (USB порт).
В такой ситуации подойдет паттерн Адаптер. Будет создана структура adapter, которая будет:
 - реализовать тот же интерфейс, который ожидает клиент (порт Lightning); 
 - переводить запрос от клиента к адаптируемому объекту в форме, которую он ожидает. Адаптер принимает коннектор Lightning, после чего переводит его сигналы в формат USB в ноутбуке на Windows.