from aiogram.types import InlineKeyboardMarkup, InlineKeyboardButton, ReplyKeyboardMarkup, KeyboardButton

contacts = KeyboardButton("Контакты")
mailing = KeyboardButton("Рассылки")
messages = KeyboardButton("Сообщения")
refr_token = KeyboardButton("Обновить токен")
stat = KeyboardButton("Статистика")
main_menu = ReplyKeyboardMarkup(resize_keyboard=True).add(refr_token).row(contacts,mailing,messages).add(stat)