package main

import (
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const TOKEN = "token"

var bot *tgbotapi.BotAPI
var chatId int64

var fortuneTellerNames = [3]string{"4el", "Krutoy", "krutoy"}

var answers = []string{
	"Так",
	"Ні",
	"Що таке парадигма програмування і назвіть основні парадигми програмування. (ООП, процедурне, функціональне, логічне, декларативне)",

	"Яка різниця між компіляцією та інтерпретацією програмного коду? (Компіляція перетворює вихідний код у виконуваний файл, тоді як інтерпретація виконує код по одній команді в разі запуску)",

	"Що таке змінна (variable) в програмуванні? (Змінна - це ім'я або посилання на місце в пам'яті, де зберігається значення)",

	"Що означає термін `рекурсія` в програмуванні? (Рекурсія - це техніка, коли функція викликає саму себе)",

	"Які основні типи даних існують в програмуванні? (Цілі числа, дійсні числа, рядки, булеві значення, масиви, об'єкти тощо)",

	"Що таке алгоритм в програмуванні? (Алгоритм - це набір кроків для вирішення конкретної задачі)",

	"Яка різниця між стеком (stack) і чергою (queue)? (Стек - це структура даних, що працює за принципом `останній ввійшов - перший вийшов`, а черга - за принципом `перший ввійшов - перший вийшов`)",

	"Що таке API і чому воно важливе для програмування? (API - це набір правил та протоколів, що дозволяють різним програмам взаємодіяти між собою)",

	"Яка роль версійного контролю (version control) у розробці програмного забезпечення? (Дозволяє зберігати та відстежувати зміни в коді, а також спільно працювати над проектом)",

	"Що таке тестування програмного забезпечення і чому воно важливе? (Тестування - це процес перевірки програми на відповідність вимогам та виявлення помилок перед випуском в експлуатацію)",
}

func connectWithTelegram() {
	var err error
	if bot, err = tgbotapi.NewBotAPI(TOKEN); err != nil {
		panic("Cannot connect to Telegram")
	}
}

func sendMessage(msg string) {
	msgConfig := tgbotapi.NewMessage(chatId, msg)
	bot.Send(msgConfig)
}

func isMessageForFortuneTeller(update *tgbotapi.Update) bool {
	if update.Message == nil || update.Message.Text == "" {
		return false
	}

	msgInLowerCase := strings.ToLower(update.Message.Text)
	for _, name := range fortuneTellerNames {
		if strings.Contains(msgInLowerCase, name) {
			return true
		}
	}
	return false
}

func getFortuneTellerAnswer() string {
	index := rand.Intn(len(answers))
	return answers[index]
}

func sendAnswer(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(chatId, getFortuneTellerAnswer())
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func main() {
	connectWithTelegram()

	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatId = update.Message.Chat.ID

			sendMessage("Привіт! Я Krutoy 4el, чим можу допомогти? \"Так\" чи \"Ні\". Наприклад, \"Anton, чи можна читати інший код?\"")
		}

		if isMessageForFortuneTeller(&update) {
			sendAnswer(&update)
		}
	}
}
