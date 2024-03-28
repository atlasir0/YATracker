name: Deploy to Yandex Cloud Function

on:
  push:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v2

      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: '1.17'  

      - name: Install make
        run: sudo apt-get install make  # Установка make, если еще не установлен

      - name: Build
        run: |
          mkdir -p bin  # Создаем папку bin, если ее еще нет
          go build -o bin/main ./cmd/...  # Сборка проекта и помещение исполняемого файла в bin/main

      - name: Deploy
        runs-on: ubuntu-latest

        steps:
          - name: Checkout code
            uses: actions/checkout@v2

          - name: Set up yc-cli
            uses: yandex-cloud/yc-cli-action@v1
            with:
              yc-token: y0_AgAAAAAXBtATAATuwQAAAAEAG_rLAAChm-kV2qBHyLjAgwKtv43dr_yZ7Q

          - name: Build and deploy
            run: |
              # Деплой функции в Yandex Cloud
              yc serverless function version create \
                --function-name my-function \
                --runtime go113 \
                --entrypoint my-function \
                --memory 256m \
                --execution-timeout 5s \
                --source-path ./bin/main  # Указываем путь к исполняемому файлу в папке bin/main

              # Публикация версии функции
              yc serverless function version publish \
                --function-name my-function

          - name: Show deployment status
            run: |
           
              yc serverless function version list \
                --function-name my-function
