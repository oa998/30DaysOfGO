### Repeated Phrases

Given a `corpus` of sentences (like a wikipedia article, book chapter, etc), look at all subphrases of each sentence and keep a count of all encountered phrases. Print out the Top 5 most observed phrases.

Phrases can not span multiple sentences:

> The cat was black. Black cats are unlucky.

`"was black black cats"` is **not** a valid phrase

`"cat was black"` is a valid phrase

example output:

> Top 5 phrases:
>
> - in the (10)
> - such as (9)
> - insertion sort (4)
> - used in (4)
> - of the (4)
