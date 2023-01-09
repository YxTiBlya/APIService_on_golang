import logging, requests
import config as cfg
from aiogram.dispatcher.filters.state import State, StatesGroup
from aiogram.contrib.fsm_storage.memory import MemoryStorage
from aiogram.dispatcher import FSMContext
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

# init bot, storage, dispatcher
bot = Bot(config.Token)
storage = MemoryStorage()   
dp = Dispatcher(bot, storage=storage)

class Form(StatesGroup):
    get_id = State()

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
        
        elif message.text == "Контакты":
            try:
                r = requests.get(f"{config.Addr}/api/contact/get", headers={"token": config.JWT})
                response = r.json()["response"]

                msg = "Результат:\n\n"
                for contact in response:
                    msg += f"id: {contact['id']}\nномер: {contact['number']}\nкод оператора: {contact['operator_code']}\nтег: {contact['tag']}\nчасовой пояс: {contact['time_zone']}\n\n"

                await bot.send_message(message.chat.id, msg)

            except Exception as err:
                logger.info(err)
                await bot.send_message(message.chat.id, "Произошла ошибка, попробуйте поменять токен")

        elif message.text == "Рассылки":
            try:
                r = requests.get(f"{config.Addr}/api/mailing/get", headers={"token": config.JWT})
                response = r.json()["response"]

                msg = "Результат:\n\n"
                for mailing in response:
                    msg += f"id: {mailing['id']}\nстарт: {mailing['start_time']}\nсообщение: {mailing['message']}\nфильтры: {mailing['filters']}\nконец: {mailing['end_time']}\n\n"

                await bot.send_message(message.chat.id, msg)
            except Exception as err:
                logger.info(err)
                await bot.send_message(message.chat.id, "Произошла ошибка, попробуйте поменять токен")

        elif message.text == "Сообщения":
            try:
                r = requests.get(f"{config.Addr}/api/message/get", headers={"token": config.JWT})
                response = r.json()["response"]

                msg = "Результат:\n\n"
                for rmessage in response:
                    msg += f"id: {rmessage['id']}\nвремя отправки: {rmessage['datetime']}\nстатус: {rmessage['status']}\nid рассылки: {rmessage['mailing_id']}\nid контакта: {rmessage['contact_id']}\n\n"

                await bot.send_message(message.chat.id, msg)
            except Exception as err:
                logger.info(err)
                await bot.send_message(message.chat.id, "Произошла ошибка, попробуйте поменять токен")
        
        elif message.text == "Статистика":
            await bot.send_message(message.chat.id, "Введите id рассылки", reply_markup=nav.cancel)
            await Form.get_id.set()

        else:
            await bot.send_message(message.chat.id, "Такая команда отсутствует")

@dp.message_handler(state=Form.get_id)
async def get_id(message: types.Message, state: FSMContext):
    id = message.text

    if id == "Отмена":
        await state.reset_state()
        await bot.send_message(message.chat.id, "Отменено!", reply_markup=nav.main_menu)
        return

    try:
        r = requests.get(f"{config.Addr}/api/stat/get/{id}", headers={"token": config.JWT})
        response = r.json()["response"]

        msg = f"Результат\nСообщения отправленные по рассылке {id}:\n\n"
        for rmessage in response:
            msg += f"id: {rmessage['id']}\nвремя отправки: {rmessage['datetime']}\nстатус: {rmessage['status']}\nid рассылки: {rmessage['mailing_id']}\nid контакта: {rmessage['contact_id']}\n\n"

        await bot.send_message(message.chat.id, msg, reply_markup=nav.main_menu)
    except Exception as err:
        logger.info(err)
        await bot.send_message(message.chat.id, "Произошла ошибка, попробуйте поменять токен", reply_markup=nav.main_menu)

    await state.finish()

if __name__ == "__main__":
    executor.start_polling(dp, skip_updates = True) 