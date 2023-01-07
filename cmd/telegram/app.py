import logging
from aiogram import Bot, Dispatcher, executor, types


logger = logging.getLogger("telegram-bot")
logger.setLevel(logging.INFO)
file_handler = logging.FileHandler("./logger/telegram.log", mode="a")
console_handler = logging.StreamHandler()
formatter = logging.Formatter("%(name)s %(asctime)s %(levelname)s %(message)s")

file_handler.setFormatter(formatter)
logger.addHandler(file_handler)
logger.addHandler(console_handler)

token = ""

bot = Bot(token)
dp = Dispatcher(bot)


@dp.message_handler(commands=['start'])
async def start(message:types.Message):
    if message.chat.type == 'private':
        #await bot.send_message(message.chat.id, "Привет👋")
        #await bot.send_message(message.chat.id, 'Выбери сервис ниже.')
        pass
