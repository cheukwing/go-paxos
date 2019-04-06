# go-paxos
Simple implementation of the Paxos consensus algorithm in Go to aid in my learning of the language, based on material and pseudocode prepared by [Naranker Dulay](http://wp.doc.ic.ac.uk/nd/).

## Debrief
To conclude my initial learning of Go, I decided to implement the Paxos consensus algorithm with inspiration from my Distributed Algorithms module earlier in the year.

Go is easy enough to write and its usage of struct methods offers an interesting balance between OOP and non-OOP, whose basics I found no trouble adapting to in this project.
Having previously worked with Elixir in making an implementation of MultiPaxos as part of the aforementioned Distributed Algorithms module, adapting my knowledge with working with Elixir process mailboxes to Go channels was fun, though I think I will need to practice further to fully understand the nuances of unbuffered channels (which I did not use in this project).

A problem I had with this project was the idea of exposed functions being a separate idea from the idea of encapsulation in OOP, and this may be potentially reflected in my code as I was unsure of which functions would need to be exposed if this package were to be an actual library.
This follows directly from my unsureness of the nuances of the struct methods, as the use of them encourages me to use an OOP style of writing and causes me to ignore the fact that in the end they are only an abstraction of a datatype, rather than an actual object - perhaps as a language it is too general and this makes it difficult for me to write in it without thinking about other language writing styles, or perhaps it is just the fault of my inexperience.

Another issue I had was with understanding the method of actually starting a Go project.
It is such a radically different approach (having a dedicated set of directories) to other languages that it caused me quite a bit of frustration when I first tried to start.

