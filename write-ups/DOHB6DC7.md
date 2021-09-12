# 1Password Challenge Write Up

### Members

s3inlc  
winxp5421  
blazer  
hops

### Hardware


| GPU                        | count |
|----------------------------|------:|
| NVIDIA GeForce GTX 1070    | 10    |
| NVIDIA GeForce GTX 1080    | 4     |
| NVIDIA GeForce GTX 1080 Ti | 5     |
| NVIDIA GeForce GTX 980     | 1     |
| NVIDIA GeForce GTX 750 Ti  | 1     |
| Total                      | 21    |

*Not all resources were used continuously* 

We were able to sustain an average speed of 209.85 kH/s and cracked the PBKDF-HMAC-SHA256 100,000 round hash in 17d 16:33:53 after covering 18.71% of the whole keyspace.

### Software used
Hashtopolis  
Hashcat 4.1.0  
Candidate generator (PHP/C)

## Writeup 

After seeing that the second bit hint was released for the challenge, we calculated that worse case it would take us under 100 days to recover a single challenge hash. This made the challenge feasible and economically viable. There are many compute instances available online which offer enormous amounts of computer power such as those equipped with Tesla V100s; however, the cost of using these instances made them unappealing and simply not cost effective for this challenge. We used a total of 6 compute nodes Windows/Linux mix, which connected in remotely from the United States, Switzerland, and Australia. The group was able to achieve a sustained cracking speed of 209.85 kH/s. Since we cluster regular consumer grade hardware, we are able to keep operational costs low. These rigs costed us approximately $16.24 per day to run. In the worst case scenario it would cost us $1,640 USD to crack a single challenge hash. This may seem like a rather large difference in cost vs reward but, you must assess the risk of the challenge.
Our group decided to select a single hash to attack instead of multiple. We are “putting all of our eggs into one basket” it's totally possible another group would find the same hash we are attacking and end up with nothing for our efforts. The calculations above do not account for the hardware we have previously invested in for previous projects. We have existing hardware so the up-front cost was a non-issue. However, if this was accounted for, it would have been approximately $11,550 USD just for the GPUs excluding other components such as motherboard, processors, storage, power supplies, memory and cabling. Password cracking has always been an expensive hobby. As enthusiasts we love this community and those that are in it for the right reasons are happy to donate time, money, electricity as well as “Intellectual property” in the hopes of promoting better security.

S3in!c (one of our group members) is the developer of Hashtopolis, a distributive hashcat wrapper. Hashtopolis allows distributing hashcat workloads across the web to ‘agents’ while providing a very friendly user interface to manage our attacks. Hashtopolis allowed us to collectively keep an eye on progress, as well as alert the group of any problems that may arise with our equipment. One of the major benefits of using Hashtopolis allowed the nodes to be taken offline at any time e.g. for other tasks or maintenance and re-join the instance when it was ready and resume where it left off seamlessly. The brilliance of using Hashtopolis was that all the heavy lifting was done by an automated system. For instance, it would dispatch all the required files to the agent and spawn the necessary processes and track which candidates were tested thus preventing resource wastage by both eliminating idle resources and duplicated work. To run this specific task with Hashtopolis some small changes to the client were made. These changes mainly dealt with integrating our candidate generator. The server was used without any modifications.

For this challenge, we decided to code a custom candidate generator. The role of the 
candidate generator was to read the wordlists, concatenate the words from each list to produce 3-word phrases, hash the concatenated strings with SHA256, check whether the SHA256 hash matched the 2 bit hint provided for the PBKDF2 hash and finally, the generator would write the matching candidates to stdout. All pre-processed candidates from our generator were then piped to hashcat, where they were hashed and verified against our target hash. The original candidate generator was coded in PHP as that could be quickly thrown together to prototype things. However, when it was placed into production across the nodes, we observed that some GPUs were not fully saturated with work. This was mainly an issue for nodes that had many GPUs paired with a weaker processor. To resolve this bottleneck, the candidate generator was quickly re-written in C. Using a faster candidate generator with much less overhead allowed us to saturate all the GPUs used during the challenge. Since we were attacking hashes with a slow hashing routine of 100,000 rounds, I/O was not a bottleneck as even cards connected to 1x PCI-E risers were able to be fully utilized. 

We were quite lucky to have cracked this hash so quickly (18 days). This project really demonstrates the power of open-source software and just how powerful distributed computing with consumer grade hardware can be. At the same time, this also demonstrates how difficult it is to crack the hashed password phrases, without the bit hints it would have been near impossible to crack one of these hashes in a reasonable time frame while operating on a budget.

We want to thank Agilebits for hosting this challenge. It was certainly a fun project to work on. We appreciate everything you guys do for the password community. Go 1Password!!

In veneration,
s3inlc, winxp5421, blazer, hops
