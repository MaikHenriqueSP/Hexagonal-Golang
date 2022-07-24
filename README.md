:construction: This is a working in progress :construction:

:art:

### Business Rules
This implementation aims to simulate a library.

Each book must have an author, ISBN, title, publisher and publication date.

Each book may be loaned by a user.

Each user is identified by a full name and id.

For simplicity's sake, we'll assume that the library has an infinite number of books.

Each loan is identified by an user, a book, a laon date, a return date and an id. Which implies that only registered users are allowed to loan books.
By default, the return date is 15 days, but it can be extended throguh a request.

A user may return a loaned book.

The library also may receive request for registering new books.