## Tasks

1. **Database**

   - [x] make db singleton
   - [x] service approach for db
   - [x] redis server for saving email verification tokens
   - [x] smtp server for emailing:
     - [ ] email check domain
   - [x] modify `Order` model to add price (with usd / uah / eur types)
   - [x] multiple categories

2. **Frontend**

   - [ ] popular orders:
     - [ ] image carousel
     - [ ] link to full order page
   - [ ] user panel organizer:
     - [x] create new sell product popup:
       - [x] form
       - [x] react hooks: inputs, selected categories hover, selected number near `Categories *`
       - [x] muatation
       - [ ] photos
     - [ ] my orders (selled / buyed) page??

3. **Controllers / API**

   - [x] middleware
     - [x] login user mutation
     - [x] register user mutation
   - [x] email server:
     - [x] email verification
     - [x] code generator
     - [x] code saver to **redis**

4. **Tests**

   - [x] db connect test
   - [ ] smtp server tests

5. **Features**

   - [ ] close on `Esc` button or click by mouse out of popup main area:

     - [ ] tracking mouse and keboard in `main` / `app` component
     - [ ] state for showing popup and give to props in other components

6. **Refactoring**

   - Patterns (3/5):

     - [x] singleton mongodb server connection
     - [x] builder redis server connection
     - [x] service layout for resolvers
     - [x] dependency injection in some services
     - [ ] dependency injection initialization of all servers (and mutations mb)

   - [x] 5 code smells (lab):

     - Bloaters:
       - [x] distribute lage methods in classes (dir `graph/resolvers` => `order_resolver`)
       - [x] distribute classes for feature (dir `libs` and `utils`)
       - [x] custom types (dir `constants` => `constants_order.go`)
     - Dispensables:
       - [x] check for dead code (deleted: dir `models` => `session_model`, dir `services` => `auth_service`)
     - [x] tests

   - [ ] 10 code smells / 10 techniques (project):

     - [ ] 10 techniques (explained):

       - composing methods:

         - [ ] extract method for smtp server config ()

     - [ ] tests:

       - [x] dir `~/tests`

     - [ ] 10 code smells (explained):

       - [x] 5 code smells from previous laboratory work ([commit link](<https://github.com/plxgwalker/e-commerce/blob/main/tasks.md#:~:text=5%20code%20smells%20(lab)%3A>))
       - Bloaters:

         - [ ] distribute one large component into few components ([commit link](https://github.com/plxgwalker/e-commerce/commit/30a1858e71d30ca8d78d55aac9684c35de8ed2fa))
         - [ ] distribute large method smtp server config for error handling ()

       - Dispensables:
         - [x] dead code: image creation ([commit link](https://github.com/plxgwalker/e-commerce/commit/691be520bccf5c0374979bce5fc8d68f6126051e))
         - [x] naming: files from `libs` move to `utils` because of same idea between folders ([commit link](https://github.com/plxgwalker/e-commerce/commit/66df879ac5a835f47e34632f8805bb89437a0f22))

     - [x] structure and docs:
       - [x] project architecture ([`README.md`](https://github.com/plxgwalker/e-commerce/blob/main/README.md))
