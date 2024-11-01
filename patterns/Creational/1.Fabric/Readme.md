# Фабричный метод

Порождающий паттерн проектирования, который определеяет общий интерфейс для создания объектов в суперклассе, позволяя подклассам изменять тип собздаваемы объектов.

При добавлении новых классов может оказаться, что надо изменить большую часть кода программы, и получится код, наполненный условными операторами, которые выполняют то или иное действие, в зависимости от класса.
Паттерн Фабричный метод предлагает создавать объекты не напрямую, через new, а через вызов фабричного метода, который скрывает в себе вызов new. Теперь можно переопределить фабричный метод в подклассе, чтобы изменить тип создаваемого продукта. Чтобы все работало, все созвращаемые объекты должны иметь общий интерфейс. Подклассы смогут производить объекты различных классов, следующих одному и тому же интерфейсу.

# Структура

Продукт определяет общий интерфейс объектов, которые может произвести создатель и его подклассы.

Конкретные продукты содержат код различных продуктов. Продукты будут отличаться реализацией, но интерфейс у них будет общий.

Создатель объявляет фабричный метод, который должен возвращать новые объекты продуктов. Важно, чтобы тип результата совпадал с общим интерфейсом продуктов.

Зачастую фабричный метод объявляют абстрактным, чтобы заставить все подклассы реализовать его по-своему. Но он может возвращать и некий стандартный продукт.

Несмотря на название, важно понимать, что создание продуктов не является единственной функцией создателя. Обычно он содержит и другой полезный код работы с продуктом. Аналогия: большая софтверная компания может иметь центр подготовки программистов, но основная задача компании — создавать программные продукты, а не готовить программистов.

Конкретные создатели по-своему реализуют фабричный метод, производя те или иные конкретные продукты.

Фабричный метод не обязан всё время создавать новые объекты. Его можно переписать так, чтобы возвращать существующие объекты из какого-то хранилища или кэша.

```java
// Паттерн Фабричный метод применим тогда, когда в программе
// есть иерархия классов продуктов.
interface Button is
    method render()
    method onClick(f)

class WindowsButton implements Button is
    method render(a, b) is
        // Отрисовать кнопку в стиле Windows.
    method onClick(f) is
        // Навесить на кнопку обработчик событий Windows.

class HTMLButton implements Button is
    method render(a, b) is
        // Вернуть HTML-код кнопки.
    method onClick(f) is
        // Навесить на кнопку обработчик события браузера.


// Базовый класс фабрики. Заметьте, что "фабрика" — это всего
// лишь дополнительная роль для класса. Скорее всего, он уже
// имеет какую-то бизнес-логику, в которой требуется создание
// разнообразных продуктов.
class Dialog is
    method render() is
        // Чтобы использовать фабричный метод, вы должны
        // убедиться в том, что эта бизнес-логика не зависит от
        // конкретных классов продуктов. Button — это общий
        // интерфейс кнопок, поэтому все хорошо.
        Button okButton = createButton()
        okButton.onClick(closeDialog)
        okButton.render()

    // Мы выносим весь код создания продуктов в особый метод,
    // который назвают "фабричным".
    abstract method createButton():Button


// Конкретные фабрики переопределяют фабричный метод и
// возвращают из него собственные продукты.
class WindowsDialog extends Dialog is
    method createButton():Button is
        return new WindowsButton()

class WebDialog extends Dialog is
    method createButton():Button is
        return new HTMLButton()


class Application is
    field dialog: Dialog

    // Приложение создаёт определённую фабрику в зависимости от
    // конфигурации или окружения.
    method initialize() is
        config = readApplicationConfigFile()

        if (config.OS == "Windows") then
            dialog = new WindowsDialog()
        else if (config.OS == "Web") then
            dialog = new WebDialog()
        else
            throw new Exception("Error! Unknown operating system.")

    // Если весь остальной клиентский код работает с фабриками и
    // продуктами только через общий интерфейс, то для него
    // будет не важно, какая фабрика была создана изначально.
    method main() is
        this.initialize()
        dialog.render()
```

# Шаги реализации
1. Привести все создаваемые продукты к общему интерфейсу.
2. В классе, который производит продукты, создать пустой фабричный метод. В качестве возвращаемого типа указать общий интерфейс продукта.
3. Затем пройтись по коду класса и найти все участки, создающие продукты. Поочерёдно заменить эти участки вызовами фабричного метода, перенося в него код создания различных продуктов. В фабричный метод, возможно, придётся добавить несколько параметров, контролирующих, какой из продуктов нужно создать. На этом этапе фабричный метод, скорее всего, будет выглядеть удручающе. В нём будет жить большой условный оператор, выбирающий класс создаваемого продукта.
4. Для каждого типа продуктов завести подкласс и переопределеить в нём фабричный метод. Переместить туда код создания соответствующего продукта из суперкласса. 
5. Если создаваемых продуктов слишком много для существующих подклассов создателя, можно подумать о введении параметров в фабричный метод, которые позволят возвращать различные продукты в пределах одного подкласса. Например, есть класс Почта с подклассами АвиаПочта и НаземнаяПочта, а также классы продуктов Самолёт, Грузовик и Поезд. Авиа соответствует Самолётам, но для НаземнойПочты есть сразу два продукта. Можно создать новый подкласс почты для поездов, но проблему можно решить и по-другому. Клиентский код может передавать в фабричный метод НаземнойПочты аргумент, контролирующий тип создаваемого продукта.
6. Если после всех перемещений фабричный метод стал пустым, можно его абстрактным. Если в нём что-то осталось — не беда, это будет его реализацией по умолчанию.

# Преимущества и недостатки

1. Избавляет класс от привязки к конкретным классам продуктов. 
2. Выделяет код производства продуктов в одно место, упрощая поддержку кода. 
3. Упрощает добавление новых продуктов в программу. 
4. Реализует принцип открытости/закрытости.

Но, может привести к созданию больших параллельных иерархий классов, так как для каждого класса продукта надо создать свой подкласс создателя.

# Отношения с другими паттернами
1. Многие архитектуры начинаются с применения Фабричного метода (более простого и расширяемого через подклассы) и эволюционируют в сторону Абстрактной фабрики, Прототипа или Строителя (более гибких, но и более сложных). 
2. Классы Абстрактной фабрики чаще всего реализуются с помощью Фабричного метода, хотя они могут быть построены и на основе Прототипа. 
3. Фабричный метод можно использовать вместе с Итератором, чтобы подклассы коллекций могли создавать подходящие им итераторы. 
4. Прототип не опирается на наследование, но ему нужна сложная операция инициализации. Фабричный метод, наоборот, построен на наследовании, но не требует сложной инициализации. 
5. Фабричный метод можно рассматривать как частный случай Шаблонного метода. Кроме того, Фабричный метод нередко бывает частью большого класса с Шаблонными методами.

# Концептуальный пример
В Go невозможно реализовать классический вариант паттерна Фабричный метод, поскольу в языке отсутствуют возможности ООП, в том числе классы и наследственность. Несмотря на это, мы все же можем реализовать базовую версию этого паттерна — Простая фабрика.

В этом примере мы будем создавать разные типы оружия при помощи структуры фабрики.

Сперва, мы создадим интерфейс iGun, который определяет все методы будущих пушек. Также имеем структуру gun (пушка), которая применяет интерфейс iGun. Две конкретных пушки — ak47 и musket — обе включают в себя структуру gun и не напрямую реализуют все методы от iGun.

gunFactory служит фабрикой, которая создает пушку нужного типа в зависимости от аргумента на входе. Клиентом служит main.go . Вместо прямого взаимодействия с объектами ak47 или musket, он создает экземпляры различного оружия при помощи gunFactory, используя для контроля изготовления только параметры в виде строк.