# form-to-mail-service

This service is used for restaurant reservations.

The service allows you to receive an application from the registration form and notify the administrator to the selected mail.

The service is implemented using the ```smtp``` package.

Implemented support for yandex mail.

## Run
Before running you need to place ```.env``` file and your TLS certificates ```certificate.crt```, ```private.key``` in the root directory:

Sample file:
**.env**
```(e)
DOMAIN=localhost
SMTP_EMAIL=mail@ya.ru
SMTP_PASSWORD=password
```

Then use command ```make run``` to run module

## API

Base url: ```<DOMAIN>:8001/api/```

# Email sample
![image](https://user-images.githubusercontent.com/64990498/204975828-9b71eebd-f20a-4c4e-ac0d-8df81288e130.png)
