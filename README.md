# hn
Hacker News Viewer, In Go!

## Usage

### CLI Mode

#### Top 25 Stories List

You can retrieve and print the Top Stories from Hacker News. If no argument is
provided to `hn top`, it will assume you would like the top 25 stories.

```shell
~/g/s/g/k/hn git:master ❯❯❯ hn top 25                                        ✱ ◼
1 Congratulations You’ve Been Fired
2 How a Car Engine Works
3 Domino's: Pizza and Payments
4 Support of OpenBSD pledge(2) in programming languages
5 Frog and Toad are Cofounders
6 Alphabet X new bipedal robot can climb stairs and overcome obstacles
7 The rotational North Pole is moving east at 14 cm per year
8 Aetna's CEO pays workers up to $500 to sleep
9 How to Avoid Empathy Burnout
10 STOMP – The Simple Text Oriented Messaging Protocol
11 The Blue State Model: How the Democrats Created a "Liberalism of the Rich"
12 Afghanistan’s only PC manufacturer
13 Rendering React without browser JavaScript
14 Declaring Reality: An Introduction to Datalog
15 Efficient Integer Overflow Checking in LLVM
16 When Free Software Depends on Nonfree
17 Bobby Fischer: from prodigy to pariah (2011)
18 Knuth versus Email (1999)
19 Steve Case: fewer tiny startups in the third wave
20 Not an ex-parrot
21 Ask HN: Is it worth paying for a Coursera course?
22 Qualcomm server chips now available to ARM developers through cloud service
23 How the Commodore Amiga Powered the Prevue Channel
24 Why babies are sleeping in boxes
25 Legal Moonshiner and University Battle Over Rights to ‘Kentucky’
```

#### Read A Top Story

To read, if it is a post which has local text a Hacker News Story, simply tell the tool which story you would like to
read, like so:

```shell
~/g/s/g/k/hn git:master ❯❯❯ hn top story 21                                  ✱ ◼
Ask HN: Is it worth paying for a Coursera course?
By: eecks, When: 2016-04-09 10:32:11 -0400 EDT

There are lots of free courses online but some offer the option of paying and
receiving a certificate on completion.

Does anyone here have one of those certs? Do they hold any value? Is it worth
paying for the cert when you can still do the full course for free?
```
