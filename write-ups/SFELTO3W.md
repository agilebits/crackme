# BUSTing challenge <span>#</span>2

### Members

winxp5421  
blazer  
hops  
m33x  
milzo  
gearjunkie  
s3in!c  

### Writeup

Having successfully completed the first challenge we decided to give the second hash a shot. Although the method was the same software-wise we tried a slightly different approach because this time around we had prize money from winning first place. After seeing jpgoldberg’s cost breakdown and the difficulty of accurately calculating physical hardware costs, we decided to use part of the prize money towards attacking the second challenge. Instead of solely relying on our existing infrastructure (with hardware already bought and paid for by our team members), we decided to hit this challenge from a “we don't have any existing hardware” angle. We incorporated renting third party GPUs to demonstrate how a team without pre-existing infrastructure might go about attacking this challenge. We continued using the hardware we personally own but, will price these the same as if we had rented the GPUs through a third party. This should somewhat normalize the difficult calculation of “total cost”. This may not be the best way to tackle the absolute cost factor but it is the way we are going to do it.

Cloud Computing Services like Google Cloud Platform, Amazon Web Services, and Microsoft Azure offer GPUs but are extremely expensive to run at scale. It simply does not make financial sense and therefore not a practical choice in this case. e into play here after considering the cost.

||Google Cloud Platform|AWS (spot) ~$7/hr|AWS On-Demand|Azure|
|---|---|---|---|---|
|Card type|2x Tesla P100|8x V100|8x V100|4x V100|
|$/Chunk|$3.20|$0.69|$2.41|$2.44|
|Full Keyspace Cost|$58,995|$12,720|$44,375|$44,983|

*All prices include bit hint optimizations*

Fortunately, we were able to find a service that allows individuals to post compute machines up for rent. Dense compute machines as high as high as 11 x `1080 Ti's` for extremely cheap in comparison to your usual cloud computing partners were available to us. Of course, this adds a certain untrustworthiness as you have no idea where your data is going or who may have the ability to see that data. However, If cost is your only limitation this is a very cost-effective (in comparison) and legal way to amass what we would label as a substantial amount of computing resources. Seen as the general idea of this challenge is to get an estimation of how an attacker may go about attacking a particular vector of the 1Password software. We figured if this was being done as a malicious attack the attacker(s) would not really care about the security aspects of offboarding compute to a 3rd party and favor the cheapest option available. Keep in mind we are renting consumer grade GPUs, with no uptime or reliability guarantees. 

The total number of chunks we searched was `17,440` each containing `18436^2` keys totaling to `5,927,613,514,240` or roughly 95% of the complete search space of `6,266,140,065,856` (18,436 chunks). It is important to note that our wordlist varies slightly to that used to generate the actual passphrases for the contest and we also do not have ours in the same order as the one supplied. Our wordlist contains `18,436` items while the one supplied contains `18,328`. Our list contains `108` more lines than the one supplied, meaning we wasted work on phrases that were not used. Moreover, our list also is missing 36 words from the supplied one, so there was a minute chance that the passphrase was not in our search space and we would have to cover the difference of very small keyspace of `36,278,883,072`. In order to keep the calculations simple, we decided not to alter our task as we could easily run the difference after if it was required to do so.

We were able to cover the keyspace of `5,927,613,514,240` in 24 days 10 hours and 34 minutes. This yields an average speed of `1,135,999` checks (hashes) a second.

### Speeds

||Minimum|Average|Burst|Max|
|---|---|---|---|---|
|Speed|250 kH/s|1135 kH/s|1284 kH/s|1528.10 kH/s|
|GPU # Equivalent|26 GTX 1080 Ti|30 GTX 1080 Ti|70 GTX 1080 Ti|87 GTX 1080 Ti|

Due to the unpredictable availability of the rental instances, the number used fluctuated throughout the entire duration. The number of instances deployed was more of a gradual build-up. At the beginning we were inexperienced with rental compute nodes, as we progressed and became more familiarised with the system, we added more and more instances to push through towards the end.

### Actual Cost
| |3rd Party|Team Owned|Total Actual Cost|Total Adjusted Cost|
|---|---|---|---|---|
|Chunks Completed|12,258|5,182|17,440|17,440|
|Cost|$1943|$459|$2,402|$2,764|
|$/Chunk|$0.158508|$0.08857|$0.13772|$0.158508|

*Above costs are accounted for 95% of the keyspace being run*

### Whoops

It was at this point when we realized there could be a glitch with our distributed platform as we would have had to be extremely unlucky for the passphrase to be found in the last 5% of our search space. While verifying the task we were running was executed correctly. One of our members noticed that the hash we were attacking in the hashtopolis database was not imported into the database correctly. 

### Correct

|Hash|Salt|Password|
|---|---|---|
|sha256:ZmZCDobLoq58n2c1jwDSWA==:A3WMD8tTiARsK8lUr+KurbjyrFKW3wDDXdWuqST6MUY=|||

### Incorrect (The Mistake)

|Hash|Salt|Password|
|---|---|---|
|sha256|ZmZCDobLoq58n2c1jwDSWA==:A3WMD8tTiARsK8lUr+KurbjyrFKW3wDDXdWuqST6MUY=||

The agent would download and pass the hash correctly to hashcat, however, once the password was cracked the client would send the result formatted like:

|Hash|Salt|Password|
|---|---|---|
|sha256:ZmZCDobLoq58n2c1jwDSWA==:A3WMD8tTiARsK8lUr+KurbjyrFKW3wDDXdWuqST6MUY=||faint bust perturb|

The hash `sha256:ZmZCDobLoq58n2c1jwDSWA==:A3WMD8tTiARsK8lUr+KurbjyrFKW3wDDXdWuqST6MUY=` does not equal the hash value of `sha256` that was stored in the database and thus the found would be discarded and machines would continue to work through the keyspace. 

Once this mistake was identified, we combed through the Hashtopolis logs to see if we could determine which chunk the hash was cracked on. The Hashtopolis server logs hashcat exit status codes so identifying which chunk the hash was cracked on was fairly easy. Worst case we would have to rerun that particular chunk of keyspace. Lucky for us the agent that was issued the “winning” chunk had the client debug flag enabled which logs all input and output from both hashcat and Hashtopolis server. The hash was grep’ed out of the log file and promptly submitted.

When did we actually crack the hash? The hash was found in chunk #3371 approximately 11 days after we started work. ~13 days were wasted looking for a result we already found. We would like to thank Agilebits and the 1Password team for putting this contest on and we hope you guys enjoy the results of our second attack method.


In veneration,  
winxp5421, blazer, hops, m33x, milzo, gearjunkie, s3in!c

