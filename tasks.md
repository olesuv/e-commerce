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

   - Refactoring (project):

     - [ ] `libs` move to `utils` dir
     - Dispensables:
       - [x] dead code: image creation [link commit](https://github.com/plxgwalker/e-commerce/commit/691be520bccf5c0374979bce5fc8d68f6126051e)
