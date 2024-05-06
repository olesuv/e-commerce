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
   - [ ] auth tests
   - [ ] smtp server tests

5. **Refactoring**

   - [ ] implement patterns (3/5):

     - [x] singleton mongodb server connection
     - [x] builder redis server connection
     - [x] service layout for resolvers
     - [ ] dependency injection initialization of all servers (and mutations mb)

   - [ ] close on `Esc` button or click by mouse out of popup main area:
     - [ ] tracking mouse and keboard in `main` / `app` component
     - [ ] state for showing popup and give to props in other components

- [ ] 5 code smells:
  - Bloaters:
    - [ ] distribute lage methods in class
    - [ ] distribute classes for feature
    - [x] custom types (instead of `1`, `0`)
  - Object-Orientation Abusers:
    - [x] Temporary Field (used in resolvers inputs)
  - Dispensables:
    - [ ] check for dead code
  - [ ] tests
