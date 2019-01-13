# 1Password Challenge Write Up

### ID

5BSLBTKR

### Members

groozavu                     
ninjaslikecheese           

### Hardware

| GPU                     | count |
|-------------------------|------:|
| NVIDIA GeForce GTX 1070 | 6     |

* Note: The costs are roughly rounded.

### Costs

| Cost                       | Amount    |
|----------------------------|----------:|
| 6 GPU mining rig: GTX 1070 | ~3000 EUR |
| 5 months Electricity bill  | ~500 EUR  |


### Software used
* Hashcat 4.2.1 & 5.0.0
* OS: EthOS
* Candidate generators: Python/Golang

With this hardware during the first couple of months we had around 47 kH/s, but after we figured out it was bottlenecking (due to cheap CPU and RAM), we upgraded it to AMD FX-8350 CPU + 16GB RAM, and it went to a steady 57 kH/sec.

### Writeup

To give you guys a bit of background, we are two guys who own a basic mining rig, comprised of 6 NVIDIA GeForce GTX 1070 video cards, without HDD, just a USB flash drive running ethOS.

The whole story started around mid August 2018. I saw the challenge on Bugcrowd, so since we already had the mining rig, we decided to try something new, even though it was a very small chance for us to win. Basically the software set-up that we used 90% of the cracking time took around 2-3 weeks to build - mostly around 1 hour after work / day. When we decided to give this challenge a try, we thought that for fun and profit we should just focus on hazard. Due to the fact that the keyspace is humongous, we reasoned that randomly generating candidates would yield faster results than generating them one by one sorted alphabetically. We just had to make sure there was enough entropy so that we don't have a high duplicate rate. According to our calculations, we would need 1250 days to run through all the possible combinations: 18329^3 = 6,157,668,625,289 => 6157668625289/57000/86400 = 1250 days (keyspace/kH/seconds in a day).

First we experimented with different PRNGs, including BlumBlumShub, but then we realized that Python and any other scripting language are not fast enough to feed our hungry hashcat. At that point you guys made the hints available. Since we needed a way to generate passwords for a specific hint bit, modifying the actual Go code you guys provided looked like a good idea. Because it's a compiled language it should be quite fast, plus we'd also learn a bit of Go along the way.

Since we didn't have an HDD on the rig, we had to use another server, which was mainly generating the candidates and was running a NFS share which got mounted on the hashcat machine. Along the way we tweaked the generating speed by adding Haveged, because generating candidates was quickly drying up the entropy pool. Since we generated ~ 800GB lists per round, we removed the duplicates from the same lists, but due to lack of disk space to store all files, we had to discard files on every run, in turn we surely had duplicated tries along the way. We reasoned that it's negligible, since the duplication rate / same wordlist was very low ~ 0.0011% (the amount of duplicates/total candidate count of a given list).

When you guys released the second hints, we decided to focus on the hashes that have hint of `0b00`, since we had limited disk space, and the same wordlist might hold they answer to one or both of them - or not. The rest is history, not much to tell. I will however note the costs that we had along this cracking period, also including the mining rig hardware for reference, though it can be excluded from the total since we already owned it prior to starting.

I'm still amazed that we were successful in enough time to actually win, so I will share with you how we initially nicknamed our operation: HOB - Highly Opportunistic Bruteforce. Since the keyspace is so huge, and I would like to note that the key here was that we had all the precise details. Including the wordlist from which the passwords were generated. Right at the beginning of the process, when we didn't read up all the docs, we thought the wordlist wasn't provided. Therefore, we tried to recreate it using lists of English words found on the internet, which proved to be huge, and it didn't provide a guarantee that it includes all the words. The longer the seed list, the bigger keyspace gets, and I saw it can grow pretty fast. This proves how difficult it is to crack this kind of hash in a real world case scenario and how important the seed list is.

We gained a lot of experience in the process and we want to thank Agilebits & 1Password for the challenge, for their commitement in improving the current password practices.

Cheerio,       
groozavu, ninjaslikecheese         
