# Cost calculations based on contest results

This directory contains some workspace that I (Jeffrey Goldberg) have been using to create tables for estimating cracking costs based on what we've learned through the challenge.

This guts of what I have is in the [R Markdown](https://rmarkdown.rstudio.com) file, [costs.Rmd](.costs.Rmd).
The result of knitting that file is in [costs.html](./cost.html), which is provided as I can't presume that most readers have the tools to build that from its source.

## Pitch for R Markdown

[Literate programming](https://en.wikipedia.org/wiki/Literate_programming) is the right way to go. It's fashionable to say that well written code should be its own documentation.  Literate programming can be taken either as the antithesis of that fashion or as the way to do it: The code and the text tell a story and produce a result.

There is no question that the simple calculations I've done here could just as easily (or perhaps more easily) be done with a spread sheet. But by putting this all is a textual file, it is far more reproducible, reusable, and sharable. It can be copied, forked, etc.