# Question 1

## Description
You are in charge of a display advertising program. Your ads are displayed
on websites all over the internet. You have some CSV input data that counts
how many times that users have clicked on an ad on each individual domain.

Every line consists of a click count and a domain name, like this:
```
counts = [ "900,google.com",
"60,mail.yahoo.com",
"10,mobile.sports.yahoo.com",
"40,sports.yahoo.com",
"300,yahoo.com",
"10,stackoverflow.com",
"20,overflow.com",
"5,com.com",
"2,en.wikipedia.org",
"1,m.wikipedia.org",
"1,mobile.sports",
"1,google.co.uk"]
```

Write a function that takes this input as a parameter and returns a data
structure containing the number of clicks that were recorded on each domain
AND each subdomain under it. For example, a click on "mail.yahoo.com"
counts toward the totals for "mail.yahoo.com", "yahoo.com", and "com".
(Subdomains are added to the left of their parent domain. So "mail" and
"mail.yahoo" are not valid domains. Note that "mobile.sports" appears as a
separate domain near the bottom of the input.)

## Examples
Sample output (in any order/format):
```javascript
calculateClicksByDomain(counts) =>
    com: 1345
    google.com: 900
    stackoverflow.com: 10
    overflow.com: 20
    yahoo.com: 410
    mail.yahoo.com: 60
    mobile.sports.yahoo.com: 10
    sports.yahoo.com: 50
    com.com: 5
    org: 3
    wikipedia.org: 3
    en.wikipedia.org: 2
    m.wikipedia.org: 1
    mobile.sports: 1
    sports: 1
    uk: 1
    co.uk: 1
    google.co.uk: 1
```

Clarifications
- The input will consist of domain names only, never paths. All domains will contain at
least two labels and one dot.
- The input will never deviate from the specified format; no counts or domain names
which appear in any element of the input will be null or empty. (Domains might not
be strictly alphabetical.)
- There will be no exact repetition of domains in the input (e.g., "yahoo.com" will not
appear twice, but "mail.yahoo.com" and "yahoo.com" may both appear)
- Numbers are all small enough that there will be no integer overflow issues, either in
the inputs or in the sum totals.
- Domain names have already been normalized to lowercase.
- In accordance with the “completeness over optimality” mantra, a candidate that
proposed a trie-based approach should be given a warning to consider something
more straightforward (but ultimately, they may proceed with that approach if they
wish).
- Count numbers will be greater than zero.

Example Solution
- Use a count map and split function

