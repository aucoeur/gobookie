# gobookie
[![Go Report Card](https://goreportcard.com/badge/github.com/aucoeur/gobookie)](https://goreportcard.com/report/github.com/aucoeur/gobookie)  

> gobookie aims to help users quickly catalog, check prices for resale, donation, etc in one shot. Take a photo of your bookshelves and Bookie will grab title/author information using Google's Cloud Vision API and return additional metadata and check their current resale value.

> **INT 2.3:** [Spring Intensive](./docs/INT2.3-SpringIntensive.md) | [Proposal](./docs/INT2.3-Proposal.md)  
> **BEW 2.5:** [MakeUtilityProject](./docs/BEW2.5-MakeUtilityProject.md) | [Proposal](./docs/BEW2.5-Proposal.md)  

## To Run:
(Requires GCP credentials)
```shell
$ export GOOGLE_APPLICATION_CREDENTIALS=GOOGLE_SA_CREDENTIALS.json
$ go build -o ./gobookie
$ ./gobookie scan images/book_1.jpg
```
