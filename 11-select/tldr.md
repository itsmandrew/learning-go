## Wrapping Up

### Select
- helps you wait on mutliple channels
- sometimes include time.After in one of your cases to prevent system for blocking forever

### httptest
- convenient way of creating test servers so you have have reliable and controllablee tests
- uses the same interfaces as the "real" net/http servers which is consistent and less for you to learn