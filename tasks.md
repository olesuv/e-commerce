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

   - [x] 10 code smells / 10 techniques (project):

     - [ ] 10 techniques (explained):

       - composing methods:

         - [x] extract method for smtp server config ([commit link](https://github.com/plxgwalker/e-commerce/commit/fa1b89e05d0cdee69b3a6e34dc1b4e423049647c))
         - [x] merged to emailing methods in one to make `CreateUser` resolver more clear ([commit link](https://github.com/plxgwalker/e-commerce/commit/5239f1c8d3a0c0d7079d06e405c02f79e1d0a37f))
         - [x] replace temp with query (`setName`) and error handling on `login` resolver ([commit link](https://github.com/plxgwalker/e-commerce/commit/f88643a67fb44981f8adf2b467c554e237085538))

       - simplifying method calls:

         - [x] correct file naming for auth helpers ([commit link](https://github.com/plxgwalker/e-commerce/commit/547479e026a0225e72d9f3dfaeea51fbd82e530e))

       - organazing data:

         - [x] enums for `OrderCategory` and `OrderStatus` ([commit link](https://github.com/plxgwalker/e-commerce/commit/9efd1bea00fa5adec70b12392e491815e5e5417f))
         - [x] enums for `OrderCurrency` ([commit link](https://github.com/plxgwalker/e-commerce/commit/f4bffc7fcc19a38d3299e1f0558b8395772b9ef0))

     - [x] tests:

       - [x] dir `~/tests`
       - [x] utils tests

     - [x] 10 code smells (explained):

       - [x] 5 code smells from previous laboratory work ([commit link](<https://github.com/plxgwalker/e-commerce/blob/main/tasks.md#:~:text=5%20code%20smells%20(lab)%3A>))

       - Bloaters:

         - [x] distribute one large component into few components ([commit link](https://github.com/plxgwalker/e-commerce/commit/30a1858e71d30ca8d78d55aac9684c35de8ed2fa))
         - [x] distribute one large method for smtp server config and error handling ([commit link](https://github.com/plxgwalker/e-commerce/commit/fa1b89e05d0cdee69b3a6e34dc1b4e423049647c))
         - [x] distribute `CreateUser` user resolver with separated error handling ([commit link](https://github.com/plxgwalker/e-commerce/commit/5239f1c8d3a0c0d7079d06e405c02f79e1d0a37f))

       - Dispensables:
         - [x] dead code: image creation ([commit link](https://github.com/plxgwalker/e-commerce/commit/691be520bccf5c0374979bce5fc8d68f6126051e))
         - [x] naming: files from `libs` move to `utils` because of same idea between folders ([commit link](https://github.com/plxgwalker/e-commerce/commit/66df879ac5a835f47e34632f8805bb89437a0f22))

     - [x] structure and docs:
       - [x] project architecture ([`README.md`](https://github.com/plxgwalker/e-commerce/blob/main/README.md))
