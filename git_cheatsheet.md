🧭 GIT CHEAT SHEET — основные команды

=== 🔹 1. Создать новый репозиторий и залить в GitHub ===
git init
git add .
git commit -m "Первый коммит"
git branch -M main
git remote add origin https://github.com/asamults/ИМЯ_РЕПОЗИТОРИЯ.git
git push -u origin main

=== 🔹 2. Стандартное обновление проекта ===
git add .
git commit -m "описание изменений"
git push

=== 🔹 3. Получить изменения с GitHub ===
git pull origin main

=== 🔹 4. Обновить (с rebase) если были локальные изменения ===
git pull --rebase origin main

=== 🔹 5. Проверить статус и историю ===
git status
git log --oneline

=== 🔹 6. Работа с ветками ===
git branch              # показать ветки
git checkout -b new-branch   # создать новую ветку
git checkout main       # перейти на основную ветку
git merge new-branch    # объединить ветки

=== 🔹 7. Удалить или заменить репозиторий ===
git remote -v           # показать текущий репозиторий
git remote remove origin
git remote add origin https://github.com/asamults/new_repository.git

