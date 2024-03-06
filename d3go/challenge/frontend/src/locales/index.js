// index.js
import { createI18n } from 'vue-i18n'
import zh from './zh-CN'
import en from './en'


const messages = {
    en,
    zh,
};

const language = (navigator.language || 'en').toLocaleLowerCase();
const i18n = createI18n({
    locale: localStorage.getItem('lang') || language.split('-')[0] || 'en',
    fallbackLocale: 'en',
    globalInjection: true,
    messages,
});

export default i18n;