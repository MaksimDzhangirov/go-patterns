# Шаблон проектирования "Шаблонный метод" в Go

[Оригинал](https://golangbyexample.com/template-method-design-pattern-golang/)

## Введение

Шаблон "Шаблонный метод" - это поведенческий паттерн проектирования, который 
позволяет вам определять шаблон или алгоритм для конкретной операции. Давайте
попытаемся понять принцип работы шаблона на примере.

Рассмотрим пример с одноразовым паролем или OTP (One Time Password). Существуют 
различные типы одноразовых паролей, которые могут быть посланы через SMS или 
электронную почту. Но не зависимо от способа передачи, сам процесс генерации 
OTP всегда одинаков. Он состоит из следующих шагов:

* сгенерировать случайное n-разрядное число.
* сохранить его в кэше для последующей проверки.
* подготовить содержимое уведомления.
* послать уведомление.
* сохранить метрики.

Даже если в будущем будет введена отправка OTP через push-уведомления, в ней
всё равно будут использоваться описанные выше шаги.

В таких случаях, когда шаги конкретной операции одинаковы, но шаги могут быть 
реализованы по-разному различными разработчиками, стоит использовать шаблон 
проектирования "Шаблонный метод". Мы определяем шаблон или алгоритм, который
состоит из фиксированного количества шагов. Тот, кто будет реализовывать 
операции, переопределяет методы шаблона.

Посмотрите на приведенный ниже код.

* **iOtp** представляет собой интерфейс, определяющий набор методов, которые должен 
реализовывать любой `otp`.
* **sms** и **email** реализуют **iOtp** интерфейс.
* **otp** - это структура, в которой определен шаблонный метод **genAndSendOTP()**. В **otp** 
  встроен интерфейс **iOtp**.
  
Обратите внимание: комбинация интерфейса `iOtp` и структуры `otp` позволяет реализовать
абстрактный класс в Go. Смотри реализацию по [ссылке](https://github.com/MaksimDzhangirov/oop-go/blob/master/docs/abstract_class.md).

## Пример:

**interfaces/otp.go**

```go
type IOtp interface {
    GenRandomOTP(int) string
    SaveOTPCache(string)
    GetMessage(string) string
    SendNotification(string) error
    PublishMetric()
}
```

**otp/otp.go**

```go
type otp struct {
    iOtp interfaces.IOtp
}

func NewOTP(iOtp interfaces.IOtp) *otp {
    return &otp{
        iOtp: iOtp,
    }
}

func (o *otp) GenAndSendOTP(otpLength int) error {
    otp := o.iOtp.GenRandomOTP(otpLength)
    o.iOtp.SaveOTPCache(otp)
    message := o.iOtp.GetMessage(otp)
    err := o.iOtp.SendNotification(message)
    if err != nil {
        return err
    }
    o.iOtp.PublishMetric()
    return nil
}
```

**otp/sms.go**

```go
type sms struct {
    otp
}

func NewSms() *sms {
    return &sms{}
}

func (s *sms) GenRandomOTP(len int) string {
    randomOTP := "1234"
    fmt.Printf("SMS: generating random otp %s\n", randomOTP)
    return randomOTP
}

func (s *sms) SaveOTPCache(otp string) {
    fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *sms) GetMessage(otp string) string {
    return "SMS OTP for login is " + otp
}

func (s *sms) SendNotification(message string) error {
    fmt.Printf("SMS: sending sms: %s\n", message)
    return nil
}

func (s *sms) PublishMetric() {
    fmt.Printf("SMS: publishing metric\n")
}
```

**otp/email.go**

```go
type email struct {
    otp
}

func NewEmail() *email {
    return &email{}
}

func (s *email) GenRandomOTP(len int) string {
    randomOTP := "1234"
    fmt.Printf("EMAIL: generaing random otp %s\n", randomOTP)
    return randomOTP
}

func (s *email) SaveOTPCache(otp string) {
    fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (s *email) GetMessage(otp string) string {
    return "EMAIL OTP for login is " + otp
}

func (s *email) SendNotification(message string) error {
    fmt.Printf("EMAIL: sending email: %s\n", message)
    return nil
}

func (s *email) PublishMetric() {
    fmt.Printf("EMAIL: publishing metrics\n")
}
```

**main.go**

```go
func main() {
    smsOTP := otp.NewSms()
    o := otp.NewOTP(smsOTP)
    err := o.GenAndSendOTP(4)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("")
    emailOTP := otp.NewEmail()
    o = otp.NewOTP(emailOTP)
    err = o.GenAndSendOTP(4)
    if err != nil {
        log.Fatalln(err)
    }
}
```

Результат в терминале:

```shell
go run main.go
SMS: generating random otp 1234
SMS: saving otp: 1234 to cache
SMS: sending sms: SMS OTP for login is 1234
SMS: publishing metric

EMAIL: generaing random otp 1234
EMAIL: saving otp: 1234 to cache
EMAIL: sending email: EMAIL OTP for login is 1234
EMAIL: publishing metrics
```