import logging, requests
import config as cfg
from aiogram import Bot, Dispatcher, executor, types
import markup as nav


# init logger
logger = logging.getLogger("telegram-bot")
logger.setLevel(logging.INFO)
file_handler = logging.FileHandler("./logger/telegram.log", mode="a")
console_handler = logging.StreamHandler()
formatter = logging.Formatter("%(name)s %(asctime)s %(levelname)s %(message)s")

file_handler.setFormatter(formatter)
console_handler.setFormatter(formatter)
logger.addHandler(file_handler)
logger.addHandler(console_handler)

# get config
config = cfg.get_config()

# init bot and dispatcher
bot = Bot(config.Token)
dp = Dispatcher(bot)

@dp.message_handler(commands=['start'])
async def start(message:types.Message):
    if message.chat.type == 'private':
        if message.from_user.id in config.Admins:
            await bot.send_message(message.chat.id, "Добро пожаловать!", reply_markup=nav.main_menu)
        else:
            await bot.send_message(message.chat.id, "Отказано")

@dp.message_handler()
async def main(message:types.Message):
    if message.chat.type == 'private':
        if message.text == "Обновить токен":
            r = requests.get(f"{config.Addr}/api/getJWT")
            config.update_jwt(r.json()['token'])

            await bot.send_message(message.chat.id, "Успешно")


if __name__ == "__main__":
    executor.start_polling(dp, skip_updates = True) 