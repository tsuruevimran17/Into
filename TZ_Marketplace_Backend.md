# Техническое задание (ТЗ) на разработку backend маркетплейса

## 1. Введение
**Цель:** разработать backend маркетплейса универсальных товаров (B2C/C2C) для одной страны/языка/валюты. Реализация только backend, без фронтенда и внешних интеграций на первом этапе. Командная разработка: 3 бэкендера.

**Ключевые решения:**
- Язык: Go (1.22+).
- HTTP-фреймворк: Gin.
- База данных: PostgreSQL 14+.
- ORM не используется: только чистые SQL-запросы, prepared statements.
- Архитектура: модульный монолит.
- Платежная модель: площадка принимает оплату, далее рассчитывается с продавцом (комиссионная модель).
- Интеграции с внешними платежными/логистическими системами — заглушки (интерфейсы для будущей реализации).
- Масштаб: до 100k пользователей.

---

## 2. Область и ограничения
**Входят в scope:**
- Аутентификация и профили пользователей.
- Онбординг продавцов.
- Каталог, категории, атрибуты, товары, SKU.
- Поиск и фильтрация.
- Корзина, оформление, заказы.
- Платежи (stub), расчёт комиссии, ledger.
- Отзывы.
- Админская часть (минимальные операции).
- Логи, метрики, аудит.

**Не входят в scope:**
- Фронтенд.
- Внешние платежные шлюзы.
- Реальная доставка.
- Продвинутые модерации.

---

## 3. Роли и доступы
- **Buyer**: просмотр каталога, оформление заказа, отзывы.
- **Seller**: управление товарами, просмотры заказов, баланс.
- **Admin**: управление пользователями/продавцами, просмотр аудита.

RBAC реализуется через middleware на уровне API.

---

## 4. Нефункциональные требования
- p95 latency: <= 300 мс для чтения, <= 600 мс для записи.
- Доступность: 99.5% на старте.
- Логи: структурированные JSON (zap/logrus).
- Метрики: Prometheus.
- Аудит: запись критических действий в audit_log.

---

## 5. Архитектура
**Модульный монолит** с явными доменными модулями:
- auth
- users
- sellers
- catalog
- search
- cart
- orders
- payments
- reviews
- admin
- notifications (stub)

### Общие требования к коду
- Разделение слоев: handler -> service -> repository.
- Только SQL-миграции.
- Обязательная идемпотентность в checkout/payments.

---

## 6. База данных (минимальный набор таблиц)
**Core:**
- users, user_profiles, user_addresses
- sellers, seller_profiles
- categories, attributes, category_attributes
- products, skus, product_media
- inventory
- carts, cart_items
- orders, order_items
- payments, payment_events
- ledger_entries, payouts
- reviews, product_ratings
- audit_log

Индексы:
- GIN индексы для search (текстовый поиск).
- Индексы на foreign keys.

---

# 7. ТЗ по каждому backend разработчику

## Backend Developer 1 — Auth, Users, Sellers, Admin
### Ответственность
- Аутентификация, профили пользователей.
- Онбординг продавцов.
- Базовая админка.

### Требования
1. **Auth**
   - Регистрация email/phone + пароль.
   - Login/logout/refresh.
   - Reset password.
2. **Users**
   - CRUD профиля.
   - Адреса доставки.
3. **Sellers**
   - Регистрация и статусы продавцов (pending/active/suspended).
4. **Admin**
   - Блокировка пользователей/продавцов.
   - Просмотр audit_log.

### Основные API
- POST /auth/register
- POST /auth/login
- POST /auth/logout
- POST /auth/refresh
- POST /auth/password/reset
- GET /users/me
- PUT /users/me
- POST /sellers
- GET /sellers/me
- PUT /sellers/me
- GET /admin/users
- PATCH /admin/users/{id}/status
- GET /admin/sellers
- PATCH /admin/sellers/{id}/status
- GET /admin/audit

### Acceptance Criteria
- Безопасное хранение паролей.
- Токены с истечением.
- RBAC для админских endpoint.

---

## Backend Developer 2 — Catalog, Search, Reviews
### Ответственность
- Каталог товаров и категории.
- Поиск.
- Отзывы.

### Требования
1. **Catalog**
   - CRUD категорий и атрибутов (admin).
   - CRUD товаров и SKU (seller).
2. **Search**
   - Фильтры по категориям, цене, атрибутам.
   - Текстовый поиск (GIN).
3. **Reviews**
   - Только после завершенного заказа.

### Основные API
- GET /categories
- POST /admin/categories
- POST /admin/attributes
- POST /seller/products
- GET /products/{id}
- GET /products
- PUT /seller/products/{id}
- POST /seller/products/{id}/skus
- GET /search
- POST /products/{id}/reviews
- GET /products/{id}/reviews

### Acceptance Criteria
- Продавец может создавать товары и SKU.
- Поиск работает с фильтрами.
- Отзывы разрешены только после delivered.

---

## Backend Developer 3 — Cart, Orders, Payments/Ledger
### Ответственность
- Корзина, оформление заказа.
- Платежи и ledger.

### Требования
1. **Cart**
   - CRUD корзины.
2. **Orders**
   - Статусы заказа: created -> paid -> shipped -> delivered -> closed.
   - Отмена/возврат (manual).
3. **Payments**
   - Payment intent, capture (stub).
   - Расчет комиссии.
   - Ledger операций.

### Основные API
- GET /cart
- POST /cart/items
- DELETE /cart/items/{id}
- POST /checkout
- GET /orders
- GET /orders/{id}
- POST /orders/{id}/cancel
- POST /payments/intent
- POST /payments/capture
- GET /seller/payouts

### Acceptance Criteria
- Идемпотентный checkout.
- Резервирование и возврат товара корректны.
- Ledger отражает комиссию и балансы.

---

## 8. Тестирование
- Unit tests для бизнес-логики.
- Integration tests с PostgreSQL.
- Контрактные тесты API.
- Минимум 5 e2e сценариев (регистрация → покупка → отзыв).

---

## 9. Сопутствующие артефакты
- OpenAPI спецификация.
- SQL миграции.
- Документация по развёртыванию (docker-compose).

---

## 10. Итог
ТЗ предназначено для распределённой разработки между 3 backend инженерами, каждая зона ответственности определена, зависимости минимизированы. Документ является базовой спецификацией для реализации backend маркетплейса на Go + Gin + PostgreSQL без ORM.
