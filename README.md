# gobookie - fun with Google Cloud Vision API :]
[![Go Report Card](https://goreportcard.com/badge/github.com/aucoeur/gobookie)](https://goreportcard.com/report/github.com/aucoeur/gobookie)  

> gobookie is an experimental project that aims to help users quickly catalog, check prices for resale, donation, etc in one shot. Take a photo of your bookshelves and Bookie will grab title/author information using Google's Cloud Vision API and return additional metadata and check their current resale value.

**INT 2.3:** [Spring Intensive](./docs/INT2.3-SpringIntensive.md) | [Proposal](./docs/INT2.3-Proposal.md)    
**BEW 2.5:** [MakeUtilityProject](./docs/BEW2.5-MakeUtilityProject.md) | [Proposal](./docs/BEW2.5-Proposal.md)  


## Some Experimental Processing
<img src="./sample/book1_processed.png" width="350" alt="edges" />    
<img src="./sample/books_1-annotated.png" width="350" alt="annotated" />    

## To Run:
(Requires GCP credentials)
```shell
$ export GOOGLE_APPLICATION_CREDENTIALS=GOOGLE_SA_CREDENTIALS.json
$ go build -o ./gobookie
$ ./gobookie scan images/book_1.jpg
```

# Resources
- A Technique to Detect Books from Library Bookshelf Image - https://imruljubair.github.io/papers/bookICCC2013.pdf
- Comparing Edge Detection Methods - https://medium.com/@nikatsanka/comparing-edge-detection-methods-638a2919476e
