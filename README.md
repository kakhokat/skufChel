## Запуск бэк части 
docker compose up(скорее всего придется перезапускать compose и во второй раз запускать бэк часть, потому что Кафка не успевает запуститься до запуска сервиса Нотификаций -> сервис Нотификаций не запускается -> не приходят уведомления на почту)